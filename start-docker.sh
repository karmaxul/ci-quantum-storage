#!/bin/bash
# ─────────────────────────────────────────────────────────────────────────────
# HealChain Docker Stack Launcher
# Builds images, starts services, deploys contract, updates config.
# ─────────────────────────────────────────────────────────────────────────────
set -e

GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

log()     { echo -e "${CYAN}$1${NC}"; }
success() { echo -e "${GREEN}✅ $1${NC}"; }
warn()    { echo -e "${YELLOW}⚠  $1${NC}"; }
error()   { echo -e "${RED}❌ $1${NC}"; exit 1; }

# ── Configuration ─────────────────────────────────────────────────────────────
DEPLOYER_ADDRESS="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
DEPLOYER_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
FUND_AMOUNT_HEX="0x8AC7230489E80000"  # 10 ETH
GETH_PORT=8545
SERVICE_PORT=8080
GETH_TIMEOUT=60
SERVICE_TIMEOUT=30
PROJECT_DIR=~/ci-sha-project
ENV_FILE="$PROJECT_DIR/.env"

echo ""
echo -e "${GREEN}🐳 Starting HealChain Docker Stack...${NC}"
echo ""

# ── 0. Check geth-custom binary exists ───────────────────────────────────────
cd $PROJECT_DIR

if [ ! -f "./geth-custom" ]; then
  error "geth-custom binary not found in $PROJECT_DIR.\nBuild it first:\n  cd $PROJECT_DIR && go build -o geth-custom ./cmd/geth"
fi

# ── 1. Write baseline .env file ───────────────────────────────────────────────
# Docker Compose reads this automatically — keeps config in one place
log "Writing .env file..."
cat > "$ENV_FILE" <<EOF
GETH_URL=http://geth:8545
STORE_PRIVATE_KEY=b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291
CONTRACT_ADDRESS=0x0000000000000000000000000000000000000000
SEPOLIA_RPC_URL=https://eth-sepolia.g.alchemy.com/v2/1caYEBX_3TJfgMsFXR0Kk
SEPOLIA_CONTRACT_ADDRESS=0xBffb39930647DBB5971690ea6e9F0CF393307e64
ORACLE_PRIVATE_KEY=33993463fcb385615d67a1e1f703af73bf2f75fc1b7fa08fc3638f223749b205
EOF
success ".env written (contract address will be updated after deploy)"

# ── 2. Bring down any existing stack ─────────────────────────────────────────
log "Stopping existing containers..."
docker compose down 2>/dev/null || true

# ── 3. Build and start services ──────────────────────────────────────────────
log "Building and starting services..."
docker compose up --build -d

# ── 4. Wait for Geth ─────────────────────────────────────────────────────────
log "Waiting for Geth to be ready (timeout: ${GETH_TIMEOUT}s)..."
GETH_START=$SECONDS
GETH_READY=false

while [ $((SECONDS - GETH_START)) -lt $GETH_TIMEOUT ]; do
  if curl -sf http://localhost:$GETH_PORT \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}' \
    > /dev/null 2>&1; then
    GETH_READY=true
    success "Geth ready in $((SECONDS - GETH_START))s"
    break
  fi
  echo "  Waiting... ($((SECONDS - GETH_START))/${GETH_TIMEOUT}s)"
  sleep 2
done

[ "$GETH_READY" = false ] && error "Geth did not start. Check: docker compose logs geth"

# ── 5. Fund the deployer wallet ───────────────────────────────────────────────
log "Funding deployer wallet..."

DEV_ACCOUNT=$(curl -s -X POST http://localhost:$GETH_PORT \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' \
  | grep -o '"0x[0-9a-fA-F]*"' | head -1 | tr -d '"')

[ -z "$DEV_ACCOUNT" ] && error "Could not get dev account from Geth."

echo "  Dev account: $DEV_ACCOUNT"
echo "  Funding:     $DEPLOYER_ADDRESS"

curl -s -X POST http://localhost:$GETH_PORT \
  -H "Content-Type: application/json" \
  -d "{
    \"jsonrpc\": \"2.0\",
    \"method\": \"eth_sendTransaction\",
    \"params\": [{
      \"from\": \"$DEV_ACCOUNT\",
      \"to\": \"$DEPLOYER_ADDRESS\",
      \"value\": \"$FUND_AMOUNT_HEX\"
    }],
    \"id\": 1
  }" > /dev/null

sleep 2
success "Deployer wallet funded."

# ── 6. Deploy contract ────────────────────────────────────────────────────────
log "Deploying HealChainStorage contract..."
cd $PROJECT_DIR/foundry

DEPLOY_OUTPUT=$(forge script script/Deploy.s.sol \
  --rpc-url http://localhost:$GETH_PORT \
  --private-key $DEPLOYER_KEY \
  --broadcast 2>&1)

# Try specific label first, fall back to regex
CONTRACT=$(echo "$DEPLOY_OUTPUT" | grep "Contract Address:" | awk '{print $3}' | head -1)

if [ -z "$CONTRACT" ]; then
  CONTRACT=$(echo "$DEPLOY_OUTPUT" | grep -o '0x[0-9a-fA-F]\{40\}' | tail -1)
fi

[ -z "$CONTRACT" ] && { echo "$DEPLOY_OUTPUT"; error "Failed to extract contract address."; }

success "Contract deployed at: $CONTRACT"
cd $PROJECT_DIR

# ── 7. Update .env and restart service cleanly ────────────────────────────────
log "Updating .env with new contract address..."
cat > "$ENV_FILE" <<EOF
GETH_URL=http://geth:8545
STORE_PRIVATE_KEY=b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291
CONTRACT_ADDRESS=$CONTRACT
SEPOLIA_RPC_URL=https://eth-sepolia.g.alchemy.com/v2/1caYEBX_3TJfgMsFXR0Kk
SEPOLIA_CONTRACT_ADDRESS=0xBffb39930647DBB5971690ea6e9F0CF393307e64
ORACLE_PRIVATE_KEY=33993463fcb385615d67a1e1f703af73bf2f75fc1b7fa08fc3638f223749b205
EOF

success ".env updated with contract address: $CONTRACT"

# Restart just the service — Geth keeps running untouched
log "Restarting HealChain service with new contract address..."
docker compose up -d --no-deps rs-service

# ── 8. Wait for service ───────────────────────────────────────────────────────
log "Waiting for HealChain service (timeout: ${SERVICE_TIMEOUT}s)..."
SVC_START=$SECONDS
SVC_READY=false

while [ $((SECONDS - SVC_START)) -lt $SERVICE_TIMEOUT ]; do
  if curl -sf http://localhost:$SERVICE_PORT/health > /dev/null 2>&1; then
    SVC_READY=true
    success "Service ready in $((SECONDS - SVC_START))s"
    break
  fi
  echo "  Waiting... ($((SECONDS - SVC_START))/${SERVICE_TIMEOUT}s)"
  sleep 2
done

[ "$SVC_READY" = false ] && error "Service did not start. Check: docker compose logs rs-service"

# ── Done ──────────────────────────────────────────────────────────────────────
echo ""
echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${GREEN}  HealChain Docker Stack Ready!${NC}"
echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo ""
echo "  Geth:     http://localhost:$GETH_PORT"
echo "  Service:  http://localhost:$SERVICE_PORT"
echo "  Contract: $CONTRACT"
echo "  Env file: $ENV_FILE"
echo ""
echo "  Container logs:"
echo "    docker compose logs -f geth"
echo "    docker compose logs -f rs-service"
echo ""
echo -e "${CYAN}  Quick test:${NC}"
echo "    curl -s http://localhost:$SERVICE_PORT/health | python3 -m json.tool"
echo ""
echo -e "${CYAN}  Store data:${NC}"
echo "    curl -X POST http://localhost:$SERVICE_PORT/storeOnChain \\"
echo "      -H 'Content-Type: application/json' \\"
echo "      -d '{\"data\": \"4865616c436861696e2074657374\", \"label\": \"docker test\"}'"
echo ""

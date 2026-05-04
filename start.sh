#!/bin/bash
# ─────────────────────────────────────────────────────────────────────────────
# HealChain Stack Launcher
# Starts Geth, funds deployer, deploys contract, rebuilds and starts service.
# ─────────────────────────────────────────────────────────────────────────────
set -e

# ── Configuration ─────────────────────────────────────────────────────────────
# Update these if you change keys or network
DEPLOYER_ADDRESS="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
DEPLOYER_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
FUND_AMOUNT_HEX="0x8AC7230489E80000"   # 10 ETH in hex wei
GETH_PORT=8545
SERVICE_PORT=8080
GETH_TIMEOUT=30                         # seconds to wait for Geth
SERVICE_TIMEOUT=15                      # seconds to wait for service

# ── Colors ────────────────────────────────────────────────────────────────────
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

log()     { echo -e "${CYAN}$1${NC}"; }
success() { echo -e "${GREEN}✅ $1${NC}"; }
warn()    { echo -e "${YELLOW}⚠  $1${NC}"; }
error()   { echo -e "${RED}❌ $1${NC}"; exit 1; }

echo ""
echo -e "${GREEN}🚀 Starting HealChain Full Stack...${NC}"
echo ""

# ── 0. Clean up any existing processes ───────────────────────────────────────
log "Stopping any existing processes..."
pkill -f geth-custom 2>/dev/null && warn "Stopped existing Geth" || true
pkill -f healchain-service 2>/dev/null && warn "Stopped existing service" || true
sleep 2

# ── 1. Start Geth ─────────────────────────────────────────────────────────────
log "Starting Geth (dev mode)..."
~/ci-sha-project/geth-custom \
  --dev \
  --http \
  --http.api eth,web3,net,personal,debug \
  --http.corsdomain "*" \
  --http.addr 0.0.0.0 \
  --http.port $GETH_PORT \
  &> /tmp/geth.log &

GETH_PID=$!
echo "  Geth PID: $GETH_PID"

# ── 2. Wait for Geth with elapsed time tracking ───────────────────────────────
log "Waiting for Geth to be ready (timeout: ${GETH_TIMEOUT}s)..."
GETH_START=$SECONDS
GETH_READY=false

while [ $((SECONDS - GETH_START)) -lt $GETH_TIMEOUT ]; do
  if curl -s http://localhost:$GETH_PORT \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}' \
    > /dev/null 2>&1; then
    GETH_READY=true
    success "Geth ready in $((SECONDS - GETH_START))s"
    break
  fi
  echo "  Waiting... ($((SECONDS - GETH_START))/${GETH_TIMEOUT}s)"
  sleep 1
done

if [ "$GETH_READY" = false ]; then
  error "Geth did not start within ${GETH_TIMEOUT}s. Check /tmp/geth.log for details."
fi

# ── 3. Fund the deployer wallet ───────────────────────────────────────────────
log "Funding deployer wallet..."

# Detect dev account dynamically — works across Geth versions
DEV_ACCOUNT=$(curl -s -X POST http://localhost:$GETH_PORT \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' \
  | grep -o '"0x[0-9a-fA-F]*"' | head -1 | tr -d '"')

if [ -z "$DEV_ACCOUNT" ]; then
  error "Could not get dev account from Geth."
fi

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

# ── 4. Deploy contract ────────────────────────────────────────────────────────
log "Deploying HealChainStorage contract..."
cd ~/ci-sha-project/foundry

DEPLOY_OUTPUT=$(forge script script/Deploy.s.sol \
  --rpc-url http://localhost:$GETH_PORT \
  --private-key $DEPLOYER_KEY \
  --broadcast 2>&1)

# Try specific label first (most reliable)
CONTRACT=$(echo "$DEPLOY_OUTPUT" | grep "Contract Address:" | awk '{print $3}' | head -1)

# Fallback: last 40-char hex address in output
if [ -z "$CONTRACT" ]; then
  CONTRACT=$(echo "$DEPLOY_OUTPUT" | grep -o '0x[0-9a-fA-F]\{40\}' | tail -1)
fi

if [ -z "$CONTRACT" ]; then
  echo "$DEPLOY_OUTPUT"
  error "Failed to extract contract address. See deploy output above."
fi

success "Contract deployed at: $CONTRACT"

# ── 5. Update service config ──────────────────────────────────────────────────
log "Updating service config..."

sed -i "s|contractAddress = getEnv(\"CONTRACT_ADDRESS\", \".*\")|contractAddress = getEnv(\"CONTRACT_ADDRESS\", \"$CONTRACT\")|" \
  ~/ci-sha-project/healchain-service.go

# Verify the sed actually worked
if ! grep -q "$CONTRACT" ~/ci-sha-project/healchain-service.go; then
  error "sed failed — contract address not updated in healchain-service.go.\nUpdate CONTRACT_ADDRESS manually to: $CONTRACT"
fi

success "Contract address updated in healchain-service.go"

# ── 6. Rebuild the service ────────────────────────────────────────────────────
log "Rebuilding Go service..."
cd ~/ci-sha-project
go build -o healchain-service healchain-service.go
success "Service built."

# ── 7. Start the service ──────────────────────────────────────────────────────
log "Starting HealChain service..."
./healchain-service &> /tmp/healchain-service.log &
SVC_PID=$!
echo "  Service PID: $SVC_PID"

# Wait for service with elapsed time tracking
SVC_START=$SECONDS
SVC_READY=false

while [ $((SECONDS - SVC_START)) -lt $SERVICE_TIMEOUT ]; do
  if curl -s http://localhost:$SERVICE_PORT/health > /dev/null 2>&1; then
    SVC_READY=true
    success "Service ready in $((SECONDS - SVC_START))s"
    break
  fi
  echo "  Waiting for service... ($((SECONDS - SVC_START))/${SERVICE_TIMEOUT}s)"
  sleep 1
done

if [ "$SVC_READY" = false ]; then
  error "Service did not start within ${SERVICE_TIMEOUT}s. Check /tmp/healchain-service.log for details."
fi

# ── Done ──────────────────────────────────────────────────────────────────────
echo ""
echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${GREEN}  HealChain Stack Ready!${NC}"
echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo ""
echo "  Geth:     http://localhost:$GETH_PORT"
echo "  Service:  http://localhost:$SERVICE_PORT"
echo "  Contract: $CONTRACT"
echo ""
echo "  Logs:"
echo "    Geth:    tail -f /tmp/geth.log"
echo "    Service: tail -f /tmp/healchain-service.log"
echo ""
echo -e "${CYAN}  Quick test:${NC}"
echo "    curl -s http://localhost:$SERVICE_PORT/health | python3 -m json.tool"
echo ""
echo -e "${CYAN}  Store data:${NC}"
echo "    curl -X POST http://localhost:$SERVICE_PORT/storeOnChain \\"
echo "      -H 'Content-Type: application/json' \\"
echo "      -d '{\"data\": \"4865616c436861696e2074657374\", \"label\": \"test\"}'"
echo ""
# ⛓ HealChain — Self-Healing Blockchain Storage

> Reed-Solomon erasure coding meets on-chain storage on Ethereum.

HealChain is a custom Ethereum stack that adds **Reed-Solomon precompiles** directly to the EVM, enabling data to be encoded into redundant shards and stored on-chain. Even with partial shard loss, the original data can be fully recovered.

**Verified end-to-end flow:**
```
Raw data → RS Encode (EVM precompile) → On-chain shards → Retrieve → RS Decode → Original data ✅
```

---

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                        Browser UI                           │
│              frontend/index.html (dark/light)               │
└──────────────────────────┬──────────────────────────────────┘
                           │ HTTP
┌──────────────────────────▼──────────────────────────────────┐
│                  HealChain REST API                         │
│               healchain-service.go (:8080)                  │
│   /storeOnChain  /retrieve  /listRecords  /getMetadata      │
│   /delete  /health  /stats  /encode  /decode                │
└──────────┬───────────────────────────┬───────────────────────┘
           │ eth_call / sendTx         │ HTTP bridge
           │                           │
┌──────────▼───────────┐   ┌──────────▼───────────────────────┐
│   Custom Geth Node   │   │     Reed-Solomon RS Service       │
│   (healchain-geth)   │   │   (called by EVM precompiles)     │
│                      │   │                                   │
│  Precompiles:        │◄──┤  /encode  /decode                 │
│  0x0400 encode       │   │  /stabilize  /stats               │
│  0x0401 decode       │   └───────────────────────────────────┘
│  0x0402 stabilize    │
│  0x0403 stats        │   ┌───────────────────────────────────┐
│                      │   │    HealChainStorage Contract       │
│  Chain ID: 1337      ├──►│    (Solidity + Foundry)           │
│  Prague/Osaka forks  │   │    store / retrieve / delete      │
└──────────────────────┘   └───────────────────────────────────┘
```

---

## Quick Start

### Docker (Recommended)

```bash
git clone https://github.com/karmaxul/ci-quantum-storage.git
cd ci-quantum-storage
./start-docker.sh
```

This single command:
- Builds and starts custom Geth with RS precompiles
- Funds the deployer wallet
- Deploys `HealChainStorage` contract
- Starts the REST API service
- Confirms everything is healthy before exiting

Open `frontend/index.html` in your browser and point it at `http://localhost:8080`.

### Local (Non-Docker)

```bash
# Requires: Go 1.24+, Foundry, custom geth-custom binary
./start.sh
```

### Environment Variables

| Variable | Default | Description |
|---|---|---|
| `CONTRACT_ADDRESS` | `0x5FbDB...` | Deployed contract address |
| `STORE_PRIVATE_KEY` | dev key | Signing key for transactions |
| `GETH_URL` | `http://localhost:8545` | Geth RPC endpoint |
| `LISTEN_ADDR` | `:8080` | Service listen address |

---

## API Reference

### Store & Retrieve

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/storeOnChain` | Encode data with RS and store on-chain |
| `GET` | `/retrieve?id=N` | Retrieve and decode record by ID |
| `GET` | `/getMetadata?id=N` | Get full record metadata without decoding |
| `GET` | `/listRecords?page=0&limit=10` | Paginated record listing |
| `GET` | `/delete?id=N` | Delete record (owner only) |

### Service Info

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/health` | Geth status, last block, chain ID |
| `GET` | `/stats` | Total records, contract address |

### RS Precompile Bridge

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/encode` | Called internally by EVM encode precompile |
| `POST` | `/decode` | Called internally by EVM decode precompile |

### Example

```bash
# Store data (hex-encoded)
curl -X POST http://localhost:8080/storeOnChain \
  -H "Content-Type: application/json" \
  -d '{"data": "0x4865616c436861696e2074657374", "label": "my record"}'

# Response
{
  "status": "success",
  "recordId": "0",
  "tx": "0xabc...",
  "originalSize": "14",
  "encodedSize": "485",
  "retrieveUrl": "http://localhost:8080/retrieve?id=0"
}

# Retrieve
curl "http://localhost:8080/retrieve?id=0"

# Response
{
  "status": "success",
  "recordId": "0",
  "text": "HealChain test",
  "data": "0x4865616c436861696e2074657374",
  "bytes": 14
}
```

---

## Project Structure

```
ci-quantum-storage/
├── geth-custom                    # Compiled custom Geth binary
├── healchain-service.go           # Main Go REST API (v2.4)
├── start.sh                       # Local one-command launcher
├── start-docker.sh                # Docker one-command launcher
├── docker-compose.yml             # Docker Compose stack
├── Dockerfile.geth                # Custom Geth container
├── Dockerfile.service             # REST API container
├── frontend/
│   └── index.html                 # Browser UI (dark/light, search, pagination)
├── foundry/
│   ├── src/
│   │   ├── HealChainStorage.sol   # Main storage contract
│   │   └── HealRS.sol             # RS precompile library
│   ├── script/
│   │   └── Deploy.s.sol           # Deployment script
│   └── test/
│       └── HealChainStorage.t.sol # Contract tests
├── binding/
│   └── healchain_storage.go       # Generated Go ABI bindings
├── contracts/                     # Contract source copies
└── healchain-geth/                # Custom go-ethereum fork (submodule)
    └── core/vm/precompiles/
        └── healchain/
            ├── heal_rs.go         # RS precompile implementation
            └── ...
```

---

## Tech Stack

| Layer | Technology |
|---|---|
| EVM | Custom go-ethereum fork with RS precompiles |
| Smart Contracts | Solidity 0.8.20 + Foundry |
| Backend | Go 1.24 REST API |
| RS Encoding | Reed-Solomon via healchain package |
| Deployment | Docker Compose |
| Frontend | Vanilla HTML/JS (no framework) |

---

## Precompile Addresses

| Address | Function | Signature |
|---|---|---|
| `0x0400` | Encode | `encode(bytes,uint8,uint8) → bytes` |
| `0x0401` | Decode | `decode(bytes,uint8,uint8) → bytes` |
| `0x0402` | Stabilize | `stabilize(bytes,uint8,uint8) → bytes` |
| `0x0403` | Stats | `stats() → bytes` |

Default configuration: **10 data shards + 4 parity shards** (tolerates up to 4 shard failures).

---

## Roadmap

### Near-term
- [ ] Public testnet deployment (Sepolia) via wrapper contract architecture
- [ ] API key authentication
- [ ] Data compression before RS encoding
- [ ] Monitoring dashboard

### Medium-term
- [ ] IPFS-hosted frontend
- [ ] WalletConnect integration
- [ ] Cross-chain shard distribution
- [ ] Rate limiting and abuse protection

### Long-term
- [ ] Community governance
- [ ] Incentivized shard storage network
- [ ] Mobile client

---

## Development

```bash
# Build custom Geth
cd healchain-geth
go build -o ../geth-custom ./cmd/geth

# Build REST service
cd ..
go build -o healchain-service healchain-service.go

# Deploy contracts
cd foundry
forge script script/Deploy.s.sol \
  --rpc-url http://localhost:8545 \
  --private-key <key> \
  --broadcast

# Run contract tests
forge test --fork-url http://localhost:8545 -vvv
```

---

## Contributing

Contributions are welcome. Please open an issue before submitting a large PR so we can discuss the approach.

---

## License

MIT License — see [LICENSE](LICENSE) for details.

---

*Built for resilient, self-healing decentralized storage.*

# 🌌 HealChain — Self-Healing Blockchain Storage

**Reed-Solomon Erasure Coding + On-Chain Storage on Ethereum**

A complete hybrid self-healing storage system that encodes data with Reed-Solomon, stores encoded shards on-chain, and can recover data even if multiple shards are lost or corrupted.

---

## Current Status (May 2026)

**Fully Operational End-to-End**

- Custom Geth with Reed-Solomon precompiles (`0x0400`–`0x0403`)
- Go REST API service with full contract integration
- Solidity `HealChainStorage` contract (store, retrieve, metadata, delete)
- Docker-first deployment with one-command startup
- Working browser UI (store, retrieve, list, delete)

**Verified Flow**: Raw data → RS Encode (precompile) → On-chain storage → Retrieve + Decode ✅

---

## 🚀 Quick Start

### Docker (Recommended)

```bash
git clone https://github.com/karmaxul/ci-quantum-storage.git
cd ci-quantum-storage
./start-docker.sh

Local (Non-Docker)bash

cd ~/ci-sha-project
./start.sh

Quick Testbash

# Store data
curl -X POST http://localhost:8080/storeOnChain \
  -H "Content-Type: application/json" \
  -d '{"data": "4865616c436861696e2074657374", "label": "test"}'

# Retrieve data
curl "http://localhost:8080/retrieve?id=0"

Available EndpointsMethod
Endpoint
Description
POST
/storeOnChain
Encode + store on-chain
GET
/retrieve?id=XX
Retrieve + full decode
GET
/getMetadata?id=XX
Get full record metadata
GET
/listRecords
List all records (with preview)
DELETE
/delete?id=XX
Delete record (owner only)
GET
/health
System + Geth health status
GET
/stats
Statistics and total records

Project Structure

ci-quantum-storage/
├── geth-custom                    # Custom Geth with precompiles
├── healchain-service.go           # Main Go REST API
├── start.sh                       # Local one-command launcher
├── start-docker.sh                # Docker one-command launcher
├── docker-compose.yml
├── Dockerfile.geth
├── Dockerfile.service
├── foundry/                       # Solidity contracts + tests
├── frontend/                      # Browser UI
└── healchain/                     # Core Reed-Solomon library

Tech StackCustom Geth — Modified client-go with Reed-Solomon precompiles
Solidity + Foundry — Smart contract development and deployment
Go — High-performance REST service and contract integration
Docker — Containerized, reproducible deployment
Reed-Solomon — Erasure coding for data resilience

Potential RoadmapShort-term Final polish of browser UI (search, filters, dark/light mode)
Add pagination and improved UX to record listing
Better input validation and rate limiting

Medium-term Deploy to public testnet (Sepolia)
Add API key authentication
Data compression before encoding
Basic monitoring dashboard

Long-termFully decentralized frontend (IPFS + WalletConnect)
Cross-chain shard distribution
Community governance and incentives

ContributingContributions are welcome! Feel free to open issues or submit pull requests.LicenseMIT License — see LICENSE for details.Made with  for resilient decentralized storage


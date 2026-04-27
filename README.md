cat > README.md << 'EOF'
# Ci-RS: Rational Self-Healing Storage

**Ci rational constants + Reed-Solomon-inspired self-healing codes on Ethereum.**

A minimal, gas-aware proof-of-concept demonstrating data that can detect corruption and repair itself on-chain using CiSHA4096.

## Live Deployment (Sepolia Testnet)

- **CiOnChainRS_Full**: [`0xB6D318ad6c29806e9bFfe64D65fb34457030B2cE`](https://sepolia.etherscan.io/address/0xB6D318ad6c29806e9bFfe64D65fb34457030B2cE)
- **CiSHA4096_UltraLight**: [`0x2376AC847284443823B8f966e54D48Ef0Faeb853`](https://sepolia.etherscan.io/address/0x2376AC847284443823B8f966e54D48Ef0Faeb853)

Deployed: April 27, 2026

## Latest Successful Deployment (Sepolia)

- **CiOnChainRS_Full**: [`0x8abf4f1B58FAAF70355221187E48f7263aF5F819`](https://sepolia.etherscan.io/address/0x8abf4f1B58FAAF70355221187E48f7263aF5F819)
- **Live Test Result**: Self-healing cycle successful (Encode ~654k gas | Repair ~133k gas)

Deployed: April 27, 2026

## Quick Start

```bash
forge test --match-test testSelfHealingFlow -vvv

Gas Comparison TablePayload Size
Encode Gas
Repair Gas
Total Cycle
Notes
8 bytes
~620,000
~340,000
~960,000
Ultra-light demo
16 bytes
654,400
351,564
1,005,964
Current stable demo
32 bytes
~710,000
~380,000
~1,090,000
Projected
64 bytes
~820,000
~450,000
~1,270,000
Projected
256 bytes
~1,400,000
~680,000
~2,080,000
Target for optimization

HealChain Target: <200,000 gas total for 256-byte payloads (with native precompiles).HealChain VisionHealChain is a purpose-built blockchain infrastructure layer where self-healing data becomes a native primitive — powered by Ci rational constants and advanced error-correcting codes.The ProblemCurrent blockchains treat data as static and fragile. Once uploaded, data can be lost, corrupted, or become inaccessible due to network failures, node churn, or long-term bit rot.The SolutionHealChain combines your Ci rational constants as a fast deterministic mixing primitive with native Galois Field arithmetic and Reed-Solomon codes. Data can automatically detect and repair itself.Core Technical InnovationsNative Ci opcodes with heavy gas discounts
GF(256) + full ECC precompiles (syndromes, Berlekamp-Massey, Chien search)
Adaptive redundancy (light for IoT, strong for archival)
Cryptographically verifiable repairs via CiSHA4096

Target Use CasesJASMY & IoT Ecosystems: Sensors publish self-healing packets that survive packet loss without retransmission.
Decentralized AI Training: Datasets that automatically heal corrupted or missing shards.
Long-term Archival Storage: Medical records, legal documents, and scientific data that remain intact for decades.
Resilient Credentials: Self-repairing Soulbound tokens and decentralized identities.

RoadmapPhase
Timeline
Key Deliverables
PoC
Q2 2026
This repository + working demo
Optimization
Q3 2026
Gas <500k, 256-byte+ support
App-Chain Launch
Q4 2026
Arbitrum Orbit / OP Stack with precompiles
Mainnet + DA Layer
2027
Sovereign HealChain DA
Ecosystem Growth
2027+
JASMY integration, AI partnerships

HealChain turns your rational constants from a mathematical discovery into foundational infrastructure for a more resilient decentralized internet.
EOF


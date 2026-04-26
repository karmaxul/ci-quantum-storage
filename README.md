# Ci Quantum-Inspired Storage

**Powered by Ci-SHA4096 v2.4 (rational constant Ci = 85/27)**

A practical demonstration of rational balance in cryptography and self-healing storage.

## Live Links
- **Web Demo**: https://ci-quantum-storage-5gk9.onrender.com
- **GitHub**: https://github.com/karmaxul/ci-quantum-storage
- **Live Contract (Sepolia)**: `0x6Db61C27F196704519c7Eb6a6FaB1E017B7e0514`

## Gas Variants

| Variant            | Rounds | States | Gas (gasBenchmark) | Best For                          |
|--------------------|--------|--------|--------------------|-----------------------------------|
| **Moderate** (default) | 96   | 16     | ~2.3M – 2.9M      | Security-critical applications    |
| Aggressive         | 64     | 16     | ~1.96M            | General on-chain use              |
| **Ultra-Light**    | 32     | 8      | **~529k**         | High-frequency / cost-sensitive   |

All variants are available in the `variants/` folder.

## On-Chain Self-Healing (Reed-Solomon)

We have built fully on-chain Reed-Solomon prototypes that enable **no-oracle, trustless self-healing storage**:

| Contract                  | Description                              | Total Cycle Gas     | Status                  |
|---------------------------|------------------------------------------|---------------------|-------------------------|
| `CiOnChainRS.sol`         | Simple placeholder RS + Ci verification  | ~2.29M             | Stable                  |
| `CiOnChainRS_Full.sol`    | Realistic GF(256) polynomial math        | ~2.9M+             | Educational / research  |

**Philosophy**  
Traditional hashes (SHA-3, BLAKE3, etc.) are designed for maximum chaos. Ci-SHA4096 deliberately creates **structured repeating patterns** using the rational constant Ci = 85/27. These patterns pair exceptionally well with Reed-Solomon codes, enabling unusually efficient error correction and self-healing on-chain.

This is a practical demonstration of rational balance improving real-world primitives.

## How to Build & Test

```bash
forge clean
forge build
forge test --match-test testGasBenchmark -vvv
forge test --match-test testOnChainRSFullGas -vvv

Deployment (Example)bash

forge script script/Deploy.s.sol \
  --rpc-url https://rpc.sepolia.org \
  --private-key <your_test_key> \
  --broadcast \
  --verify

Created in collaboration with Grok (xAI) — April 2026


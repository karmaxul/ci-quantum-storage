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

## CiSHA4096 vs SHA-3 vs BLAKE3

| Aspect                    | SHA-3 (Keccak)           | BLAKE3                     | CiSHA4096 (Variants)             | Notes |
|---------------------------|--------------------------|----------------------------|----------------------------------|-------|
| Output Size               | 256/512 bit              | 256 bit (configurable)     | 4096 bit                         | Ci is much larger |
| Gas Cost (on-chain)       | ~150k–300k               | ~80k–200k (est.)           | 529k – 2.9M                      | BLAKE3 cheapest |
| Speed (software)          | Fast                     | Extremely fast             | Slower                           | BLAKE3 wins |
| Avalanche Quality         | Excellent                | Excellent                  | Good (~47–53%)                   | All strong |
| Repeating Patterns        | Avoids                   | Avoids                     | Deliberate “double-helix”        | Ci advantage |
| Error Correction Synergy  | None                     | None                       | Strong with Reed-Solomon (~39–43% savings) | **Ci wins** |
| Constants                 | Irrational               | ChaCha-derived             | Rational Ci = 85/27              | Philosophical edge |

## How to Build & Test

```bash
forge clean
forge build
forge test --match-test testGasBenchmark -vvv

Deployment (Example)bash

forge script script/Deploy.s.sol \
  --rpc-url https://rpc.sepolia.org \
  --private-key <your_test_key> \
  --broadcast \
  --verify

Philosophy
This project explores replacing irrational constants with rational Ci = 85/27. The repeating patterns create structured redundancy that enables efficient Reed-Solomon error correction while maintaining respectable avalanche properties.
Created in collaboration with Grok (xAI) — April 2026


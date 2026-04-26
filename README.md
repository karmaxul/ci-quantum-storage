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

## CiSHA4096 vs SHA-3 Comparison

| Aspect                    | SHA-3 (Keccak)                  | CiSHA4096 (Current Variants)              | Notes |
|---------------------------|---------------------------------|-------------------------------------------|-------|
| Output Size               | 256 / 512 bit                   | 4096 bit                                  | Ci much larger |
| Gas Cost (on-chain)       | ~150k–300k                      | 529k – 2.9M                               | SHA-3 cheaper |
| Avalanche Quality         | Excellent (~50%)                | Good (~47–53%)                            | Very close |
| Repeating Patterns        | Avoids them                     | Deliberate “double-helix” patterns        | Ci advantage |
| Error Correction Synergy  | None                            | Excellent with Reed-Solomon (~39–43% savings) | **Ci wins** |
| Constants                 | Irrational roots                | Rational Ci = 85/27                       | Philosophical difference |

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
This project explores replacing irrational constants with rational Ci = 85/27. The resulting repeating patterns create structured redundancy that enables efficient Reed-Solomon error correction while maintaining respectable avalanche properties.
Created in collaboration with Grok (xAI) — April 2026


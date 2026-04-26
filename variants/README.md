# CiSHA4096 Variants + On-Chain Reed-Solomon

## Core Hash Variants

| Variant            | Rounds | States | Gas (gasBenchmark) | Recommended Use Case                     |
|--------------------|--------|--------|--------------------|------------------------------------------|
| **Moderate** (default) | 96   | 16     | ~2.3M – 2.9M      | Security-critical, high-value operations |
| Aggressive         | 64     | 16     | ~1.96M            | General on-chain use                     |
| **Ultra-Light**    | 32     | 8      | **~529k**         | High-frequency / cost-sensitive calls    |

## On-Chain Reed-Solomon Prototypes

| Contract                  | Description                              | Total Cycle Gas     | Status                  |
|---------------------------|------------------------------------------|---------------------|-------------------------|
| `CiOnChainRS.sol`         | Simple placeholder RS + Ci verification  | ~2.29M             | Stable & working        |
| `CiOnChainRS_Full.sol`    | Realistic GF(256) polynomial math        | ~2.9M+ (placeholder)| Educational / research  |

**Philosophy**  
These contracts explore **completely trustless, no-oracle self-healing storage** on Ethereum. The structured repeating patterns from rational Ci = 85/27 give Reed-Solomon a unique practical advantage compared to traditional hashes.

All variants are ready for experimentation.

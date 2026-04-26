# CiSHA4096 Variants + On-Chain Reed-Solomon

## Gas Variants (Core Hash)

| Variant            | Rounds | States | Gas (gasBenchmark) | Best For                          |
|--------------------|--------|--------|--------------------|-----------------------------------|
| **Moderate** (default) | 96   | 16     | ~2.3M – 2.9M      | Security-critical applications    |
| Aggressive         | 64     | 16     | ~1.96M            | General on-chain use              |
| **Ultra-Light**    | 32     | 8      | **~529k**         | High-frequency / cost-sensitive   |

## On-Chain Reed-Solomon Prototypes

| Contract                  | Description                              | Total Cycle Gas     | Notes |
|---------------------------|------------------------------------------|---------------------|-------|
| `CiOnChainRS.sol`         | Simple placeholder RS + Ci verification  | ~2.29M             | Fastest prototype |
| `CiOnChainRS_Full.sol`    | Realistic GF(256) math + polynomial structure | ~2.91M        | Educational / closer to real RS |

**All variants maintain the rational Ci = 85/27 influence and repeating patterns.**

### Philosophy
These contracts explore **completely trustless self-healing storage** on Ethereum. The combination of Ci-SHA4096 structured patterns + Reed-Solomon gives a unique advantage for verifiable, repairable on-chain data.

Choose based on gas budget and desired realism.

# CiSHA4096 Gas Variants

This folder contains three optimized versions of the CiSHA4096 contract, offering different trade-offs between security strength, diffusion quality, and gas cost.

## Variant Comparison

| Variant            | Rounds | States | Gas (gasBenchmark) | Diffusion Strength | Recommended Use Case                     |
|--------------------|--------|--------|--------------------|--------------------|------------------------------------------|
| **Moderate** (default) | 96   | 16     | ~2.3M – 2.9M      | Highest            | Security-critical, high-value operations |
| **Aggressive**     | 64     | 16     | ~1.96M            | Strong             | General on-chain applications            |
| **Ultra-Light**    | 32     | 8      | **~529k**         | Good               | High-frequency calls, cost-sensitive     |

### Quick Notes
- All variants maintain the core rational Ci = 85/27 influence and repeating patterns.
- `src/CiSHA4096.sol` is a copy of the Moderate variant (default for easy testing/deployment).
- Ultra-Light returns `bytes32[8]` (2048-bit output) for maximum speed.

Choose based on your gas budget and security requirements.

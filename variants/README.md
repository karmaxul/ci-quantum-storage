# CiSHA4096 Gas Variants

This folder contains three optimized versions of CiSHA4096, balancing security strength vs gas cost.

## Variant Comparison

| Variant            | Rounds | States | Approx. Gas (gasBenchmark) | Avalanche Strength | Recommended Use Case                     |
|--------------------|--------|--------|----------------------------|--------------------|------------------------------------------|
| **Moderate** (default) | 96   | 16     | ~2.3M – 2.9M              | Highest            | Security-critical, high-value operations |
| **Aggressive**     | 64     | 16     | ~1.9M – 2.0M              | Strong             | General on-chain use                     |
| **Ultra-Light**    | 32     | 8      | **~500k – 550k**          | Good               | High-frequency calls, cost-sensitive     |

### Quick Start
- `src/CiSHA4096.sol` → Moderate (default)
- All variants support `ciSha4096()`, `gasBenchmark()`, and `verify()`
- Ultra-Light returns `bytes32[8]` (2048-bit) for maximum speed

### Philosophy
All variants maintain the core rational Ci = 85/27 influence and repeating patterns that enable strong Reed-Solomon efficiency in the Python/web layer.

Choose based on your gas budget and security needs.

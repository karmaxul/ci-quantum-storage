# HealChain EVM Precompiles Specification

**Version:** 0.1 (April 30, 2026)  
**Status:** Planning

## 1. Precompile Addresses (Proposed)

| Address          | Name                    | Purpose                              |
|------------------|-------------------------|--------------------------------------|
| `0x0000000000000000000000000000000000000400` | `HealRS.Encode`     | Reed-Solomon Encode + Compress       |
| `0x0000000000000000000000000000000000000401` | `HealRS.Decode`     | Reed-Solomon Decode + Decompress     |
| `0x0000000000000000000000000000000000000402` | `HealRS.Stabilize`  | Run stabilizer / healing simulation  |
| `0x0000000000000000000000000000000000000403` | `HealRS.Stats`      | Return overhead, shard info, timing  |

## 2. Input / Output Formats (ABI)

**Encode**
```solidity
function healEncode(
    bytes calldata data,
    uint8 dataShards,
    uint8 parityShards
) external view returns (bytes memory encoded, uint256 gasUsed);

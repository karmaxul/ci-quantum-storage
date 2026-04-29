// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract HealChainRS {
    uint8 public constant VERSION = 1;
    uint256 public constant HASH_SIZE = 32;
    address constant SHA256_PRECOMPILE = 0x0000000000000000000000000000000000000002;

    event HealChainVerified(bool success, uint256 originalLength);
    event HealChainCorruptionDetected(uint256 badShardCount);

    // Note: Not marked as `view` because it emits events
    function verify(bytes calldata encoded) external returns (bool success, uint256 originalLength) {
        if (encoded.length < 9) {
            return (false, 0);
        }

        uint8 version = uint8(encoded[0]);
        if (version != VERSION) {
            return (false, 0);
        }

        uint32 origLen = uint32(bytes4(encoded[1:5]));
        uint16 dataShards = uint16(bytes2(encoded[5:7]));
        uint16 parityShards = uint16(bytes2(encoded[7:9]));

        uint256 totalShards = uint256(dataShards) + uint256(parityShards);
        uint256 shardSize = calculateShardSize(origLen, dataShards);

        uint256 minSize = 9 + (totalShards * shardSize) + (totalShards * HASH_SIZE);
        if (encoded.length < minSize) {
            return (false, 0);
        }

        uint256 badShards = countBadShards(encoded, 9, shardSize, totalShards);

        bool isValid = badShards <= uint256(parityShards);

        if (isValid) {
            emit HealChainVerified(true, uint256(origLen));
        } else {
            emit HealChainCorruptionDetected(badShards);
        }

        return (isValid, uint256(origLen));
    }

    function calculateShardSize(uint32 originalLength, uint16 dataShards) internal pure returns (uint256) {
        return (uint256(originalLength) + dataShards - 1) / dataShards;
    }

    function countBadShards(
        bytes calldata encoded,
        uint256 dataStart,
        uint256 shardSize,
        uint256 totalShards
    ) internal view returns (uint256 badCount) {
        uint256 hashOffset = dataStart + (totalShards * shardSize);

        for (uint256 i = 0; i < totalShards; i++) {
            uint256 start = dataStart + i * shardSize;
            uint256 end = start + shardSize;
            if (end > encoded.length) {
                end = encoded.length;
            }

            bytes calldata shard = encoded[start:end];

            (bool ok, bytes memory hashBytes) = SHA256_PRECOMPILE.staticcall(shard);
            if (!ok || hashBytes.length != 32) {
                badCount++;
                continue;
            }

            bool hashMatch = true;
            for (uint256 j = 0; j < 32; j++) {
                if (hashBytes[j] != encoded[hashOffset + i * 32 + j]) {
                    hashMatch = false;
                    break;
                }
            }
            if (!hashMatch) {
                badCount++;
            }
        }
    }
}

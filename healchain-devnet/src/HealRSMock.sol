// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "./HealRS.sol";

/**
 * @title HealRSMock - Reliable mock with correct roundtrip
 */
contract HealRSMock is HealRS {

    function encode(
        bytes calldata data,
        uint8 dataShards,
        uint8 parityShards
    ) external view override returns (bytes memory encoded, uint256 gasUsed) {
        require(data.length > 0, "Empty data");
        require(dataShards > 0 && parityShards > 0, "Invalid shards");

        bytes memory simulated = abi.encodePacked(
            bytes("ENCODED_v1_"),
            data,
            bytes("_DS"),
            abi.encodePacked(uint8(dataShards), uint8(parityShards))
        );

        return (simulated, 42069);
    }

    function decode(
        bytes calldata encoded,
        uint8, // dataShards
        uint8  // parityShards
    ) external view override returns (bytes memory original, uint256 gasUsed) {
        require(encoded.length > 15, "Invalid encoded data");

        bytes memory prefix = bytes("ENCODED_v1_");
        uint256 start = prefix.length;

        // Find the position of "_DS" suffix
        bytes memory suffix = bytes("_DS");
        uint256 suffixPos = 0;
        for (uint i = start; i < encoded.length - suffix.length; i++) {
            if (encoded[i] == suffix[0] && 
                encoded[i+1] == suffix[1] && 
                encoded[i+2] == suffix[2]) {
                suffixPos = i;
                break;
            }
        }

        uint256 end = (suffixPos > start) ? suffixPos : encoded.length - 4;

        bytes memory originalData = new bytes(end - start);
        for (uint i = 0; i < originalData.length; i++) {
            originalData[i] = encoded[start + i];
        }

        return (originalData, 33777);
    }

    function stabilize(
        bytes calldata,
        uint8,
        uint8,
        uint8
    ) external pure override returns (bool success, uint256 healTimeMs) {
        return (true, 1245);
    }

    function getStats() external pure override returns (uint256 overheadPercent, uint256 totalBlocks) {
        return (42, 1337);
    }
}

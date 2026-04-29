// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface CiSHA4096 {
    function ciSha4096(bytes calldata data) external view returns (bytes32[16] memory);
}

interface CiRSRepair {
    function encode(bytes calldata payload, uint8 redundancy) 
        external view returns (bytes memory encoded, uint256 gasUsed);
    
    function decodeAndRepair(bytes calldata encoded) 
        external view returns (bytes memory recovered, bool success, uint256 gasUsed);
}

contract TestHealChain {
    CiSHA4096 public immutable ciSha = CiSHA4096(0x0000000000000000000000000000000000000c17);
    CiRSRepair public immutable repair = CiRSRepair(0x0000000000000000000000000000000000000c18);

    event SelfHealingTest(
        bytes original,
        bytes encoded,
        bytes recovered,
        bool success,
        uint256 encodeGas,
        uint256 repairGas
    );

    // Non-view because it emits an event
    function testSelfHealing(bytes calldata payload, uint8 redundancy) 
        external 
        returns (
            bytes memory recovered, 
            bool success, 
            uint256 encodeGas, 
            uint256 repairGas
        ) 
    {
        // Encode
        bytes memory encoded;
        (encoded, encodeGas) = repair.encode(payload, redundancy);

        // Simulate 1-byte corruption
        if (encoded.length > 10) {
            encoded[10] = bytes1(uint8(encoded[10]) ^ 0xFF);
        }

        // Repair
        (recovered, success, repairGas) = repair.decodeAndRepair(encoded);

        emit SelfHealingTest(payload, encoded, recovered, success, encodeGas, repairGas);
    }

    function hashWithCiSHA(bytes calldata data) external view returns (bytes32[16] memory) {
        return ciSha.ciSha4096(data);
    }
}

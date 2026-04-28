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

contract HealChainRS {
    CiSHA4096 public immutable ciSha = CiSHA4096(0x0000000000000000000000000000000000000c17);
    CiRSRepair public immutable repair = CiRSRepair(0x0000000000000000000000000000000000000c18);

    event SelfHealed(
        bytes original,
        bytes recovered,
        bool success,
        uint256 encodeGas,
        uint256 repairGas
    );

    // Main function: Encode + optional corruption simulation + Repair
    function testSelfHealing(bytes calldata payload, uint8 redundancy) 
        external 
        returns (bytes memory recovered, bool success, uint256 encodeGas, uint256 repairGas) 
    {
        // Encode
        bytes memory encoded;
        (encoded, encodeGas) = repair.encode(payload, redundancy);

        // Simulate corruption (real use case would be natural data corruption)
        if (encoded.length > 40) {
            encoded[40] = bytes1(uint8(encoded[40]) ^ 0xFF);
        }

        // Repair
        (recovered, success, repairGas) = repair.decodeAndRepair(encoded);

        emit SelfHealed(payload, recovered, success, encodeGas, repairGas);
        return (recovered, success, encodeGas, repairGas);
    }

    function encode(bytes calldata payload, uint8 redundancy) 
        external view returns (bytes memory, uint256) 
    {
        return repair.encode(payload, redundancy);
    }

    function decodeAndRepair(bytes calldata encoded) 
        external view returns (bytes memory, bool, uint256) 
    {
        return repair.decodeAndRepair(encoded);
    }

    function hash(bytes calldata data) external view returns (bytes32[16] memory) {
        return ciSha.ciSha4096(data);
    }
}

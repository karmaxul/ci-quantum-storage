// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/HealRSMock.sol";

contract HealRSTest is Test {
    HealRSMock public healrs;

    function setUp() public {
        healrs = new HealRSMock();
    }

    function testContractDeployment() public view {
        console.log("HealRSMock contract deployed at:", address(healrs));
        console.log("Encode precompile address:", healrs.HEAL_RS_ENCODE());
        console.log("Decode precompile address:", healrs.HEAL_RS_DECODE());
    }

    function testEncodeDecodeRoundtrip() public {
        bytes memory original = bytes("HealChain Self-Healing Storage Test Payload - Roundtrip Success!");

        console.log("Original data length:", original.length);

        // Test Encode
        (bytes memory encoded, uint256 gasEncode) = healrs.encode(original, 10, 4);
        console.log("Encoded length:", encoded.length);
        console.log("Gas used for encode:", gasEncode);

        // Test Decode
        (bytes memory decoded, uint256 gasDecode) = healrs.decode(encoded, 10, 4);
        console.log("Decoded length:", decoded.length);
        console.log("Gas used for decode:", gasDecode);

        // Verify roundtrip
        assertEq(keccak256(original), keccak256(decoded));
        console.log("Roundtrip encode -> decode successful!");
    }

    function testStabilize() public view {
        bytes memory sample = bytes("sample data");
        (bool success, uint256 healTime) = healrs.stabilize(sample, 10, 4, 2);
        
        assertTrue(success);
        console.log("Stabilizer heal time (mock):", healTime, "ms");
    }

    function testStats() public view {
        (uint256 overhead, uint256 blocks) = healrs.getStats();
        console.log("Overhead percent:", overhead, "%");
        console.log("Total blocks:", blocks);
    }
}

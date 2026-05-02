// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/HealRSMock.sol";

contract HealRSMockTest is Test {
    HealRSMock public healrs;

    function setUp() public {
        healrs = new HealRSMock();
    }

    function testEncodeDecodeRoundtrip() public {
        bytes memory original = bytes("This is a test payload for HealChain self-healing storage. It should survive encode -> decode!");

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

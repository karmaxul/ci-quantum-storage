// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../HealChainInterface.sol";

contract HealChainTest is Test {
    HealChainInterface public hc;

    function setUp() public {
        hc = new HealChainInterface();
    }

    function test_FullSelfHealingFlow() public {
        bytes memory original = bytes("Harmony HealChain self-healing test payload - should recover perfectly");

        // 1. Encode
        bytes memory encoded = hc.encode(original);
        console.log("Encoded length:", encoded.length);

        // 2. Corrupt (simulate transmission error)
        encoded[100] ^= 0xFF;

        // 3. Decode with healing
        bytes memory recovered = hc.decode(encoded);

        assertEq(recovered, original, "Self-healing failed");
        console.log("SUCCESS: Perfect recovery!");
    }

    function test_Verify() public {
        bytes memory original = bytes("Test payload");
        bytes memory encoded = hc.encode(original);

        (bool success, uint256 len) = hc.verify(encoded);
        assertTrue(success);
        assertEq(len, original.length);
        console.log("SUCCESS: Verify passed");
    }
}

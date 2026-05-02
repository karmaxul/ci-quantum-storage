// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";

contract HealChainPrecompileTest is Test {
    address constant HEAL_ENCODE = 0x0000000000000000000000000000000000000400;

    bytes constant TEST_DATA = hex"4865616c436861696e20746573742064617461";

    function test_Encode() public {
        (bool success, bytes memory encoded) = HEAL_ENCODE.call(abi.encodeWithSignature(
            "encode(bytes,uint8,uint8)",
            TEST_DATA,
            10,
            4
        ));

        require(success, "Call failed");
        require(encoded.length > 0, "Empty result");

        console.log("Encode success! Length:", encoded.length);
        console.logBytes(encoded);
    }
}

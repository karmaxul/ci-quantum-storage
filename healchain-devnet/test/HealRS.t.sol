// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/HealRS.sol";

contract HealRSTest is Test {
    HealRS public healrs;

    function setUp() public {
        healrs = new HealRS();
    }

    function testContractDeployment() public view {
        console.log("HealRS contract deployed at:", address(healrs));
        console.log("Encode precompile address:", healrs.HEAL_RS_ENCODE());
        console.log("Decode precompile address:", healrs.HEAL_RS_DECODE());
    }

    function testEncodeDecodeInterface() public {
        bytes memory testData = bytes("HealChain Self-Healing Storage Test Payload");

        // Expect revert because precompiles don't exist yet in normal Foundry/Anvil
        vm.expectRevert();
        healrs.encode(testData, 10, 4);
        
        console.log("HealRS interface test passed (precompile calls expected to fail for now)");
    }
}

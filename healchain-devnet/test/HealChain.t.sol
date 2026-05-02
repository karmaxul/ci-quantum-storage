// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "forge-std/Test.sol";

contract HealChainTest is Test {
    function testBasicMath() public {
        uint256 a = 2;
        uint256 b = 2;
        uint256 result = a + b;
        
        assertEq(result, 4);                    // Explicit uint256
        assertEq(result, uint256(4));           // Alternative way
        
        console.log("HealChain Foundry setup is working correctly");
        console.log("Basic math test passed");
    }
}

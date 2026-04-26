// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/CiSHA4096.sol";   // Points to Moderate by default

contract CiTest is Test {
    CiSHA4096 public ci;

    function setUp() public {
        ci = new CiSHA4096();
    }

    function testGasBenchmark() public view {
        bytes memory message = "test message for Ci-SHA4096";

        uint256 start = gasleft();
        bytes32[16] memory hash = ci.ciSha4096(message);
        uint256 gasUsed = start - gasleft();

        console.log("Gas used (Default - Moderate):", gasUsed);
        console.log("First 4 hashes (of 16):");
        for (uint256 i = 0; i < 4; i++) {
            console.logBytes32(hash[i]);
        }
    }
}

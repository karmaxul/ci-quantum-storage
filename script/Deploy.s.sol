// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/CiSHA4096.sol";

contract Deploy is Script {
    function run() external {
        vm.startBroadcast();

        CiSHA4096 ci = new CiSHA4096();

        console.log("CiSHA4096 (Moderate - Default) deployed at:", address(ci));

        vm.stopBroadcast();
    }
}

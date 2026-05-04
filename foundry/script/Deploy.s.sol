// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/HealChainStorage.sol";

contract Deploy is Script {
    function run() external {
        vm.startBroadcast();

        HealChainStorage store = new HealChainStorage();

        console.log("HealChainStorage deployed at:", address(store));

        vm.stopBroadcast();
    }
}

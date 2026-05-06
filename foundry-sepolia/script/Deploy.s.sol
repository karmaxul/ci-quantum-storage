// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/HealChainStorage.sol";

contract Deploy is Script {
    function run() external {
        // Oracle address — the wallet that will fulfill RS requests
        // This should be the address corresponding to STORE_PRIVATE_KEY
        // in your healchain-service configuration
        address oracle1 = vm.envAddress("ORACLE_ADDRESS");

        address[] memory initialOracles = new address[](1);
        initialOracles[0] = oracle1;

        vm.startBroadcast();

        HealChainStorage store = new HealChainStorage(initialOracles);

        console.log("HealChainStorage (multi-oracle) deployed at:", address(store));
        console.log("Initial oracle:", oracle1);
        console.log("Chain ID:", block.chainid);
        console.log("Owner:", msg.sender);

        vm.stopBroadcast();
    }
}

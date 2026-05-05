// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/HealChainStorage.sol";

contract Deploy is Script {
    function run() external {
        // Oracle address — the wallet that will fulfill RS requests
        // This should be the address corresponding to STORE_PRIVATE_KEY
        // in your healchain-service configuration
        address oracle = vm.envAddress("ORACLE_ADDRESS");

        vm.startBroadcast();

        HealChainStorage store = new HealChainStorage(oracle);

        console.log("HealChainStorage deployed at:", address(store));
        console.log("Oracle address:", oracle);
        console.log("Chain ID:", block.chainid);

        vm.stopBroadcast();
    }
}

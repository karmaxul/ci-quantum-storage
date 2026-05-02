// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import "../src/HealRSMock.sol";

contract DeployHealChainDevnet is Script {
    function run() external {
        vm.startBroadcast();

        HealRSMock healrs = new HealRSMock();
        
        console.log("HealChain Devnet with Mock Precompiles Ready!");
        console.log("HealRSMock deployed at:", address(healrs));
        console.log("Encode precompile address:", healrs.HEAL_RS_ENCODE());
        console.log("Decode precompile address:", healrs.HEAL_RS_DECODE());

        vm.stopBroadcast();
    }
}

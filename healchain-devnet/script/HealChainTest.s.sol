// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import "../src/HealRS.sol";

contract HealChainTest is Script {
    HealRS public healrs = HealRS(0x0000000000000000000000000000000000000400);

    function run() external {
        vm.startBroadcast();

        bytes memory data = bytes("HealChain Self-Healing Test Payload");

        console.log("Original length:", data.length);

        (bytes memory encoded, uint256 gasE) = healrs.encode(data, 10, 4);
        console.log("Encoded length:", encoded.length);

        (bytes memory decoded, uint256 gasD) = healrs.decode(encoded, 10, 4);
        console.log("Decoded length:", decoded.length);

        console.log("Roundtrip success:", keccak256(data) == keccak256(decoded));

        vm.stopBroadcast();
    }
}

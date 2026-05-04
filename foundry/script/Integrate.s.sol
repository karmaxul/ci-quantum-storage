// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/HealChainStorage.sol";

contract Integrate is Script {
    function run() external {
        vm.startBroadcast();

        HealChainStorage store = new HealChainStorage();
        console.log("Deployed HealChainStorage at:", address(store));

        // Test 1: Basic store + retrieve
        bytes memory data = bytes("HealChain integration test payload");
        uint256 id = store.store(data, "integration test");
        console.log("Stored record ID:", id);

        bytes memory restored = store.retrieve(id);
        console.log("Restored:", string(restored));
        require(keccak256(restored) == keccak256(data), "Data mismatch");
        console.log("PASS: store + retrieve");

        // Test 2: Verified retrieve
        uint256 id2 = store.store(bytes("verified data"), "verified");
        bytes memory restored2 = store.retrieveVerified(id2);
        require(keccak256(restored2) == keccak256(bytes("verified data")), "Verified mismatch");
        console.log("PASS: retrieveVerified");

        // Test 3: Custom shards
        uint256 id3 = store.store(bytes("custom shards test"), 6, 3, "custom");
        bytes memory restored3 = store.retrieve(id3);
        require(keccak256(restored3) == keccak256(bytes("custom shards test")), "Custom shards mismatch");
        console.log("PASS: custom shards");

        // Test 4: Metadata
        (, uint256 origSize, uint256 encSize,,,,,) = store.getMetadata(id);
        require(origSize == data.length, "originalSize wrong");
        require(encSize > origSize, "encoded not larger");
        console.log("PASS: metadata - original:", origSize, "encoded:", encSize);

        // Test 5: Total records
        require(store.totalRecords() >= 3, "totalRecords wrong");
        console.log("PASS: totalRecords > 3");

        console.log("All integration tests passed!");

        vm.stopBroadcast();
    }
}
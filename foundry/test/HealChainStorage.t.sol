// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/HealChainStorage.sol";
import "../src/HealRS.sol";

/**
 * @title HealChainStorageTest
 * @notice Integration tests for HealChainStorage + HealRS precompiles.
 *
 * Run against your custom devnet:
 *   forge test --fork-url http://localhost:8545 -vvv
 *
 * All tests that call encode/decode require the precompile service to be running.
 */
contract HealChainStorageTest is Test {

    HealChainStorage public store;

    // Test accounts
    address public alice = makeAddr("alice");
    address public bob   = makeAddr("bob");

    // Fund accounts and deploy contract before each test
    function setUp() public {
        vm.deal(alice, 10 ether);
        vm.deal(bob,   10 ether);
        vm.etch(address(0x0400), hex"00");
        vm.etch(address(0x0401), hex"00");
        vm.etch(address(0x0402), hex"00");
        vm.etch(address(0x0403), hex"00");
        store = new HealChainStorage();
    }

    // ── Basic store + retrieve ────────────────────────────────────────────────

    function test_StoreAndRetrieve() public {
        bytes memory data = bytes("HealChain test");

        vm.prank(alice);
        uint256 id = store.store(data, "first record");

        bytes memory restored = store.retrieve(id);

        assertEq(restored, data, "restored data should match original");
        console.log("ID:", id);
        console.log("Restored:", string(restored));
    }

    function test_StoreAndRetrieveVerified() public {
        bytes memory data = bytes("verified integrity test");

        vm.prank(alice);
        uint256 id = store.store(data, "verified record");

        bytes memory restored = store.retrieveVerified(id);
        assertEq(restored, data, "verified restore should match original");
    }

    function test_StoreWithCustomShards() public {
        bytes memory data = bytes("custom shard config");

        vm.prank(alice);
        uint256 id = store.store(data, 6, 3, "custom shards");

        bytes memory restored = store.retrieve(id);
        assertEq(restored, data, "custom shard restore should match");
    }

    // ── Multiple records ──────────────────────────────────────────────────────

    function test_MultipleRecords() public {
        bytes memory dataA = bytes("record alpha");
        bytes memory dataB = bytes("record beta");
        bytes memory dataC = bytes("record gamma");

        vm.startPrank(alice);
        uint256 idA = store.store(dataA, "alpha");
        uint256 idB = store.store(dataB, "beta");
        uint256 idC = store.store(dataC, "gamma");
        vm.stopPrank();

        assertEq(store.retrieve(idA), dataA);
        assertEq(store.retrieve(idB), dataB);
        assertEq(store.retrieve(idC), dataC);
        assertEq(store.totalRecords(), 3);
    }

    function test_MultipleOwners() public {
        bytes memory aliceData = bytes("alice's secret");
        bytes memory bobData   = bytes("bob's secret");

        vm.prank(alice);
        uint256 aliceId = store.store(aliceData, "alice record");

        vm.prank(bob);
        uint256 bobId = store.store(bobData, "bob record");

        assertEq(store.retrieve(aliceId), aliceData);
        assertEq(store.retrieve(bobId),   bobData);
    }

    // ── Metadata ──────────────────────────────────────────────────────────────

    function test_Metadata() public {
        bytes memory data  = bytes("metadata check");
        string memory label = "my label";

        vm.prank(alice);
        uint256 id = store.store(data, label);

        (
            bytes32 dataHash,
            uint256 originalSize,
            uint256 encodedSize,
            uint8   dataShards,
            uint8   parityShards,
            address owner,
            uint256 timestamp,
            string memory gotLabel
        ) = store.getMetadata(id);

        assertEq(dataHash,     keccak256(data));
        assertEq(originalSize, data.length);
        assertGt(encodedSize,  data.length,   "encoded should be larger than original");
        assertEq(dataShards,   10);
        assertEq(parityShards, 4);
        assertEq(owner,        alice);
        assertGt(timestamp,    0);
        assertEq(gotLabel,     label);

        console.log("Original size:", originalSize);
        console.log("Encoded size: ", encodedSize);
    }

    function test_GetEncoded() public {
        bytes memory data = bytes("raw shard blob");

        vm.prank(alice);
        uint256 id = store.store(data, "blob test");

        bytes memory encoded = store.getEncoded(id);
        assertGt(encoded.length, 0, "encoded blob should not be empty");
        assertGt(encoded.length, data.length, "encoded should be larger than raw");

        console.log("Encoded length:", encoded.length);
    }

    function test_TotalRecords() public {
        assertEq(store.totalRecords(), 0);

        vm.startPrank(alice);
        store.store(bytes("a"), "a");
        store.store(bytes("b"), "b");
        store.store(bytes("c"), "c");
        vm.stopPrank();

        assertEq(store.totalRecords(), 3);
    }

    // ── Delete ────────────────────────────────────────────────────────────────

    function test_OwnerCanDelete() public {
        vm.prank(alice);
        uint256 id = store.store(bytes("to delete"), "delete me");

        vm.prank(alice);
        store.remove(id);

        vm.expectRevert();
        store.retrieve(id);
    }

    function test_NonOwnerCannotDelete() public {
        vm.prank(alice);
        uint256 id = store.store(bytes("alice's data"), "owned");

        vm.prank(bob);
        vm.expectRevert(abi.encodeWithSelector(HealChainStorage.NotOwner.selector, id));
        store.remove(id);
    }

    function test_TotalRecordsDoesNotDecrementOnDelete() public {
        vm.prank(alice);
        uint256 id = store.store(bytes("temp"), "temp");

        vm.prank(alice);
        store.remove(id);

        // totalRecords counts all ever created, not current active
        assertEq(store.totalRecords(), 1);
    }

    // ── Error cases ───────────────────────────────────────────────────────────

    function test_EmptyDataReverts() public {
        vm.prank(alice);
        vm.expectRevert(HealChainStorage.EmptyData.selector);
        store.store(bytes(""), "empty");
    }

    function test_EmptyDataWithShardsReverts() public {
        vm.prank(alice);
        vm.expectRevert(HealChainStorage.EmptyData.selector);
        store.store(bytes(""), 10, 4, "empty custom");
    }

    function test_RetrieveNonexistentReverts() public {
        vm.expectRevert(abi.encodeWithSelector(HealChainStorage.RecordNotFound.selector, 9999));
        store.retrieve(9999);
    }

    function test_GetMetadataNonexistentReverts() public {
        vm.expectRevert(abi.encodeWithSelector(HealChainStorage.RecordNotFound.selector, 9999));
        store.getMetadata(9999);
    }

    function test_GetEncodedNonexistentReverts() public {
        vm.expectRevert(abi.encodeWithSelector(HealChainStorage.RecordNotFound.selector, 9999));
        store.getEncoded(9999);
    }

    function test_RemoveNonexistentReverts() public {
        vm.expectRevert(abi.encodeWithSelector(HealChainStorage.RecordNotFound.selector, 9999));
        store.remove(9999);
    }

    // ── Events ────────────────────────────────────────────────────────────────

    function test_StoreEmitsEvent() public {
        bytes memory data = bytes("event test");
        bytes32 expectedHash = keccak256(data);

        vm.prank(alice);
        vm.expectEmit(true, true, false, false);
        emit HealChainStorage.Stored(0, alice, expectedHash, data.length, 0, "event label");

        store.store(data, "event label");
    }

    function test_DeleteEmitsEvent() public {
        vm.prank(alice);
        uint256 id = store.store(bytes("bye"), "bye");

        vm.prank(alice);
        vm.expectEmit(true, true, false, false);
        emit HealChainStorage.RecordDeleted(id, alice);

        store.remove(id);
    }

    // ── Larger data ───────────────────────────────────────────────────────────

    function test_LargerPayload() public {
        // 200 bytes of data
        bytes memory data = new bytes(200);
        for (uint i = 0; i < 200; i++) {
        // forge-lint: disable-next-line(unsafe-typecast)
            data[i] = bytes1(uint8(i % 256));
        }

        vm.prank(alice);
        uint256 id = store.store(data, "large payload");

        bytes memory restored = store.retrieve(id);
        assertEq(restored, data, "large payload should restore correctly");
        console.log("Large payload restored, length:", restored.length);
    }
}

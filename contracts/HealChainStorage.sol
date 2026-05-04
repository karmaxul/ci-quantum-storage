// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./HealRS.sol";

/**
 * @title HealChainStorage
 * @notice Store data on-chain as Reed-Solomon shards.
 *         Any stored record can survive up to `parityShards` corrupted shards
 *         and still be fully recovered via the decode precompile.
 *
 * Flow:
 *   store(data)  → encodes via precompile → saves shards + metadata on-chain
 *   retrieve(id) → decodes via precompile → returns original bytes
 *   verify(id)   → re-decodes and checks hash matches original
 */
contract HealChainStorage {

    // ── Types ─────────────────────────────────────────────────────────────────

    struct Record {
        bytes   encoded;        // RS shard blob
        bytes32 dataHash;       // keccak256 of original data
        uint256 originalSize;   // original byte length
        uint8   dataShards;
        uint8   parityShards;
        address owner;
        uint256 timestamp;
        string  label;          // optional human-readable tag
    }

    // ── Storage ───────────────────────────────────────────────────────────────

    mapping(uint256 => Record) private _records;
    uint256 private _nextId;

    // ── Events ────────────────────────────────────────────────────────────────

    event Stored(
        uint256 indexed id,
        address indexed owner,
        bytes32 dataHash,
        uint256 originalSize,
        uint256 encodedSize,
        string  label
    );

    event Retrieved(
        uint256 indexed id,
        address indexed caller,
        bool    verified
    );

    event RecordDeleted(uint256 indexed id, address indexed owner);

    // ── Errors ────────────────────────────────────────────────────────────────

    error NotOwner(uint256 id);
    error RecordNotFound(uint256 id);
    error VerificationFailed(uint256 id, bytes32 expected, bytes32 got);
    error EmptyData();

    // ── Write functions ───────────────────────────────────────────────────────

    /**
     * @notice Encode and store data with custom shard config.
     * @param data          Raw bytes to store.
     * @param dataShards    RS data shards (e.g. 10).
     * @param parityShards  RS parity shards (e.g. 4). Tolerates this many failures.
     * @param label         Optional label for this record.
     * @return id           Record ID for future retrieval.
     */
    function store(
        bytes calldata data,
        uint8 dataShards,
        uint8 parityShards,
        string calldata label
    ) external returns (uint256 id) {
        if (data.length == 0) revert EmptyData();

        bytes memory encoded = HealRS.encode(data, dataShards, parityShards);

        id = _nextId++;

        _records[id] = Record({
            encoded:      encoded,
            dataHash:     keccak256(data),
            originalSize: data.length,
            dataShards:   dataShards,
            parityShards: parityShards,
            owner:        msg.sender,
            timestamp:    block.timestamp,
            label:        label
        });

        emit Stored(id, msg.sender, keccak256(data), data.length, encoded.length, label);
    }

    /**
     * @notice Store with default shard config (10 data, 4 parity).
     */
    function store(
        bytes calldata data,
        string calldata label
    ) external returns (uint256 id) {
        if (data.length == 0) revert EmptyData();

        bytes memory encoded = HealRS.encode(data);

        id = _nextId++;

        _records[id] = Record({
            encoded:      encoded,
            dataHash:     keccak256(data),
            originalSize: data.length,
            dataShards:   HealRS.DEFAULT_DATA_SHARDS,
            parityShards: HealRS.DEFAULT_PARITY_SHARDS,
            owner:        msg.sender,
            timestamp:    block.timestamp,
            label:        label
        });

        emit Stored(id, msg.sender, keccak256(data), data.length, encoded.length, label);
    }

    /**
     * @notice Delete a record. Only the original owner can delete.
     */
    function remove(uint256 id) external {
        Record storage rec = _records[id];
        if (rec.owner == address(0)) revert RecordNotFound(id);
        if (rec.owner != msg.sender) revert NotOwner(id);

        emit RecordDeleted(id, msg.sender);
        delete _records[id];
    }

    // ── Read functions ────────────────────────────────────────────────────────

    /**
     * @notice Decode and return the original data for a record.
     * @param id  Record ID returned by store().
     * @return    Original bytes.
     */
    function retrieve(uint256 id) external returns (bytes memory) {
        Record storage rec = _records[id];
        if (rec.owner == address(0)) revert RecordNotFound(id);

        bytes memory decoded = HealRS.decode(
            rec.encoded,
            rec.dataShards,
            rec.parityShards
        );

        emit Retrieved(id, msg.sender, keccak256(decoded) == rec.dataHash);

        return decoded;
    }

    /**
     * @notice Decode and verify integrity against stored hash.
     * @param id  Record ID.
     * @return    Original bytes (reverts if hash mismatch).
     */
    function retrieveVerified(uint256 id) external returns (bytes memory) {
        Record storage rec = _records[id];
        if (rec.owner == address(0)) revert RecordNotFound(id);

        bytes memory decoded = HealRS.decode(
            rec.encoded,
            rec.dataShards,
            rec.parityShards
        );

        bytes32 got = keccak256(decoded);
        if (got != rec.dataHash) {
            revert VerificationFailed(id, rec.dataHash, got);
        }

        emit Retrieved(id, msg.sender, true);

        return decoded;
    }

    /**
     * @notice Return record metadata without decoding.
     */
    function getMetadata(uint256 id) external view returns (
        bytes32 dataHash,
        uint256 originalSize,
        uint256 encodedSize,
        uint8   dataShards,
        uint8   parityShards,
        address owner,
        uint256 timestamp,
        string memory label
    ) {
        Record storage rec = _records[id];
        if (rec.owner == address(0)) revert RecordNotFound(id);

        return (
            rec.dataHash,
            rec.originalSize,
            rec.encoded.length,
            rec.dataShards,
            rec.parityShards,
            rec.owner,
            rec.timestamp,
            rec.label
        );
    }

    /**
     * @notice Return the raw encoded shard blob for a record.
     *         Useful for off-chain reconstruction or shard inspection.
     */
    function getEncoded(uint256 id) external view returns (bytes memory) {
        Record storage rec = _records[id];
        if (rec.owner == address(0)) revert RecordNotFound(id);
        return rec.encoded;
    }

    /**
     * @notice Total number of records ever created (includes deleted).
     */
    function totalRecords() external view returns (uint256) {
        return _nextId;
    }
}

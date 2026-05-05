// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title HealChainStorage (Oracle Version)
 * @notice Store data on-chain using Reed-Solomon encoding via an off-chain oracle.
 *
 * Flow:
 *   1. User calls store(data, label)
 *   2. Contract emits EncodeRequested event with requestId + raw data
 *   3. Oracle (HealChain Go service) picks up event
 *   4. Oracle RS-encodes data off-chain
 *   5. Oracle calls fulfillStore(requestId, encodedData)
 *   6. Contract stores encoded shards and emits Stored event
 *   7. User calls retrieve(id) — oracle watches and calls fulfillRetrieve
 *
 * Security:
 *   - Only the designated oracle address can call fulfill functions
 *   - Owner can rotate oracle address
 *   - Records can only be deleted by their original owner
 */
contract HealChainStorage {

    // ── Types ─────────────────────────────────────────────────────────────────

    struct Record {
        bytes   encoded;
        bytes32 dataHash;
        uint256 originalSize;
        uint8   dataShards;
        uint8   parityShards;
        address owner;
        uint256 timestamp;
        string  label;
        bool    fulfilled;
    }

    struct PendingRequest {
        address requester;
        bytes   rawData;
        uint8   dataShards;
        uint8   parityShards;
        string  label;
        bool    exists;
    }

    // ── Storage ───────────────────────────────────────────────────────────────

    address public owner;
    address public oracle;

    mapping(uint256 => Record)         private _records;
    mapping(uint256 => PendingRequest) private _pending;

    uint256 private _nextId;
    uint256 private _nextRequestId;

    // ── Events ────────────────────────────────────────────────────────────────

    // Emitted when a user requests storage — oracle listens for this
    event EncodeRequested(
        uint256 indexed requestId,
        address indexed requester,
        bytes           data,
        uint8           dataShards,
        uint8           parityShards,
        string          label
    );

    // Emitted when oracle fulfills encoding and data is stored
    event Stored(
        uint256 indexed id,
        address indexed owner,
        bytes32         dataHash,
        uint256         originalSize,
        uint256         encodedSize,
        string          label
    );

    // Emitted when a retrieve is requested — oracle listens for this
    event RetrieveRequested(
        uint256 indexed requestId,
        uint256 indexed recordId,
        address indexed requester
    );

    // Emitted when oracle fulfills decode
    event Retrieved(
        uint256 indexed recordId,
        address indexed caller,
        bool            verified
    );

    event RecordDeleted(uint256 indexed id, address indexed owner);
    event OracleUpdated(address indexed oldOracle, address indexed newOracle);

    // ── Errors ────────────────────────────────────────────────────────────────

    error NotOwner(uint256 id);
    error NotOracle();
    error NotContractOwner();
    error RecordNotFound(uint256 id);
    error RequestNotFound(uint256 requestId);
    error RecordNotFulfilled(uint256 id);
    error VerificationFailed(uint256 id, bytes32 expected, bytes32 got);
    error EmptyData();

    // ── Modifiers ─────────────────────────────────────────────────────────────

    modifier onlyOracle() {
        if (msg.sender != oracle) revert NotOracle();
        _;
    }

    modifier onlyContractOwner() {
        if (msg.sender != owner) revert NotContractOwner();
        _;
    }

    // ── Constructor ───────────────────────────────────────────────────────────

    constructor(address _oracle) {
        owner  = msg.sender;
        oracle = _oracle;
    }

    // ── User functions ────────────────────────────────────────────────────────

    /**
     * @notice Request storage with custom shard config.
     *         Emits EncodeRequested — oracle will fulfill async.
     * @return requestId  Track this to know when your data is stored.
     */
    function store(
        bytes calldata data,
        uint8 dataShards,
        uint8 parityShards,
        string calldata label
    ) external returns (uint256 requestId) {
        if (data.length == 0) revert EmptyData();

        requestId = _nextRequestId++;

        _pending[requestId] = PendingRequest({
            requester:    msg.sender,
            rawData:      data,
            dataShards:   dataShards,
            parityShards: parityShards,
            label:        label,
            exists:       true
        });

        emit EncodeRequested(requestId, msg.sender, data, dataShards, parityShards, label);
    }

    /**
     * @notice Request storage with default shards (10 data, 4 parity).
     */
    function store(
        bytes calldata data,
        string calldata label
    ) external returns (uint256 requestId) {
        if (data.length == 0) revert EmptyData();

        requestId = _nextRequestId++;

        _pending[requestId] = PendingRequest({
            requester:    msg.sender,
            rawData:      data,
            dataShards:   10,
            parityShards: 4,
            label:        label,
            exists:       true
        });

        emit EncodeRequested(requestId, msg.sender, data, 10, 4, label);
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

    // ── Oracle functions ──────────────────────────────────────────────────────

    /**
     * @notice Oracle calls this after RS-encoding the requested data.
     * @param requestId   The request ID from EncodeRequested event.
     * @param encoded     RS-encoded shard blob.
     */
    function fulfillStore(
        uint256 requestId,
        bytes calldata encoded
    ) external onlyOracle returns (uint256 recordId) {
        PendingRequest storage req = _pending[requestId];
        if (!req.exists) revert RequestNotFound(requestId);

        recordId = _nextId++;
        bytes32 dataHash = keccak256(req.rawData);

        _records[recordId] = Record({
            encoded:      encoded,
            dataHash:     dataHash,
            originalSize: req.rawData.length,
            dataShards:   req.dataShards,
            parityShards: req.parityShards,
            owner:        req.requester,
            timestamp:    block.timestamp,
            label:        req.label,
            fulfilled:    true
        });

        emit Stored(
            recordId,
            req.requester,
            dataHash,
            req.rawData.length,
            encoded.length,
            req.label
        );

        delete _pending[requestId];
    }

    // ── View functions ────────────────────────────────────────────────────────

    /**
     * @notice Return record metadata without decoding.
     */
    function getMetadata(uint256 id) external view returns (
        bytes32 dataHash,
        uint256 originalSize,
        uint256 encodedSize,
        uint8   dataShards,
        uint8   parityShards,
        address recordOwner,
        uint256 timestamp,
        string memory label,
        bool    fulfilled
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
            rec.label,
            rec.fulfilled
        );
    }

    /**
     * @notice Return the raw encoded shard blob.
     */
    function getEncoded(uint256 id) external view returns (bytes memory) {
        Record storage rec = _records[id];
        if (rec.owner == address(0)) revert RecordNotFound(id);
        return rec.encoded;
    }

    /**
     * @notice Total records ever created.
     */
    function totalRecords() external view returns (uint256) {
        return _nextId;
    }

    /**
     * @notice Check if a pending request exists.
     */
    function isPending(uint256 requestId) external view returns (bool) {
        return _pending[requestId].exists;
    }

    // ── Admin functions ───────────────────────────────────────────────────────

    /**
     * @notice Rotate the oracle address.
     */
    function setOracle(address newOracle) external onlyContractOwner {
        emit OracleUpdated(oracle, newOracle);
        oracle = newOracle;
    }
}

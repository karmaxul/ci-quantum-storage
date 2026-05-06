// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title HealChainStorage (Multi-Oracle Version)
 * @notice Store data on-chain using Reed-Solomon encoding via an off-chain oracle network.
 *
 * Flow:
 *   1. User calls store(data, label)
 *   2. Contract emits EncodeRequested event with requestId + raw data
 *   3. Any approved oracle picks up the event
 *   4. Oracle RS-encodes data off-chain
 *   5. First oracle to call fulfillStore(requestId, encodedData) wins
 *   6. Contract stores encoded shards and emits Stored event
 *
 * Security:
 *   - Only approved oracle addresses can call fulfill functions
 *   - Owner can add/remove oracles via addOracle/removeOracle
 *   - isPending check prevents duplicate fulfillment
 *   - Records can only be deleted by their original owner
 *
 * Multi-oracle design:
 *   - Multiple oracles watch the same contract
 *   - First to fulfill wins, others detect isPending=false and skip
 *   - No gas wasted on failed duplicate attempts
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

    // Oracle whitelist — any approved address can fulfill requests
    mapping(address => bool) public approvedOracles;
    address[] private _oracleList; // for enumeration

    mapping(uint256 => Record)         private _records;
    mapping(uint256 => PendingRequest) private _pending;

    uint256 private _nextId;
    uint256 private _nextRequestId;

    // ── Events ────────────────────────────────────────────────────────────────

    event EncodeRequested(
        uint256 indexed requestId,
        address indexed requester,
        bytes           data,
        uint8           dataShards,
        uint8           parityShards,
        string          label
    );

    event Stored(
        uint256 indexed id,
        address indexed owner,
        bytes32         dataHash,
        uint256         originalSize,
        uint256         encodedSize,
        string          label
    );

    event RecordDeleted(uint256 indexed id, address indexed owner);

    event OracleAdded(address indexed oracle);
    event OracleRemoved(address indexed oracle);

    // ── Errors ────────────────────────────────────────────────────────────────

    error NotOwner(uint256 id);
    error NotOracle();
    error NotContractOwner();
    error RecordNotFound(uint256 id);
    error RequestNotFound(uint256 requestId);
    error EmptyData();
    error OracleAlreadyApproved(address oracle);
    error OracleNotApproved(address oracle);

    // ── Modifiers ─────────────────────────────────────────────────────────────

    modifier onlyOracle() {
        if (!approvedOracles[msg.sender]) revert NotOracle();
        _;
    }

    modifier onlyContractOwner() {
        if (msg.sender != owner) revert NotContractOwner();
        _;
    }

    // ── Constructor ───────────────────────────────────────────────────────────

    /**
     * @param initialOracles  Array of oracle addresses to approve at deploy time.
     *                        Pass an empty array to add oracles later via addOracle().
     */
    constructor(address[] memory initialOracles) {
        owner = msg.sender;
        for (uint256 i = 0; i < initialOracles.length; i++) {
            _addOracle(initialOracles[i]);
        }
    }

    // ── User functions ────────────────────────────────────────────────────────

    /**
     * @notice Request storage with custom shard config.
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
     * @notice Any approved oracle calls this after RS-encoding the data.
     *         First oracle to call wins. Others will see isPending=false and skip.
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

    function getEncoded(uint256 id) external view returns (bytes memory) {
        Record storage rec = _records[id];
        if (rec.owner == address(0)) revert RecordNotFound(id);
        return rec.encoded;
    }

    function totalRecords() external view returns (uint256) {
        return _nextId;
    }

    function isPending(uint256 requestId) external view returns (bool) {
        return _pending[requestId].exists;
    }

    /**
     * @notice Returns all currently approved oracle addresses.
     */
    function getOracles() external view returns (address[] memory) {
        return _oracleList;
    }

    /**
     * @notice Check if an address is an approved oracle.
     */
    function isOracle(address addr) external view returns (bool) {
        return approvedOracles[addr];
    }

    // ── Admin functions ───────────────────────────────────────────────────────

    /**
     * @notice Add an oracle to the approved list.
     */
    function addOracle(address oracle) external onlyContractOwner {
        if (approvedOracles[oracle]) revert OracleAlreadyApproved(oracle);
        _addOracle(oracle);
    }

    /**
     * @notice Remove an oracle from the approved list.
     */
    function removeOracle(address oracle) external onlyContractOwner {
        if (!approvedOracles[oracle]) revert OracleNotApproved(oracle);
        approvedOracles[oracle] = false;

        // Remove from list
        for (uint256 i = 0; i < _oracleList.length; i++) {
            if (_oracleList[i] == oracle) {
                _oracleList[i] = _oracleList[_oracleList.length - 1];
                _oracleList.pop();
                break;
            }
        }

        emit OracleRemoved(oracle);
    }

    // ── Internal ──────────────────────────────────────────────────────────────

    function _addOracle(address oracle) internal {
        approvedOracles[oracle] = true;
        _oracleList.push(oracle);
        emit OracleAdded(oracle);
    }
}

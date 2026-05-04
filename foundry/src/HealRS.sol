// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title HealRS
 * @notice Library for interacting with HealChain Reed-Solomon precompiles.
 *
 * Precompile addresses:
 *   0x0400  encode(bytes data, uint8 dataShards, uint8 parityShards) → raw bytes
 *   0x0401  decode(bytes encoded, uint8 dataShards, uint8 parityShards) → raw bytes
 *   0x0402  stabilize(bytes encoded, uint8 dataShards, uint8 parityShards) → raw bytes
 *   0x0403  stats() → raw bytes
 *
 * Precompiles return raw bytes (no ABI wrapper on output).
 * Input is standard ABI-encoded: 4-byte selector + encoded args.
 */
library HealRS {

    // ── Precompile addresses ──────────────────────────────────────────────────

    address internal constant ENCODE_ADDR    = address(0x0400);
    address internal constant DECODE_ADDR    = address(0x0401);
    address internal constant STABILIZE_ADDR = address(0x0402);
    address internal constant STATS_ADDR     = address(0x0403);

    // ── Default shard configuration ───────────────────────────────────────────

    uint8 internal constant DEFAULT_DATA_SHARDS   = 10;
    uint8 internal constant DEFAULT_PARITY_SHARDS = 4;

    // ── Gas limits ────────────────────────────────────────────────────────────
    // Precompile calls go through an HTTP bridge — they need more gas than
    // typical precompiles. 5M covers large payloads comfortably.

    uint256 internal constant PRECOMPILE_GAS = 5_000_000;

    // ── Errors ────────────────────────────────────────────────────────────────

    error EncodeCallFailed();
    error DecodeCallFailed();
    error StabilizeCallFailed();
    error StatsCallFailed();
    error EmptyInput();

    // ── Encode ────────────────────────────────────────────────────────────────

    /**
     * @notice Encode data into Reed-Solomon shards.
     * @param data          Raw bytes to encode.
     * @param dataShards    Number of data shards.
     * @param parityShards  Number of parity shards.
     * @return              RS-encoded shard blob.
     */
    function encode(
        bytes memory data,
        uint8 dataShards,
        uint8 parityShards
    ) internal returns (bytes memory) {
        if (data.length == 0) revert EmptyInput();

        bytes memory callData = abi.encodeWithSignature(
            "encode(bytes,uint8,uint8)",
            data,
            dataShards,
            parityShards
        );

        (bool ok, bytes memory result) = ENCODE_ADDR.call{gas: PRECOMPILE_GAS}(callData);
        if (!ok || result.length == 0) revert EncodeCallFailed();

        return result;
    }

    /**
     * @notice Encode with default shard parameters (10 data, 4 parity).
     * @param data  Raw bytes to encode.
     * @return      RS-encoded shard blob.
     */
    function encode(bytes memory data) internal returns (bytes memory) {
        return encode(data, DEFAULT_DATA_SHARDS, DEFAULT_PARITY_SHARDS);
    }

    // ── Decode ────────────────────────────────────────────────────────────────

    /**
     * @notice Decode RS shards back to original data.
     * @param encoded       RS-encoded shard blob (output of encode).
     * @param dataShards    Number of data shards (must match encode call).
     * @param parityShards  Number of parity shards (must match encode call).
     * @return              Original data bytes.
     */
    function decode(
        bytes memory encoded,
        uint8 dataShards,
        uint8 parityShards
    ) internal returns (bytes memory) {
        if (encoded.length == 0) revert EmptyInput();

        bytes memory callData = abi.encodeWithSignature(
            "decode(bytes,uint8,uint8)",
            encoded,
            dataShards,
            parityShards
        );

        (bool ok, bytes memory result) = DECODE_ADDR.call{gas: PRECOMPILE_GAS}(callData);
        if (!ok || result.length == 0) revert DecodeCallFailed();

        return result;
    }

    /**
     * @notice Decode with default shard parameters.
     * @param encoded  RS-encoded shard blob.
     * @return         Original data bytes.
     */
    function decode(bytes memory encoded) internal returns (bytes memory) {
        return decode(encoded, DEFAULT_DATA_SHARDS, DEFAULT_PARITY_SHARDS);
    }

    // ── Stabilize ─────────────────────────────────────────────────────────────

    /**
     * @notice Attempt to repair a damaged shard set.
     * @param encoded       Possibly-damaged RS shard blob.
     * @param dataShards    Number of data shards.
     * @param parityShards  Number of parity shards.
     * @return              Repaired shard blob.
     */
    function stabilize(
        bytes memory encoded,
        uint8 dataShards,
        uint8 parityShards
    ) internal returns (bytes memory) {
        if (encoded.length == 0) revert EmptyInput();

        bytes memory callData = abi.encodeWithSignature(
            "stabilize(bytes,uint8,uint8)",
            encoded,
            dataShards,
            parityShards
        );

        (bool ok, bytes memory result) = STABILIZE_ADDR.call{gas: PRECOMPILE_GAS}(callData);
        if (!ok || result.length == 0) revert StabilizeCallFailed();

        return result;
    }

    /**
     * @notice Stabilize with default shard parameters.
     */
    function stabilize(bytes memory encoded) internal returns (bytes memory) {
        return stabilize(encoded, DEFAULT_DATA_SHARDS, DEFAULT_PARITY_SHARDS);
    }

    // ── Stats ─────────────────────────────────────────────────────────────────

    /**
     * @notice Get service stats from the RS precompile.
     * @return Raw stats bytes (JSON from the Go service).
     */
    function stats() internal returns (bytes memory) {
        bytes memory callData = abi.encodeWithSignature("stats()");

        (bool ok, bytes memory result) = STATS_ADDR.call{gas: PRECOMPILE_GAS}(callData);
        if (!ok) revert StatsCallFailed();

        return result;
    }

    // ── Availability check ────────────────────────────────────────────────────

    /**
     * @notice Check if the encode precompile is reachable.
     * @return true if the precompile responds successfully.
     */
    function isAvailable() internal returns (bool) {
        bytes memory callData = abi.encodeWithSignature(
            "encode(bytes,uint8,uint8)",
            bytes("ping"),
            DEFAULT_DATA_SHARDS,
            DEFAULT_PARITY_SHARDS
        );
        (bool ok, bytes memory result) = ENCODE_ADDR.call{gas: PRECOMPILE_GAS}(callData);
        return ok && result.length > 0;
    }
}

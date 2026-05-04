// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title HealRS
 * @notice Library for interacting with HealChain Reed-Solomon precompiles.
 *
 * Precompile addresses:
 *   0x0400  encode(bytes data, uint8 dataShards, uint8 parityShards) → bytes
 *   0x0401  decode(bytes encoded, uint8 dataShards, uint8 parityShards) → bytes
 *   0x0402  stabilize(bytes encoded, uint8 dataShards, uint8 parityShards) → bytes
 *   0x0403  stats() → bytes
 *
 * The precompiles accept standard ABI-encoded input (function selector + args)
 * and return raw bytes (no ABI wrapper on output).
 *
 * Usage:
 *   using HealRS for bytes;
 *   bytes memory shards = HealRS.encode(myData, 10, 4);
 *   bytes memory restored = HealRS.decode(shards, 10, 4);
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

    // ── Errors ────────────────────────────────────────────────────────────────

    error EncodeCallFailed();
    error DecodeCallFailed();
    error StabilizeCallFailed();
    error StatsCallFailed();
    error EmptyInput();
    error EmptyResult();

    // ── Core functions ────────────────────────────────────────────────────────

    /**
     * @notice Encode data into Reed-Solomon shards.
     * @param data          Raw bytes to encode.
     * @param dataShards    Number of data shards.
     * @param parityShards  Number of parity shards.
     * @return encoded      RS-encoded shard blob.
     */
    function encode(
        bytes memory data,
        uint8 dataShards,
        uint8 parityShards
    ) internal returns (bytes memory encoded) {
        if (data.length == 0) revert EmptyInput();

        bytes memory callData = abi.encodeWithSignature(
            "encode(bytes,uint8,uint8)",
            data,
            dataShards,
            parityShards
        );

        (bool ok, bytes memory result) = ENCODE_ADDR.call(callData);
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

    /**
     * @notice Decode RS shards back to original data.
     * @param encoded       RS-encoded shard blob (output of encode).
     * @param dataShards    Number of data shards (must match encode call).
     * @param parityShards  Number of parity shards (must match encode call).
     * @return decoded      Original data bytes.
     */
    function decode(
        bytes memory encoded,
        uint8 dataShards,
        uint8 parityShards
    ) internal returns (bytes memory decoded) {
        if (encoded.length == 0) revert EmptyInput();

        bytes memory callData = abi.encodeWithSignature(
            "decode(bytes,uint8,uint8)",
            encoded,
            dataShards,
            parityShards
        );

        (bool ok, bytes memory result) = DECODE_ADDR.call(callData);
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

    /**
     * @notice Attempt to repair a damaged shard set.
     * @param encoded       Possibly-damaged RS shard blob.
     * @param dataShards    Number of data shards.
     * @param parityShards  Number of parity shards.
     * @return stabilized   Repaired shard blob.
     */
    function stabilize(
        bytes memory encoded,
        uint8 dataShards,
        uint8 parityShards
    ) internal returns (bytes memory stabilized) {
        if (encoded.length == 0) revert EmptyInput();

        bytes memory callData = abi.encodeWithSignature(
            "stabilize(bytes,uint8,uint8)",
            encoded,
            dataShards,
            parityShards
        );

        (bool ok, bytes memory result) = STABILIZE_ADDR.call(callData);
        if (!ok || result.length == 0) revert StabilizeCallFailed();

        return result;
    }

    /**
     * @notice Get service stats from the RS precompile.
     * @return Raw stats bytes (JSON from the Go service).
     */
    function stats() internal returns (bytes memory) {
        bytes memory callData = abi.encodeWithSignature("stats()");

        (bool ok, bytes memory result) = STATS_ADDR.call(callData);
        if (!ok) revert StatsCallFailed();

        return result;
    }

    // ── View variants (staticcall) ────────────────────────────────────────────
    // Use these when calling from a view/pure context or when you don't need
    // to modify state. Note: precompile calls don't modify EVM state anyway,
    // but Solidity requires staticcall for view functions.

    function encodeView(
        bytes memory data,
        uint8 dataShards,
        uint8 parityShards
    ) internal view returns (bytes memory) {
        if (data.length == 0) revert EmptyInput();

        bytes memory callData = abi.encodeWithSignature(
            "encode(bytes,uint8,uint8)",
            data,
            dataShards,
            parityShards
        );

        (bool ok, bytes memory result) = ENCODE_ADDR.staticcall(callData);
        if (!ok || result.length == 0) revert EncodeCallFailed();

        return result;
    }

    function decodeView(
        bytes memory encoded,
        uint8 dataShards,
        uint8 parityShards
    ) internal view returns (bytes memory) {
        if (encoded.length == 0) revert EmptyInput();

        bytes memory callData = abi.encodeWithSignature(
            "decode(bytes,uint8,uint8)",
            encoded,
            dataShards,
            parityShards
        );

        (bool ok, bytes memory result) = DECODE_ADDR.staticcall(callData);
        if (!ok || result.length == 0) revert DecodeCallFailed();

        return result;
    }

    // ── Helpers ───────────────────────────────────────────────────────────────

    /**
     * @notice Check if the encode precompile is reachable.
     * @return true if the precompile responds.
     */
    function isAvailable() internal returns (bool) {
        bytes memory callData = abi.encodeWithSignature(
            "encode(bytes,uint8,uint8)",
            bytes("ping"),
            DEFAULT_DATA_SHARDS,
            DEFAULT_PARITY_SHARDS
        );
        (bool ok,) = ENCODE_ADDR.call(callData);
        return ok;
    }
}

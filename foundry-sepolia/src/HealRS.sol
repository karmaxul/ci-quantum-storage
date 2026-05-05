// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title HealRS
 * @notice Reed-Solomon interface library for HealChain.
 *
 * On devnet:  calls precompiles at 0x0400/0x0401 directly.
 * On Sepolia: uses oracle pattern — emits events, oracle fulfills off-chain.
 *
 * This file is the Sepolia/oracle version.
 * The devnet version (with direct precompile calls) lives in contracts/HealRS.sol.
 */
library HealRS {

    // ── Default shard configuration ───────────────────────────────────────────

    uint8 internal constant DEFAULT_DATA_SHARDS   = 10;
    uint8 internal constant DEFAULT_PARITY_SHARDS = 4;

    // ── Errors ────────────────────────────────────────────────────────────────

    error EmptyInput();
}

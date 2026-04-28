<<<<<<< Updated upstream
# Ci-RS — On-Chain Self-Healing Data (CiSHA4096)

**Ci On-Chain Reed-Solomon** — A Solidity implementation of ultra-light self-healing data using the custom **CiSHA4096** hash.

Deployed & Verified on **Sepolia** (April 2026):

- **CiOnChainRS_Full**: [`0xcBa6D4606A311f0a99fE62A24b94a3DBAC39f189`](https://sepolia.etherscan.io/address/0xcBa6D4606A311f0a99fE62A24b94a3DBAC39f189)
- **CiSHA4096_UltraLight**: [`0xcB6Fe282444FBBff89bE49298f91Fea85ccf48a3`](https://sepolia.etherscan.io/address/0xcB6Fe282444FBBff89bE49298f91Fea85ccf48a3)

---

## Overview

This project demonstrates **on-chain self-healing** for small payloads using a custom 4096-bit cryptographic hash (`CiSHA4096`) combined with lightweight Reed-Solomon-inspired repair logic.

The goal is to push the limits of what is practically possible for **self-repairing data** directly on Ethereum (and future HealChain layers).

## Current Capabilities (Phase 2 Plateau)

| Payload Size | Recovery Success | Recovery Quality      | Encode Gas (approx) | Repair Gas (approx) | Status          |
|--------------|------------------|-----------------------|---------------------|---------------------|-----------------|
| ≤ 16 bytes   | Reliable         | Clean / Exact         | ~650k               | ~4.9k               | **Production-ready** |
| 19–20 bytes  | Partial          | Garbled but partial match | ~881k            | ~478k               | Limit reached   |
| 22+ bytes    | Unreliable       | Frequent failure      | High                | High                | Not viable      |

**Key Insight**: CiSHA4096's extreme sensitivity makes 16 bytes the sweet spot for reliable on-chain self-healing in pure Solidity.

## Repository Structure

ci-solidity/
├── src/
│   ├── CiOnChainRS_Full.sol          # Main self-healing contract
│   └── CiSHA4096_UltraLight.sol      # Core 4096-bit hash
├── script/
│   ├── DeployCiRS.s.sol
│   └── InteractCiRS.s.sol            # Test scripts
└── README.md

## Quick Start

### 1. Deploy (already done)
=======
## Live Demo (Web App)
**Current URL:** `https://ci-quantum-storage-5gk9.onrender.com` (or your latest Render URL)

## How to Run Locally
>>>>>>> Stashed changes

```bash
forge script script/DeployCiRS.s.sol --rpc-url https://ethereum-sepolia-rpc.publicnode.com --broadcast --verify

2. Test Self-Healingbash

<<<<<<< Updated upstream
forge script script/InteractCiRS.s.sol \
  --rpc-url https://ethereum-sepolia-rpc.publicnode.com \
  --broadcast -vvvv

Technical DetailsHash: CiSHA4096 (custom ultra-light 4096-bit hash)
Repair Mechanism: Reed-Solomon style with on-chain brute-force search in constrained space
EVM Compatibility: Works on Sepolia (and Ethereum mainnet with higher gas)
Verification: Fully verified on Sourcify

Next Phase: HealChain Architecture
We have now reached the practical limit of pure Solidity on existing EVMs.
The next stage will explore:Custom precompiles for CiSHA4096
Higher payload capacity (64–256+ bytes)
Native HealChain layer-1 / layer-2 design
Integration prospectives (JASMY, etc.)
=======
Also update the **Live Contract** section if needed.

---

### 3. Full Updated Sharing Text (Ready to Use)

**For GitHub / X / LinkedIn:**

```text
Ci Quantum-Inspired Storage — Exploring rational cryptography with Ci = 85/27.

Features:
• Repeating "double-helix" patterns in hashes
• Reed-Solomon protected web storage
• Live Solidity contract on Sepolia (16 states, ~2M gas)

Live Web Demo: https://ci-quantum-storage-5gk9.onrender.com
GitHub: https://github.com/karmaxul/ci-quantum-storage
Live Contract: 0x6Db61C27F196704519c7Eb6a6FaB1E017B7e0514

Curious about what rational constants could enable in primitives. Feedback welcome. 🌱

## Philosophy

This project explores the 9-bit → 12-bit precision upgrade using rational constants for better balance and stability in computational primitives.

### Motivation & History of Rational Constants

Traditional cryptographic hashes (SHA-256, SHA-512, etc.) rely almost exclusively on **irrational constants** (square roots of small primes) to maximize chaos and avalanche. The goal has historically been to eliminate detectable patterns.

This project deliberately explores the opposite direction: using **rational constants** (specifically Ci = 85/27, derived from physical scaling principles) to create structured, repeating "double-helix" patterns. While this might seem counter-intuitive for security, it opens the door to better error correction, stabilizer-like behavior, and philosophical alignment with natural balance.

Rational constants have appeared sparingly in cryptography history (early mechanical ciphers, some stream cipher designs, and certain coding theory constructions), but they remain underexplored in modern hash functions. This work continues the spirit of the original 9-bit → 12-bit fraction concept, asking whether structured redundancy can be a feature rather than a bug when combined with Reed-Solomon and other tools.

The result is a hash that maintains respectable avalanche (~47–53%) while producing clear, analyzable repeating patterns — something that may prove useful for verification, storage, or quantum-inspired applications.


## Features
- **Ci-SHA4096 v2.4**: 4096-bit output with 100/100 Ci Signature Score
- Reed-Solomon error correction for protected, self-healing storage
- Persistent storage with Gzip compression
- Flask web interface (Store, Retrieve, Test Erasure, Stabilizer Demo, Download)
- **Solidity on-chain PoC** (16 states, Ci influence, deployed on Sepolia)

### How Ci-SHA4096 + Reed-Solomon Work Together

- **Ci-SHA4096** creates a strong, structured 4096-bit fingerprint (hash) of your data.
- **Reed-Solomon** adds powerful error-correction parity symbols, allowing recovery even if parts of the stored block are damaged or lost.
- When retrieving: Reed-Solomon first attempts to repair the data, then the system re-computes the Ci-SHA4096 hash and verifies it matches the stored hash.
- Result: Self-verifying, self-healing storage with structured redundancy from the rational constant Ci = 85/27.

This combination is one of the core innovations of the project.

**Created in collaboration with Grok (xAI) — April 2026**

## Features
- **Ci-SHA4096 v2.4**: 4096-bit output with 100/100 Ci Signature Score
- Reed-Solomon error correction for protected, self-healing storage
- Persistent storage with Gzip compression
- Flask web interface (Store, Retrieve, Test Erasure, Stabilizer Demo, Download)
- **Solidity on-chain PoC** (16 states, Ci influence, deployed on Sepolia)

### Security & Gas Summary
- **Gas Benchmark**: ~2.06 million gas (optimized version with 8 active states + simplified diffusion)
- **Security Notes**: Fully deterministic view function with safe arithmetic (full uint256 masking). Rational Ci = 85/27 adds structured redundancy (beneficial for error correction when paired with Reed-Solomon). Includes on-chain `verify()` helper for cross-checking against Python reference. Not intended for cryptographic signing or key derivation — best used as a verifiable hash + storage primitive.
- Reed-Solomon typically adds ~39–43% parity overhead but enables recovery from multiple errors per block.

## Live Contract (Sepolia)
**Address:** `0x6Db61C27F196704519c7Eb6a6FaB1E017B7e0514`  
[View on Sepolia Etherscan](https://sepolia.etherscan.io/address/0x6Db61C27F196704519c7Eb6a6FaB1E017B7e0514)

>>>>>>> Stashed changes


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

```bash
forge script script/DeployCiRS.s.sol --rpc-url https://ethereum-sepolia-rpc.publicnode.com --broadcast --verify

2. Test Self-Healingbash

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


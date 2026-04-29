# Ci Quantum-Inspired Storage

**Powered by Ci-SHA4096 (rational Ci = 85/27) + Self-Healing Reed-Solomon**

A demonstration project exploring structured redundancy in cryptography. By replacing irrational constants with rational ones (Ci = 85/27), we create repeating "double-helix" patterns that improve compatibility with error correction while maintaining strong avalanche properties.

## Current Status (April 2026)

- **Web Demo**: Fully functional Flask application with Store (RS-protected), Retrieve, Test Erasure, Stabilizer Demo, and Download Block.  
  This is the most polished and user-facing component.

- **Python Reference**: Ci-SHA4096 v2.4 — 4096-bit output, excellent stabilizer patterns, 100/100 Ci Signature Score.

- **Solidity PoC (Sepolia Testnet)**: Reliable self-healing for **≤16-byte payloads**  
  - Contract: `0x37A93F327aE1aB5cb3D8Eb253BB5DeC9297caCa7` (CiOnChainRS_Full)  
  - Encode gas: ~650,000  
  - Repair gas: ~4,900 (very efficient)  
  - Recovery: Clean and successful on live testnet

- **Go Library** (`healchain/`): Clean, standalone Reed-Solomon engine with per-shard SHA256 integrity checks and automatic self-healing. Works reliably for larger payloads in off-chain / service mode.

### Current Plateau Summary

| Payload Size | Status          | Encode Gas | Repair Gas | Recovery Quality | Notes                     |
|--------------|-----------------|------------|------------|------------------|---------------------------|
| ≤ 16 bytes   | Stable          | ~650k      | ~5k        | Clean            | Reliable on Sepolia       |
| 24+ bytes    | Overflow (0x11) | N/A        | N/A        | Not reachable    | CiSHA4096 sensitivity     |
| 64–512+ bytes| Target          | <200k goal | <50k goal  | High             | Requires precompiles      |

## Features

- Rational Ci = 85/27 constants creating natural structured redundancy
- Reed-Solomon-inspired self-healing with per-shard integrity verification
- Flask web interface for easy testing and demonstration
- On-chain proof-of-concept on Sepolia
- Modular Go implementation with Cosmos SDK / EVM precompile support

## Quick Start – Web App

```bash
cd ~/ci-sha-project
source ci_venv/bin/activate
pip install -r requirements.txt
python app.py

Open → http://127.0.0.1:5000Quick Start – Go Librarybash

go run demo.go

go

import "ci-sha-test/healchain"

rs, err := healchain.New(10, 4)        // dataShards, parityShards
encoded, _ := rs.Encode(data)
recovered, _ := rs.Decode(encoded)

Tuning Guide (Reed-Solomon Configuration)Payload Size
dataShards
parityShards
Approx. Overhead
Recommended Use Case
32–64 bytes
8
4
~50%
Small IoT/sensor packets
64–128 bytes
10
4
~40%
Balanced default
128–256 bytes
12
5
~42%
General purpose
256–512 bytes
16
6
~38%
Archival / AI data chunks
512+ bytes
20+
8+
~40%
High redundancy scenarios

Tip: Higher parity improves correction capability but increases storage overhead. Test with your expected corruption patterns.Architecture OverviewEncoding Flow:Split data into data shards
Generate parity shards (Reed-Solomon)
Compute SHA256 hash per shard for integrity
Add header + shards + hashes → self-healing blob

Decoding / Self-Healing:Detect corrupted shards via hash mismatch
Mark bad shards
Reconstruct using Reed-Solomon
Return original data

PhilosophyThis project continues the exploration of rational constants (Ci = 85/27, derived from physical scaling principles) to create computational primitives with natural structured redundancy — offering better synergy with error correction than purely chaotic traditional hashes.RoadmapSee ROADMAP.md for current priorities (Phase 1: Precompiles & Scaling to larger payloads).Security NoteThis library provides error correction, not cryptographic security. Always combine with proper encryption and signatures when confidentiality or authenticity is required.


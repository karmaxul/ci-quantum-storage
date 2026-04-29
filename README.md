# HealChain Self-Healing System

A lightweight Reed-Solomon based self-healing data encoding library designed for blockchain environments (Harmony, Cosmos SDK, EVM).

## Overview

HealChain encodes data with redundancy so that corruption can be detected and repaired automatically.

### Key Features

- Reed-Solomon encoding with per-shard SHA256 integrity checks
- Cosmos SDK / EVM Precompile support
- Solidity helper interface
- Lightweight on-chain verification contract
- Clean, modular Go implementation

## Project Structure

ci-sha-test/
├── healchain/                  # Core library
│   ├── healchain.go
│   └── precompile.go
├── app.go                      # Bootstrap & testing
├── demo.go                     # Quick demo
├── HealChainInterface.sol      # Solidity interface
├── HealChainRS.sol             # On-chain verifier
├── test/HealChain.t.sol        # Foundry tests
└── README.md

## Quick Start

```bash
go run demo.go
go run app.go

Usage ExamplesGogo

import "ci-sha-test/healchain"

rs, err := healchain.New(10, 4)
encoded, _ := rs.Encode(data)
recovered, _ := rs.Decode(encodedData)

Precompile (Cosmos SDK)See healchain/precompile.go for registration details.Soliditysolidity

HealChainInterface hc = new HealChainInterface();
bytes memory encoded = hc.encode(originalData);
bytes memory recovered = hc.decode(encoded);

Tuning GuidePayload Size
dataShards
parityShards
Approximate Overhead
Notes
32–64 bytes
8
4
~50%
Small payloads
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
Larger data
512+ bytes
20+
8
~40%
High redundancy

Choose values based on your acceptable overhead and expected corruption rate.Architecture Diagram (Text)

Original Data
      ↓
[ Split into Data Shards ]
      ↓
[ Generate Parity Shards ]
      ↓
[ Compute SHA256 per shard ]
      ↓
[ Header + Shards + Hashes ]
      ↓
   Encoded Blob (self-healing)

On Receive:
   Decode → Check Hashes → Mark Bad Shards → Reconstruct → Output

Gas Cost Estimates (Approximate)These are rough observations from Remix / local testing:encode(): ~25k–60k gas (depending on size)
decode() (healing): ~40k–90k gas
verify(): ~15k–35k gas

Note: Actual gas usage depends heavily on payload size, shard configuration, and chain implementation.Security ConsiderationsThis library provides error correction, not cryptographic security.
Always combine with proper encryption/signatures when confidentiality or authenticity is required.
The precompile runs in the EVM execution environment — review carefully before mainnet deployment.
Header parsing and length fields should be treated as untrusted input.
Test thoroughly with your specific corruption patterns.

Roadmap (Current Ideas)Support for optional compression before encoding
Dynamic/adaptive shard sizing
Integration examples with common Cosmos modules
More comprehensive test coverage
Community feedback and improvements

(No specific timelines or guarantees are provided.)Contribution GuidelinesContributions are welcome. Please:Open an issue first for major changes
Keep code style consistent
Add tests for new functionality
Update documentation as needed

Feel free to submit Pull Requests.StatusCore functionality is implemented and tested in local environments.


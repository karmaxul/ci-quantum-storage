# HealChain / Ci-RS Roadmap

**Self-Healing Data Infrastructure**

**Last Updated**: April 28, 2026

---

## Phase 0 – Foundation (Completed ✅)

- [x] CiSHA4096_UltraLight implementation
- [x] CiOnChainRS_Full with on-chain repair
- [x] Sepolia deployment + Sourcify verification
- [x] Phase 2 plateau documentation (16-byte reliable sweet spot)
- [x] HealChain Architecture Draft v0.1.1

---

## Phase 1 – Precompiles & Scaling (Q2–Q3 2026)

- [ ] Finalize CiSHA4096 & CiRSRepair precompile specifications
- [ ] Implement precompiles (Go/Rust) in local geth fork
- [ ] Launch HealChain Devnet (single sequencer)
- [ ] Achieve reliable **256-byte** self-healing payloads
- [ ] Gas benchmarking & optimization
- [ ] Port existing contracts to use precompiles

**Target**: 50–200x speedup + much higher payload capacity

---

## Phase 2 – HealChain L1 Devnet → Testnet (Q3–Q4 2026)

- [ ] Full custom EVM chain launch (op-geth or Reth based)
- [ ] Native HealBlock + Self-Healing Transaction types
- [ ] Economic incentives (repair participation rewards)
- [ ] Basic explorer + RPC endpoints
- [ ] Security audit (pre-audit + formal audit)

**Target**: Public testnet with 1 KB+ reliable payloads

---

## Phase 3 – Ecosystem & Mainnet (2027+)

- [ ] Mainnet launch
- [ ] Optimistic / ZK Heal-Rollups
- [ ] Integration exploration with complementary projects (IoT, storage layers, etc.)
- [ ] Developer SDKs (JS, Python, full Solidity library)
- [ ] Production use cases (decentralized identity, sensor data, documents)

---

## Open Items / Nice-to-Haves

- Dynamic redundancy levels
- Multi-block repair for large payloads
- Tokenomics design
- Governance model
- JASMY ecosystem compatibility exploration (non-binding)

---

**Status**: Phase 0 Complete → Entering Phase 1

We move fast but carefully. Contributions and feedback welcome.

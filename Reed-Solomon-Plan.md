# Reed-Solomon Implementation Plan for HealChain Precompile

**Date**: April 28, 2026  
**Goal**: Achieve reliable self-healing for **64–256+ byte** payloads with 1–4 byte corruption tolerance.

---

## 1. Current Limitations

- Simple majority + parity + syndrome methods work well up to ~32 bytes.
- At 48+ bytes, 1-byte corruption often causes noticeable degradation.
- We need mathematically correct error correction.

---

## 2. Target: Lightweight Reed-Solomon

We will implement a **simplified Reed-Solomon** variant suitable for EVM precompiles:

- **Field**: GF(256) (using primitive polynomial `0x11D`)
- **Code Parameters**: RS(n=chunkSize+parity, k=chunkSize) with 4–8 parity symbols
- **Correction Capability**: Up to 2–4 symbol errors per chunk
- **Approach**: Systematic encoding + syndrome calculation + Berlekamp-Massey / Chien search for error location/correction (simplified)

---

## 3. Implementation Phases

### Phase 1 (Immediate – Lightweight)
- Pre-computed GF(256) log/exp tables
- Simple polynomial evaluation for parity
- Syndrome-based error detection
- Basic error correction for 1–2 symbols

### Phase 2 (Next)
- Full Berlekamp-Massey algorithm for error locator polynomial
- Chien search for error positions
- Forney algorithm for error value correction

### Phase 3 (Future)
- Dynamic redundancy levels
- Multi-chunk / multi-block repair
- Integration with data sharding

---

## 4. Key Components to Add in Precompile

1. **GF(256) Tables** (precomputed at init)
2. **Polynomial Math** (multiply, divide, evaluate)
3. **Encoding**: Generate parity symbols
4. **Decoding**: Syndrome → Error Locator → Correction

---

## 5. Next Immediate Actions

1. Add GF(256) tables and basic polynomial functions to the precompile
2. Implement systematic encoding with 4–6 parity symbols
3. Add syndrome calculation + simple single-error correction
4. Test with 64-byte payloads
5. Benchmark gas vs current version

---

## 6. Expected Outcomes

- Reliable recovery for **64-byte** payloads with 1–2 byte corruption
- Usable recovery for **128-byte** payloads
- Foundation for production-grade self-healing storage

---

**Status**: Moving from heuristic repair → mathematically correct error correction.

This is the real technical leap for HealChain.

---

*Approved by project lead — April 28, 2026*

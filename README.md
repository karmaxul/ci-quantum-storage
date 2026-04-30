# 🌌 Ci Quantum-Inspired Storage / HealChain

**Hybrid Self-Healing Storage System**  
Flask Frontend + Go Reed-Solomon Backend

## Current Status (April 30, 2026)

**Phase 1 — Hybrid Architecture: Stable & Feature Complete**

### Key Features
- Dynamic Reed-Solomon shards (configurable data + parity)
- Gzip compression support before encoding
- Live Storage Summary dashboard
- Large payload testing (1KB / 5KB / 10KB)
- Self-healing with real healing time measurement
- Full block management (Retrieve, Test Erasure, Stabilizer Demo, Download, Delete)
- Anchor navigation (returns to same position)
- Comprehensive logging (`healchain.log`)
- Clean, responsive UI with mobile support
- Green/Red service status indicator

### Performance Highlights
- ~40–41% overhead at 50KB–120KB payloads
- Sub-2ms healing times
- Excellent scaling behavior

## Docker Setup (Recommended)

```bash
# Build and run
docker compose up --build

# Run in background
docker compose up -d

# View logs
docker compose logs -f

Open → http://127.0.0.1:5000

To stop:bash

docker compose down




Project Structurehealchain-service.go → Go self-healing HTTP service
app.py → Flask web interface + hybrid logic
healchain/ → Core Go library
healchain.log → Runtime logs
web_blocks.json → Persistent storage



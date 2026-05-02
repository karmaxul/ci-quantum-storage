# 🌌 HealChain — Quantum-Inspired Self-Healing Storage

**Hybrid Self-Healing Storage System**  
Flask Frontend + Go Reed-Solomon Backend

---

## Current Status (April 30, 2026)

Current StatusPrecompiles: Encode, Decode, Stabilize, Stats working
Service: HTTP bridge with RS
Devnet: 2 nodes connected
UI: Basic tester page
Tests: All passing

### Key Features
- Dynamic Reed-Solomon erasure coding (configurable data + parity shards)
- Gzip compression before encoding
- Live Global Storage Summary dashboard
- Large payload testing (1KB / 5KB / 10KB+)
- Realistic self-healing with timing measurements
- Full block lifecycle: Store, Retrieve, Test Erasure, Stabilizer Demo, Download, Delete
- Search, copy ID, floating navigation buttons
- Persistent storage (`web_blocks.json`)
- Real-time service health indicators (Green/Red)
- Comprehensive logging

### Performance Highlights
- ~40–41% storage overhead at 50–120KB payloads
- Sub-2ms healing simulation times
- Excellent scaling with shard count

---

## 🚀 Quick Start — Docker (Recommended)

### One-command deployment

```bash
# 1. Clone the repo (if not already done)
git clone https://github.com/karmaxul/ci-quantum-storage.git
cd ci-quantum-storage

# 2. Start the full system
docker compose up --build -d

Open → http://127.0.0.1:5000Docker Commandsbash

# Start in background
docker compose up --build -d

# View live logs
docker compose logs -f

# Restart
docker compose restart

# Stop
docker compose down

# Full reset (deletes volumes)
docker compose down -v

Ports:5000 → Flask Web UI
8080 → Go Self-Healing Service

Manual Setup (Alternative)PrerequisitesGo 1.24+
Python 3.12+
pip install flask requests gunicorn

Run Servicesbash

# Terminal 1 - Go Backend
go run healchain-service.go

# Terminal 2 - Flask UI
python app.py

Project Structure

.
├── Dockerfile                  # Multi-stage build
├── docker-compose.yml
├── start-docker.sh
├── requirements.txt
├── app.py                      # Flask UI + frontend logic
├── healchain-service.go        # Go HTTP service
├── healchain/                  # Core Go library (Reed-Solomon, stabilizer, etc.)
├── ci_sha4096_v2_4.py
├── ci_rs_wrapper.py
├── web_blocks.json             # Persistent block storage
├── healchain.log               # Runtime logs
├── ROADMAP.md
├── HealChain-Architecture-Draft.md
└── Precompile-Spec.md

Architecture OverviewGo Backend: High-performance Reed-Solomon encoding/decoding + self-healing engine
Flask UI: Modern, responsive interface with real-time stats
Hybrid Mode: UI calls Go service via HTTP for heavy operations
Persistence: web_blocks.json + volume mounts in Docker
Future: EVM Precompiles + HealChain Devnet (Phase 2)

Next Steps (Phase 2)Integrate Reed-Solomon as EVM precompiles in custom Geth fork
Deploy HealChain Devnet
On-chain stabilizer and healing proofs
Gas optimization variants

See ROADMAP.md and HealChain-Architecture-Draft.md for details.ContributingContributions welcome! Feel free to open issues or PRs.Made with  for resilient decentralized storageLast updated: April 30, 2026


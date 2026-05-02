#!/bin/bash
echo "🚀 Starting HealChain Hybrid System..."

# Start Go service in background
echo "Starting Go Self-Healing Service on port 8080..."
go run healchain-service.go > go-service.log 2>&1 &
GO_PID=$!

sleep 2

# Start Flask
echo "Starting Flask Web UI on port 5000..."
PYTHONPATH=. ci_venv/bin/python app.py

# Cleanup on exit
kill $GO_PID 2>/dev/null

#!/bin/bash
set -e

echo "🚀 Starting HealChain Hybrid System..."

# Start Go backend
echo "Starting Go self-healing service on :8080..."
healchain-service > healchain.log 2>&1 &

# Wait a few seconds
sleep 4

echo "✅ Go service started (background)"
echo "Starting Flask UI on :5000..."

# Start Flask with Gunicorn
exec gunicorn --bind 0.0.0.0:5000 \
              --workers 2 \
              --timeout 180 \
              --access-logfile - \
              --error-logfile - \
              app:app

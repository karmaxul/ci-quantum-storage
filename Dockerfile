# ==================== Go Builder Stage ====================
FROM golang:1.24-alpine AS go-builder

WORKDIR /app

# Copy Go modules
COPY go.mod go.sum ./
RUN go mod download && go mod tidy

# Force download missing dependencies
RUN go get github.com/ethereum/go-ethereum/common \
           github.com/ethereum/go-ethereum/core/vm \
           github.com/klauspost/reedsolomon

# Copy source code
COPY healchain/ ./healchain/
COPY healchain-service.go .

# Build the Go service
RUN CGO_ENABLED=0 GOOS=linux go build -o /healchain-service healchain-service.go

# ==================== Final Python Stage ====================
FROM python:3.12-slim

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Copy Go binary from builder
COPY --from=go-builder /healchain-service /usr/local/bin/healchain-service

# Python dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy Flask app and helpers
COPY app.py .
COPY ci_sha4096_v2_4.py ci_rs_wrapper.py ./

# Create data directory
RUN mkdir -p /data && chmod 777 /data

EXPOSE 5000 8080

# Copy entrypoint
COPY start-docker.sh .
RUN chmod +x start-docker.sh

ENTRYPOINT ["./start-docker.sh"]

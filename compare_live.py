#!/usr/bin/env python3
"""
Ci-SHA4096 - Python vs Live Solidity (Sepolia) Comparison
"""
from ci_sha4096_v2_4 import ci_sha4096_v2_4
import subprocess

CONTRACT = "0x6Db61C27F196704519c7Eb6a6FaB1E017B7e0514"
RPC = "https://ethereum-sepolia-rpc.publicnode.com"

def call_solidity(message: bytes):
    """Call the live contract on Sepolia"""
    try:
        result = subprocess.run([
            "cast", "call", CONTRACT,
            "gasBenchmark(bytes)(uint256,bytes32[16])",
            "0x" + message.hex(),
            "--rpc-url", RPC
        ], capture_output=True, text=True, timeout=30)
        
        if result.returncode == 0:
            print("✅ Solidity call successful")
            return result.stdout.strip()[:300] + "..."  # Truncated for display
        else:
            print("❌ Cast error:", result.stderr.strip()[:200])
            return None
    except Exception as e:
        print(f"Error calling contract: {e}")
        return None

def main():
    messages = [
        b"",
        b"test",
        b"test message for Ci-SHA4096",
        b"Hello from Ci Quantum Storage!",
        b"Plant the seeds..."
    ]

    print("=== Ci-SHA4096 Python vs Live Solidity (Sepolia) ===\n")
    
    for msg in messages:
        text = msg.decode('utf-8', errors='ignore') or "<empty>"
        print(f"Message: {text}")
        
        # Python version
        py_hash = ci_sha4096_v2_4(msg, rounds=128)
        print("Python first 64 hex:", py_hash.hex()[:64])
        
        # Live Solidity version
        sol_result = call_solidity(msg)
        if sol_result:
            print("Solidity response received")
        
        print("-" * 90)

if __name__ == "__main__":
    main()

#!/usr/bin/env python3
"""
Ci-SHA4096 - Python vs Solidity Comparison Helper
"""
from ci_sha4096_v2_4 import ci_sha4096_v2_4

def main():
    messages = [
        b"",
        b"test",
        b"test message for Ci-SHA4096",
        b"Hello from Ci Quantum Storage!"
    ]

    print("=== Ci-SHA4096 Python Reference ===")
    for msg in messages:
        print(f"\nMessage: {msg}")
        py_hash = ci_sha4096_v2_4(msg, rounds=128)
        print("Python first 64 hex:", py_hash.hex()[:64])

    print("\n=== Next Steps ===")
    print("1. Run `anvil` in another terminal for local testing")
    print("2. Use `cast call` to compare with Solidity")
    print("3. Full comparison script with anvil coming soon")

if __name__ == "__main__":
    main()

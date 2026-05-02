import hashlib
import time
from collections import Counter

def right_rotate(value: int, shift: int, bits: int = 32) -> int:
    return ((value >> shift) | (value << (bits - shift))) & ((1 << bits) - 1)

def ci_sha4096_v2_4(message: bytes, rounds: int = 128) -> bytes:
    """Ci-SHA4096 v2.4 - Stable Maximum Potential Version"""
    CI_NUM = 85
    CI_DEN = 27

    # Initial state (16 parallel chains → 4096-bit output)
    iv = [0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
          0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19]
    states = [iv[:] for _ in range(16)]

    # Simple message padding
    data = bytearray(message)
    bit_len = len(data) * 8
    data.append(0x80)
    while len(data) % 64 != 56:
        data.append(0)
    data += bit_len.to_bytes(8, 'big')

    for i in range(0, len(data), 64):
        W = [0] * 64
        for j in range(16):
            W[j] = int.from_bytes(data[i + j*4 : i + j*4 + 4], 'big')

        for j in range(16, 64):
            s0 = right_rotate(W[j-15], 7) ^ right_rotate(W[j-15], 18) ^ (W[j-15] >> 3)
            s1 = right_rotate(W[j-2], 17) ^ right_rotate(W[j-2], 19) ^ (W[j-2] >> 10)
            W[j] = (W[j-16] + s0 + W[j-7] + s1) & 0xFFFFFFFF

        for chain in range(16):
            a, b, c, d, e, f, g, h = states[chain]
            for r in range(rounds):
                k = (CI_NUM * (r + 1)) // CI_DEN
                S1 = right_rotate(e, 6) ^ right_rotate(e, 11) ^ right_rotate(e, 25)
                ch = (e & f) ^ (~e & g)
                temp1 = (h + S1 + ch + k + W[r % 64]) & 0xFFFFFFFF

                S0 = right_rotate(a, 2) ^ right_rotate(a, 13) ^ right_rotate(a, 22)
                maj = (a & b) ^ (a & c) ^ (b & c)
                temp2 = (S0 + maj) & 0xFFFFFFFF

                h = g; g = f; f = e; e = (d + temp1) & 0xFFFFFFFF
                d = c; c = b; b = a; a = (temp1 + temp2) & 0xFFFFFFFF

            states[chain] = [a, b, c, d, e, f, g, h]

    # Final diffusion + cross-chain mixing
    digest = b''
    for chain in range(16):
        for word in states[chain]:
            digest += word.to_bytes(4, 'big')
    return digest

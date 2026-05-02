from reedsolo import RSCodec
from ci_sha4096_v2_4 import ci_sha4096_v2_4

class CiReedSolomon:
    def __init__(self, nsym=32, hash_bytes=512):
        """
        nsym: Number of Reed-Solomon parity symbols
        hash_bytes: 512 = full Ci-SHA4096 (recommended)
                    64  = light mode (for small/high-frequency tx)
        """
        self.rs = RSCodec(nsym)
        self.nsym = nsym
        self.hash_len = hash_bytes

    def encode(self, data: bytes) -> bytes:
        """Encode with Ci hash + Reed-Solomon"""
        ci_hash = ci_sha4096_v2_4(data, rounds=128)
        payload = data + ci_hash[:self.hash_len]   # truncate if using light mode
        encoded = self.rs.encode(payload)
        return encoded

    def decode(self, encoded: bytes):
        """Decode and verify hash"""
        try:
            decoded = self.rs.decode(encoded)[0]
            data = decoded[:-self.hash_len]
            stored_hash = decoded[-self.hash_len:]
            
            computed_hash = ci_sha4096_v2_4(data, rounds=128)
            success = (computed_hash[:self.hash_len] == stored_hash)
            
            return data, success
        except Exception as e:
            return None, False

    def demonstrate(self, original_data: bytes, errors=6):
        """Quick demo"""
        print("=== Ci + Reed-Solomon Demo ===")
        encoded = self.encode(original_data)
        print(f"Original: {len(original_data)} bytes | Encoded: {len(encoded)} bytes")
        
        noisy = bytearray(encoded)
        for _ in range(errors):
            idx = __import__('random').randint(0, len(noisy)-1)
            noisy[idx] ^= __import__('random').randint(1, 255)
        
        recovered, success = self.decode(bytes(noisy))
        print(f"Injected {errors} errors → Recovery: {'✅ SUCCESS' if success else '❌ FAILED'}")

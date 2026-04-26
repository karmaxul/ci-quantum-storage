// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title CiSHA4096 - Aggressive Gas Variant
 * @notice 64 rounds, 16 states, heavy optimization
 *         Expected gas: ~900k - 1.1M
 */
contract CiSHA4096_Aggressive {
    uint256 constant ROUNDS = 64;
    uint256 constant CI_NUM = 85;
    uint256 constant CI_DEN = 27;

    function ciSha4096(bytes calldata data) public pure returns (bytes32[16] memory) {
        uint256[16] memory state = [
            uint256(0x6a09e667bb67ae85), uint256(0x3c6ef372a54ff53a),
            uint256(0x510e527f9b05688c), uint256(0x1f83d9ab5be0cd19),
            uint256(0x5be0cd191f83d9ab), uint256(0x9b05688c510e527f),
            uint256(0xa54ff53a3c6ef372), uint256(0xbb67ae856a09e667),
            uint256(0x3c6ef372a54ff53a), uint256(0x510e527f9b05688c),
            uint256(0x1f83d9ab5be0cd19), uint256(0x5be0cd191f83d9ab),
            uint256(0x9b05688c510e527f), uint256(0xa54ff53a3c6ef372),
            uint256(0xbb67ae856a09e667), uint256(0x6a09e667bb67ae85)
        ];

        uint256 len = data.length;
        uint256 i = 0;

        while (i < len) {
            uint256 word = 0;
            uint256 maxB = i + 8 < len ? i + 8 : len;
            for (uint256 j = i; j < maxB; j++) {
                word = (word << 8) | uint8(data[j]);
            }
            i = maxB;

            uint256 ciMix = (word * CI_NUM) / CI_DEN;

            for (uint256 r = 0; r < ROUNDS; r++) {
                uint256 temp = state[0];
                state[0] = (temp + ((temp >> 2) ^ (temp >> 13) ^ (temp >> 22)) + ciMix) & type(uint256).max;

                uint256 carry = state[0];
                for (uint256 s = 1; s < 16; s++) {
                    carry = (state[s] ^ carry + ciMix) & type(uint256).max;
                    state[s] = carry;
                }
            }
        }

        bytes32[16] memory output;
        for (uint256 s = 0; s < 16; s++) {
            output[s] = bytes32(state[s]);
        }
        return output;
    }

    function gasBenchmark(bytes calldata message) external view returns (uint256 gasUsed, bytes32[16] memory hash) {
        uint256 start = gasleft();
        hash = ciSha4096(message);
        gasUsed = start - gasleft();
    }

    function verify(bytes calldata data, bytes32[16] calldata expected) public pure returns (bool) {
        bytes32[16] memory computed = ciSha4096(data);
        for (uint256 i = 0; i < 16; i++) {
            if (computed[i] != expected[i]) return false;
        }
        return true;
    }
}

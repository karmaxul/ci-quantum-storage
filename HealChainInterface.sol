// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title HealChain Self-Healing Interface
 * @notice Clean interface to call your HealChain Precompile
 * @dev Precompile at 0x0000000000000000000000000000000000000099
 */
contract HealChainInterface {
    address public constant HEALCHAIN_PRECOMPILE = 0x0000000000000000000000000000000000000099;

    event HealChainEncoded(bytes encodedData);
    event HealChainDecoded(bytes recoveredData);
    event HealChainVerified(bool success, uint256 originalLength);

    function encode(bytes calldata data) external returns (bytes memory encoded) {
        (bool success, bytes memory result) = HEALCHAIN_PRECOMPILE.call(
            abi.encodePacked(uint8(0x01), data)
        );
        require(success, "HealChain: Encode failed");
        encoded = result;
        emit HealChainEncoded(encoded);
        return encoded;
    }

    function decode(bytes calldata encoded) external returns (bytes memory recovered) {
        (bool success, bytes memory result) = HEALCHAIN_PRECOMPILE.call(
            abi.encodePacked(uint8(0x02), encoded)
        );
        require(success, "HealChain: Decode failed");
        recovered = result;
        emit HealChainDecoded(recovered);
        return recovered;
    }

    function verify(bytes calldata encoded) external returns (bool success, uint256 originalLength) {
        (bool callSuccess, bytes memory result) = HEALCHAIN_PRECOMPILE.call(
            abi.encodePacked(uint8(0x03), encoded)
        );
        require(callSuccess, "HealChain: Verify call failed");

        // Safe extraction without range slicing on memory
        success = false;
        originalLength = 0;

        if (result.length >= 32) {
            bytes32 firstWord;
            assembly {
                firstWord := mload(add(result, 32))
            }
            success = (uint256(firstWord) == 1);
        }

        if (result.length >= 64) {
            bytes32 secondWord;
            assembly {
                secondWord := mload(add(result, 64))
            }
            originalLength = uint256(secondWord);
        }

        emit HealChainVerified(success, originalLength);
        return (success, originalLength);
    }
}

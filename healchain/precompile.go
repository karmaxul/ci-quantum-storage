package healchain

import (
    "encoding/binary"
    "fmt"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/vm"
)

const HealChainPrecompileAddress = "0x0000000000000000000000000000000000000099"

type Precompile struct {
    rs *HealChainRS
}

var _ vm.PrecompiledContract = (*Precompile)(nil)

// NewPrecompile creates a new precompile instance
func NewPrecompile(dataShards, parityShards int) (*Precompile, error) {
    rs, err := New(dataShards, parityShards)
    if err != nil {
        return nil, err
    }
    return &Precompile{rs: rs}, nil
}

func (p *Precompile) Address() common.Address {
    return common.HexToAddress(HealChainPrecompileAddress)
}

func (p *Precompile) RequiredGas(input []byte) uint64 {
    // Rough estimate - can be tuned later
    return 80000 + uint64(len(input))*120
}

func (p *Precompile) Run(input []byte) (ret []byte, err error) {
    if len(input) == 0 {
        return nil, fmt.Errorf("empty input")
    }

    method := input[0]

    switch method {
    case 0x01: // Encode
        return p.rs.Encode(input[1:])
    case 0x02: // Decode + Self-Heal
        return p.rs.Decode(input[1:])
    case 0x03: // Verify only (no repair)
        ok, origLen := p.verifyOnly(input[1:])
        result := make([]byte, 64)
        if ok {
            result[0] = 1
        }
        binary.BigEndian.PutUint32(result[32:], origLen)
        return result, nil
    default:
        return nil, fmt.Errorf("unknown method: %d", method)
    }
}

func (p *Precompile) Name() string {
    return "HealChainRS"
}

// verifyOnly checks header and returns success + original length
func (p *Precompile) verifyOnly(encoded []byte) (bool, uint32) {
    if len(encoded) < HeaderSize || encoded[0] != Version {
        return false, 0
    }
    return true, binary.BigEndian.Uint32(encoded[1:5])
}

// Convenience wrappers
func (p *Precompile) Encode(data []byte) ([]byte, error) { return p.rs.Encode(data) }
func (p *Precompile) Decode(encoded []byte) ([]byte, error) { return p.rs.Decode(encoded) }

package healchainrs

import (
    "bytes"
    "crypto/sha256"
    "encoding/binary"
    "fmt"

    "github.com/klauspost/reedsolomon"
)

const (
    Version     byte = 0x01
    HeaderSize       = 1 + 4 + 2 + 2 // version + len + dataShards + parityShards
)

type HealChainRS struct {
    dataShards   int
    parityShards int
    enc          reedsolomon.Encoder
}

func New(dataShards, parityShards int) (*HealChainRS, error) {
    enc, err := reedsolomon.New(dataShards, parityShards)
    if err != nil {
        return nil, err
    }
    return &HealChainRS{
        dataShards:   dataShards,
        parityShards: parityShards,
        enc:          enc,
    }, nil
}

func (h *HealChainRS) Encode(data []byte) ([]byte, error) {
    if len(data) == 0 {
        return nil, fmt.Errorf("empty data")
    }

    shards, err := h.enc.Split(data)
    if err != nil {
        return nil, err
    }
    if err = h.enc.Encode(shards); err != nil {
        return nil, err
    }

    // Per-shard hashes
    hashes := make([][]byte, len(shards))
    for i, s := range shards {
        h := sha256.Sum256(s)
        hashes[i] = h[:]
    }

    // Header
    header := make([]byte, HeaderSize)
    header[0] = Version
    binary.BigEndian.PutUint32(header[1:], uint32(len(data)))
    binary.BigEndian.PutUint16(header[5:], uint16(h.dataShards))
    binary.BigEndian.PutUint16(header[7:], uint16(h.parityShards))

    var buf bytes.Buffer
    buf.Write(header)
    for _, s := range shards {
        buf.Write(s)
    }
    for _, h := range hashes {
        buf.Write(h)
    }

    return buf.Bytes(), nil
}

func (h *HealChainRS) Decode(encoded []byte) ([]byte, error) {
    if len(encoded) < HeaderSize {
        return nil, fmt.Errorf("data too short")
    }

    version := encoded[0]
    if version != Version {
        return nil, fmt.Errorf("unsupported version")
    }

    origLen := int(binary.BigEndian.Uint32(encoded[1:5]))
    dataShards := int(binary.BigEndian.Uint16(encoded[5:7]))
    parityShards := int(binary.BigEndian.Uint16(encoded[7:9]))

    if h.dataShards != dataShards || h.parityShards != parityShards {
        var err error
        h.enc, err = reedsolomon.New(dataShards, parityShards)
        if err != nil {
            return nil, err
        }
        h.dataShards = dataShards
        h.parityShards = parityShards
    }

    shardSize := (origLen + dataShards - 1) / dataShards
    totalShards := dataShards + parityShards

    // Extract shards
    shardStart := HeaderSize
    shards := make([][]byte, totalShards)
    for i := 0; i < totalShards; i++ {
        end := shardStart + shardSize
        if end > len(encoded) {
            end = len(encoded)
        }
        shards[i] = append([]byte{}, encoded[shardStart:end]...)
        shardStart = end
    }

    // Check hashes & mark bad shards
    hashStart := HeaderSize + totalShards*shardSize
    for i := range shards {
        if hashStart+32 > len(encoded) {
            break
        }
        stored := encoded[hashStart : hashStart+32]
        hashStart += 32

        if len(shards[i]) == 0 {
            shards[i] = nil
            continue
        }

        if current := sha256.Sum256(shards[i]); !bytes.Equal(current[:], stored) {
            shards[i] = nil
        }
    }

    if err := h.enc.Reconstruct(shards); err != nil {
        return nil, err
    }

    var buf bytes.Buffer
    if err := h.enc.Join(&buf, shards, origLen); err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}

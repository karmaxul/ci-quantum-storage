package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	bindingsepolia "ci-sha-test/binding-sepolia"
	"ci-sha-test/healchain"
)

// ── Configuration ─────────────────────────────────────────────────────────────

// OracleConfig holds all configuration for the Sepolia oracle.
type OracleConfig struct {
	SepoliaRPC      string
	ContractAddress string
	PrivateKey      string
	ChainID         int64
	PollInterval    time.Duration
	Confirmations   uint64 // blocks to wait before processing
	StateFile       string // path to persist last processed block
}

// ── State persistence ─────────────────────────────────────────────────────────

type oracleState struct {
	LastProcessedBlock uint64 `json:"lastProcessedBlock"`
}

func loadState(path string) oracleState {
	data, err := os.ReadFile(path)
	if err != nil {
		return oracleState{}
	}
	var s oracleState
	if err := json.Unmarshal(data, &s); err != nil {
		return oracleState{}
	}
	return s
}

func saveState(path string, s oracleState) {
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Oracle: failed to marshal state:", err)
		return
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		fmt.Println("Oracle: failed to save state:", err)
	}
}

// ── Oracle ────────────────────────────────────────────────────────────────────

// Oracle watches for EncodeRequested events and fulfills them.
type Oracle struct {
	cfg      OracleConfig
	client   *ethclient.Client
	instance *bindingsepolia.HealChainStorage
	privKey  *ecdsa.PrivateKey
}

// NewOracle creates and connects an Oracle instance.
func NewOracle(cfg OracleConfig) (*Oracle, error) {
	client, err := ethclient.Dial(cfg.SepoliaRPC)
	if err != nil {
		return nil, fmt.Errorf("oracle: failed to connect to Sepolia: %w", err)
	}

	contractAddr := common.HexToAddress(cfg.ContractAddress)
	instance, err := bindingsepolia.NewHealChainStorage(contractAddr, client)
	if err != nil {
		return nil, fmt.Errorf("oracle: failed to load contract: %w", err)
	}

	privKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("oracle: invalid private key: %w", err)
	}

	oracleAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	fmt.Printf("🔮 Oracle initialized\n")
	fmt.Printf("   Contract: %s\n", cfg.ContractAddress)
	fmt.Printf("   Oracle:   %s\n", oracleAddr.Hex())
	fmt.Printf("   Chain ID: %d\n", cfg.ChainID)
	fmt.Printf("   Confirmations required: %d\n", cfg.Confirmations)
	fmt.Printf("   State file: %s\n", cfg.StateFile)

	return &Oracle{
		cfg:      cfg,
		client:   client,
		instance: instance,
		privKey:  privKey,
	}, nil
}

// Start begins polling for EncodeRequested events.
func (o *Oracle) Start(ctx context.Context) {
	fmt.Printf("🔮 Oracle started — polling every %s\n", o.cfg.PollInterval)

	// Load persisted state — resume from last processed block
	state := loadState(o.cfg.StateFile)

	// If no saved state, start from current block minus a small buffer
	if state.LastProcessedBlock == 0 {
		current, err := o.client.BlockNumber(ctx)
		if err != nil {
			fmt.Println("Oracle: failed to get initial block:", err)
		} else {
			// Start from 100 blocks back to catch any missed events
			if current > 100 {
				state.LastProcessedBlock = current - 100
			}
		}
		fmt.Printf("🔮 No saved state — starting from block %d\n", state.LastProcessedBlock)
	} else {
		fmt.Printf("🔮 Resuming from saved block %d\n", state.LastProcessedBlock)
	}

	ticker := time.NewTicker(o.cfg.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("🔮 Oracle stopped — saving state")
			saveState(o.cfg.StateFile, state)
			o.client.Close()
			return

		case <-ticker.C:
			currentBlock, err := o.client.BlockNumber(ctx)
			if err != nil {
				fmt.Println("Oracle: failed to get block number:", err)
				continue
			}

			// Only process blocks that have enough confirmations
			if currentBlock < o.cfg.Confirmations {
				continue
			}
			safeBlock := currentBlock - o.cfg.Confirmations

			if safeBlock <= state.LastProcessedBlock {
				continue
			}

			fromBlock := state.LastProcessedBlock + 1
			toBlock := safeBlock

			// Cap range to avoid overly large queries
			if toBlock-fromBlock > 9 {
				toBlock = fromBlock + 9
			}

			fmt.Printf("🔮 Scanning blocks %d→%d (confirmed up to %d)\n",
				fromBlock, toBlock, safeBlock)

			opts := &bind.FilterOpts{
				Start:   fromBlock,
				End:     &toBlock,
				Context: ctx,
			}

			iter, err := o.instance.FilterEncodeRequested(opts, nil, nil)
			if err != nil {
				fmt.Printf("Oracle: filter error: %v\n", err)
				continue
			}

			for iter.Next() {
				event := iter.Event
				fmt.Printf("🔮 EncodeRequested: requestId=%s requester=%s dataLen=%d label=%s\n",
					event.RequestId.String(),
					event.Requester.Hex(),
					len(event.Data),
					event.Label,
				)
				// Process synchronously to avoid race conditions on nonce
				o.handleEncodeRequest(ctx, event)
			}
			iter.Close()

			// Save progress
			state.LastProcessedBlock = toBlock
			saveState(o.cfg.StateFile, state)
		}
	}
}

// handleEncodeRequest processes a single EncodeRequested event.
func (o *Oracle) handleEncodeRequest(ctx context.Context, event *bindingsepolia.HealChainStorageEncodeRequested) {
	requestId := event.RequestId
	data := event.Data
	dataShards := int(event.DataShards)
	parityShards := int(event.ParityShards)

	// ── Duplicate protection: check isPending before doing any work ───────────
	isPending, err := o.instance.IsPending(nil, requestId)
	if err != nil {
		fmt.Printf("Oracle: isPending check failed for requestId=%s: %v\n", requestId, err)
		return
	}
	if !isPending {
		fmt.Printf("Oracle: requestId=%s already fulfilled or does not exist — skipping\n", requestId)
		return
	}

	fmt.Printf("🔮 Processing requestId=%s (%d bytes, shards %d+%d)\n",
		requestId.String(), len(data), dataShards, parityShards)

	// ── RS encode off-chain ───────────────────────────────────────────────────
	rs, err := healchain.New(dataShards, parityShards)
	if err != nil {
		fmt.Printf("Oracle: RS init failed for requestId=%s: %v\n", requestId, err)
		return
	}

	encoded, err := rs.Encode(data)
	if err != nil {
		fmt.Printf("Oracle: RS encode failed for requestId=%s: %v\n", requestId, err)
		return
	}

	fmt.Printf("🔮 Encoded requestId=%s: %d → %d bytes\n",
		requestId.String(), len(data), len(encoded))

	// ── Call fulfillStore with retries ────────────────────────────────────────
	var tx *types.Transaction
	var txErr error

	for attempt := 1; attempt <= 3; attempt++ {
		// Fresh transactor each attempt — picks up latest nonce
		auth, err2 := bind.NewKeyedTransactorWithChainID(o.privKey, big.NewInt(o.cfg.ChainID))
		if err2 != nil {
			fmt.Printf("Oracle: transactor error attempt %d: %v\n", attempt, err2)
			time.Sleep(time.Duration(attempt) * 2 * time.Second)
			continue
		}
		auth.GasLimit = 3_000_000

		tx, txErr = o.instance.FulfillStore(auth, requestId, encoded)
		if txErr == nil && tx != nil {
			break
		}
		fmt.Printf("Oracle: fulfillStore attempt %d failed: %v\n", attempt, txErr)
		time.Sleep(time.Duration(attempt) * 2 * time.Second)
	}

	if txErr != nil || tx == nil {
		fmt.Printf("Oracle: fulfillStore failed after retries for requestId=%s: %v\n", requestId, txErr)
		return
	}

	fmt.Printf("🔮 fulfillStore submitted: requestId=%s tx=%s\n",
		requestId.String(), tx.Hash().Hex())

	// ── Wait for confirmation ─────────────────────────────────────────────────
	receipt, err := bind.WaitMined(ctx, o.client, tx)
	if err != nil {
		fmt.Printf("Oracle: WaitMined failed for requestId=%s: %v\n", requestId, err)
		return
	}

	if receipt.Status == 0 {
		fmt.Printf("Oracle: fulfillStore reverted for requestId=%s\n", requestId)
		return
	}

	fmt.Printf("✅ Oracle fulfilled requestId=%s | block=%d | tx=%s\n",
		requestId.String(), receipt.BlockNumber.Uint64(), tx.Hash().Hex())
}

// ── Sepolia store helper ──────────────────────────────────────────────────────

// sepoliaStoreOnChain sends a store request to the Sepolia contract.
// Returns the tx hash and requestId. The oracle fulfills async.
func sepoliaStoreOnChain(
	data []byte,
	label string,
) (txHash string, requestID string, err error) {

	sepoliaURL := getEnv("SEPOLIA_RPC_URL", "")
	contractAddr := getEnv("SEPOLIA_CONTRACT_ADDRESS", "")
	privKeyHex := getEnv("ORACLE_PRIVATE_KEY", getEnv("STORE_PRIVATE_KEY",
		"b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291"))

	if sepoliaURL == "" || contractAddr == "" {
		return "", "", fmt.Errorf("SEPOLIA_RPC_URL and SEPOLIA_CONTRACT_ADDRESS must be set")
	}

	client, err := ethclient.Dial(sepoliaURL)
	if err != nil {
		return "", "", fmt.Errorf("failed to connect to Sepolia: %w", err)
	}
	defer client.Close()

	addr := common.HexToAddress(contractAddr)
	instance, err := bindingsepolia.NewHealChainStorage(addr, client)
	if err != nil {
		return "", "", fmt.Errorf("failed to load contract: %w", err)
	}

	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return "", "", fmt.Errorf("invalid private key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(11155111))
	if err != nil {
		return "", "", fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.GasLimit = 200_000

	tx, err := instance.Store0(auth, data, label)
	if err != nil {
		return "", "", fmt.Errorf("store tx failed: %w", err)
	}

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return "", "", fmt.Errorf("WaitMined failed: %w", err)
	}

	if receipt.Status == 0 {
		return "", "", fmt.Errorf("store transaction reverted")
	}

	// Parse EncodeRequested event to get requestId
	var reqID *big.Int
	for _, log := range receipt.Logs {
		event, parseErr := instance.ParseEncodeRequested(*log)
		if parseErr == nil {
			reqID = event.RequestId
			break
		}
	}

	if reqID == nil {
		return tx.Hash().Hex(), "", fmt.Errorf("could not parse requestId from event")
	}

	fmt.Printf("Store submitted to Sepolia: tx=%s requestId=%s\n",
		tx.Hash().Hex(), reqID.String())

	return tx.Hash().Hex(), reqID.String(), nil
}

// filterLogs is a helper for manual log filtering.
func filterLogs(client *ethclient.Client, ctx context.Context,
	contractAddr common.Address, fromBlock, toBlock uint64,
	topic common.Hash) ([]types.Log, error) {

	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []common.Address{contractAddr},
		Topics:    [][]common.Hash{{topic}},
	}
	return client.FilterLogs(ctx, query)
}

// suppress unused import
var _ = hex.EncodeToString
var _ = filepath.Join

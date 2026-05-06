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
	"sync"
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
	LastProcessedBlock    uint64 `json:"lastProcessedBlock"`
	RequestsObserved      uint64 `json:"requestsObserved,omitempty"`
	FulfillmentsSucceeded uint64 `json:"fulfillmentsSucceeded,omitempty"`
	FulfillmentsFailed    uint64 `json:"fulfillmentsFailed,omitempty"`
	FulfillmentsSkipped   uint64 `json:"fulfillmentsSkipped,omitempty"`
	GasUsedTotal          uint64 `json:"gasUsedTotal,omitempty"`
	GasSpentWei           string `json:"gasSpentWei,omitempty"` // big.Int as decimal string
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

	mu      sync.Mutex
	metrics oracleMetricsState
}

// oracleMetricsState is the live in-memory metrics. All access must hold mu.
type oracleMetricsState struct {
	StartedAt             time.Time
	LastScanAt            time.Time
	LastFulfillmentAt     time.Time
	LastErrorAt           time.Time
	CurrentBlock          uint64
	LastProcessedBlock    uint64
	RequestsObserved      uint64
	FulfillmentsSucceeded uint64
	FulfillmentsFailed    uint64
	FulfillmentsSkipped   uint64
	InFlight              uint64
	GasUsedTotal          uint64
	GasSpentWei           *big.Int
	LastTxHash            string
	LastError             string
	OracleAddress         string
}

// OracleMetrics is the JSON-friendly snapshot returned by Oracle.Metrics().
type OracleMetrics struct {
	Enabled                     bool    `json:"enabled"`
	OracleAddress               string  `json:"oracleAddress,omitempty"`
	ContractAddress             string  `json:"contractAddress,omitempty"`
	ChainID                     int64   `json:"chainId,omitempty"`
	StartedAt                   string  `json:"startedAt,omitempty"`
	UptimeSeconds               int64   `json:"uptimeSeconds"`
	LastScanAt                  string  `json:"lastScanAt,omitempty"`
	SecondsSinceLastScan        int64   `json:"secondsSinceLastScan"`
	LastFulfillmentAt           string  `json:"lastFulfillmentAt,omitempty"`
	SecondsSinceLastFulfillment int64   `json:"secondsSinceLastFulfillment"`
	PollIntervalSeconds         int64   `json:"pollIntervalSeconds"`
	CurrentBlock                uint64  `json:"currentBlock"`
	LastProcessedBlock          uint64  `json:"lastProcessedBlock"`
	BlocksBehind                uint64  `json:"blocksBehind"`
	RequestsObserved            uint64  `json:"requestsObserved"`
	FulfillmentsSucceeded       uint64  `json:"fulfillmentsSucceeded"`
	FulfillmentsFailed          uint64  `json:"fulfillmentsFailed"`
	FulfillmentsSkipped         uint64  `json:"fulfillmentsSkipped"`
	InFlight                    uint64  `json:"inFlight"`
	SuccessRate                 float64 `json:"successRate"`
	GasUsedTotal                uint64  `json:"gasUsedTotal"`
	GasSpentWei                 string  `json:"gasSpentWei"`
	GasSpentEth                 string  `json:"gasSpentEth"`
	LastTxHash                  string  `json:"lastTxHash,omitempty"`
	LastError                   string  `json:"lastError,omitempty"`
	LastErrorAt                 string  `json:"lastErrorAt,omitempty"`
	Healthy                     bool    `json:"healthy"`
	HealthReason                string  `json:"healthReason,omitempty"`
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
		metrics: oracleMetricsState{
			OracleAddress: oracleAddr.Hex(),
			GasSpentWei:   new(big.Int),
		},
	}, nil
}

// Start begins polling for EncodeRequested events.
func (o *Oracle) Start(ctx context.Context) {
	fmt.Printf("🔮 Oracle started — polling every %s\n", o.cfg.PollInterval)

	// Load persisted state — resume from last processed block
	state := loadState(o.cfg.StateFile)

	// Hydrate in-memory metrics from persisted counters
	o.hydrateMetrics(state)

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

	o.recordLastProcessedBlock(state.LastProcessedBlock)

	ticker := time.NewTicker(o.cfg.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("🔮 Oracle stopped — saving state")
			saveState(o.cfg.StateFile, o.snapshotState(state.LastProcessedBlock))
			o.client.Close()
			return

		case <-ticker.C:
			currentBlock, err := o.client.BlockNumber(ctx)
			if err != nil {
				fmt.Println("Oracle: failed to get block number:", err)
				o.recordError(fmt.Errorf("blockNumber: %w", err))
				continue
			}

			o.recordScan(currentBlock)

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
				o.recordError(fmt.Errorf("filter: %w", err))
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
				o.recordObserved()
				// Process synchronously to avoid race conditions on nonce
				o.handleEncodeRequest(ctx, event)
			}
			iter.Close()

			// Save progress (block + cumulative counters)
			state.LastProcessedBlock = toBlock
			o.recordLastProcessedBlock(toBlock)
			saveState(o.cfg.StateFile, o.snapshotState(toBlock))
		}
	}
}

// handleEncodeRequest processes a single EncodeRequested event.
func (o *Oracle) handleEncodeRequest(ctx context.Context, event *bindingsepolia.HealChainStorageEncodeRequested) {
	requestId := event.RequestId
	data := event.Data
	dataShards := int(event.DataShards)
	parityShards := int(event.ParityShards)

	o.recordInFlight(true)
	defer o.recordInFlight(false)

	// ── Duplicate protection: check isPending before doing any work ───────────
	isPending, err := o.instance.IsPending(nil, requestId)
	if err != nil {
		fmt.Printf("Oracle: isPending check failed for requestId=%s: %v\n", requestId, err)
		o.recordError(fmt.Errorf("isPending: %w", err))
		o.recordFailed()
		return
	}
	if !isPending {
		fmt.Printf("Oracle: requestId=%s already fulfilled or does not exist — skipping\n", requestId)
		o.recordSkipped()
		return
	}

	fmt.Printf("🔮 Processing requestId=%s (%d bytes, shards %d+%d)\n",
		requestId.String(), len(data), dataShards, parityShards)

	// ── RS encode off-chain ───────────────────────────────────────────────────
	rs, err := healchain.New(dataShards, parityShards)
	if err != nil {
		fmt.Printf("Oracle: RS init failed for requestId=%s: %v\n", requestId, err)
		o.recordError(fmt.Errorf("RS init: %w", err))
		o.recordFailed()
		return
	}

	encoded, err := rs.Encode(data)
	if err != nil {
		fmt.Printf("Oracle: RS encode failed for requestId=%s: %v\n", requestId, err)
		o.recordError(fmt.Errorf("RS encode: %w", err))
		o.recordFailed()
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
		o.recordError(fmt.Errorf("fulfillStore: %w", txErr))
		o.recordFailed()
		return
	}

	fmt.Printf("🔮 fulfillStore submitted: requestId=%s tx=%s\n",
		requestId.String(), tx.Hash().Hex())

	// ── Wait for confirmation ─────────────────────────────────────────────────
	receipt, err := bind.WaitMined(ctx, o.client, tx)
	if err != nil {
		fmt.Printf("Oracle: WaitMined failed for requestId=%s: %v\n", requestId, err)
		o.recordError(fmt.Errorf("WaitMined: %w", err))
		o.recordFailed()
		return
	}

	if receipt.Status == 0 {
		fmt.Printf("Oracle: fulfillStore reverted for requestId=%s\n", requestId)
		o.recordError(fmt.Errorf("fulfillStore reverted (requestId=%s)", requestId))
		// Reverted txs still consume gas — count it
		o.recordSucceededOrReverted(tx, receipt, false)
		return
	}

	fmt.Printf("✅ Oracle fulfilled requestId=%s | block=%d | tx=%s\n",
		requestId.String(), receipt.BlockNumber.Uint64(), tx.Hash().Hex())
	o.recordSucceededOrReverted(tx, receipt, true)
}

// ── Metrics helpers ───────────────────────────────────────────────────────────

func (o *Oracle) hydrateMetrics(state oracleState) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.metrics.GasSpentWei == nil {
		o.metrics.GasSpentWei = new(big.Int)
	}
	o.metrics.StartedAt = time.Now()
	o.metrics.LastProcessedBlock = state.LastProcessedBlock
	o.metrics.RequestsObserved = state.RequestsObserved
	o.metrics.FulfillmentsSucceeded = state.FulfillmentsSucceeded
	o.metrics.FulfillmentsFailed = state.FulfillmentsFailed
	o.metrics.FulfillmentsSkipped = state.FulfillmentsSkipped
	o.metrics.GasUsedTotal = state.GasUsedTotal
	if state.GasSpentWei != "" {
		if g, ok := new(big.Int).SetString(state.GasSpentWei, 10); ok {
			o.metrics.GasSpentWei = g
		}
	}
}

// snapshotState builds a persistable oracleState from current metrics.
func (o *Oracle) snapshotState(lastProcessedBlock uint64) oracleState {
	o.mu.Lock()
	defer o.mu.Unlock()
	gas := "0"
	if o.metrics.GasSpentWei != nil {
		gas = o.metrics.GasSpentWei.String()
	}
	return oracleState{
		LastProcessedBlock:    lastProcessedBlock,
		RequestsObserved:      o.metrics.RequestsObserved,
		FulfillmentsSucceeded: o.metrics.FulfillmentsSucceeded,
		FulfillmentsFailed:    o.metrics.FulfillmentsFailed,
		FulfillmentsSkipped:   o.metrics.FulfillmentsSkipped,
		GasUsedTotal:          o.metrics.GasUsedTotal,
		GasSpentWei:           gas,
	}
}

func (o *Oracle) recordScan(currentBlock uint64) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.metrics.LastScanAt = time.Now()
	o.metrics.CurrentBlock = currentBlock
}

func (o *Oracle) recordLastProcessedBlock(b uint64) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.metrics.LastProcessedBlock = b
}

func (o *Oracle) recordObserved() {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.metrics.RequestsObserved++
}

func (o *Oracle) recordInFlight(active bool) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if active {
		o.metrics.InFlight = 1
	} else {
		o.metrics.InFlight = 0
	}
}

func (o *Oracle) recordSkipped() {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.metrics.FulfillmentsSkipped++
}

func (o *Oracle) recordFailed() {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.metrics.FulfillmentsFailed++
}

func (o *Oracle) recordError(err error) {
	if err == nil {
		return
	}
	o.mu.Lock()
	defer o.mu.Unlock()
	o.metrics.LastError = err.Error()
	o.metrics.LastErrorAt = time.Now()
}

// recordSucceededOrReverted accounts gas for a mined tx and bumps the
// success or failure counter depending on whether the receipt status is OK.
func (o *Oracle) recordSucceededOrReverted(tx *types.Transaction, receipt *types.Receipt, ok bool) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.metrics.GasSpentWei == nil {
		o.metrics.GasSpentWei = new(big.Int)
	}
	if receipt != nil {
		o.metrics.GasUsedTotal += receipt.GasUsed
		gasPrice := receipt.EffectiveGasPrice
		if gasPrice == nil && tx != nil {
			gasPrice = tx.GasPrice()
		}
		if gasPrice != nil {
			cost := new(big.Int).Mul(
				new(big.Int).SetUint64(receipt.GasUsed),
				gasPrice,
			)
			o.metrics.GasSpentWei.Add(o.metrics.GasSpentWei, cost)
		}
	}

	if tx != nil {
		o.metrics.LastTxHash = tx.Hash().Hex()
	}

	if ok {
		o.metrics.FulfillmentsSucceeded++
		o.metrics.LastFulfillmentAt = time.Now()
	} else {
		o.metrics.FulfillmentsFailed++
	}
}

// Metrics returns a JSON-friendly snapshot of the oracle's metrics.
func (o *Oracle) Metrics() OracleMetrics {
	o.mu.Lock()
	defer o.mu.Unlock()

	m := o.metrics
	now := time.Now()

	var snap OracleMetrics
	snap.Enabled = true
	snap.OracleAddress = m.OracleAddress
	snap.ContractAddress = o.cfg.ContractAddress
	snap.ChainID = o.cfg.ChainID
	snap.PollIntervalSeconds = int64(o.cfg.PollInterval / time.Second)

	if !m.StartedAt.IsZero() {
		snap.StartedAt = m.StartedAt.UTC().Format(time.RFC3339)
		snap.UptimeSeconds = int64(now.Sub(m.StartedAt).Seconds())
	}
	if !m.LastScanAt.IsZero() {
		snap.LastScanAt = m.LastScanAt.UTC().Format(time.RFC3339)
		snap.SecondsSinceLastScan = int64(now.Sub(m.LastScanAt).Seconds())
	}
	if !m.LastFulfillmentAt.IsZero() {
		snap.LastFulfillmentAt = m.LastFulfillmentAt.UTC().Format(time.RFC3339)
		snap.SecondsSinceLastFulfillment = int64(now.Sub(m.LastFulfillmentAt).Seconds())
	}
	if !m.LastErrorAt.IsZero() {
		snap.LastErrorAt = m.LastErrorAt.UTC().Format(time.RFC3339)
	}

	snap.CurrentBlock = m.CurrentBlock
	snap.LastProcessedBlock = m.LastProcessedBlock
	if m.CurrentBlock > m.LastProcessedBlock {
		snap.BlocksBehind = m.CurrentBlock - m.LastProcessedBlock
	}

	snap.RequestsObserved = m.RequestsObserved
	snap.FulfillmentsSucceeded = m.FulfillmentsSucceeded
	snap.FulfillmentsFailed = m.FulfillmentsFailed
	snap.FulfillmentsSkipped = m.FulfillmentsSkipped
	snap.InFlight = m.InFlight

	finalized := m.FulfillmentsSucceeded + m.FulfillmentsFailed
	if finalized > 0 {
		snap.SuccessRate = float64(m.FulfillmentsSucceeded) / float64(finalized)
	}

	snap.GasUsedTotal = m.GasUsedTotal
	if m.GasSpentWei != nil {
		snap.GasSpentWei = m.GasSpentWei.String()
		eth := new(big.Float).Quo(
			new(big.Float).SetInt(m.GasSpentWei),
			new(big.Float).SetInt(big.NewInt(1_000_000_000_000_000_000)),
		)
		snap.GasSpentEth = eth.Text('f', 6)
	} else {
		snap.GasSpentWei = "0"
		snap.GasSpentEth = "0.000000"
	}

	snap.LastTxHash = m.LastTxHash
	snap.LastError = m.LastError

	// Health: scanning recently, no recent error spike
	snap.Healthy = true
	switch {
	case m.LastScanAt.IsZero():
		snap.Healthy = false
		snap.HealthReason = "oracle has not completed a scan yet"
	case now.Sub(m.LastScanAt) > 3*o.cfg.PollInterval:
		snap.Healthy = false
		snap.HealthReason = fmt.Sprintf("no successful scan in %ds (poll interval %s)",
			int(now.Sub(m.LastScanAt).Seconds()), o.cfg.PollInterval)
	case !m.LastErrorAt.IsZero() && now.Sub(m.LastErrorAt) < 60*time.Second:
		snap.Healthy = false
		snap.HealthReason = "recent error: " + m.LastError
	}

	return snap
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

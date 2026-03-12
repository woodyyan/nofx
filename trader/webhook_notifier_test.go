package trader

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"nofx/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildWebhookPayload(t *testing.T) {
	at := &AutoTrader{
		id:             "trader-test-001",
		name:           "Test Trader",
		initialBalance: 1000.0,
	}

	record := &store.DecisionRecord{
		TraderID:    "trader-test-001",
		CycleNumber: 5,
		Timestamp:   time.Now(),
		Decisions: []store.DecisionAction{
			{
				Symbol:     "BTCUSDT",
				Action:     "open_long",
				Leverage:   5,
				StopLoss:   80000,
				TakeProfit: 90000,
				Confidence: 82,
				Reasoning:  "Strong breakout",
				Success:    true,
			},
		},
		AccountState: store.AccountSnapshot{
			TotalBalance:          1100.0,
			AvailableBalance:      800.0,
			TotalUnrealizedProfit: 50.0,
			PositionCount:         1,
			MarginUsedPct:         20.0,
		},
	}

	payload := at.buildWebhookPayload(record, at.initialBalance)

	assert.Equal(t, "trader-test-001", payload.TraderID)
	assert.Equal(t, "Test Trader", payload.TraderName)
	assert.Equal(t, 5, payload.CycleNum)
	assert.Len(t, payload.Decisions, 1)

	// PnL calculations
	assert.InDelta(t, 100.0, payload.Account.TotalPnL, 0.001)
	assert.InDelta(t, 10.0, payload.Account.TotalPnLPct, 0.001)
	assert.Equal(t, 1100.0, payload.Account.TotalEquity)
	assert.Equal(t, 800.0, payload.Account.AvailableBalance)
	assert.Equal(t, 1000.0, payload.Account.InitialBalance)
	assert.Equal(t, 50.0, payload.Account.UnrealizedPnL)

	// Decision fields
	d := payload.Decisions[0]
	assert.Equal(t, "BTCUSDT", d.Symbol)
	assert.Equal(t, "open_long", d.Action)
	assert.Equal(t, 5, d.Leverage)
	assert.Equal(t, 82, d.Confidence)
	assert.True(t, d.Success)
}

func TestBuildWebhookPayload_ZeroInitialBalance(t *testing.T) {
	at := &AutoTrader{
		id:             "trader-zero",
		name:           "Zero Balance Trader",
		initialBalance: 0,
	}

	record := &store.DecisionRecord{
		TraderID:    "trader-zero",
		CycleNumber: 1,
		Decisions:   []store.DecisionAction{},
		AccountState: store.AccountSnapshot{
			TotalBalance: 0,
		},
	}

	payload := at.buildWebhookPayload(record, at.initialBalance)
	// Should not panic with zero initial balance (division by zero guard)
	assert.Equal(t, 0.0, payload.Account.TotalPnLPct)
}

func TestSendWebhookPayload_Success(t *testing.T) {
	var received CycleWebhookPayload
	var receivedSig string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		receivedSig = r.Header.Get("X-Nofx-Signature")
		err := json.NewDecoder(r.Body).Decode(&received)
		require.NoError(t, err)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	payload := CycleWebhookPayload{
		TraderID:   "t001",
		TraderName: "Bot",
		CycleNum:   3,
		Timestamp:  time.Now(),
		Decisions:  []WebhookDecision{},
		Account: WebhookAccount{
			TotalPnL: 42.0,
		},
	}

	err := sendWebhookPayload(server.URL, "", payload)
	require.NoError(t, err)
	assert.Equal(t, "t001", received.TraderID)
	assert.Equal(t, 3, received.CycleNum)
	assert.Empty(t, receivedSig) // No secret = no signature
}

func TestSendWebhookPayload_WithSecret(t *testing.T) {
	var receivedSig string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedSig = r.Header.Get("X-Nofx-Signature")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	payload := CycleWebhookPayload{TraderID: "t001", CycleNum: 1, Decisions: []WebhookDecision{}}
	err := sendWebhookPayload(server.URL, "mysecret", payload)
	require.NoError(t, err)
	assert.NotEmpty(t, receivedSig)
	assert.Contains(t, receivedSig, "sha256=")
}

func TestSendWebhookPayload_NonSuccessStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	payload := CycleWebhookPayload{Decisions: []WebhookDecision{}}
	err := sendWebhookPayload(server.URL, "", payload)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "500")
}

func TestSendWebhookPayload_InvalidURL(t *testing.T) {
	payload := CycleWebhookPayload{Decisions: []WebhookDecision{}}
	err := sendWebhookPayload("http://127.0.0.1:1", "", payload)
	assert.Error(t, err) // connection refused
}

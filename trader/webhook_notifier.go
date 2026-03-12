package trader

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"nofx/logger"
	"nofx/store"
	"time"
)

// CycleWebhookPayload is the payload sent to the webhook on each trading cycle
type CycleWebhookPayload struct {
	// Meta info
	TraderID   string    `json:"trader_id"`
	TraderName string    `json:"trader_name"`
	CycleNum   int       `json:"cycle_number"`
	Timestamp  time.Time `json:"timestamp"`

	// Decisions in this cycle
	Decisions []WebhookDecision `json:"decisions"`

	// Account snapshot at time of this cycle
	Account WebhookAccount `json:"account"`
}

// WebhookDecision summarises a single trading decision
type WebhookDecision struct {
	Symbol          string  `json:"symbol"`
	Action          string  `json:"action"`
	Leverage        int     `json:"leverage,omitempty"`
	PositionSizeUSD float64 `json:"position_size_usd,omitempty"`
	StopLoss        float64 `json:"stop_loss,omitempty"`
	TakeProfit      float64 `json:"take_profit,omitempty"`
	Confidence      int     `json:"confidence,omitempty"`
	Reasoning       string  `json:"reasoning,omitempty"`
	Success         bool    `json:"success"`
	Error           string  `json:"error,omitempty"`
}

// WebhookAccount summarises the account state for this cycle
type WebhookAccount struct {
	TotalEquity      float64 `json:"total_equity"`
	AvailableBalance float64 `json:"available_balance"`
	InitialBalance   float64 `json:"initial_balance"`
	TotalPnL         float64 `json:"total_pnl"`
	TotalPnLPct      float64 `json:"total_pnl_pct"`
	UnrealizedPnL    float64 `json:"unrealized_pnl"`
	PositionCount    int     `json:"position_count"`
	MarginUsedPct    float64 `json:"margin_used_pct"`
}

// notifyWebhook fires an async webhook notification after each cycle.
// It reads the webhook config fresh from the DB each time so config
// changes take effect immediately without a trader restart.
func (at *AutoTrader) notifyWebhook(record *store.DecisionRecord) {
	if at.store == nil {
		return
	}

	traderCfg, err := at.store.Trader().GetByID(at.id)
	if err != nil || traderCfg == nil {
		return
	}
	if !traderCfg.WebhookEnabled || traderCfg.WebhookURL == "" {
		return
	}

	payload := at.buildWebhookPayload(record, traderCfg.InitialBalance)
	url := traderCfg.WebhookURL
	secret := traderCfg.WebhookSecret

	go func() {
		if err := sendWebhookPayload(url, secret, payload); err != nil {
			logger.Infof("⚠️ [%s] Webhook send failed (cycle #%d): %v", at.name, payload.CycleNum, err)
		} else {
			logger.Infof("📡 [%s] Webhook sent (cycle #%d, decisions=%d)", at.name, payload.CycleNum, len(payload.Decisions))
		}
	}()
}

// buildWebhookPayload converts a DecisionRecord into the public webhook payload.
func (at *AutoTrader) buildWebhookPayload(record *store.DecisionRecord, initialBalance float64) CycleWebhookPayload {
	decisions := make([]WebhookDecision, 0, len(record.Decisions))
	for _, d := range record.Decisions {
		decisions = append(decisions, WebhookDecision{
			Symbol:     d.Symbol,
			Action:     d.Action,
			Leverage:   d.Leverage,
			StopLoss:   d.StopLoss,
			TakeProfit: d.TakeProfit,
			Confidence: d.Confidence,
			Reasoning:  d.Reasoning,
			Success:    d.Success,
			Error:      d.Error,
		})
	}

	acct := record.AccountState
	totalPnL := acct.TotalBalance - initialBalance
	totalPnLPct := 0.0
	if initialBalance > 0 {
		totalPnLPct = totalPnL / initialBalance * 100
	}

	return CycleWebhookPayload{
		TraderID:   record.TraderID,
		TraderName: at.name,
		CycleNum:   record.CycleNumber,
		Timestamp:  time.Now().UTC(),
		Decisions:  decisions,
		Account: WebhookAccount{
			TotalEquity:      acct.TotalBalance,
			AvailableBalance: acct.AvailableBalance,
			InitialBalance:   initialBalance,
			TotalPnL:         totalPnL,
			TotalPnLPct:      totalPnLPct,
			UnrealizedPnL:    acct.TotalUnrealizedProfit,
			PositionCount:    acct.PositionCount,
			MarginUsedPct:    acct.MarginUsedPct,
		},
	}
}

// sendWebhookPayload serialises the payload and POSTs it to the target URL.
// If secret is non-empty, an HMAC-SHA256 signature is added as X-Nofx-Signature.
func sendWebhookPayload(url, secret string, payload CycleWebhookPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "nofx-webhook/1.0")

	// Optional HMAC-SHA256 signature for receiver verification
	if secret != "" {
		mac := hmac.New(sha256.New, []byte(secret))
		mac.Write(body)
		sig := hex.EncodeToString(mac.Sum(nil))
		req.Header.Set("X-Nofx-Signature", "sha256="+sig)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

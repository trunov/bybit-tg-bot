package bybit

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type BybitClient struct {
	APIKey    string
	APISecret string
	Client    *http.Client
	BaseURL   string
}

func NewBybitClient(apiKey, apiSecret string) *BybitClient {
	return &BybitClient{
		APIKey:    apiKey,
		APISecret: apiSecret,
		Client:    &http.Client{},
		BaseURL:   "https://api2.bybit.com/fiat/otc/item/online",
	}
}

func (b *BybitClient) generateSignature(payload map[string]interface{}) (string, error) {
	keys := make([]string, 0, len(payload))
	for k := range payload {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	queryString := ""
	for i, k := range keys {
		queryString += fmt.Sprintf("%s=%v", k, payload[k])
		if i < len(keys)-1 {
			queryString += "&"
		}
	}

	mac := hmac.New(sha256.New, []byte(b.APISecret))
	_, err := mac.Write([]byte(queryString))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func (b *BybitClient) FetchOTCAds() (BybitResponse, error) {
	// TODO: Add extra request for sell
	payload := map[string]interface{}{
		"currencyId":  "RUB",
		"itemRegion":  1,
		"makerUserId": "25901077",
		"page":        "1",
		"side":        "1", // 1 for buy ads, 0 for sell ads
		"size":        "2000",
		"tokenId":     "USDT",
		"userId":      183223363,
	}

	payload["api_key"] = b.APIKey
	payload["timestamp"] = strconv.FormatInt(time.Now().UnixMilli(), 10)

	var response BybitResponse

	signature, err := b.generateSignature(payload)
	if err != nil {
		return response, fmt.Errorf("failed to generate signature: %w", err)
	}
	payload["sign"] = signature

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return response, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", b.BaseURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return response, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := b.Client.Do(req)
	if err != nil {
		return response, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("error: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response, fmt.Errorf("failed to decode response: %w", err)
	}

	return response, nil
}

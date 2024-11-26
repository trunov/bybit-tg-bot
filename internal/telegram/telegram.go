package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TelegramBot struct {
	BotToken  string
	ChannelID string
}

type MessagePayload struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func NewTelegramBot(botToken, channelID string) *TelegramBot {
	return &TelegramBot{
		BotToken:  botToken,
		ChannelID: channelID,
	}
}

func (t *TelegramBot) SendText(amount, rate, bank string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.BotToken)

	link := "https://www.bybit.com/fiat/trade/otc/profile/25901077/USDT/RUB/item"

	// TODO: add green/red circle logic based on method
	message := fmt.Sprintf(
		"🔴 Новая заявка на %s USDT\n✅ Курс: %s\n🏦 Банк: %s\n🔗 Ссылка на ордер: %s",
		amount, rate, bank, link,
	)

	payload := MessagePayload{
		ChatID: t.ChannelID,
		Text:   message,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	return nil
}

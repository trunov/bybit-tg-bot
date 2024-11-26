package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
	"github.com/trunov/bybit-tg-bot/internal/bybit"
	"github.com/trunov/bybit-tg-bot/internal/telegram"
)

type Worker struct {
	bybitClient *bybit.BybitClient
	tgBot       *telegram.TelegramBot
	history     map[string]bool
}

func (w *Worker) clearHistory() {
	w.history = make(map[string]bool)
}

func NewWorker(apiKey, apiSecret, botToken, channelID string) *Worker {
	client := bybit.NewBybitClient(apiKey, apiSecret)
	tgBot := telegram.NewTelegramBot(botToken, channelID)

	return &Worker{bybitClient: client, tgBot: tgBot, history: make(map[string]bool)}
}

func (w *Worker) RunWorker(ctx context.Context) error {
	c := cron.New()

	c.AddFunc("*/1 * * * *", func() {
		ads, err := w.bybitClient.FetchOTCAds()
		if err != nil {
			fmt.Println(err)
		}

		if ads.Result.Count > 0 {
			rate := ads.Result.Items[0].Price
			amount := ads.Result.Items[0].LastQuantity

			messageKey := fmt.Sprintf("%s_%s", rate, amount)
			if w.history[messageKey] {
				return
			}

			// TODO: add api list call (data for payments can be taken from ads.Result.Items[0].Payments[0])
			bank := "Сбер"

			err := w.tgBot.SendText(amount, rate, bank)
			if err != nil {
				log.Println("Failed to send Telegram message:", err)
				return
			}

			w.history[messageKey] = true
		} else {
			fmt.Println("No new ads:", ads)
		}
	})

	c.AddFunc("@midnight", func() {
		log.Println("Clearing message history")
		w.clearHistory()
	})

	c.Start()

	<-ctx.Done()
	c.Stop()

	return ctx.Err()
}

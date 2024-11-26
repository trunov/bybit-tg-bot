package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/trunov/bybit-tg-bot/internal/config"
	"github.com/trunov/bybit-tg-bot/internal/worker"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Failed to read config", err)
	}

	if cfg.ApiKey == "" || cfg.ApiSecret == "" {
		log.Fatal("API key or secret is missing in the configuration")
	}

	if cfg.BotToken == "" || cfg.ChannelID == "" {
		log.Fatal("BotToken or ChannelID is missing in the configuration")
	}

	w := worker.NewWorker(cfg.ApiKey, cfg.ApiKey, cfg.BotToken, cfg.ChannelID)
	go func() {
		if err := w.RunWorker(ctx); err != nil {
			log.Println("Worker stopped with error:", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutdown signal received, stopping worker...")

	cancel()
}

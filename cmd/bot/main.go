package main

import (
	"log"
	"tour-guide-bot/internal/bot"
	"tour-guide-bot/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	bot, err := bot.NewBot(cfg.TelegramToken, cfg.ChannelName, cfg.GuideUrl)
	if err != nil {
		log.Fatalf("Failed to init bot: %v", err)
	}

	bot.Start()
}

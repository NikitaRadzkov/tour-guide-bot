package main

import (
	"log"
	"tour-guide-bot/internal/bot"
	"tour-guide-bot/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("Prepare bot...")

	botInstance, err := bot.NewBot(cfg.TelegramToken, cfg.ChannelName, cfg.GuideUrl)
	if err != nil {
		log.Fatalf("Failed to init bot: %v", err)
	}

	botInstance.Start()
}
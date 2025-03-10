package main

import (
	"log"
	"net/http"
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

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

		port := cfg.Port

		log.Printf("Starting server on port %s", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatalf("Failed to start health-check server: %v", err)
		}
	}()

	log.Printf("Prepare bot...")

	botInstance, err := bot.NewBot(cfg.TelegramToken, cfg.ChannelName, cfg.GuideUrl)
	if err != nil {
		log.Fatalf("Failed to init bot: %v", err)
	}

	botInstance.Start()
}
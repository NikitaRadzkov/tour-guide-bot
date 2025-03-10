package config

import (
	"fmt"
	"os"
)

type Config struct {
	TelegramToken string
	ChannelName string
	GuideUrl string
	Port string
}

func LoadConfig() (*Config, error) {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("TELEGRAM_TOKEN is not set")
	}

	channelName := os.Getenv("TELEGRAM_CHANNEL_NAME")
	if channelName == "" {
		return nil, fmt.Errorf("TELEGRAM_CHANNEL_NAME is not set")
	}

	guideUrl := os.Getenv("GUIDE_URL")
	if guideUrl == "" {
		return nil, fmt.Errorf("GUIDE_URL is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("PORT is not set")
	}

	return &Config{
		TelegramToken: token,
		ChannelName: channelName,
		GuideUrl: guideUrl,
		Port: port,
	}, nil
}

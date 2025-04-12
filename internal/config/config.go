package config

import (
	"fmt"
	"os"
)

type Config struct {
	TelegramToken string
	ChannelName string
	GuideUrl string
	TopDealUrl string
	ChecklistUrl string
	SearchUrl string
	AboutUrl string
	ContactUser string
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

	topDealUrl := os.Getenv("TOP_DEALS_URL")
	if guideUrl == "" {
		return nil, fmt.Errorf("TOP_DEALS_URL is not set")
	}

	checklistUrl := os.Getenv("CHECKLIST_URL")
	if guideUrl == "" {
		return nil, fmt.Errorf("CHECKLIST_URL is not set")
	}

	searchUrl := os.Getenv("SEARCH_URL")
	if guideUrl == "" {
		return nil, fmt.Errorf("SEARCH_URL is not set")
	}

	aboutUrl := os.Getenv("ABOUT_URL")
	if guideUrl == "" {
		return nil, fmt.Errorf("ABOUT_URL is not set")
	}

	contactUser := os.Getenv("CONTACT_USER")
	if guideUrl == "" {
		return nil, fmt.Errorf("CONTACT_USER is not set")
	}

	return &Config{
		TelegramToken: token,
		ChannelName: channelName,
		GuideUrl: guideUrl,
		TopDealUrl: topDealUrl,
		ChecklistUrl: checklistUrl,
		SearchUrl: searchUrl,
		AboutUrl: aboutUrl,
		ContactUser: contactUser,
	}, nil
}

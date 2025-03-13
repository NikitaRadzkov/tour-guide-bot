# Telegram Bot for Rome Guide

This is a Telegram bot designed to help users get a free guide for exploring Rome in two days. The bot checks if the user is subscribed to a specific Telegram channel and provides a download link for the guide.

## Features

- **Start Command**: Welcomes the user and provides instructions.
- **Subscription Check**: Verifies if the user is subscribed to the channel.
- **Inline Button**: Provides a download link for the guide via an inline button.

## Installation

### Prerequisites

- Go 1.20 or higher
- A Telegram bot token (obtained from [BotFather](https://core.telegram.org/bots#botfather))
- A Telegram channel (for subscription check)
- A Google Drive link (or any other link) for the guide

### Steps

1. **Clone the repository**:
```bash
git clone https://github.com/NikitaRadzkov/tour-guide-bot.git
cd tour-guide-bot
```

2. **Set up environment variables**:
Create a `.env` file in the root of the project and add the following variables:

```bash
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
TELEGRAM_CHANNEL_NAME=@your_channel_name
GUIDE_URL=https://your_google_drive_link
```

3. **Install dependencies**:
Run the following command to download the required Go packages:
```bash
go mod tidy
```

4. **Run the bot**:
Start the bot using the following command:

```bash
go run cmd/bot/main.go
```


5. **Interact with the bot**:
- Open Telegram and search for your bot.
- Send the /start command to begin.
- Follow the instructions provided by the bot.

### Build
1. Build binary project on current platform
```bash
go build -o my-bot cmd/bot/main.go
```

2. Start builded project

```bash
./my-bot
```

### Project Structure
```
/tour-guide-bot
├── cmd
│   └── bot
│       └── main.go
├── internal
│   ├── bot
│   │   └── bot.go
│   ├── config
│   │   └── config.go
│   └── commands
│       └── commands.go
├── template.env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

Dependencies
- [go-telegram-bot-api](https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api/v5): A Go library for interacting with the Telegram Bot API.
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv): A Go library for loading environment variables from a .env file.
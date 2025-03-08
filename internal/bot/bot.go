package bot

import (
	"fmt"
	"log"
	"tour-guide-bot/internal/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	channelName string
	guideUrl string
}

func NewBot(token, channelName, guideUrl string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		bot: bot,
		channelName: channelName,
		guideUrl: guideUrl,
	}, nil
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	log.Println("Bot is running...")

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case commands.Start:
				b.handleStart(update.Message)
			case commands.Confirm:
				b.handleConfirmation(update.Message)
			default:
				b.handleUnknownCommand(update.Message)
			}
		}
	}
}

func (b *Bot) handleStart(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! 👋\nЭтот бот поможет вам получить мои лучшие путеводители и полезные материалы. Нажмите кнопку «Начать», чтобы запустить! ⬇️")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Начать"),
		),
	)
	b.bot.Send(msg)
}

func (b *Bot) handleConfirmation(message *tgbotapi.Message) {
	isSubscribed, err := b.isUserSubscribed(message.Chat.ID)
	if err != nil {
		log.Printf("Failed to check subscribe: %v", err)
		return
	}

	if isSubscribed {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Отлично! 🎉\nВы подписаны, и теперь гайд ваш! Нажмите на кнопку ниже, чтобы скачать его и начать планировать идеальное путешествие. ✈️")

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("Скачать гайд", b.guideUrl),
			),
		)

		msg.ReplyMarkup = inlineKeyboard

		if _, err := b.bot.Send(msg); err != nil {
			log.Printf("Failed to send message: %v\n", err)
		}
} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Пожалуйста, подпишитесь на канал, чтобы получить гайд.")
		b.bot.Send(msg)
	}
}


func (b *Bot) isUserSubscribed(userID int64) (bool, error) {
	chatID , err := b.getChatIDByUsername(b.channelName)
	if err != nil {
		return false, fmt.Errorf("failed to get ChatID: %v", err)
	}


	chatMember, err := b.bot.GetChatMember(tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID: chatID,
			UserID: userID,
		},
	})
	if err != nil {
		return false, err
	}

	return chatMember.IsMember || chatMember.IsAdministrator() || chatMember.IsCreator(), nil
}

func (b *Bot) getChatIDByUsername(username string) (int64, error) {
	chat, err := b.bot.GetChat(tgbotapi.ChatInfoConfig{
		ChatConfig: tgbotapi.ChatConfig{
			SuperGroupUsername: username,
		},
	})
	if err != nil {
		return 0, err
	}
	return chat.ID, nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда. Пожалуйста, используйте /start для начала.")
	b.bot.Send(msg)
}

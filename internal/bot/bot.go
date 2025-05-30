package bot

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	channelName string
	guideUrl    string
	topDealsUrl string
	checklistUrl string
	searchUrl   string
	aboutUrl    string
	contactUser string
}

func NewBot(token, channelName, guideUrl, topDealsUrl, checklistUrl, searchUrl, aboutUrl, contactUser string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		bot:         bot,
		channelName: channelName,
		guideUrl:    guideUrl,
		topDealsUrl: topDealsUrl,
		checklistUrl: checklistUrl,
		searchUrl:   searchUrl,
		aboutUrl:    aboutUrl,
		contactUser: contactUser,
	}, nil
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	log.Println("Bot is running...")

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					b.handleStart(update.Message)
				default:
					b.handleUnknownCommand(update.Message)
				}
			} else {
				b.handleTextMessage(update.Message)
			}
		} else if update.CallbackQuery != nil {
			b.handleCallbackQuery(update.CallbackQuery)
		}
	}
}

func (b *Bot) handleStart(message *tgbotapi.Message) {
	text := `Привет! Это многофункциональный бот Веры Агеенковой ✨

Организую ваш отдых «как на облачке» в лучшем соотношение «цена-качество».
Здесь ты можешь забрать бесплатный 🎁, посмотреть 🔝 выгодных предложений, подобрать тур или связаться со мной лично.

Выбирай ниже 👇`

	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🎁 Забрать подарок"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("📋 Чек-лист на подбор тура"),
			tgbotapi.NewKeyboardButton("🔎 Поиск тура"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("ℹ️ Обо мне"),
			tgbotapi.NewKeyboardButton("💬 Связаться со мной"),
		),
	)

	keyboard.OneTimeKeyboard = false
	keyboard.ResizeKeyboard = true

	msg.ReplyMarkup = keyboard
	b.bot.Send(msg)
}

func (b *Bot) handleTextMessage(message *tgbotapi.Message) {
	switch message.Text {
	case "🎁 Забрать подарок":
		b.handleGift(message)
	case "📋 Чек-лист на подбор тура":
		b.handleChecklist(message)
	case "🔎 Поиск тура":
		b.handleSearch(message)
	case "ℹ️ Обо мне":
		b.handleAbout(message)
	case "💬 Связаться со мной":
		b.handleContact(message)
	default:
		b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleCallbackQuery(callbackQuery *tgbotapi.CallbackQuery) {
	switch callbackQuery.Data {
	case "check_subscription":
		b.handleSubscriptionCheck(callbackQuery)
	default:
		log.Printf("Unknown callback data: %s\n", callbackQuery.Data)
	}

	callback := tgbotapi.NewCallback(callbackQuery.ID, "")
	if _, err := b.bot.Request(callback); err != nil {
		log.Printf("Failed to answer callback: %v\n", err)
	}
}

func (b *Bot) handleGift(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	isSubscribed, err := b.isUserSubscribed(message.From.ID)
	if err != nil {
		log.Printf("Failed to check subscription: %v", err)
		msg.Text = "Произошла ошибка при проверке подписки. Пожалуйста, попробуйте позже."
		b.bot.Send(msg)
		return
	}

	if isSubscribed {
		msg.Text = "🎉 Отлично! Вот твой подарок."
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("📥 Скачать гайд", b.guideUrl),
			),
		)
		msg.ReplyMarkup = keyboard
	} else {
		msg.Text = "Увы, я не вижу твою подписку."
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("🔄 Подписаться", fmt.Sprintf("https://t.me/%s", strings.TrimPrefix(b.channelName, "@"))),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("✅ Я подписался", "check_subscription"),
			),
		)
		msg.ReplyMarkup = keyboard
	}

	b.bot.Send(msg)
}

func (b *Bot) handleTopDeals(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	isSubscribed, err := b.isUserSubscribed(message.From.ID)
	if err != nil {
		log.Printf("Failed to check subscription: %v", err)
		msg.Text = "Произошла ошибка при проверке подписки. Пожалуйста, попробуйте позже."
		b.bot.Send(msg)
		return
	}

	if isSubscribed {
		msg.Text = "🔝 выгодных предложений этой недели, подобранных лично мной:"
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("👉 Посмотреть предложения", b.topDealsUrl),
			),
		)
		msg.ReplyMarkup = keyboard
	} else {
		msg.Text = "Увы, я не вижу твою подписку."
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("🔄 Подписаться", fmt.Sprintf("https://t.me/%s", strings.TrimPrefix(b.channelName, "@"))),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("✅ Я подписался", "check_subscription"),
			),
		)
		msg.ReplyMarkup = keyboard
	}

	b.bot.Send(msg)
}

func (b *Bot) handleChecklist(message *tgbotapi.Message) {
	text := `Если вы цените индивидуальный подход и внимание к деталям, для вас важны качество и высокий уровень сервиса

Заполни чек-лист на подбор тура — и я предложу варианты, подходящие именно вам.`

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("✅ Заполнить", b.checklistUrl),
		),
	)
	msg.ReplyMarkup = keyboard
	b.bot.Send(msg)
}

func (b *Bot) handleSearch(message *tgbotapi.Message) {
	text := `🏖 Если вы любите выбирать самостоятельно или уже знаете отель — воспользуйтесь формой поиска тура на сайте:

1. Выбирай тур
2. Напиши мне детали и стоимость
3. Мы актуализируем цену и наличие мест и забронируем отдых`

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🔎 Искать тур", b.searchUrl),
		),
	)
	msg.ReplyMarkup = keyboard
	b.bot.Send(msg)
}

func (b *Bot) handleAbout(message *tgbotapi.Message) {
	text := `👩‍💼 Меня зовут Вера Агеенкова.
Я турагент, эксперт с 12-летним опытом, та, кто превращает ваши мечты в бронирования и делает «конфетку» из ваших «хочу».

🌍Подробнее обо мне и что я делаю для вас — на сайте.`

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Узнать", b.aboutUrl),
		),
	)
	msg.ReplyMarkup = keyboard
	b.bot.Send(msg)
}

func (b *Bot) handleContact(message *tgbotapi.Message) {
	text := `📩 Есть вопросы? Пиши прямо мне в Telegram — я на связи!`

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("👉 Написать", fmt.Sprintf("https://t.me/%s", strings.TrimPrefix(b.contactUser, "@"))),
		),
	)
	msg.ReplyMarkup = keyboard
	b.bot.Send(msg)
}

func (b *Bot) handleSubscriptionCheck(callbackQuery *tgbotapi.CallbackQuery) {
	isSubscribed, err := b.isUserSubscribed(callbackQuery.From.ID)
	if err != nil {
		log.Printf("Failed to check subscription: %v", err)
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Произошла ошибка при проверке подписки. Пожалуйста, попробуйте позже.")
		b.bot.Send(msg)
		return
	}

	if isSubscribed {
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🎉 Отлично! Теперь ты подписан, вот твой подарок!")
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("📥 Скачать гайд", b.guideUrl),
			),
		)
		msg.ReplyMarkup = keyboard
		b.bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Кажется, ты еще не подписался. Пожалуйста, подпишись на канал и попробуй снова.")
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("🔄 Подписаться", fmt.Sprintf("https://t.me/%s", strings.TrimPrefix(b.channelName, "@"))),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("✅ Я подписался", "check_subscription"),
			),
		)
		msg.ReplyMarkup = keyboard
		b.bot.Send(msg)
	}
}

func (b *Bot) isUserSubscribed(userID int64) (bool, error) {
	chatID, err := b.getChatIDByUsername(b.channelName)
	if err != nil {
		return false, fmt.Errorf("failed to get chat ID: %v", err)
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

	return chatMember.Status == "member" ||
	       chatMember.Status == "administrator" ||
	       chatMember.Status == "creator", nil
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
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда. Пожалуйста, используйте кнопки меню для навигации.")
	b.bot.Send(msg)
}

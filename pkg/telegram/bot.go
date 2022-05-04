package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

const answer = "Я не знаю такой команды, введите /help"

type Message struct {
	ChatId     int64
	Text       string
	ReceiveAt  time.Time
	ResponseAt time.Time

	From struct {
		FirstName    string
		LastName     string
		UserName     string
		LanguageCode string
	}
	Location *struct {
		Longitude float64
		Latitude  float64
	}
}

type Bot struct {
	Bot *tgbotapi.BotAPI

	handleMessageFunction  func(msg Message) string
	handleCommandFunctions map[string]func() string
}

func NewBot(apiToken string) *Bot {
	tg, e := tgbotapi.NewBotAPI(apiToken)
	if e != nil {
		log.Panic(e)
	}

	bot := &Bot{
		Bot: tg,
		handleMessageFunction: func(msg Message) string {
			return msg.Text
		},
		handleCommandFunctions: map[string]func() string{},
	}
	bot.Bot.Debug = true

	log.Printf("Authorized on account %s", bot.Bot.Self.UserName)

	return bot
}

func (b *Bot) StartListening() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.Bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				b.handleCommand(update.Message)
				continue
			}
			msg := update.Message
			_, err := b.Bot.Send(tgbotapi.NewMessage(msg.Chat.ID, b.handleMessageFunction(Message{
				ChatId: msg.Chat.ID,
				From: struct {
					FirstName    string
					LastName     string
					UserName     string
					LanguageCode string
				}{
					FirstName:    msg.From.FirstName,
					LastName:     msg.From.LastName,
					UserName:     msg.From.UserName,
					LanguageCode: msg.From.LanguageCode,
				},
				Location: &struct {
					Longitude float64
					Latitude  float64
				}{
					Longitude: msg.Location.Longitude,
					Latitude:  msg.Location.Latitude,
				},
				Text:       msg.Text,
				ReceiveAt:  msg.Time(),
				ResponseAt: time.Time{},
			})))
			if err != nil {
				panic(err)
			}
		}
	}
}

func (b *Bot) RegisterMessageHandler(fn func(msg Message) string) {
	b.handleMessageFunction = fn
}

func (b *Bot) RegisterCommand(cmd string, fn func() string) {
	b.handleCommandFunctions[cmd] = fn
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, answer)
	if fn, ok := b.handleCommandFunctions[message.Command()]; ok {
		msg = tgbotapi.NewMessage(message.Chat.ID, fn())
	}

	if _, e := b.Bot.Send(msg); e != nil {
		log.Print("Error while sending message", e)
	}
}

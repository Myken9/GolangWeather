package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	api *tgbotapi.BotAPI

	handleMessageFunction  func(msg string) string
	handleCommandFunctions map[string]func() string
}

func NewBot(apiToken string) *Bot {
	tg, e := tgbotapi.NewBotAPI(apiToken)
	if e != nil {
		log.Panic(e)
	}

	bot := &Bot{
		api: tg,
		handleMessageFunction: func(msg string) string {
			return msg
		},
		handleCommandFunctions: map[string]func() string{},
	}
	bot.api.Debug = true

	log.Printf("Authorized on account %s", bot.api.Self.UserName)

	return bot
}

func (b *Bot) StartListening() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				b.handleCommand(update.Message)
				continue
			}
			msg := update.Message
			//from := msg.From.UserName
			b.api.Send(tgbotapi.NewMessage(msg.Chat.ID, b.handleMessageFunction(msg.Text)))
		}
	}
}

func (b *Bot) RegisterMessageHandler(fn func(msg string) string) {
	b.handleMessageFunction = fn
}

func (b *Bot) RegisterCommand(cmd string, fn func() string) {
	b.handleCommandFunctions[cmd] = fn
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")
	if fn, ok := b.handleCommandFunctions[message.Command()]; ok {
		msg = tgbotapi.NewMessage(message.Chat.ID, fn())
	}

	if _, e := b.api.Send(msg); e != nil {
		log.Print("Error while sending message", e)
	}
}

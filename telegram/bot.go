package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const answer = "Я не знаю такой команды, введите /help"

type Bot struct {
	Bot *tgbotapi.BotAPI

	HandleMessageFunction  func(msg tgbotapi.Message) string
	HandleCommandFunctions map[string]func() string
}

func NewBot(apiToken string) *Bot {
	tg, e := tgbotapi.NewBotAPI(apiToken)
	if e != nil {
		log.Panic(e)
	}

	bot := &Bot{
		Bot: tg,
		HandleMessageFunction: func(msg tgbotapi.Message) string {
			return msg.Text
		},
		HandleCommandFunctions: map[string]func() string{},
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
			b.Bot.Send(tgbotapi.NewMessage(msg.Chat.ID, b.HandleMessageFunction(*msg)))
		}
	}
}

func (b *Bot) RegisterMessageHandler(fn func(msg tgbotapi.Message) string) {
	b.HandleMessageFunction = fn
}

func (b *Bot) RegisterCommand(cmd string, fn func() string) {
	b.HandleCommandFunctions[cmd] = fn
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, answer)
	if fn, ok := b.HandleCommandFunctions[message.Command()]; ok {
		msg = tgbotapi.NewMessage(message.Chat.ID, fn())
	}

	if _, e := b.Bot.Send(msg); e != nil {
		log.Print("Error while sending message", e)
	}
}

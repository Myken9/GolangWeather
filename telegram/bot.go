package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (w *Bot) Start(weather *Weather) error {
	log.Printf("Authorized on account %s", w.bot.Self.UserName)

	updates := w.initUpdatesChannel()

	w.handleUpdates(updates, weather)

	return nil
}

func (w *Bot) handleUpdates(updates tgbotapi.UpdatesChannel, weather *Weather) {
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				w.handleCommand(update.Message)
				continue
			}

			w.handleMessage(update.Message, weather)
		}
	}
}

func (w *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return w.bot.GetUpdatesChan(u)
}

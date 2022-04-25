package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const commandStart = "start"

func (w *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")

	switch message.Command() {
	case commandStart:
		msg.Text = "Я Sebastian - бот погоды, приятно познакомитсья! Напишите название города, в котором хотите узнать погоду."
		_, err := w.bot.Send(msg)
		if err != nil {
			return
		}
	default:
		_, err := w.bot.Send(msg)
		if err != nil {
			return
		}
	}
}

func (w *Bot) handleMessage(message *tgbotapi.Message, weather *Weather) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, w.tellWeather(message, weather))
	_, err := w.bot.Send(msg)
	if err != nil {
		return
	}
}

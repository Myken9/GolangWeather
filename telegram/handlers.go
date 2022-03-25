package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")

	switch message.Command() {
	case commandStart:
		msg.Text = "Я Sebastian - бот погоды, приятно познакомитсья! Напишите название города, в котором хотите узнать погоду."
		_, err := b.bot.Send(msg)
		if err != nil {
			return
		}
	default:
		_, err := b.bot.Send(msg)
		if err != nil {
			return
		}
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, TellWeather(message))
	_, err := b.bot.Send(msg)
	if err != nil {
		return
	}
}

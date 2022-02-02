package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")

	switch message.Command() {
	case commandStart:
		msg.Text = "Я Sebastian - бот погоды, приятно познакомитсья! Напишите название города, в котором хотите узнать погоду."
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}

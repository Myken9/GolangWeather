package telegram

import (
	"GolangWeather/weather"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
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

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, weather.TellWeather(message.Text))
	b.bot.Send(msg)
}
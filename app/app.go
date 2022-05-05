package app

import (
	"GolangWeather/pkg/telegram"
)

type Application struct {
	bot     *telegram.Bot
	handler *handler
}

func NewApplication(bot *telegram.Bot, handler *handler) *Application {
	return &Application{
		bot:     bot,
		handler: handler,
	}
}

func (a *Application) Run() {
	a.bot.RegisterCommand("start", func() string {
		answer := "Я Sebastian - бот погоды, приятно познакомитсья!\n\n" +
			"Отправь /help для помощи."
		return answer
	})
	a.bot.RegisterCommand("help", func() string {
		answer := "Напишите название города, в котором хотите узнать погоду, или отправьте мне свою текущую геопозицию."
		return answer
	})
	a.bot.RegisterMessageHandler(a.handler.handleTelegramMessage)

	a.bot.StartListening()
}

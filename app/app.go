package app

import (
	"GolangWeather/pkg/telegram"
)

type Application struct {
	*telegram.Bot
	*Weather
}

func NewApplication(bot *telegram.Bot, weather *Weather) *Application {
	return &Application{
		bot,
		weather,
	}
}

func (a *Application) Run() {
	a.RegisterCommand("start", func() string {
		return "Я Sebastian - бот погоды, приятно познакомитсья!\n\n" +
			"Отправь /help для помощи."
	})
	a.RegisterCommand("help", func() string {
		return "Напишите название города, в котором хотите узнать погоду, или отправьте мне свою текущую геопозицию."
	})
	a.RegisterMessageHandler(a.handleTelegramMessage)

	a.StartListening()
}

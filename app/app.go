package app

import (
	"GolangWeather/pkg/telegram"
)

type Application struct {
	*telegram.Bot
	*app
}

func NewApplication(bot *telegram.Bot, weather *app) *Application {
	return &Application{
		bot,
		weather,
	}
}

func (a *Application) Run() {
	a.RegisterCommand("start", func() string {
		answer := "Я Sebastian - бот погоды, приятно познакомитсья!\n\n" +
			"Отправь /help для помощи."
		return answer
	})
	a.RegisterCommand("help", func() string {
		answer := "Напишите название города, в котором хотите узнать погоду, или отправьте мне свою текущую геопозицию."
		return answer
	})
	a.RegisterMessageHandler(a.handleTelegramMessage)

	a.StartListening()
}

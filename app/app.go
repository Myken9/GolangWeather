package app

import (
	"GolangWeather/pkg/storage"
	"GolangWeather/pkg/telegram"
	"github.com/briandowns/openweathermap"
	"github.com/jackc/pgx/v4"
)

type Application struct {
	bot     *telegram.Bot
	handler *handler
}

func NewApplication(bot *telegram.Bot, weather *openweathermap.CurrentWeatherData, conn *pgx.Conn) *Application {
	repository := storage.NewStorage(conn)
	handler := NewHandler(weather, repository)
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

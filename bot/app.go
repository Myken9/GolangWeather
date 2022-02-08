package bot

import "GolangWeather/telegram"

type Application struct {
	bot *telegram.Bot
}

func NewApplication(bot *telegram.Bot) *Application {
	return &Application{bot: bot}
}

func (a *Application) Run() {
	a.bot.RegisterCommand("start", func() string {
		return "Я Sebastian - бот погоды, приятно познакомитсья!\n\n" +
			"Отправь /help для помощи."
	})
	a.bot.RegisterCommand("help", func() string {
		return "Напишите название города, в котором хотите узнать погоду."
	})
	a.bot.RegisterMessageHandler(a.handleTelegramMessage)

	a.bot.StartListening()
}

func (a *Application) handleTelegramMessage(msg string) string {
	return "output message " + msg
}

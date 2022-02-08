package bot

import (
	"GolangWeather/telegram"
	"fmt"
	"github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const needWeatherFor3HoursInterval = 8

type Application struct {
	bot     *telegram.Bot
	weather *openweathermap.ForecastWeatherData
}

func NewApplication(bot *telegram.Bot, weather *openweathermap.ForecastWeatherData) *Application {
	return &Application{
		bot:     bot,
		weather: weather,
	}
}

func (a *Application) Run() {
	a.bot.RegisterCommand("start", func() string {
		return "Я Sebastian - бот погоды, приятно познакомитсья!\n\n" +
			"Отправь /help для помощи."
	})
	a.bot.RegisterCommand("help", func() string {
		return "Напишите название города, в котором хотите узнать погоду, или отправьте мне свою текущую геопозицию."
	})
	a.bot.RegisterMessageHandler(a.handleTelegramMessage)

	a.bot.StartListening()
}

// todo don't delete this line it's reminder for me: lock in weather library
func (a *Application) handleTelegramMessage(msg tgbotapi.Message) string {
	var e error
	if msg.Location != nil {
		e = a.weather.DailyByCoordinates(&openweathermap.Coordinates{
			Longitude: msg.Location.Longitude,
			Latitude:  msg.Location.Latitude,
		}, needWeatherFor3HoursInterval)
	} else {
		e = a.weather.DailyByName(msg.Text, needWeatherFor3HoursInterval)
	}
	if e != nil {
		log.Print(e)
	}

	data, ok := a.weather.ForecastWeatherJson.(*openweathermap.Forecast5WeatherData)
	if !ok {
		log.Fatal("Invalid forecast type")
	}

	responseString := "Погода в " + data.City.Name + ", " + data.City.Country + ":\n"
	for _, item := range data.List {
		responseString = responseString + "\n" +
			item.DtTxt.Format("2006-01-02 15:04") +
			" - температура воздуха " + fmt.Sprintf("%.1f", item.Main.Temp) + "℃"
	}

	return responseString
}

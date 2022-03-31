package telegram

import (
	"github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"

	owm "github.com/briandowns/openweathermap"
)

type Weather struct {
	w *openweathermap.CurrentWeatherData
}

func NewWeather(w *openweathermap.CurrentWeatherData) *Weather {
	return &Weather{w: w}
}

func (w *Bot) tellWeather(message *tgbotapi.Message, weather *Weather) (text string) {
	if message.Text != "" {
		return weatherByCity(message.Text, weather)
	} else if message.Location.Latitude != 0 && message.Location.Longitude != 0 {
		return weatherByGeopos(message.Location.Latitude, message.Location.Longitude, weather.w)
	} else {
		log.Print("Location is not found")
		return "Попробуйте чуть позже"
	}
}

func weatherByCity(messageText string, weather *Weather) (text string) {

	if err := weather.w.CurrentByName(messageText); err != nil {
		return "Я не знаю такого города."
	}
	text = "Погода в г." + weather.w.Name + ": " + strconv.Itoa(int(weather.w.Main.Temp)) + "C"
	return text
}

func weatherByGeopos(messageLocationLatitude, messageLocationLongitude float64, w *openweathermap.CurrentWeatherData) (text string) {

	err := w.CurrentByCoordinates(
		&owm.Coordinates{
			Longitude: messageLocationLongitude,
			Latitude:  messageLocationLatitude,
		},
	)
	if err != nil {
		return "Где это вы находитесь?"
	}
	text = "Температура в вашем месте расположения : " + strconv.Itoa(int(w.Main.Temp)) + "C"
	return text
}

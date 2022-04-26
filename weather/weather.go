package weather

import (
	"github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

type Weather struct {
	W *openweathermap.CurrentWeatherData
}

func NewWeather(W *openweathermap.CurrentWeatherData) *Weather {
	return &Weather{W: W}
}

func (w *Weather) HandleTelegramMessage(msg tgbotapi.Message) string {
	var e error
	if msg.Location != nil {
		e = w.W.CurrentByCoordinates(&openweathermap.Coordinates{
			Longitude: msg.Location.Longitude,
			Latitude:  msg.Location.Latitude,
		})
	} else {
		if e = w.W.CurrentByName(msg.Text); e != nil {
			return "Я не знаю такого города."
		}
	}
	if e != nil {
		log.Print(e)
	}

	switch {
	case msg.Text == "":
		return "Температура в вашем месте расположения : " + strconv.Itoa(int(w.W.Main.Temp)) + "C"

	default:
		return "Погода в г." + msg.Text + ": " + strconv.Itoa(int(w.W.Main.Temp)) + "C"

	}

}

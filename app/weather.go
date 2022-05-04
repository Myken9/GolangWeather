package app

import (
	"GolangWeather/storage"
	"fmt"
	"github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

type weather struct {
	*openweathermap.CurrentWeatherData
	*storage.Storage
}

func NewWeather(W *openweathermap.CurrentWeatherData, S *storage.Storage) *weather {
	return &weather{W, S}
}

func (w *weather) handleTelegramMessage(msg tgbotapi.Message) string {
	var e error
	if msg.Location != nil {
		e = w.CurrentByCoordinates(&openweathermap.Coordinates{
			Longitude: msg.Location.Longitude,
			Latitude:  msg.Location.Latitude,
		})
	} else {
		if e = w.CurrentByName(msg.Text); e != nil {
			return "Я не знаю такого города."
		}
	}
	if e != nil {
		log.Print(e)
	}

	switch {
	case msg.Text == "":
		answer := "Температура в вашем месте расположения : " + strconv.Itoa(int(w.Main.Temp)) + "C"
		err := w.SaveUserLocation(msg, answer)
		if err != nil {
			fmt.Println(err)
		}
		return answer

	default:
		answer := "Погода в г." + msg.Text + ": " + strconv.Itoa(int(w.Main.Temp)) + "C"
		err := w.SaveUserMessage(msg, answer)
		if err != nil {
			fmt.Println(err)
		}
		return answer

	}

}

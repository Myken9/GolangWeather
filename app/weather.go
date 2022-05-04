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
	w *openweathermap.CurrentWeatherData
	s *storage.Storage
}

func NewWeather(w *openweathermap.CurrentWeatherData, s *storage.Storage) *weather {
	return &weather{w: w, s: s}
}

func (w *weather) handleTelegramMessage(msg tgbotapi.Message) string {
	var e error
	if msg.Location != nil {
		e = w.w.CurrentByCoordinates(&openweathermap.Coordinates{
			Longitude: msg.Location.Longitude,
			Latitude:  msg.Location.Latitude,
		})
	} else {
		if e = w.w.CurrentByName(msg.Text); e != nil {
			return "Я не знаю такого города."
		}
	}
	if e != nil {
		log.Print(e)
	}

	switch {
	case msg.Text == "":
		answer := "Температура в вашем месте расположения : " + strconv.Itoa(int(w.w.Main.Temp)) + "C"
		err := w.s.SaveUserLocation(msg, answer)
		if err != nil {
			fmt.Println(err)
		}
		return answer

	default:
		answer := "Погода в г." + msg.Text + ": " + strconv.Itoa(int(w.w.Main.Temp)) + "C"
		err := w.s.SaveUserMessage(msg, answer)
		if err != nil {
			fmt.Println(err)
		}
		return answer

	}

}

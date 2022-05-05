package app

import (
	"GolangWeather/pkg/storage"
	"GolangWeather/pkg/telegram"
	"github.com/briandowns/openweathermap"
	"log"
	"strconv"
	"time"
)

type handler struct {
	w *openweathermap.CurrentWeatherData
	s *storage.Storage
}

func NewHandler(w *openweathermap.CurrentWeatherData, s *storage.Storage) *handler {
	return &handler{w: w, s: s}
}

func (h *handler) handleTelegramMessage(msg telegram.Message) (answer string) {
	msg.ResponseAt = int(time.Now().Unix())
	if msg.Location != nil {
		if e := h.w.CurrentByCoordinates(&openweathermap.Coordinates{
			Longitude: msg.Location.Longitude,
			Latitude:  msg.Location.Latitude,
		}); e != nil {
			return e.Error()
		}
		answer = "Температура в вашем месте расположения : " + strconv.Itoa(int(h.w.Main.Temp)) + "C"
		if e := h.s.SaveUserLocation(msg, answer); e != nil {
			log.Fatal(e.Error())
		}
	} else {
		if e := h.w.CurrentByName(msg.Text); e != nil {
			return "Я не знаю такого города."
		}
		answer = "Погода в г." + msg.Text + ": " + strconv.Itoa(int(h.w.Main.Temp)) + "C"
		if e := h.s.SaveUserMessage(msg, answer); e != nil {
			log.Fatal(e.Error())
		}
	}

	return answer
}

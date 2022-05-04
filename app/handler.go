package app

import (
	"GolangWeather/pkg/storage"
	"GolangWeather/pkg/telegram"
	"github.com/briandowns/openweathermap"
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

func (h *handler) handleTelegramMessage(msg telegram.Message) string {
	var answer string
	if msg.Location != nil {
		if e := h.w.CurrentByCoordinates(&openweathermap.Coordinates{
			Longitude: msg.Location.Longitude,
			Latitude:  msg.Location.Latitude,
		}); e != nil {
			return e.Error()
		}
		answer = "Температура в вашем месте расположения : " + strconv.Itoa(int(h.w.Main.Temp)) + "C"
	} else {
		if e := h.w.CurrentByName(msg.Text); e != nil {
			return e.Error()
		}
		answer = "Погода в г." + msg.Text + ": " + strconv.Itoa(int(h.w.Main.Temp)) + "C"
	}
	msg.ResponseAt = time.Now()
	if e := h.s.SaveUserMessage(msg, answer); e != nil {
		return e.Error()
	}

	return answer
}

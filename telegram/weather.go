package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"

	owm "github.com/briandowns/openweathermap"
)

func (b *Bot) tellWeather(message *tgbotapi.Message) (text string) {
	if message.Text != "" {
		return b.weatherByCity(message.Text)
	} else if message.Location.Latitude != 0 && message.Location.Longitude != 0 {
		return b.weatherByGeopos(message.Location.Latitude, message.Location.Longitude)
	} else {
		log.Print("Location is not found")
		return "Попробуйте чуть позже"
	}
}

func (b *Bot) weatherByCity(messageText string) (text string) {

	w, err := owm.NewCurrent("C", "ru", b.tokenWeather)
	if err != nil {
		log.Fatalln(err)
	}

	if err = w.CurrentByName(messageText); err != nil {
		return "Я не знаю такого города."
	}
	text = "Погода в г." + w.Name + ": " + strconv.Itoa(int(w.Main.Temp)) + "C"
	return text
}

func (b *Bot) weatherByGeopos(messageLocationLatitude, messageLocationLongitude float64) (text string) {

	w, err := owm.NewCurrent("C", "ru", b.tokenWeather)
	if err != nil {
		log.Fatalln(err)
	}

	err = w.CurrentByCoordinates(
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

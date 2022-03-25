package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"

	owm "github.com/briandowns/openweathermap"
)

func TellWeather(message *tgbotapi.Message) (text string) {
	if message.Text != "" {
		return weatherByCity(message.Text)
	} else if message.Location.Latitude != 0 {
		return weatherByGeopos(message.Location.Latitude, message.Location.Longitude)
	} else {
		log.Print("Location is not found")
		return "Попробуйте чуть позже"
	}
}

func weatherByCity(messageText string) (text string) {

	w, err := owm.NewCurrent("C", "ru", tokenWeather)
	if err != nil {
		log.Fatalln(err)
	}

	if err = w.CurrentByName(messageText); err != nil {
		return "Я не знаю такого города."
	}
	text = "Погода в г." + w.Name + ": " + strconv.Itoa(int(w.Main.Temp)) + "C"
	return text
}

func weatherByGeopos(messageLocationLatitude, messageLocationLongitude float64) (text string) {

	w, err := owm.NewCurrent("C", "ru", tokenWeather)
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

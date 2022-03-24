package weather

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"

	owm "github.com/briandowns/openweathermap"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func TellWeather(message *tgbotapi.Message) (text string) {
	if message.Text != "" {
		return weatherByCity(message.Text)
	} else if message.Location.Latitude != 0 {
		return weatherByGeopos(message.Location.Latitude, message.Location.Longitude)
	} else {
		return "Я не понимаю чего ты хочешь!"
	}
}

func weatherByCity(messageText string) (text string) {
	token, _ := os.LookupEnv("WEATHER_API_KEY")

	w, err := owm.NewCurrent("C", "ru", token)
	if err != nil {
		log.Fatalln(err)
	}

	if err = w.CurrentByName(messageText); err != nil {
		log.Fatalln(err)
	}
	text = "Погода в г." + w.Name + ": " + strconv.Itoa(int(w.Main.Temp)) + "C"
	return text
}

func weatherByGeopos(messageLocationLatitude, messageLocationLongitude float64) (text string) {
	token, _ := os.LookupEnv("WEATHER_API_KEY")

	w, err := owm.NewCurrent("C", "ru", token)
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByCoordinates(
		&owm.Coordinates{
			Longitude: messageLocationLongitude,
			Latitude:  messageLocationLatitude,
		},
	)
	text = "Температура в вашем месте расположения : " + strconv.Itoa(int(w.Main.Temp)) + "C"
	return text
}

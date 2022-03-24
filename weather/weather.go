package weather

import (
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

func TellWeather(messageText string) (text string) {
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

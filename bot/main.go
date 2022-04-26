package main

import (
	"GolangWeather/app"
	"GolangWeather/telegram"
	"GolangWeather/weather"
	owm "github.com/briandowns/openweathermap"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	botWeather, err := owm.NewCurrent("C", "ru", os.Getenv("WEATHER_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}

	weather := weather.NewWeather(botWeather)

	telegramBot := telegram.NewBot(os.Getenv("TOKEN"))

	application := app.NewApplication(telegramBot, weather)
	application.Run()
}

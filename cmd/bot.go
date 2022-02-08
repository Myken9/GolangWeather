package main

import (
	"GolangWeather/bot"
	"GolangWeather/telegram"
	owm "github.com/briandowns/openweathermap"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	telegramBot := telegram.NewBot(os.Getenv("TOKEN"))

	openWeatherMap, err := owm.NewForecast("5", "C", "ru", os.Getenv("WEATHER_API_KEY"))
	if err != nil {
		log.Panic(err)
	}

	application := bot.NewApplication(telegramBot, openWeatherMap)
	application.Run()
}

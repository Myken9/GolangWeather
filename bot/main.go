package main

import (
	"GolangWeather/telegram"
	owm "github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	token := os.Getenv("TOKEN")
	tokenWeather := os.Getenv("WEATHER_API_KEY")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	botweather, err := owm.NewCurrent("C", "ru", tokenWeather)
	if err != nil {
		log.Fatalln(err)
	}

	weather := telegram.NewWeather(botweather)

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(weather); err != nil {
		log.Fatal(err)
	}
}

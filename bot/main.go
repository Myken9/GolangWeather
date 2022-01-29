package main

import (
	"GolangWeather/telegram"
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	token, exists := os.LookupEnv("TOKEN")
	if exists {
		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Panic(err)
		}

		bot.Debug = true

		telegramBot := telegram.NewBot(bot)
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}

}

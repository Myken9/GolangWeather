package main

import (
	"GolangWeather/bot"
	"GolangWeather/telegram"
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
	application := bot.NewApplication(telegramBot)
	application.Run()
}

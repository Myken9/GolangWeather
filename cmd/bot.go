package main

import (
	"GolangWeather/app"
	"GolangWeather/pkg/storage"
	"GolangWeather/pkg/telegram"
	"context"
	owm "github.com/briandowns/openweathermap"
	"github.com/jackc/pgx/v4"
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

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	botWeather, err := owm.NewCurrent("C", "ru", os.Getenv("WEATHER_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}

	weather := app.NewHandler(botWeather, storage.NewStorage(conn))

	telegramBot := telegram.NewBot(os.Getenv("TOKEN"))

	application := app.NewApplication(telegramBot, weather)
	application.Run()
}

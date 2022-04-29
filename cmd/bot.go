package main

import (
	"GolangWeather/app"
	"GolangWeather/pkg/telegram"
	"context"
	"fmt"
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
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	botWeather, err := owm.NewCurrent("C", "ru", os.Getenv("WEATHER_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}

	weather := app.NewWeather(botWeather)

	telegramBot := telegram.NewBot(os.Getenv("TOKEN"))

	application := app.NewApplication(telegramBot, weather)
	application.Run()
}

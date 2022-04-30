package main

import (
	"GolangWeather/app"
	"GolangWeather/pkg/telegram"
	"GolangWeather/storage"
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

	Conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	newStorage := storage.NewStorage(Conn)

	defer func(newStorage *storage.Storage, ctx context.Context) {
		err := newStorage.Close(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(newStorage, context.Background())

	botWeather, err := owm.NewCurrent("C", "ru", os.Getenv("WEATHER_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}

	weather := app.NewWeather(botWeather, newStorage)

	telegramBot := telegram.NewBot(os.Getenv("TOKEN"))

	application := app.NewApplication(telegramBot, weather)
	application.Run()
}
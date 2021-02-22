package main

import (
	"log"
	"os"

	"github.com/Vlad1slavZhuk/TelegramBot/pkg/telegram"
	"github.com/Vlad1slavZhuk/TelegramBot/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../TelegramBot/.env")
	if err != nil {
		log.Panic("Error:", err)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic("Error:", err)
	}

	list := storage.NewList(bot)

	bot.Debug = true

	telegramBot := telegram.NewBot(bot, list)

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

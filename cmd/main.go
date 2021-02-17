package main

import (
	"fmt"
	"log"
	"os"

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

	TelegramToken := os.Getenv("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(TelegramToken)
	if err != nil {
		log.Panic("Error:", err)
	}

	list := storage.NewList(bot)

	bot.Debug = true

	log.Printf("Authorized on account %s (http://t.me/%s)", bot.Self.UserName, bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal("Error:", err)
	}

	for update := range updates {
		// Callback
		if update.CallbackQuery != nil {
			log.Printf("[%s] %v", update.CallbackQuery.From.UserName, update.CallbackQuery.Message)
			switch update.CallbackQuery.Data {
			case "del":
				if e := list.Delete(update.CallbackQuery); e != nil {
					log.Fatal("Error:", e)
				}
			case "done":
				if e := list.Check(update.CallbackQuery); e != nil {
					log.Fatal("Error:", e)
				}
			}
		}

		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.FirstName, update.Message.Text)

		// Command
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				msg.Text = fmt.Sprintf("Hello %v", update.Message.From.FirstName)
				if _, e := bot.Send(msg); e != nil {
					log.Fatal("Error:", e)
				}

				// Get all products
				// for _, pr := range mem.GetAll() {
				// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, pr.Name)
				// 	msg.ReplyMarkup = inlineKeyboard
				// 	if _, e := bot.Send(msg); e != nil {
				// 		log.Fatal("Error:", e)
				// 	}
				// }
			}
		}

		if len(update.Message.Text) >= 2 && !update.Message.IsCommand() {
			if e := list.Append(update.Message); e != nil {
				log.Fatal("Error:", e)
			}
		}

	}

}

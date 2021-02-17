package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Vlad1slavZhuk/TelegramBot/model"
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

var mem = storage.NewMemory()

// var Keyboard = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("/add"),
// 		tgbotapi.NewKeyboardButton("/list"),
// 	),
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("/completed"),
// 		tgbotapi.NewKeyboardButton("/backlog"),
// 	),
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("/help"),
// 		tgbotapi.NewKeyboardButton("/exit"),
// 	),
// )
var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("✅", "done"),
		tgbotapi.NewInlineKeyboardButtonData("❌", "del"),
	),
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	TelegramToken := os.Getenv("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(TelegramToken)
	if err != nil {
		log.Panic("Error:", err)
	}

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
			case "del", "done":
				if err := mem.RemoveProduct(update.CallbackQuery.Message.Text); err != nil {
					log.Fatal("Error:", err)
				}
				msgDel := tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
				if _, e := bot.Send(msgDel); e != nil {
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
				msg.Text = fmt.Sprintf("Hello @%v", update.Message.From.UserName)
				if _, e := bot.Send(msg); e != nil {
					log.Fatal("Error:", e)
				}

				// Get all products
				for _, pr := range mem.GetAll() {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, pr.Name)
					msg.ReplyMarkup = inlineKeyboard
					if _, e := bot.Send(msg); e != nil {
						log.Fatal("Error:", e)
					}
				}
			}
		}

		if len(update.Message.Text) >= 2 && !update.Message.IsCommand() {
			name := update.Message.Text
			if err := mem.AddProduct(&model.Product{Name: name}); err != nil {
				log.Fatal("Error: ", err)
			}
			msgDel := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			if _, e := bot.Send(msgDel); e != nil {
				log.Fatal("Error:", e)
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, name)
			msg.ReplyMarkup = inlineKeyboard
			if _, e := bot.Send(msg); e != nil {
				log.Fatal("Error:", e)
			}
		}

	}

}

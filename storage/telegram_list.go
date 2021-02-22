package storage

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("✅", "done"),
		tgbotapi.NewInlineKeyboardButtonData("❌", "del"),
	),
)

type Commands interface {
	Append(msg tgbotapi.Message) error
	Delete(cbq *tgbotapi.CallbackQuery) error
	Check(cbq *tgbotapi.CallbackQuery) error
}

type List struct {
	Commands *Commands
	bot      *tgbotapi.BotAPI
}

func NewList(bot *tgbotapi.BotAPI) *List {
	return &List{
		Commands: new(Commands),
		bot:      bot,
	}
}

// Append new product
func (l *List) Append(msg *tgbotapi.Message) error {
	name := msg.Text

	msgDel := tgbotapi.NewDeleteMessage(msg.Chat.ID, msg.MessageID)
	if _, e := l.bot.Send(msgDel); e != nil {
		log.Println("Error:", e)
		return e
	}

	newMsg := tgbotapi.NewMessage(msg.Chat.ID, name)

	newMsg.ReplyMarkup = inlineKeyboard
	if _, e := l.bot.Send(newMsg); e != nil {
		log.Println("Error:", e)
		return e
	}

	return nil
}

// Delete product
func (l *List) Delete(cbq *tgbotapi.CallbackQuery) error {
	msgDel := tgbotapi.NewDeleteMessage(cbq.Message.Chat.ID, cbq.Message.MessageID)
	_, err := l.bot.Send(msgDel)
	return err
}

// Check product
func (l *List) Check(cbq *tgbotapi.CallbackQuery) error {
	msgDel := tgbotapi.NewDeleteMessage(cbq.Message.Chat.ID, cbq.Message.MessageID)
	_, err := l.bot.Send(msgDel)
	return err
}

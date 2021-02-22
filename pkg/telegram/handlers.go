package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart string = "start"
	commandHelp  string = "help"
)

const (
	callbackQueryDone   string = "done"
	callbackQueryDelete string = "del"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleCommandStart(message)
	case commandHelp:
		return b.handleCommandHelp(message)
	default:
		return b.handleCommandUnknown(message)
	}
}

func (b *Bot) handleCallbackQuery(cbq *tgbotapi.CallbackQuery) error {
	switch cbq.Data {
	case callbackQueryDelete:
		return b.handleCBQDel(cbq)
	case callbackQueryDone:
		return b.handleCBQDDone(cbq)
	default:
		return nil
	}
}

// -------------------------------- COMMANDS ---------------------------------------

func (b *Bot) handleCommandStart(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Hello %v", message.From.FirstName))
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommandHelp(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Help...")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommandUnknown(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Command is unknown.")
	_, err := b.bot.Send(msg)
	return err
}

// -------------------------------- COMMANDS ---------------------------------------

// -------------------------------- MESSAGE ---------------------------------------

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	return b.list.Append(message)
}

// -------------------------------- MESSAGE ---------------------------------------

// -------------------------------- CALLBACKQUERY ---------------------------------------

func (b *Bot) handleCBQDel(cbq *tgbotapi.CallbackQuery) error {
	return b.list.Delete(cbq)
}

func (b *Bot) handleCBQDDone(cbq *tgbotapi.CallbackQuery) error {
	return b.list.Check(cbq)
}

// -------------------------------- CALLBACKQUERY ---------------------------------------

package telegram

import (
	"log"

	"github.com/Vlad1slavZhuk/TelegramBot/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot  *tgbotapi.BotAPI
	list *storage.List
}

func NewBot(bot *tgbotapi.BotAPI, list *storage.List) *Bot {
	return &Bot{bot: bot, list: list}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s (http://t.me/%s)", b.bot.Self.UserName, b.bot.Self.UserName)
	updates, err := b.getUpdatesChannel()
	if err != nil {
		return err
	}

	if err = b.handleUpdates(updates); err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.CallbackQuery != nil {
			if err := b.handleCallbackQuery(update.CallbackQuery); err != nil {
				return err
			}
			continue
		}

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				return err
			}
			continue
		}

		if err := b.handleMessage(update.Message); err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) getUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

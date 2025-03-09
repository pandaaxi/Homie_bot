package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var BotAPI *tgbotapi.BotAPI

// StartBot initializes the Telegram bot and starts processing updates.
func StartBot(token string) {
	var err error
	BotAPI, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Failed to create Telegram bot: %v", err)
	}
	log.Printf("Authorized on account %s", BotAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := BotAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			go HandleUpdate(update)
		}
	}
}

package geeksonator

import (
	"log"

	"geeksonator/internal/observer"
	"geeksonator/internal/provider/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(tgBotToken string, tgTimeoutSeconds int, debugMode bool) {
	if debugMode {
		log.Print("Debug mode running")
	}

	botAPI, err := tgbotapi.NewBotAPI(tgBotToken)
	if err != nil {
		panic(err)
	}
	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	telegramService := telegram.NewService(botAPI)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = tgTimeoutSeconds // long polling

	updatesChan := botAPI.GetUpdatesChan(updateConfig)

	observer := observer.NewManager(telegramService, updatesChan, debugMode)
	go observer.Run()
	log.Print("Bot command watching started")
}

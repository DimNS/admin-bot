package observer

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotProvider interface {
	GetChatAdministrators(chatConfig tgbotapi.ChatConfig) ([]tgbotapi.ChatMember, error)
	NewMessage(chatID int64, text string) tgbotapi.MessageConfig
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

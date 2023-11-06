package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Service struct {
	bot BotAPI
}

func NewService(bot BotAPI) *Service {
	return &Service{
		bot: bot,
	}
}

func (s *Service) GetChatAdministrators(chatConfig tgbotapi.ChatConfig) ([]tgbotapi.ChatMember, error) {
	return s.bot.GetChatAdministrators(
		tgbotapi.ChatAdministratorsConfig{
			ChatConfig: chatConfig,
		},
	)
}

func (s *Service) NewMessage(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, text)
}

func (s *Service) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	return s.bot.Send(c)
}

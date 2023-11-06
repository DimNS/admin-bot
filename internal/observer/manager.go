package observer

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Manager struct {
	bot         BotProvider
	chanUpdates tgbotapi.UpdatesChannel
	debug       bool
}

func NewManager(bot BotProvider, chanUpdates tgbotapi.UpdatesChannel, debug bool) *Manager {
	return &Manager{
		bot:         bot,
		chanUpdates: chanUpdates,
		debug:       debug,
	}
}

func (m *Manager) Run() {
	for update := range m.chanUpdates {
		message := getMessageText(update.Message.Text)
		if message == "" {
			continue
		}
		if m.debug {
			log.Printf("Output: %s", message)
		}

		canContinue := m.processingUpdate(update.Message)
		if !canContinue {
			continue
		}

		err := m.sendMessage(update.Message, message)
		if err != nil {
			log.Printf("sendMessage error: %v", err)
		}
	}
}

func (m *Manager) processingUpdate(message *tgbotapi.Message) bool {
	if message == nil {
		return false
	}
	if m.debug {
		log.Printf("Message: \"%s\"", message.Text)
	}

	if !m.debug {
		admins, err := m.bot.GetChatAdministrators(message.Chat.ChatConfig())
		if err != nil {
			log.Printf("GetChatAdministrators error: %v", err)
			return false
		}
		if !authorIsAdmin(admins, message.From.ID) {
			return false
		}
	}

	return true
}

func (m *Manager) sendMessage(updateMsg *tgbotapi.Message, message string) error {
	msg := m.bot.NewMessage(updateMsg.Chat.ID, message)
	msg.ParseMode = "html"
	msg.DisableWebPagePreview = true

	if updateMsg.ReplyToMessage != nil {
		msg.ReplyToMessageID = updateMsg.ReplyToMessage.MessageID

		if updateMsg.ReplyToMessage.From != nil && updateMsg.ReplyToMessage.From.UserName != "" {
			msg.Text = "@" + updateMsg.ReplyToMessage.From.UserName + " " + msg.Text
		}
	}

	_, err := m.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("[%s] Send message error: %w", msg.Text, err)
	}

	return nil
}

func authorIsAdmin(admins []tgbotapi.ChatMember, userID int64) bool {
	for _, admin := range admins {
		if admin.User != nil && admin.User.ID == userID {
			return true
		}
	}

	return false
}

func getMessageText(text string) string {
	switch text {
	case "/help", "/хелп":
		return `БОТ РАБОТАЕТ ТОЛЬКО У АДМИНОВ.

Команды можно писать обычным сообщением и ответом на сообщение.

Список доступных команд:
[<code>/help</code>, <code>/хелп</code>] Список доступных команд бота
[<code>/php</code>, <code>/пхп</code>] @phpGeeks - Best PHP chat
[<code>/jun</code>, <code>/джун</code>] @phpGeeksJunior - Группа для новичков. Не стесняйтесь задавать вопросы по php.
[<code>/go</code>, <code>/го</code>] @golangGeeks - Приветствуем всех в нашем гетеросексуальном чате гоферов!
[<code>/db</code>, <code>/дб</code>] @dbGeeks - Чат про базы данных, их устройство и приемы работы с ними.
[<code>/lara</code>, <code>/лара</code>] @laravel_pro - Официальный чат для всех Laravel программистов.
[<code>/js</code>, <code>/жс</code>] @jsChat - Чат посвященный программированию на языке JavaScript.
[<code>/hr</code>, <code>/хр</code>] @jobGeeks - Топ вакансии (250 000+ р/мес).
[<code>/fl</code>, <code>/фл</code>] @freelanceGeeks - IT фриланс, ищем исполнителей и заказчиков, делимся опытом и проблемами связанными с фрилансом.
[<code>/job</code>, <code>/раб</code>] Объединяет сразу две команды: <code>/hr</code> и <code>/fl</code>.
[<code>/code</code>, <code>/код</code>] Код в нашем чате <a href="https://t.me/phpGeeks/1318040">ложут</a> на pastebin.org, gist.github.com или любой аналогичный ресурс (с)der_Igel
[<code>/nometa</code>, <code>/номета</code>] nometa.xyz`
	case "/php", "/пхп":
		return "@phpGeeks - Best PHP chat"
	case "/jun", "/джун":
		return "@phpGeeksJunior - Группа для новичков. Не стесняйтесь задавать вопросы по php."
	case "/go", "/го":
		return "@golangGeeks - Приветствуем всех в нашем гетеросексуальном чате гоферов!"
	case "/db", "/бд":
		return "@dbGeeks - Чат про базы данных, их устройство и приемы работы с ними."
	case "/lara", "/лара":
		return "@laravel_pro - Официальный чат для всех Laravel программистов."
	case "/js", "/жс":
		return "@jsChat - Чат посвященный программированию на языке JavaScript."
	case "/hr", "/хр":
		return "@jobGeeks - Топ вакансии (250 000+ р/мес)."
	case "/fl", "/фл":
		return "@freelanceGeeks - IT фриланс, ищем исполнителей и заказчиков, делимся опытом и проблемами связанными с фрилансом."
	case "/job", "/раб":
		return `@jobGeeks - Топ вакансии (250 000+ р/мес).
@freelanceGeeks - IT фриланс, ищем исполнителей и заказчиков, делимся опытом и проблемами связанными с фрилансом.`
	case "/code", "/код":
		return "Код в нашем чате <a href=\"https://t.me/phpGeeks/1318040\">ложут</a> на pastebin.org, gist.github.com или любой аналогичный ресурс (с)der_Igel"
	case "/nometa", "/номета":
		return "nometa.xyz"
	}

	return ""
}

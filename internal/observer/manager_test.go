package observer

import (
	"testing"

	"geeksonator/internal/observer/mocks"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewManager(t *testing.T) {
	type args struct {
		bot         BotProvider
		chanUpdates tgbotapi.UpdatesChannel
		debug       bool
	}
	tests := []struct {
		name string
		args args
		want *Manager
	}{
		{
			name: "Success",
			args: args{
				bot:         nil,
				chanUpdates: nil,
				debug:       false,
			},
			want: &Manager{
				bot:         nil,
				chanUpdates: nil,
				debug:       false,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := NewManager(tt.args.bot, tt.args.chanUpdates, tt.args.debug)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestManager_processingUpdate(t *testing.T) {
	type args struct {
		message *tgbotapi.Message
	}
	tests := []struct {
		name string
		man  func() *Manager
		args args
		want bool
	}{
		{
			name: "Message is nil",
			man: func() *Manager {
				return &Manager{}
			},
			args: args{
				message: nil,
			},
			want: false,
		},
		{
			name: "Author is admin",
			man: func() *Manager {
				botProvider := mocks.NewBotProviderMock(t)

				botProvider.EXPECT().
					GetChatAdministrators(tgbotapi.ChatConfig{}).
					Return(
						[]tgbotapi.ChatMember{
							{
								User: &tgbotapi.User{
									ID: 100500,
								},
							},
						},
						nil,
					)

				return &Manager{
					bot:   botProvider,
					debug: false,
				}
			},
			args: args{
				message: &tgbotapi.Message{
					Chat: &tgbotapi.Chat{},
					From: &tgbotapi.User{
						ID: 100500,
					},
					Text: "test",
				},
			},
			want: true,
		},
		{
			name: "Author is not admin",
			man: func() *Manager {
				botProvider := mocks.NewBotProviderMock(t)

				botProvider.EXPECT().
					GetChatAdministrators(tgbotapi.ChatConfig{}).
					Return(
						[]tgbotapi.ChatMember{
							{
								User: &tgbotapi.User{
									ID: 100500,
								},
							},
						},
						nil,
					)

				return &Manager{
					bot:   botProvider,
					debug: false,
				}
			},
			args: args{
				message: &tgbotapi.Message{
					Chat: &tgbotapi.Chat{},
					From: &tgbotapi.User{
						ID: 100501,
					},
					Text: "test",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := tt.man().processingUpdate(tt.args.message)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestManager_sendMessage(t *testing.T) {
	type args struct {
		updateMsg *tgbotapi.Message
		message   string
	}
	tests := []struct {
		name    string
		man     func() *Manager
		args    args
		wantErr error
	}{
		{
			name: "Success without reply to message",
			man: func() *Manager {
				botProvider := mocks.NewBotProviderMock(t)

				botProvider.EXPECT().
					NewMessage(int64(100500), "message text").
					Return(tgbotapi.MessageConfig{
						BaseChat: tgbotapi.BaseChat{
							ChatID: 100500,
						},
						Text: "message text",
					})

				botProvider.EXPECT().
					Send(
						tgbotapi.MessageConfig{
							BaseChat: tgbotapi.BaseChat{
								ChatID: 100500,
							},
							Text:                  "message text",
							ParseMode:             "html",
							DisableWebPagePreview: true,
						},
					).
					Return(tgbotapi.Message{}, nil)

				return &Manager{
					bot:         botProvider,
					chanUpdates: make(<-chan tgbotapi.Update, 100),
					debug:       false,
				}
			},
			args: args{
				updateMsg: &tgbotapi.Message{
					Chat: &tgbotapi.Chat{
						ID: 100500,
					},
				},
				message: "message text",
			},
			wantErr: nil,
		},
		{
			name: "Success with reply to message",
			man: func() *Manager {
				botProvider := mocks.NewBotProviderMock(t)

				botProvider.EXPECT().
					NewMessage(int64(100500), "message text").
					Return(tgbotapi.MessageConfig{
						BaseChat: tgbotapi.BaseChat{
							ChatID: 100500,
						},
						Text: "message text",
					})

				botProvider.EXPECT().
					Send(
						tgbotapi.MessageConfig{
							BaseChat: tgbotapi.BaseChat{
								ChatID:           100500,
								ReplyToMessageID: 100,
							},
							Text:                  "@username message text",
							ParseMode:             "html",
							DisableWebPagePreview: true,
						},
					).
					Return(tgbotapi.Message{}, nil)

				return &Manager{
					bot:         botProvider,
					chanUpdates: make(<-chan tgbotapi.Update, 100),
					debug:       false,
				}
			},
			args: args{
				updateMsg: &tgbotapi.Message{
					Chat: &tgbotapi.Chat{
						ID: 100500,
					},
					ReplyToMessage: &tgbotapi.Message{
						MessageID: 100,
						From: &tgbotapi.User{
							UserName: "username",
						},
						Text: "reply to message text",
					},
				},
				message: "message text",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.man().sendMessage(tt.args.updateMsg, tt.args.message)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_authorIsAdmin(t *testing.T) {
	type args struct {
		admins []tgbotapi.ChatMember
		userID int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Author is admin",
			args: args{
				admins: []tgbotapi.ChatMember{
					{
						User: &tgbotapi.User{
							ID: 100500,
						},
					},
				},
				userID: 100500,
			},
			want: true,
		},
		{
			name: "Author is not admin",
			args: args{
				admins: []tgbotapi.ChatMember{
					{
						User: &tgbotapi.User{
							ID: 100500,
						},
					},
				},
				userID: 100501,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := authorIsAdmin(tt.args.admins, tt.args.userID)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getMessageText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				text: "/lara",
			},
			want: "@laravel_pro - Официальный чат для всех Laravel программистов.",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := getMessageText(tt.args.text)
			assert.Equal(t, tt.want, got)
		})
	}
}

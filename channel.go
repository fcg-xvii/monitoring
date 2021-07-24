package monitoring

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Channel interface {
	Log(title, message string)
}

// channel telegram

type ChannelTelegram struct {
	botAPI *tgbotapi.BotAPI
	ChatID int64
}

func (s *ChannelTelegram) Log(title, message string) {
	chatMessage := fmt.Sprintf(`<b>%v</b>\n%v`, title, message)
	tMsg := tgbotapi.NewMessage(s.ChatID, chatMessage)
	s.botAPI.Send(tMsg)
}

func (s *ChannelTelegram) JSONField(fieldName string) (ptr interface{}, err error) {
	var token string
	switch fieldName {
	case "token":
		ptr = &token
	case "chat_id":
		ptr = &s.ChatID
	}
	return
}

func (s *ChannelTelegram) JSONFinish() (err error) {
	if 
}

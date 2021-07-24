package monitoring

import (
	"errors"
	"fmt"

	"github.com/fcg-xvii/go-tools/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ConstructorChannelTelegram(m json.Map) (res Channel, err error) {
	token := m.String("token", "")
	if len(token) == 0 {
		err = errors.New("Constructor channel telegram error :: [token] is not defined")
		return
	}
	chatID := m.Int("token", 0)
	if chatID == 0 {
		err = errors.New("Constructor channel telegram error :: [chat_id] is not defined")
		return
	}
	tgChannel := &ChannelTelegram{
		chatID: chatID,
	}
	if tgChannel.api, err = tgbotapi.NewBotAPI(token); err == nil {
		res = tgChannel
	}
	return
}

type ChannelTelegram struct {
	api    *tgbotapi.BotAPI
	chatID int64
}

func (s *ChannelTelegram) Log(title, message string) {
	chatMessage := fmt.Sprintf(`<b>%v</b>\n%v`, title, message)
	tMsg := tgbotapi.NewMessage(s.chatID, chatMessage)
	s.api.Send(tMsg)
}

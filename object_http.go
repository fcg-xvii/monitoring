package monitoring

import (
	"errors"
	"net/http"

	"github.com/fcg-xvii/go-tools/json"
)

func ConstructorObjectHttpGet(m json.Map) (res Object, err error) {
	url := m.String("url", "")
	if len(url) == 0 {
		err = errors.New("HttpGet constructor error :: url is not defined")
		return
	}
	channels, err := ChannelsFromList(m.Slice("channels", nil))
	if err != nil {
		return nil, err
	}
	res = &HttpGet{
		url:      url,
		channels: channels,
	}
	return
}

type HttpGet struct {
	url      string
	channels []Channel
}

func (s *HttpGet) Request() {
	_, err := http.Get(s.url)
	for _, ch := range s.channels {
		ch.Log("ALERT!!!", err.Error())
	}
}

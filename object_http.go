package monitoring

import (
	"errors"
	"log"
	"net/http"

	"github.com/fcg-xvii/go-tools/json"
)

func ConstructorObjectHttpGet(m json.Map) (res Object, err error) {
	url := m.String("url", "")
	if len(url) == 0 {
		err = errors.New("HttpGet constructor error :: url is not defined")
		return
	}
	res = &HttpGet{
		URL: url,
	}
	return
}

type HttpGet struct {
	URL string
}

func (s *HttpGet) Request() {
	resp, err := http.Get(s.URL)
	log.Println(resp, err)
}

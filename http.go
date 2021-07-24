package monitoring

import (
	"log"
	"net/http"
)

type HttpGet struct {
	URL string
}

func (s *HttpGet) Request() {
	resp, err := http.Get(s.URL)
	log.Println(resp, err)
}

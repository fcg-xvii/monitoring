package monitoring

import (
	"context"
	"log"
	"testing"
	"time"
)

type dObject struct {
	i int
}

func (s *dObject) Request() {
	log.Println("request", s.i)
}

func TestMonitoring(t *testing.T) {
	objList := []Object{
		&HttpGet{
			URL: "https://gm6.ru",
		},
	}
	m := New(context.Background(), objList, time.Second*10)
	go func() {
		time.Sleep(time.Second * 50)
		m.Cancel()
		log.Println("Canceled...")
	}()
	time.Sleep(time.Minute * 2)
}
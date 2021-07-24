package monitoring

import (
	"context"
	"log"
	"time"

	"github.com/fcg-xvii/go-tools/json"
)

func New(ctx context.Context, objects []Object, interval time.Duration) *Monitoring {
	cctx, cancel := context.WithCancel(ctx)
	m := &Monitoring{
		objects:   objects,
		ctx:       cctx,
		cancel:    cancel,
		mInterval: interval,
	}
	go m.start()
	return m
}

func FromMap(ctx context.Context, m json.Map) *Monitoring {
	items := m.Slice("monitoring", nil)
	res := &Monitoring{
		mInterval: time.Minute * time.Duration(m.Int("interval", 5)),
		objects: make([]Object, 0, len(items))
	}
	for _, item := range items {
		if iMap, check := item.(map[string]interface{}); check {
			mm := json.Map(iMap)
			log.Println("mm", mm)
		}
	}
	return res
}

func FromConfig(ctx context.Context, fileName string) (res *Monitoring, err error) {
	var m json.Map
	if err = json.UnmarshalFile(fileName, &m); err != nil {
		return
	}
	res = FromMap(ctx, m)
	return
}

type Monitoring struct {
	objects   []Object
	ctx       context.Context
	cancel    context.CancelFunc
	mInterval time.Duration
}

func (s *Monitoring) Cancel() {
	s.cancel()
}

func (s *Monitoring) start() {
	mTimer := time.NewTicker(s.mInterval)
	for {
		select {
		case <-mTimer.C:
			{
				for _, obj := range s.objects {
					obj.Request()
				}
			}
		case <-s.ctx.Done():
			return
		}
	}
}

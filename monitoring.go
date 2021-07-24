package monitoring

import (
	"context"
	"time"
)

type Object interface {
	Request()
}

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

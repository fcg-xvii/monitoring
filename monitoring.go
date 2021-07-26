package monitoring

import (
	"context"
	"log"
	"time"

	"github.com/fcg-xvii/go-tools/containers/concurrent"
	"github.com/fcg-xvii/go-tools/json"
)

/*
{"ok":true,"result":[{"update_id":149690937,
"channel_post":{"message_id":235,"sender_chat":{"id":-1001269290416,"title":"\u0417\u0430\u043f\u0438\u0441\u043a\u0438 \u0424\u043b\u0438\u043d\u0442\u0430","type":"channel"},"chat":{"id":-1001269290416,"title":"\u0417\u0430\u043f\u0438\u0441\u043a\u0438 \u0424\u043b\u0438\u043d\u0442\u0430","type":"channel"},"date":1627307422,"text":"@SubaruMonitoringBot 123","entities":[{"offset":0,"length":20,"type":"mention"}]}}]}
*/

func New(ctx context.Context, objects []Object, interval time.Duration) *Monitoring {
	cctx, cancel := context.WithCancel(ctx)
	l := concurrent.NewList()
	for _, obj := range objects {
		l.PushBack(obj)
	}
	m := &Monitoring{
		objects:   l,
		ctx:       cctx,
		cancel:    cancel,
		mInterval: interval,
	}
	go m.start()
	return m
}

func FromMap(ctx context.Context, m json.Map) *Monitoring {
	objects, err := ObjectsFromList(m.Slice("monitoring", nil))
	if err != nil {
		log.Println(err)
	}
	cctx, cancel := context.WithCancel(ctx)
	res := &Monitoring{
		mInterval: time.Minute * time.Duration(m.Int("interval", 5)),
		objects:   objects,
		ctx:       cctx,
		cancel:    cancel,
	}
	res.start()
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
	objects   *concurrent.List
	ctx       context.Context
	cancel    context.CancelFunc
	mInterval time.Duration
}

func (s *Monitoring) Cancel() {
	s.cancel()
}

func (s *Monitoring) ObjectAdd(obj Object) {
	s.objects.PushBack(obj)
}

func (s *Monitoring) ObjectRemove(obj Object) {
	elem := s.objects.Search(obj)
	if elem != nil {
		s.objects.Remove(elem)
	}
}

func (s *Monitoring) start() {
	mTimer := time.NewTicker(s.mInterval)
	for {
		select {
		case <-mTimer.C:
			{
				for f := s.objects.First(); f != nil; f = f.Next() {
					f.Val().(Object).Request()
				}
			}
		case <-s.ctx.Done():
			return
		}
	}
}

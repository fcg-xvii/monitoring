package monitoring

import (
	"errors"
	"log"
	"sync"

	"github.com/fcg-xvii/go-tools/json"
)

func init() {
	RegisterChannelType("telegram", ConstructorChannelTelegram)
}

func ChannelsFromList(iface interface{}) ([]Channel, error) {
	ifaces, check := iface.([]interface{})
	if !check {
		return nil, errors.New("Channels constructor error :: expected list")
	}
	res := make([]Channel, 0, len(ifaces))
	for _, i := range ifaces {
		m, check := i.(map[string]interface{})
		if !check {
			log.Println("Channel constructor error :: expected map")
			continue
		}
		mm := json.Map(m)
		cType := mm.String("type", "")
		constructor, ok := channelTypes.Load(cType)
		if !ok {
			log.Printf("Channel constructor error :: unexpected type [%v]\n", cType)
			continue
		}
		ch, err := constructor.(ChannelConstructor)(mm)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		res = append(res, ch)
	}
	if len(res) == 0 {
		return nil, errors.New("Channels constructor error :: result list is empty")
	}
	return res, nil
}

type ChannelConstructor func(m json.Map) (Channel, error)

var (
	channelTypes = new(sync.Map)
)

func RegisterChannelType(typeName string, constructor ChannelConstructor) {
	channelTypes.Store(typeName, constructor)
}

func RemoveChannelType(typeName string) {
	channelTypes.Delete(typeName)
}

type Channel interface {
	Log(title, message string)
}

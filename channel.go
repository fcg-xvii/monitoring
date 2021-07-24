package monitoring

import (
	"sync"

	"github.com/fcg-xvii/go-tools/json"
)

func init() {
	RegisterChannelType("telegram", ConstructorChannelTelegram)
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

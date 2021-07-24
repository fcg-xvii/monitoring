package monitoring

import (
	"sync"

	"github.com/fcg-xvii/go-tools/json"
)

func init() {
	RegisterObjectType("http-get", ConstructorObjectHttpGet)
}

type ObjectConstructor func(m json.Map) (Object, error)

var (
	objectTypes = new(sync.Map)
)

type Object interface {
	Request()
}

func RegisterObjectType(typeName string, constructor ObjectConstructor) {
	objectTypes.Store(typeName, constructor)
}

func RemoveObjectType(typeName string) {
	objectTypes.Delete(typeName)
}

package monitoring

import (
	"errors"
	"log"
	"sync"

	"github.com/fcg-xvii/go-tools/containers/concurrent"
	"github.com/fcg-xvii/go-tools/json"
)

func init() {
	RegisterObjectType("http-get", ConstructorObjectHttpGet)
}

func ObjectsFromList(list interface{}) (*concurrent.List, error) {
	ifaces, check := list.([]interface{})
	if !check {
		return nil, errors.New("Objects constructor error :: expected list")
	}
	res := concurrent.NewList()
	for _, i := range ifaces {
		m, check := i.(map[string]interface{})
		if !check {
			log.Println("Object constructor error :: expected map")
			continue
		}
		mm := json.Map(m)
		cType := mm.String("type", "")
		constructor, ok := objectTypes.Load(cType)
		if !ok {
			log.Printf("Object constructor error :: unexpected type [%v]\n", cType)
			continue
		}
		ch, err := constructor.(ObjectConstructor)(mm)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		res.PushBack(ch)
	}
	if res.Size() == 0 {
		return res, errors.New("Objects constructor error :: result list is empty")
	}
	return res, nil
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

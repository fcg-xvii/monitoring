package monitoring

import (
	"github.com/fcg-xvii/go-tools/json"
)

func ReadConfig(fileName string) (res json.Map, err error) {
	err = json.UnmarshalFile(fileName, &res)
	return
}

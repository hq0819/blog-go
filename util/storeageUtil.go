package util

import (
	"github.com/gofiber/storage/memory"
	"sync"
)

var Instance *memory.Storage
var once sync.Once

func GetInstance() *memory.Storage {

	once.Do(func() {
		Instance = memory.New()

	})
	return Instance
}

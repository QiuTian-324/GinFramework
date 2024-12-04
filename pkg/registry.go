package pkg

import "sync"

var (
	handlerRegistry = make(map[string]interface{})
	mu              sync.RWMutex
)

func RegisterHandler(name string, handler interface{}) {
	mu.Lock()
	defer mu.Unlock()
	handlerRegistry[name] = handler
}

func GetHandler(name string) interface{} {
	mu.RLock()
	defer mu.RUnlock()
	return handlerRegistry[name]
}

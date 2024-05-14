package gateway

import (
	"context"
	"sync"
)

// CtxStorage Хранилище функций отмены контекстов
type CtxStorage struct {
	mutex sync.RWMutex
	ctxs  map[int]context.CancelFunc
}

func NewCtxStorage(size int) *CtxStorage {
	storage := &CtxStorage{}
	storage.ctxs = make(map[int]context.CancelFunc, size)
	return storage
}

func (c *CtxStorage) Get(id int) context.CancelFunc {

	c.mutex.RLock()
	value := c.ctxs[id]
	c.mutex.RUnlock()

	return value
}

func (c *CtxStorage) Set(id int, ctxFunc context.CancelFunc) {

	c.mutex.Lock()
	c.ctxs[id] = ctxFunc
	c.mutex.Unlock()
}

func (c *CtxStorage) Remove(id int) {
	c.mutex.Lock()
	delete(c.ctxs, id)
	c.mutex.Unlock()
}

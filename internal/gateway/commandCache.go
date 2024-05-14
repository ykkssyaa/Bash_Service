package gateway

import (
	"github.com/ykkssyaa/Bash_Service/internal/models"
	"sync"
)

type CommandCache struct {
	mp    map[int]models.Command
	mutex sync.RWMutex
}

func NewCommandCache(size int) *CommandCache {
	storage := &CommandCache{}
	storage.mp = make(map[int]models.Command, size)
	return storage
}

func (c *CommandCache) Get(id int) (models.Command, error) {

	c.mutex.RLock()
	value := c.mp[id]
	c.mutex.RUnlock()

	return value, nil
}

func (c *CommandCache) Set(id int, cmd models.Command) error {

	c.mutex.Lock()
	c.mp[id] = cmd
	c.mutex.Unlock()

	return nil
}

func (c *CommandCache) Remove(id int) error {
	c.mutex.Lock()
	delete(c.mp, id)
	c.mutex.Unlock()

	return nil
}

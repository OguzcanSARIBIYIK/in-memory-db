package indb

import "sync"

type InDB struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func (i *InDB) Save(key string, value interface{}) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.data[key] = value
}

func (i *InDB) Get(key string) (interface{}, bool) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	data, err := i.data[key]
	return data, err
}

func (i *InDB) GetAll() map[string]interface{} {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.data
}

func (i *InDB) Delete(key string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	delete(i.data, key)
}

func Init() *InDB {
	return &InDB{
		data: map[string]interface{}{},
		mu:   sync.RWMutex{},
	}
}

package syncmap

import (
	"sync"
)

type SyncMap[K comparable, V any] struct {
	sync.RWMutex
	Data map[K]V
}

func (syncMap *SyncMap[K, V]) Make() {
	syncMap.TryLock()
	syncMap.Data = make(map[K]V)
	syncMap.Unlock()
}

func (syncMap *SyncMap[K, V]) Get(key K) V {
	syncMap.RLock()
	value := syncMap.Data[key]
	syncMap.RUnlock()

	return value
}

func (syncMap *SyncMap[K, V]) Add(key K, value V) {
	syncMap.Lock()

	if syncMap.Data == nil {
		syncMap.Make()
	}

	syncMap.Push(key, value)
}

func (syncMap *SyncMap[K, V]) Push(key K, value V) {
	syncMap.TryLock()
	syncMap.Data[key] = value
	syncMap.Unlock()
}

func (syncMap *SyncMap[K, V]) Del(key K) {
	syncMap.Lock()
	delete(syncMap.Data, key)
	syncMap.Unlock()
}

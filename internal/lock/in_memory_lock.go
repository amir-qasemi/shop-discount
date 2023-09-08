package lock

import (
	"errors"
	"sync"
)

// inMemLockStore a implementation that keeps the locks in memory.
// No house cleaning is done. So this implementation suffers from inbounded memeory growth
// The process of getting locks from this service is serialized with mapLock.
type inMemLockStore struct {
	locks   map[string]*sync.RWMutex
	mapLock sync.Mutex
}

func (l *inMemLockStore) Lock(key string) error {
	l.mapLock.Lock()
	lock, ok := l.locks[key]
	if !ok {
		lock = &sync.RWMutex{}
		l.locks[key] = lock
	}
	l.mapLock.Unlock()

	lock.Lock()
	return nil
}

func (l *inMemLockStore) Unlock(key string) error {
	l.mapLock.Lock()
	lock, ok := l.locks[key]
	l.mapLock.Unlock()
	if !ok {
		return errors.New("Lock")
	}

	lock.Unlock()
	return nil
}

// NewInMemLockStore returns new a lock store implemented in memory
func NewInMemLockStore() LockStore {
	return &inMemLockStore{locks: make(map[string]*sync.RWMutex)}
}

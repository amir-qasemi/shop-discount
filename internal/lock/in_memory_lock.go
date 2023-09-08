package lock

import "sync"

// inMemLockStore a implementation that keeps the locks in memory.
// No house cleaning is done. So this implementation suffers from inbounded memeory growth
// The process of getting locks from this service is serialized with mapLock.
type inMemLockStore struct {
	locks   map[string]*sync.RWMutex
	mapLock sync.Mutex
}

func (l *inMemLockStore) Lock(key string) *sync.RWMutex {
	l.mapLock.Lock()
	defer l.mapLock.Unlock()

	lock, ok := l.locks[key]
	if !ok {
		lock = &sync.RWMutex{}
		l.locks[key] = lock
	}
	return lock
}

// NewInMemLockStore returns new a lock store implemented in memory
func NewInMemLockStore() LockStore {
	return &inMemLockStore{locks: make(map[string]*sync.RWMutex)}
}

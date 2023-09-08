package lock

import "sync"

// LockStore a service for implementing lock manager.
// Each lock has an application defined lock with the given key.
// If the given key does not exist already, a new lock will be created for that key.
// This service should a global service. But for demonstration purposes, it is implemented in the same process.
type LockStore interface {
	// Lock get the lock for the given key
	Lock(key string) *sync.RWMutex
}

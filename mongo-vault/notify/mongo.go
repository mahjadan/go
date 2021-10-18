package notify

import (
	"sync"
)

type (
	// Event defines an indication of a point-in-time occurrence.
	MongoEvent struct {
		// Data will hold the username and password of mongodb.
		Data map[string]interface{}
		// Error will hold errs in case of error while getting Data
		Error error
	}

	MongoEventNotifier struct {
		// Using a map with an empty struct allows us to keep the observers
		// unique while still keeping memory usage relatively low.
		observers map[Observer]struct{}
		rwMutex   sync.RWMutex
	}
)

func (m *MongoEventNotifier) Notify(event MongoEvent) {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()
	for observer, _ := range m.observers {
		observer.OnNotify(event)
	}
}

func (m *MongoEventNotifier) Register(observer Observer) {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()
	m.observers[observer] = struct{}{}
}

func (m *MongoEventNotifier) Deregister(observer Observer) {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()
	delete(m.observers, observer)
}

func New() Notifier {
	return &MongoEventNotifier{
		observers: make(map[Observer]struct{}),
		rwMutex:   sync.RWMutex{},
	}
}

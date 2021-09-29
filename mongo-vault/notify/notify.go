package notify

// Observer defines a standard interface for instances that wish to list for
// the occurrence of a specific event.
type Observer interface {
	// OnNotify allows an event to be "published" to interface implementations.
	// In the "real world", error handling would likely be implemented.
	OnNotify(MongoEvent)
}

// Notifier is the instance being observed. Publisher is perhaps another decent
// name, but naming things is hard.
type Notifier interface {
	// Register allows an instance to register itself to listen/observe
	// events.
	Register(Observer)
	// Deregister allows an instance to remove itself from the collection
	// of observers/listeners.
	Deregister(Observer)
	// Notify publishes new events to listeners. The method is not
	// absolutely necessary, as each implementation could define this itself
	// without losing functionality.
	Notify(event MongoEvent)
}

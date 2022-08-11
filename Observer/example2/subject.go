package main

// A Subject is an interface that defines the methods that an Observer can use to register and
// deregister itself, and to notify all Observers.
// @property register - This method is used to register an observer to the subject.
// @property deregister - Removes an observer from the list of observers.
// @property notifyAll - This method is used to notify all the observers when there is some change in
// the Subject state.
type Subject interface {
	register(observer ...Observer)
	deregister(observer ...Observer)
	notifyAll()
}

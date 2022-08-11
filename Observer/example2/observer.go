package main

// An Observer is an interface that has two methods: update and getID.
// @property update - This is the method that will be called when the subject's state changes.
// @property {string} getID - returns the ID of the observer
type Observer interface {
	update(string)
	getID() string
}

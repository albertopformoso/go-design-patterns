package main

import "log"

// Item is a struct with a slice of Observers, a string, and a boolean.
// @property {[]Observer} observerList - A list of observers that are watching this item.
// @property {string} name - The name of the item.
// @property {bool} inStock - This is the property that will be observed.
type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

// `newItem` is a function that takes a string and returns a pointer to an Item
func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

// This is the function that is called when the item is in stock. It prints a message to the console,
// sets the inStock property to true, and then calls the notifyAll function.
func (i *Item) updateAvailability() {
	log.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

// Adding an observer to the observerList.
func (i *Item) register(observers ...Observer) {
	i.observerList = append(i.observerList, observers...)
}

// Removing an observer from the observerList.
func (i *Item) deregister(observers ...Observer) {
	for _, observer := range observers {
		i.observerList = removeFromslice(i.observerList, observer)
	}
}

// Iterating through the observerList and calling the update function on each observer.
func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

// It takes a slice of observers and an observer to remove, and returns a new slice of observers with
// the observer to remove removed
func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

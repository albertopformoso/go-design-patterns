package main

import (
	"log"

	"github.com/gofrs/uuid"
)

func main() {
	var p Publisher = newPublisher()
	p.broadcastMsg("hello")

	s1 := newSubscriber("123")
	s2 := newSubscriber("456")
	s3 := newSubscriber("789")
	p.addSubscriber(s1, s2, s3)
	p.broadcastMsg("hello again")
	p.removeSubscriber(s1.id(), s3.id())
	p.broadcastMsg("hello one more time")

	a1 := newAutogenerateIdSubscriber()
	p.addSubscriber(a1)
	p.broadcastMsg("bye")
}

// // INTERFACES ////
// A Publisher is a type that can add and remove subscribers, and broadcast messages to them.
// @property addSubscriber - This method adds a subscriber to the publisher.
// @property removeSubscriber - This method removes a subscriber from the publisher's list of
// subscribers.
// @property broadcastMsg - This is the method that will be used to broadcast messages to all the
// subscribers.
type Publisher interface {
	addSubscriber(subscriber ...Subscriber)
	removeSubscriber(subId ...string)
	broadcastMsg(msg string)
}

// A Subscriber is a type that has an id() method that returns a string and a react(msg string) method.
// @property {string} id - a unique identifier for the subscriber
// @property react - This is the function that will be called when a message is published.
type Subscriber interface {
	id() string
	react(msg string)
}

// // PUBLISHER ////
// `publisher` is a type that has a map of strings to `Subscriber`s.
// @property subscribers - A map of Subscriber objects, where the key is the subscriber's name.
type publisher struct {
	subscribers map[string]Subscriber
}

// `newPublisher()` returns a pointer to a `publisher` struct with a `subscribers` map that has a
// string key and a `Subscriber` value
func newPublisher() *publisher {
	return &publisher{subscribers: make(map[string]Subscriber)}
}

// Adding a subscriber to the publisher.
func (p *publisher) addSubscriber(subscriber ...Subscriber) {
	for _, sub := range subscriber {
		p.subscribers[sub.id()] = sub
	}
}

// Deleting a subscriber from the map of subscribers.
func (p *publisher) removeSubscriber(subIds ...string) {
	for _, id := range subIds {
		delete(p.subscribers, id)
	}
}

// Iterating over the map of subscribers and calling the `react` method on each subscriber.
func (p *publisher) broadcastMsg(msg string) {
	for _, subscriber := range p.subscribers {
		subscriber.react(msg)
	}
}

// // SUBSCRIBER ////
// A subscriber is a struct with a subId field of type string.
// @property {string} subId - The unique identifier for the subscriber.
type subscriber struct {
	subId string
}

// NewSubscriber returns a pointer to a new subscriber with the given subId.
func newSubscriber(subId string) *subscriber {
	return &subscriber{subId: subId}
}

// This is a method that is defined on the `subscriber` type. It returns the `subId` field of the
// `subscriber` struct.
func (s *subscriber) id() string {
	return s.subId
}

// This is the function that will be called when a message is published.
func (s *subscriber) react(msg string) {
	log.Printf("ID: %v - received: %v", s.subId, msg)
}

// // AUTO-ID SUBSCRIBER ////
// It's a subscriber that automatically generates an ID for the subscription.
// @property {string} subId - The ID of the subscriber.
type autogenerateIdSubscriber struct {
	subId string
}

// It creates a new autogenerateIdSubscriber struct, generates a new UUID, and returns a pointer to the
// struct
func newAutogenerateIdSubscriber() *autogenerateIdSubscriber {
	uuid, _ := uuid.NewV4()
	return &autogenerateIdSubscriber{
		subId: uuid.String(),
	}
}

// It's a method that is defined on the `autogenerateIdSubscriber` type. It returns the `subId` field
// of the `autogenerateIdSubscriber` struct.
func (a *autogenerateIdSubscriber) id() string {
	return a.subId
}

// It's a method that is defined on the `autogenerateIdSubscriber` type. It returns the `subId` field
// of the `autogenerateIdSubscriber` struct.
func (a *autogenerateIdSubscriber) react(msg string) {
	log.Printf("ID: %s - autogen Id sub received: %s", a.subId, msg)
}

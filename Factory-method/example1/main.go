package main

import (
	"log"
)

func main() {
	farm("dog")
}

// An Animal is a type that can make a sound.
// @property {string} Sound - This is a method that returns a string.
type Animal interface {
	Sound() string
}

// `Cat` is a type that has no fields and no methods.
type Cat struct{}

// NewCat returns a pointer to a new `Cat` struct
func NewCat() *Cat { return &Cat{} }

// It's defining a method on the Cat struct.
func (c *Cat) Sound() string { return "meow" }

// A Dog is a type that has no fields.
type Dog struct{}

// NewDog returns a pointer to a Dog
func NewDog() *Dog { return &Dog{} }

// It's defining a method on the Dog struct.
func (c *Dog) Sound() string { return "woof" }

// "If the animal is a cat, create a new cat, if it's a dog, create a new dog, otherwise, log an error
// and exit."
//
// The above function is a good example of the Factory pattern
func farm(animal string) {
	var a Animal

	if animal == "cat" {
		a = NewCat()
	} else if animal == "dog" {
		a = NewDog()
	} else {
		log.Fatalln("unexistent animal")
	}

	log.Printf("%+v", a.Sound())
}

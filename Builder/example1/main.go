package main

import "log"

func main() {
	me := newHuman().withEyeColor("brown").withAge(25).withHeight(1.71)
	you := newHumanWithFields(30, 1.90, "green")

	log.Printf("%+v", me)
	log.Printf("%+v", you)

	me = me.reset()
	log.Printf("%+v", me)

	youngGiant := giant().withAge(3)
	oldGiant := giant().withAge(150)

	log.Printf("%+v", youngGiant)
	log.Printf("%+v", oldGiant)
}

// A human is a struct with three fields: age, height, and eyeColor.
// @property {int} age - int
// @property {float64} height - a float64
// @property {string} eyeColor - string
type human struct {
	age      int
	height   float64
	eyeColor string
}

// NewHuman returns a pointer to a human struct
// Empty constructor
func newHuman() *human {
	return &human{}
}

// NewHumanWithFields returns a pointer to a human with the given age, height, and eyeColor.
func newHumanWithFields(age int, height float64, eyeColor string) *human {
	return &human{
		age,
		height,
		eyeColor,
	}
}

//// Builders ////

// A method that takes a pointer to a human struct and an int. It sets the age of the human struct to
// the int and returns a pointer to the human struct.
func (h *human) withAge(age int) *human {
	h.age = age
	return h
}

// It's a method that takes a pointer to a human struct and a float64. It sets the height of the human
// struct to the float64 and returns a pointer to the human struct.
func (h *human) withHeight(height float64) *human {
	h.height = height
	return h
}

// It's a method that takes a pointer to a human struct and a string. It sets the eyeColor of the human
// struct to the string and returns a pointer to the human struct.
func (h *human) withEyeColor(color string) *human {
	h.eyeColor = color
	return h
}

// // Reset ////
// It's creating a new human struct and returning a pointer to it.
func (h *human) reset() *human {
	return newHuman()
}

//// Pre-defined Builders ////

// Create a new human with a height of 2.5 and an eye color of blue.
func giant() *human {
	return newHuman().withHeight(2.5).withEyeColor("blue")
}

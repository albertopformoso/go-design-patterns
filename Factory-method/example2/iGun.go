package main

// IGun is an interface that has two methods, setName and setPower, that take a string and an int
// respectively, and two methods, getName and getPower, that return a string and an int respectively.
// @property setName - sets the name of the gun
// @property setPower - sets the power of the gun
// @property {string} getName - returns the name of the gun
// @property {int} getPower - returns the power of the gun
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

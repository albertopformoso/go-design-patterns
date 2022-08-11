package main

// IPizza is an interface that has a method called getPrice that takes no parameters and returns an
// int.
// @property {int} getPrice - This is a method that returns the price of the pizza.
type IPizza interface {
	getPrice() int
}

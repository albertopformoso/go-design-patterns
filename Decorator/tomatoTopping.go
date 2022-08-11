package main

// TomatoTopping is a type that has a field called pizza of type IPizza.
// @property {IPizza} pizza - This is the IPizza interface that we're going to use to decorate the
// pizza.
type TomatoTopping struct {
	pizza IPizza
}

// AddTomatoTopping returns a pointer to a TomatoTopping struct that wraps the given IPizza.
func AddTomatoTopping(pizza IPizza) *TomatoTopping {
	return &TomatoTopping{pizza}
}

// getPrice returns the pizza price plus the price of the tomato topping.
func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

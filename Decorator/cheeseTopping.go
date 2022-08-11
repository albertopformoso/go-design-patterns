package main

// CheeseTopping is a type that has a field called pizza of type IPizza.
// @property {IPizza} pizza - This is the IPizza interface that we're going to use to decorate the
// pizza.
type CheeseTopping struct {
	pizza IPizza
}

// AddCheeseTopping returns a pointer to a CheeseTopping struct that wraps the given IPizza.
func AddCheeseTopping(pizza IPizza) *CheeseTopping {
	return &CheeseTopping{pizza}
}

// getPrice returns the pizza price plus the price of the cheese topping.
func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

package main

import "log"

func main() {
	pizza := NewVeggeMania()

	//Add cheese topping
	pizzaWithCheese := AddCheeseTopping(pizza)

	//Add tomato topping
	pizzaWithCheeseAndTomato := AddTomatoTopping(pizzaWithCheese)

	log.Printf("Price of pizza with tomato and cheese topping is $%d", pizzaWithCheeseAndTomato.getPrice())
}

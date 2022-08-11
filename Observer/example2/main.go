package main

func main() {
	shirtItem := newItem("Nike Shirt")

	observer1 := newCustomer("abc@gmail.com")
	observer2 := newCustomer("xyz@gmail.com")

	shirtItem.register(observer1, observer2)

	shirtItem.updateAvailability()
}

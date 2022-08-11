package main

import "log"

// @property {string} id - The unique identifier for the customer.
type Customer struct {
	id string
}

func newCustomer(id string) *Customer {
	return &Customer{
		id,
	}
}

// A method of the Customer struct.
func (c *Customer) update(itemName string) {
	log.Printf("Sending email to customer %s for item %s", c.id, itemName)
}

// A method of the Customer struct.
func (c *Customer) getID() string {
	return c.id
}

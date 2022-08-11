package main

// `VeggeMania` is a struct with no fields.
type VeggeMania struct{}

// `NewVeggeMania()` returns a pointer to a `VeggeMania` struct
func NewVeggeMania() *VeggeMania {
	return &VeggeMania{}
}

// getPrice returns the price of the veggie pizza
func (p *VeggeMania) getPrice() int {
	return 15
}

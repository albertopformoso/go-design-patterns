package main

// A Director is a struct that has a builder field.
// @property {IBuilder} builder - This is a reference to the builder interface.
type Director struct {
	builder IBuilder
}

// NewDirector returns a pointer to a Director struct with a builder field set to the value of the b
// parameter.
func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// Setting the builder field of the Director struct.
func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

// Calling the methods of the builder interface.
func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

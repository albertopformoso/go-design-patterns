package main

// NormalBuilder is a struct that has three fields: windowType, doorType, and floor.
// @property {string} windowType - The type of window to be used in the house.
// @property {string} doorType - The type of door to be used in the house.
// @property {int} floor - The number of floors the house will have.
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// This function returns a pointer to a NormalBuilder struct.
func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

// Setting the windowType field of the NormalBuilder struct to "wooden window".
func (b *NormalBuilder) setWindowType() {
	b.windowType = "wooden window"
}

// Setting the doorType field of the NormalBuilder struct to "wooden door".
func (b *NormalBuilder) setDoorType() {
	b.doorType = "wooden door"
}

// Setting the floor field of the NormalBuilder struct to 2.
func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

// Returning a House struct with the fields of the NormalBuilder struct.
func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

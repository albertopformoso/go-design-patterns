package main

// IglooBuilder is a type that has three fields, windowType, doorType, and floor, and all of them are
// strings.
// @property {string} windowType - The type of window to use.
// @property {string} doorType - The type of door to be used.
// @property {int} floor - The number of floors the igloo will have.
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// This function returns a pointer to an IglooBuilder struct.
func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

// Setting the windowType field of the IglooBuilder struct to "snow window".
func (b *IglooBuilder) setWindowType() {
	b.windowType = "snow window"
}

// Setting the doorType field of the IglooBuilder struct to "snow door".
func (b *IglooBuilder) setDoorType() {
	b.doorType = "snow door"
}

// Setting the floor field of the IglooBuilder struct to 1.
func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

// Returning a House struct with the fields doorType, windowType, and floor.
func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

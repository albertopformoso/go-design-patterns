package main

// IBuilder is an interface that has four methods: setWindowType, setDoorType, setNumFloor, and
// getHouse.
// @property setWindowType - This method sets the type of window for the house.
// @property setDoorType - This method sets the type of door to be used in the house.
// @property setNumFloor - This method sets the number of floors in the house.
// @property {House} getHouse - This method returns the house object.
type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

// "If the builderType is normal, return a new normal builder, otherwise return a new igloo builder."
// The above function is a factory function. It's a function that returns an object
func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}

	return nil
}

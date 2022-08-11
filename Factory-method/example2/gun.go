package main

// A Gun is a type that has a name and a power.
// @property {string} name - The name of the gun.
// @property {int} power - The power of the gun.
type Gun struct {
	name  string
	power int
}

// A method that sets the name of the gun.
func (g *Gun) setName(name string) {
	g.name = name
}

// A method that returns the name of the gun.
func (g *Gun) getName() string {
	return g.name
}

// A method that sets the power of the gun.
func (g *Gun) setPower(power int) {
	g.power = power
}

// A method that returns the power of the gun.
func (g *Gun) getPower() int {
	return g.power
}

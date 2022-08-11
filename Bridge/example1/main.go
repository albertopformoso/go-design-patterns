package main

import "log"

func main() {
	triangle := newTriangle(1, 2, 3)
	volumeCalculator := newVolumeCalculator(triangle)
	log.Println(volumeCalculator.Volume())

	rectangle := newRectangle(3, 3, 3)
	volumeCalculator = newVolumeCalculator(rectangle)
	log.Println(volumeCalculator.Volume())
}

// Object is an interface that has two methods, Perimeter and Height.
// @property {float64} Perimeter - The perimeter of the object
// @property {float64} Height - The height of the object.
type Object interface {
	Perimeter() float64
	Height() float64
}

// `VolumeCalculator` is a struct that has a field called `object` of type `Object`
// @property {Object} object - This is the object whose volume we want to calculate.
type VolumeCalculator struct {
	object Object
}

// `newVolumeCalculator` is a function that takes an `Object` as an argument and returns a pointer to a
// `VolumeCalculator` struct
func newVolumeCalculator(object Object) *VolumeCalculator {
	return &VolumeCalculator{object}
}

// A method that calculates the volume of the object.
func (calculator *VolumeCalculator) Volume() float64 {
	return calculator.object.Perimeter() * calculator.object.Height()
}

//// IMPLEMENTATION OF AN OBJECT: TRIANGULAR ////

// TriangularObject is a struct with three fields: height, base, and length.
// @property {float64} height - The height of the triangle
// @property {float64} base - The length of the base of the triangle.
// @property {float64} length - The length of the triangle's base
type TriangularObject struct {
	height float64
	base   float64
	length float64
}

// `newTriangle` is a function that takes three float64s as arguments and returns a pointer to a
// TriangularObject
func newTriangle(height, base, length float64) *TriangularObject {
	return &TriangularObject{height, base, length}
}

// A method that is being defined on the `TriangularObject` struct. It is a method that takes a pointer
// to a `TriangularObject` and returns a float64.
func (to *TriangularObject) Perimeter() float64 {
	return to.height * to.base / 2
}

// Defining a method on the `TriangularObject` struct. It is a method that takes a pointer
// to a `TriangularObject` and returns a float64.
func (to *TriangularObject) Height() float64 {
	return to.length
}

//// IMPLEMENTATION OF AN OBJECT: RECTANGULAR ////

// RectangularObject is a type that has three fields, x, y, and z, each of which is a float64.
// @property {float64} x - The x-coordinate of the object.
// @property {float64} y - The y-coordinate of the object.
// @property {float64} z - The height of the object
type RectangularObject struct {
	x float64
	y float64
	z float64
}

// NewRectangle returns a pointer to a RectangularObject with the given dimensions.
func newRectangle(x, y, z float64) *RectangularObject {
	return &RectangularObject{x, y, z}
}

// Defining a method on the `RectangularObject` struct. It is a method that takes a pointer
// to a `RectangularObject` and returns a float64.
func (r *RectangularObject) Perimeter() float64 {
	return r.x * r.y
}

// Defining a method on the `RectangularObject` struct. It is a method that takes a pointer
// to a `RectangularObject` and returns a float64.
func (r *RectangularObject) Height() float64 {
	return r.z
}

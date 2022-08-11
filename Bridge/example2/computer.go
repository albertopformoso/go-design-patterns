package main

// A Computer is a type that has a Print method and a SetPrinter method.
// @property Print - This is a method that prints the computer's name.
// @property SetPrinter - This is a method that takes a Printer interface as a parameter.
type Computer interface {
	Print()
	SetPrinter(Printer)
}

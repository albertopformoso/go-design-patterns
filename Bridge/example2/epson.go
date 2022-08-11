package main

import "log"

// Epson is a struct with no fields.
type Epson struct{}

// A method of the struct Epson.
func (p *Epson) PrintFile() {
	log.Println("Printing by a EPSON Printer")
}

package main

import "log"

// `Hp` is a struct with no fields.
type Hp struct{}

// A method of the struct `Hp`
func (p *Hp) PrintFile() {
	log.Println("Printing by a HP Printer")
}

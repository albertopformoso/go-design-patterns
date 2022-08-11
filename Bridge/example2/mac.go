package main

import "log"

// Mac is a struct that has a field called printer that is of type Printer.
// @property {Printer} printer - This is the interface that we will use to print the message.
type Mac struct {
	printer Printer
}

// Calling the PrintFile() method of the printer object.
func (m *Mac) Print() {
	log.Println("Print request for mac")
	m.printer.PrintFile()
}

// Setting the printer object to the printer interface.
func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

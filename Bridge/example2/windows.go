package main

import "log"

// Windows is a struct that has a field called printer that is of type Printer.
// @property {Printer} printer - This is the interface that we will use to print the message.
type Windows struct {
	printer Printer
}

// Calling the PrintFile() method of the printer struct.
func (w *Windows) Print() {
	log.Println("Print request for windows")
	w.printer.PrintFile()
}

// Setting the printer for the Windows struct.
func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

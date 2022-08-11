package main

import "log"

func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	log.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	log.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	log.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
}

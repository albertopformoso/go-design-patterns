package main

import "log"

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	log.Printf("Normal House Door Type: %q", normalHouse.doorType)
	log.Printf("Normal House Window Type: %q", normalHouse.windowType)
	log.Printf("Normal House Num Floor: %v", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	log.Printf("Igloo House Door Type: %q", iglooHouse.doorType)
	log.Printf("Igloo House Window Type: %q", iglooHouse.windowType)
	log.Printf("Igloo House Num Floor: %v", iglooHouse.floor)
}

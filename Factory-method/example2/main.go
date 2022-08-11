package main

import "log"

func main() {
	ak47, err := getGun("ak47")
	checkErr(err)
	musket, err := getGun("musket")
	checkErr(err)

	printDetails(ak47)
	printDetails(musket)

	_, err = getGun("other")
	checkErr(err)
}

// It takes a parameter of type IGun and prints the name and power of the gun
func printDetails(g IGun) {
	log.Printf("Gun: %s", g.getName())
	log.Printf("Power: %v", g.getPower())
}

// If there's an error, print it and exit the program
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

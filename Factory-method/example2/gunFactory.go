package main

import (
	"fmt"
)

// It returns an interface of type IGun and an error
func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}
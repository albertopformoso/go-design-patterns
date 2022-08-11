package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

// `single` is a struct with no fields.
type single struct{}

var singleInstance *single

// If the single instance is nil, lock the mutex, check again if the single instance is nil, and if it
// is, create it
func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

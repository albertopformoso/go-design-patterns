package main

// Printer is an interface that has a method called PrintFile.
// @property PrintFile - This is the method that will be called when the PrintFile command is executed.
type Printer interface {
	PrintFile()
}

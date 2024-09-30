package main

import (
	"fmt"
)

// Version represents the current version of the application

var Version = "0.0.4"

// PrintVersion prints the current version
func PrintVersion() {
	fmt.Println("Current Version:", Version)
}

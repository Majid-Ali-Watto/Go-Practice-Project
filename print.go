package main

import (
	"Practice_Projects/person" // Import the person package
	"fmt"
)

func printData(persons []person.Person) {
	fmt.Println("Initial persons:")
	for _, v := range persons {
		fmt.Println(v.Name, v.Age, v.ID)
	}
}

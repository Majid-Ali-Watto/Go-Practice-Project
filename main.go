package main

import (
	"Practice_Projects/person" // Import the person package
	"fmt"
)

func main() {
	makeMaps()
	// Initialize a slice of Person from the person package
	persons := []person.Person{
		{ID: 1, Name: "Majid", Age: 24},
		{ID: 2, Name: "Ali", Age: 24},
	}

	// Print details of the first two persons
	printData(persons)

	// Append third person
	persons = append(persons, person.Person{ID: 3, Name: "Watto", Age: 24})

	// Print updated slice before modification
	fmt.Println("\nPersons before Print function:")
	for _, v := range persons {
		fmt.Println(v.Name, v.Age, v.ID)
	}

	// Call the Print function from the person package to update the third person
	persons = person.Print(persons)

	// Print updated slice after modification
	fmt.Println("\nPersons after Print function:")
	for _, v := range persons {
		fmt.Println(v.Name, v.Age, v.ID)
	}
}

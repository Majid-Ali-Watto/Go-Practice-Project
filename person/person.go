package person

import "fmt"

// Person struct definition
type Person struct {
	ID   int
	Name string
	Age  int
}

// Print function to print and modify a slice of persons
func Print(p []Person) []Person {
	fmt.Println("Inside Print function:")
	fmt.Println(p) // Print the slice
	if len(p) > 2 {
		// Modify the third person (index 2)
		p[2].Age = 12
		p[2].Name = "Ali"
		p[2].ID = 123
	}
	return p
}

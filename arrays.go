package main

import (
	"Practice_Projects/person" // Import the person package
	"fmt"
)

func makeArray() {
	arr := [10]string{} // array of 10 elements
	arr[0] = "Majid"
	arr[1] = "Ali"
	arr[2] = "Watto"
	arr[3] = "Majid"
	arr[4] = "Ali"
	arr[5] = "Watto"
	arr[6] = "Majid"
	arr[7] = "Ali"
	arr[8] = "Watto"
	arr[9] = "Majid"
	// arr[10] = "Majid" //invalid argument: index 10 out of bounds [0:10]
	fmt.Println("Print array without iteration")
	fmt.Println(arr)
	// Or
	fmt.Println("Print array with iteration")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("Print array without iteration range")

	// or
	for i, v := range arr {
		fmt.Printf("%s at %d\n", v, i)
	}

	arr[5] = "Bahawalnagar"
	// arr[5] is replaced with Bahawalnagar
	fmt.Println("Print array without iteration")
	fmt.Println(arr)

	persons := [5]person.Person{}
	persons[0] = person.Person{
		ID:   1,
		Name: "Bahawalnagar",
		Age:  1900,
	}
	persons[4] = person.Person{
		ID:   1,
		Name: "Ali",
		Age:  19,
	}
	fmt.Println(persons)
	fmt.Println(persons[4])
	// fmt.Println(persons[5]) // invalid argument: index 5 out of bounds [0:5]

}

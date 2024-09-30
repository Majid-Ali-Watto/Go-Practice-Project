package main

import "fmt"

func makeArray() {
	arr := [10]string{}
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
}

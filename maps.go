package main

import "fmt"

func singleMap(name string) map[string]string {
	persons := map[string]string{}
	persons["name"] = name
	fmt.Println(persons)
	return persons
}
func makeMaps() {
	maps := []map[string]string{}
	// maps := make(map[string]string,5)
	maps = append(maps, singleMap("Majid"))
	maps = append(maps, singleMap("Ali"))
	fmt.Println(maps)
}

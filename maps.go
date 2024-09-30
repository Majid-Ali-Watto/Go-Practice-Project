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
	maps = append(maps, singleMap("Majid"))
	maps = append(maps, singleMap("Ali"))
	fmt.Println(maps)
}

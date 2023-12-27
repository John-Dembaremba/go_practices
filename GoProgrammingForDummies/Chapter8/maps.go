package main

import (
	"fmt"
)

func main() {

	heights := make(map[string]int)

	heights["John"] = 12
	heights["Josh"] = 13
	heights["Joe"] = 5
	fmt.Print("\n######Maps######\n")
	fmt.Printf("John: %v\n", heights["John"])
	fmt.Printf("Josh: %v\n", heights["Josh"])
	fmt.Printf("Joe: %v\n", heights["Joe"])

	fmt.Print("\n######Maps Initialized with values ######\n")
	ages := initMapsWithValues()
	fmt.Println(ages)

	// if given key doesn`t exists, the value o is returned
	fmt.Printf("\n `Key` doesn`t exists , value: %v is returned\n", ages["Key"])
	fmt.Println("Precision use a function")
	exists := keyExists(heights, "Key")
	fmt.Printf("`Key` doesn`t exists , value: %v is returned\n", exists)

	fmt.Print("\n######Delete keys ######\n")
	del := deleteKey(heights, "John")
	fmt.Printf("Deleted `John`, results: %v\n", del)
	fmt.Print(heights)

}

func initMapsWithValues() map[string]int {
	ages := map[string]int{
		"John": 26,
		"Josh": 25,
		"Joe":  6,
	}
	return ages
}

func keyExists(m map[string]int, key string) bool {
	if v, ok := m[key]; ok {
		fmt.Println(v)
		return true
	} else {
		fmt.Println("Key does not exist")
		return false
	}
}

func deleteKey(m map[string]int, key string) string {
	exists := keyExists(m, key)

	if exists {
		delete(m, key)
		success := "key deleted successfully."
		return success
	} else {
		failure := "Key does not exist."
		return failure
	}
}

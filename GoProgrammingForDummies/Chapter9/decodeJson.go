package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name    string
	Surname string
}

func main() {

	// jsonString := `{
	// 	"name": "John",
	// 	"surname": "Dembaremba"
	// }`
	// decodeJson(jsonString)

	people := People{
		Name:    "John",
		Surname: "Dembaremba",
	}

	encodeJson(people)

}

// func decodeJson(jsonData string) {
// 	var person People
// 	decoder := json.Unmarshal([]byte(jsonData), &person)

// 	if decoder == nil {
// 		fmt.Println(person.Name)
// 		fmt.Println(person.Surname)
// 	} else {
// 		fmt.Println(decoder)
// 	}
// }

func encodeJson(s People) {
	jsonData, val := json.Marshal(s)
	fmt.Printf("Json: %v\n", string(jsonData))
	fmt.Printf("Val: %v\n", val)
}

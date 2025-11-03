package jsons

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Skills []string `json:"skills"`
}

// readJson reads the json data from json file
func readJson() {
	file, err := os.Open("./jsons/data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var person Person
	if err := json.NewDecoder(file).Decode(&person); err != nil {
		panic(err)
	}
	fmt.Println(person)
}

// writeJson writes data to json files
func writeJson() {
	person := Person{
		Name:   "Bob",
		Age:    30,
		Skills: []string{"Golang", "AWS"},
	}

	file, err := os.Create("./jsons/output.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(person); err != nil {
		panic(err)
	}
}

func RunJsonDemo() {
	readJson()
	writeJson()
}

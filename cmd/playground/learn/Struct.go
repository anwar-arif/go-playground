package learn

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Employee struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Salary  int    `json:"salary"`
	Address Address
}

func Test(emp Employee) {
	emp.Name = "Arif"
	return
}

func StructRun() {
	fmt.Println("Running StructRun()...")
	emp := Employee{
		Name:   "Anwar",
		Age:    28,
		Salary: 1000,
		Address: Address{
			City:    "Dhaka",
			Country: "Bangladesh",
		},
	}
	empJson, err := json.Marshal(emp)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(empJson))

	empJSON, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(string(empJSON))

	Test(emp)

	empJSON1, err := json.MarshalIndent(emp, "", " ")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(empJSON1))
}

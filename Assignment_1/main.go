package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	name           string
	address        string
	job            string
	reason         string
	absence_number int
}

func main() {
	number, _ := strconv.Atoi(os.Args[1])

	var peoples = []Person{
		{name: "Vina", address: "Mawar Street", job: "Designer", reason: "Change career", absence_number: 1},
		{name: "Sera", address: "Manggrove Complex", job: "Engineer", reason: "Portfolio", absence_number: 2},
		{name: "Glen", address: "High Mountain Street", job: "Marketing", reason: "Change career to one with better payment", absence_number: 3},
		{name: "John", address: "Pretty Beach", job: "Doctor", reason: "want to work from home", absence_number: 4},
	}

	find(peoples, number)

	os.Exit(3)
}

func find(list_of_person []Person, absence_num int) {
	for i := range list_of_person {
		if list_of_person[i].absence_number == absence_num {
			printData(list_of_person[i])
			return
		}
	}
	fmt.Println("Not found a person with absence number", absence_num)
}

func printData(p Person) {
	fmt.Println("Name:", p.name)
	fmt.Println("Address:", p.address)
	fmt.Println("Job:", p.job)
	fmt.Println("Why choose Golang class:", p.reason)
}

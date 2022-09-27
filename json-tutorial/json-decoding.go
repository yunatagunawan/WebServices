package main

import (
	"fmt"
	"encoding/json"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Age int `json:"age"`
}

func main() {
	var jsonString = `
		{
			"full_name":"Airell Jordan",
			"email": "airell@gmail.com",
			"age": 23
		}
	`

	var result Employee 
	// or var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name:", result.FullName)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)

	/* to map */

	var resMap map[string]interface{}

	// var temp interface{}

	err = json.Unmarshal([]byte(jsonString), &resMap)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// resMap = temp.(map[string]interface{})

	fmt.Println("full_name:", resMap["full_name"])
	fmt.Println("email:", resMap["email"])
	fmt.Println("age:", resMap["age"])

}
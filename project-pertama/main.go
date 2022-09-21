package main

import (
	"fmt"
	"strings"
	// "math"
)

func main() {

	var currYear int = 2022

	if age := currYear - 2001; age > 18 {
		fmt.Println("You are a young adult")
	} else {
		fmt.Println("Other")
	}

	if currYear > 2020 {
		fmt.Println("New Era!")
	} else {
		fmt.Println("Other")
	}

	var array [4]int = [4]int{1, 2, 3, 4}
	// var array [2][3]int = [2][3]int{{5, 6, 7}, 1, 2, 3}

	for i := 0; i < 4; i++ {
		fmt.Println(array[i])
	}

	fmt.Println(strings.Repeat("$", 20))

	fmt.Println(array)

	var dict map[int]string = map[int]string{
		2: "dua",
		3: "dua",
		4: "dua",
	}

	value, exist := dict[3]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("not exist")
	}

	dict[1] = "satu"

	fmt.Println(dict)

	delete(dict, 1)

	fmt.Println(dict)

	quitFunc()
}

func return_mmultiple(d string) (string, string) {
	return d, "names"
}

func quitFunc() {
	fmt.Println("123")
	return
	fmt.Println("1234567890")
}

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Airell", Age: 23, Division: "IT"},
	{ID: 2, Name: "Nanda", Age: 23, Division: "Finance"},
	{ID: 3, Name: "Mailo", Age: 20, Division: "IT"},
}

var PORT = ":8080"

func main() {
	// Register an function to be APIs
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employeesV2", getEmployeesV2)

	http.HandleFunc("/employee", createEmployee)

	fmt.Println("App is listen on port", PORT)
	http.ListenAndServe(PORT, nil)
}

// POST API
func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

// GET API
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func getEmployeesV2(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employees)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

// Simple API
func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world!"
	fmt.Fprint(w, msg)
}

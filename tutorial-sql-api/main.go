package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"tutorial-sql-api/database"
	"tutorial-sql-api/models"
	"gorm.io/gorm"
	"errors"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "admin"
	dbname = "db-go-sql"
)

var (
	db *sql.DB
	err error
)

type Employee struct {
	ID int
	Full_name string
	Email string
	Age int
	Division string
}

func CreateEmployee() {
	var employee = Employee{}
	sql_query := `INSERT INTO employees (full_name, email, age, division) 
	VALUES ($1, $2, $3, $4) 
	Returning *`

	err = db.QueryRow(sql_query, "Airell Jordan", "airell@email.com", 23, "IT").Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func GetEmployees() {
	var results = []Employee{}

	sql_query := "SELECT * from employees"

	rows, err := db.Query(sql_query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)
	
		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employee datas:", results)
}

func UpdateEmployee() {
	sql_query := `UPDATE employees
	SET full_name = $2, email = $3, division = $4, age = $5
	WHERE id = $1;`

	res, err := db.Exec(sql_query, 1, "Airell Jordan Hidayat", "airellhidayat@gmail.com", "CurDevs", 24)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount:", count)
}

func DeleteEmployee() {

	sql_query := `
	DELETE from employees
	WHERE id = $1;`

	res, err := db.Exec(sql_query, 1)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data amount:", count)
}

func main() {

	/*

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// CreateEmployee()
	// GetEmployees()
	// UpdateEmployee()

	DeleteEmployee()

	*/

	// PART 2
	database.StartDB()

	// createUser("johndoe@gmail.com")
	// getUserById(1)
	// updateUserById(1, "johnnewemail@gmail.com")

	// createProduct(1, "YLO", "SmartTV")
	// getUsersWithProducts()
	deleteProductById(1)
}

func createUser(email string){
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error
	
	if err != nil {
		fmt.Println("error creating user data:", err)
		return
	}

	fmt.Println("New User Data:", User)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			fmt.Println("User data not found")
			return
		}
		fmt.Println("Error finding user:", err)
	}

	fmt.Printf("User Data: %+v \n", user)
}

func updateUserById(id uint, email string){
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updateing user data:", err)
		return
	}
	fmt.Printf("Update users's email : %+v \n", user.Email)

	// db.Table("users").Where("id = ?", 1).Updates(map[string]interface{}{"email":"newemail@gmail.com"})
}

func createProduct(userId uint, brand string, name string){
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand: brand,
		Name: name,
	}

	err := db.Create(&Product).Error
	
	if err != nil {
		fmt.Println("error creating product data:", err.Error())
		return
	}

	fmt.Println("New User Data:", Product)
}

// join statement
func getUsersWithProducts() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products:", err.Error())
		return
	}

	fmt.Println("User datas with products")
	fmt.Printf("%+v", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted", id)
}
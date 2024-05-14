package main

import (
	"employee/internal/controller"
	"employee/internal/service"
	"employee/internal/sql_data"
	"fmt"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := getDb()

	storer := sql_data.NewEmployeeStore(db)
	service := service.NewEmployeeService(storer)
	ctrl := controller.NewEmployeeController(service)

	ctrl.ProcessData()

	e.POST("/v1/employee", ctrl.EmployeeData)
	e.GET("/v1/employee", ctrl.GetEmployees)

	// Run the echo server
	e.Logger.Fatal(e.Start(":8000"))

}

func getDb() *gorm.DB {
	var db *gorm.DB
	var err error

	// Wait for MySQL to become available
	for {
		db, err = gorm.Open(mysql.Open("root:root123@tcp(localhost:3306)/employeedb"), &gorm.Config{})
		if err == nil {
			break // MySQL is available, break the loop
		}

		fmt.Println("MySQL is not available yet. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	// MySQL is now available, start your application logic here
	fmt.Println("MySQL is now available. Starting the application...")
	return db
}

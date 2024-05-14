package controller

import (
	"context"
	"employee/internal/model"
	"employee/internal/service"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	EmployeeData(c echo.Context) error
	GetEmployees(c echo.Context) error
}

type EmployeeController struct {
	Service service.EmployeeService
}

func NewEmployeeController(employeeService service.EmployeeService) Controller {
	return EmployeeController{
		Service: employeeService,
	}
}

func (e EmployeeController) EmployeeData(c echo.Context) error {

	// parse JSON request body
	var employees []model.Employee
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = json.Unmarshal(b, &employees)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Create temp file
	f, err := os.CreateTemp("", "employee_data")
	fileName := f.Name()
	defer f.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	defer os.Remove(fileName)

	// Write employee data to the file
	err = writeEmployeesToFile(fileName, employees)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err

	}

	go e.Service.CreateEmployee(context.Background(), fileName)

	c.JSON(http.StatusOK, gin.H{"message": "Employees received successfully"})
	return nil

}
func (e EmployeeController) GetEmployees(c echo.Context) error {

	resp, err := e.Service.GetEmployee(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func writeEmployeesToFile(fileName string, employees []model.Employee) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	for _, employee := range employees {
		if err := encoder.Encode(employee); err != nil {
			return err
		}
	}
	return nil
}

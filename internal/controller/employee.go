package controller

import (
	"employee/internal/service"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	EmployeeData(c echo.Context) error
	GetEmployees(c echo.Context) error
	ProcessData() error
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
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Create temp directory
	os.Mkdir("./tmp", 0777)

	fileName := fmt.Sprintf("./tmp/employee_data_%d.json", rand.Intn(1000000))
	f, err := os.Create(fileName)
	defer f.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	// Write employee data to the file
	err = writeEmployeesToFile(f.Name(), b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err

	}

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

func writeEmployeesToFile(fileName string, jsonData []byte) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	os.WriteFile(fileName, jsonData, 0644)
	return nil
}

func (e EmployeeController) ProcessData() error {
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {
			files, err := getAllFiles("./tmp/")
			if err != nil {
				return
			}
			var wg sync.WaitGroup
			wg.Add(len(files))
			for _, file := range files {
				go e.Service.ProcessFile(file, &wg)
			}
		}
	}()
	return nil
}

func getAllFiles(dirPath string) ([]string, error) {
	var fileList []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList, err
}

package service

import (
	"bufio"
	"bytes"
	"context"
	"employee/internal/model"
	"employee/internal/sql_data"
	"encoding/json"
	"log"
	"os"
	"sync"
)

type EmployeeService interface {
	GetEmployee(ctx context.Context) ([]model.Employee, error)
	CreateEmployee(ctx context.Context, fileName string) error
}

type employeeService struct {
	Data sql_data.EmployeeStorer
}

func NewEmployeeService(storer sql_data.EmployeeStorer) EmployeeService {
	return &employeeService{Data: storer}
}

func (e employeeService) GetEmployee(ctx context.Context) ([]model.Employee, error) {
	m, err := e.Data.Get(ctx)
	if err != nil {
		log.Println("error in getting employee", err)
		return []model.Employee{}, err
	}
	return m, nil
}

func (e employeeService) CreateEmployee(ctx context.Context, fileName string) error {
	// Open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			return i + 1, data[:i], nil
		}
		// If no newline found, keep scanning
		return 0, nil, nil
	})

	// Use a wait group to track goroutines
	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()

			// Parse employee data from line
			var employee model.Employee
			if err := json.Unmarshal([]byte(line), &employee); err != nil {
				// Handle parsing error (log or ignore)
				return
			}

			// Store employee data in database
			if err := e.Data.Create(ctx, &employee); err != nil {
				// Handle database error (log or retry)
				return
			}
		}(scanner.Text())
	}

	wg.Wait()
	return nil
}

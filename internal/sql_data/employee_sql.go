package sql_data

import (
	"context"
	"employee/internal/model"
	"time"

	"gorm.io/gorm"
)

type EmployeeStorer interface {
	Get(ctx context.Context) ([]model.Employee, error)
	Create(ctx context.Context, employee *model.Employee) error
}

type employeeStore struct {
	db *gorm.DB
}

func (e employeeStore) Get(ctx context.Context) ([]model.Employee, error) {
	emp := []model.Employee{}
	result := e.db.Table("employee").Select("*").Scan(&emp)
	if result.Error != nil {
		return []model.Employee{}, result.Error
	}
	return emp, nil
}

func (e employeeStore) Create(ctx context.Context, employee *model.Employee) error {
	sqlQuery := "INSERT INTO employee (name, designation,salary,insurance_id,insurance_amount,location, created_at, updated_at) VALUES (?, ?, ?, ?,?,?,?,?)"
	// Execute the raw SQL query with parameters

	result := e.db.Exec(sqlQuery, employee.Name, employee.Designation, employee.Salary, employee.InsuranceID, employee.InsuranceAmt, employee.Location, time.Now(), time.Now())
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewEmployeeStore(db *gorm.DB) EmployeeStorer {
	return &employeeStore{
		db: db,
	}
}

package model

type Employee struct {
	ID           int     `gorm:"primaryKey" json:"id"`
	Name         string  `json:"name"`
	Designation  string  `json:"designation"`
	Salary       float64 `json:"salary"`
	Location     string  `json:"location"`
	InsuranceID  int     `json:"insurance_id"`
	InsuranceAmt float64 `json:"insurance_amount"`
}

func (e Employee) TableName() string {
	return "employee"
}

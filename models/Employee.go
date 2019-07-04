package models

import (
	"time"
)

const (
	worker   = "WORKER"
	engineer = "ENGINEER"
)

// Employee represents general data, related to all employees stored in database
type Employee struct {
	EmployeeID   int64        `json:"employee_id"  		gorm:"Column:employee_id;PRIMARY_KEY"`
	Name         string       `json:"name"         		gorm:"Column:name"`
	LastName     string       `json:"last_name"			gorm:"Column:last_name"`
	MiddleName   string       `json:"middle_name" 		gorm:"Column:middle_name"`
	Position     string       `json:"position" 			gorm:"Column:position"`
	IsQuit       bool         `json:"is_quit" 			gorm:"Column:is_quit"`
	EmployeeInfo EmployeeInfo `json:"employee_info"`
}

// EmployeeInfo contains general info about employee
type EmployeeInfo struct {
	EmployeeInfoID    int64           `json:"employee_info_id"  	gorm:"Column:employee_info_id;PRIMARY_KEY"`
	EmployeeID        int64           `json:"employee_id"  			gorm:"Column:employee_id;FOREIGN_KEY"`
	EmployeeAddress   EmployeeAddress `json:"employee_address"`
	PassSeriesNum     string          `json:"pass_series_num" 		gorm:"Column:pass_series_num"`
	IdentificationNum string          `json:"identification_num"	gorm:"Column:identification_num"`
	BirthDate         time.Time       `json:"birth_date" 			gorm:"Column:birth_date"`
	JoinDate          time.Time       `json:"join_date" 			gorm:"Column:join_date"`
	QuitDate          time.Time       `json:"quit_date" 			gorm:"Column:quit_date"`
}

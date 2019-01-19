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
	EmployeeId        int64     `json:"employee_id"  		gorm:"Column:employee_id;PRIMARY_KEY"`
	Name              string    `json:"name"         		gorm:"Column:name"`
	LastName          string    `json:"last_name"			gorm:"Column:last_name"`
	MiddleName        string    `json:"middle_name" 		gorm:"Column:middle_name"`
	PassSeriesNum     string    `json:"pass_series_num" 	gorm:"Column:pass_series_num"`
	IdentificationNum string    `json:"identification_num"	gorm:"Column:identification_num"`
	BirthDate         time.Time `json:"birth_date" 			gorm:"Column:birth_date"`
	Position          string    `json:"position" 			gorm:"Column:position"`
	JoinDate          time.Time `json:"join_date" 			gorm:"Column:join_date"`
	QuitDate          time.Time `json:"quit_date" 			gorm:"Column:quit_date"`
	IsQuit            bool      `json:"is_quit" 			gorm:"Column:is_quit"`
}

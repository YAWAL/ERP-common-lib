package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	worker   = "WORKER"
	engineer = "ENGINEER"
)

// Employee represents general data, related to all employees stored in database
type Employee struct {
	ID           uuid.UUID    `gorm:"Column:employee_id;PRIMARY_KEY"`
	Name         string       `json:"name,omitempty"         		gorm:"Column:name"`
	LastName     string       `json:"last_name,omitempty"			gorm:"Column:last_name"`
	MiddleName   string       `json:"middle_name,omitempty" 		gorm:"Column:middle_name"`
	Position     string       `json:"position,omitempty" 			gorm:"Column:position"`
	IsQuit       bool         `json:"is_quit,omitempty" 			gorm:"Column:is_quit"`
	EmployeeInfo EmployeeInfo `json:"employee_info,omitempty"`
}

// EmployeeInfo contains general info about employee
type EmployeeInfo struct {
	ID                uuid.UUID       `gorm:"Column:employee_info_id;PRIMARY_KEY"`
	EmployeeAddress   EmployeeAddress `json:"employee_address,omitempty"`
	PassSeriesNum     string          `json:"pass_series_num,omitempty" 	gorm:"Column:pass_series_num"`
	IdentificationNum string          `json:"identification_num,omitempty"	gorm:"Column:identification_num"`
	BirthDate         time.Time       `json:"birth_date,omitempty" 			gorm:"Column:birth_date"`
	JoinDate          time.Time       `json:"join_date,omitempty" 			gorm:"Column:join_date"`
	QuitDate          time.Time       `json:"quit_date,omitempty" 			gorm:"Column:quit_date"`
}

// EmployeeAddress represents data about employees addresses stored in database
type EmployeeAddress struct {
	ID             uuid.UUID `gorm:"address_id;PRIMARY_KEY"`
	PhoneNum       string    `json:"phone_num,omitempty" gorm:"phone_num"`
	ResidencePlace string    `json:"residence_place,omitempty" gorm:"residence_place"`
	Street         string    `json:"street,omitempty" gorm:"street"`
	BuildingNum    string    `json:"building_num,omitempty" gorm:"building_num"`
	FlatNum        string    `json:"flat_num,omitempty" gorm:"flat_num"`
	Zip            string    `json:"zip,omitempty" gorm:"zip"`
}

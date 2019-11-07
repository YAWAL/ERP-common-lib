package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	worker   = "WORKER"
	engineer = "ENGINEER"
)

// Employee represents general data, related to all employees stored in database
type Employee struct {
	ID              primitive.ObjectID `json:"id"                         bson:"_id"`
	Name            string             `json:"name,omitempty"             bson:"name"`
	LastName        string             `json:"last_name,omitempty"        bson:"last_name"`
	MiddleName      string             `json:"middle_name,omitempty"      bson:"middle_name"`
	IsQuit          bool               `json:"is_quit,omitempty"          bson:"is_quit"`
	EmployeeInfo    EmployeeInfo       `json:"employee_info,omitempty"    bson:"employee_info"`
	EmployeeAddress []EmployeeAddress  `json:"employee_address,omitempty" bson:"employee_address"`
}

// EmployeeInfo contains general info about employee
type EmployeeInfo struct {
	Position          string    `json:"position,omitempty"           bson:"position"`
	PassSeriesNum     string    `json:"pass_series_num,omitempty"    bson:"pass_series_num"`
	IdentificationNum string    `json:"identification_num,omitempty" bson:"identification_num"`
	PhoneNums         []string  `json:"phone_nums,omitempty"         bson:"phone_nums"`
	BirthDate         time.Time `json:"birth_date,omitempty"         bson:"birth_date"`
	JoinDate          time.Time `json:"join_date,omitempty"          bson:"join_date"`
	QuitDate          time.Time `json:"quit_date,omitempty"          bson:"quit_date"`
}

// EmployeeAddress represents data about employees addresses stored in database
type EmployeeAddress struct {
	ResidencePlace string `json:"residence_place,omitempty" bson:"residence_place"`
	Street         string `json:"street,omitempty"          bson:"street"`
	BuildingNum    string `json:"building_num,omitempty"    bson:"building_num"`
	FlatNum        string `json:"flat_num,omitempty"        bson:"flat_num"`
	Zip            string `json:"zip,omitempty"             bson:"zip"`
}

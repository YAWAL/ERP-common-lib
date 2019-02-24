package models

// EmployeeAddress represents data about employees addresses stored in database
type EmployeeAddress struct {
	AddressID      int    `json:"address_id"`
	EmployeeID     int64  `json:"employee_id"`
	PhoneNum       string `json:"phone_num"`
	ResidencePlace string `json:"residence_place"`
	Street         string `json:"street"`
	BuildingNum    string `json:"building_num"`
	FlatNum        string `json:"flat_num"`
	Zip            string `json:"zip"`
}

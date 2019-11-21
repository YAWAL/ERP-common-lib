package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EquipmentType represents is this equipment for production use or non production
type EquipmentType int

const (
	// Production = 1
	Production EquipmentType = iota + 1
	// Nonproduction = 2
	Nonproduction
)

// MachinesGroup represents group for machine
type MachinesGroup int

const (
	// Turning
	First MachinesGroup = iota + 1
	Second
	Third
	Fourth
	Fifth
	Sixth
	Seventh
	Eighth
	Nineth
)

// MachinesType represents type in machine's group
type MachinesType int

const (
	Zero MachinesType = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

// Currency represents money currency for equipment's price
type Currency int

const (
	// USD - United States dollar
	USD = iota + 1
	// EUR - Euro
	EUR
	// UAH -Ukrainian hryvnia
	UAH
	// PLN - Polish zloty
	PLN
)

// Equipment represents general schema for production and non production equipment
type Equipment struct {
	ID              primitive.ObjectID `json:"_id,omitempty"              bson:"_id"`
	Type            EquipmentType
	InventoryNumber string
	StartDate       time.Time
	ProductionYear  int
	Manufacturer    string // or should be this manufacturer ID?
	Vendor          string // or should be this Vendor ID?
}

type PhysicalCharacteristic struct {
	Power  int
	Weight int
}

// Size holds information about equipment's sizes (in millimiters)
type Size struct {
	Length float32
	Height float32
	Width  float32
}

// Price represent information about equipment's cost and currency
type Price struct {
	Value    int
	Currency Currency
}

package models

// Workshop represents data about workshops stored in database
type Workshop struct {
	WorkshopId   int
	WorkshopName string
	ChairmanId   int64
}

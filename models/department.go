package models

// Department represents data about departments stored in database
type Department struct {
	DepartmentId   int
	DepartmentName string
	ChairmanId     int64
}

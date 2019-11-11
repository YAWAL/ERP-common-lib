package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// dialect for Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const postgresDBstring = "user=%s dbname=%s sslmode=%s password=%s"

// Postgres creates client to Postgres DB with provided in Postgres database config file's part
func Postgres(conf PostgresConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open(conf.Dialect, fmt.Sprintf(postgresDBstring, conf.User,
		conf.DataBaseName, conf.SSLMode, conf.Password))
	if err != nil {
		return nil, err
	}
	return db, nil
}

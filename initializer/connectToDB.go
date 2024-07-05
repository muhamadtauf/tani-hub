package initializer

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tani-hub/helper"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	//dsn := os.Getenv("DB")
	//DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	panic("Failed to connect to DB")
	//}

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbName   = "test"
	)

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	DB, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)
}

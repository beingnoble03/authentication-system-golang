package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DATABASE_CONFIG")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to Database. Stoped the execution.")
	}
}

package database

import (
	"emailn/internal/domain/campaign"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=emailn_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	//? migrate all entities table
	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})
	return db
}

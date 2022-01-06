package database

import (
	"fmt"
	"log"

	"inkass/inkassback/models"
	sm "inkass/inkassback/models/site-models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	connection, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Invalid database url")
	}
	sql, err := connection.DB()

	err = sql.Ping()
	if err != nil {
		log.Fatal("Database connected")
	}
	fmt.Println("Database connection successuful.")
	return connection
}

func InitialMigration() {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	_ = connection.AutoMigrate(&sm.HomeScreen{})
	_ = connection.AutoMigrate(&sm.AboutUs{})
	_ = connection.AutoMigrate(&sm.Contacts{})
	_ = connection.AutoMigrate(&sm.CurrencyRates{})
	_ = connection.AutoMigrate(&sm.Links{})
	_ = connection.AutoMigrate(&sm.News{})
	_ = connection.AutoMigrate(&sm.SingleNew{})
	_ = connection.AutoMigrate(&sm.Service{})
	_ = connection.AutoMigrate(&sm.SingleService{})
	_ = connection.AutoMigrate(&sm.ContactScreen{})
	_ = connection.AutoMigrate(&sm.Regions{})
	_ = connection.AutoMigrate(&sm.Branch{})
	_ = connection.AutoMigrate(&sm.Management{})
	_ = connection.AutoMigrate(&sm.Biography{})
	_ = connection.AutoMigrate(&sm.WorkTime{})
	_ = connection.AutoMigrate(&sm.Vacancies{})
	_ = connection.AutoMigrate(&sm.CityVacancies{})
	_ = connection.AutoMigrate(&models.Reviews{})

}

//closes database connection
func CloseDatabase(connection *gorm.DB) {
	sqldb, _ := connection.DB()
	sqldb.Close()
}

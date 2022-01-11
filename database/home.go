package database

import sm "inkass/inkassback/models/site-models"

func GetHomeScreen() sm.HomeScreen {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	var home sm.HomeScreen
	connection.Preload("CurrencyRates").Preload("Links").First(&home)
	return home
}

func CreateHomeScreen(home sm.HomeScreen) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&home)
}

func EditHomeScreen(home sm.HomeScreen) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Save(&home)
}

func DeleteHomeScreen(id uint) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Delete(&sm.HomeScreen{Id: id})
}

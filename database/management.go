package database

import sm "inkass/inkassback/models/site-models"

func GetManagement() []sm.Management {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	var management []sm.Management
	connection.Preload("Biography").Preload("WorkTime").Find(&management)
	return management
}

func CreateManagement(management sm.Management) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&management)
}

func EditManagement(management sm.Management) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Save(&management)
}

func DeleteManagement(id uint) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Delete(&sm.Management{Id: id})
}

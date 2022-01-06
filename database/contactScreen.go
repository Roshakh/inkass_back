package database

import sm "inkass/inkassback/models/site-models"

func GetContactScreen() []sm.ContactScreen {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	var contactScreen []sm.ContactScreen
	connection.Preload("Regions").Preload("Regions.Branch").Find(&contactScreen)
	return contactScreen
}

func CreateContactScreen(contactScreen sm.ContactScreen) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&contactScreen)
}

func EditContactScreen(contactScreen sm.ContactScreen) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Save(&contactScreen)
}

func DeleteContactScreen(id uint) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Delete(&sm.ContactScreen{Id: id})
}

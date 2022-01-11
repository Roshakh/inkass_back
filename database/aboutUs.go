package database

import sm "inkass/inkassback/models/site-models"

func GetAboutUs() sm.AboutUs {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	var aboutUs sm.AboutUs
	connection.Preload("Review").First(&aboutUs)
	return aboutUs
}

func CreateAboutUs(aboutUs sm.AboutUs) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&aboutUs)
}

func EditAboutUs(aboutUs sm.AboutUs) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Save(&aboutUs)
}

func DeleteAboutUs(id uint) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Delete(&sm.AboutUs{Id: id})
}

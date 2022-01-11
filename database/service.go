package database

import sm "inkass/inkassback/models/site-models"

func GetService() sm.Service {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	var service sm.Service
	connection.Preload("ServiceCatalog").Preload("SingleService").First(&service)
	return service
}

func CreateService(service sm.Service) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&service)
}

func EditService(service sm.Service) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Save(&service)
}

func DeleteService(id uint) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Delete(&sm.Service{Id: id})
}

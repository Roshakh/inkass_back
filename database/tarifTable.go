package database

import sm "inkass/inkassback/models/site-models"

func GetTarifTable() []sm.TarifTable {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	var tarifTable []sm.TarifTable
	connection.Find(&tarifTable)
	return tarifTable
}

func CreateTarifTable(tarifTable sm.TarifTable) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&tarifTable)
}

func EditTarifTable(tarifTable sm.TarifTable) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Save(&tarifTable)
}

func DeleteTarifTable(id uint) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Delete(&sm.TarifTable{Id: id})
}

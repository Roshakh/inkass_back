package database

import sm "inkass/inkassback/models/site-models"

func GetVacancies() []sm.Vacancies {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	var vacancies []sm.Vacancies
	connection.Preload("CityVacancies").Find(&vacancies)
	return vacancies
}

func CreateVacancies(vacancies sm.Vacancies) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&vacancies)
}

func EditVacancies(vacancies sm.Vacancies) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&vacancies)
}

func DeleteVacancies(id uint) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&sm.Vacancies{Id: id})
}

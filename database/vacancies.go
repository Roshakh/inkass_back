package database

import sm "inkass/inkassback/models/site-models"

func GetVacancies() sm.Vacancies {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	var vacancies sm.Vacancies
	connection.Preload("CityVacancies").Preload("CityVacancies.SimpleVacancies").First(&vacancies)
	return vacancies
}

func CreateVacancies(vacancies sm.Vacancies) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&vacancies)
}

func EditVacancies(vacancies *sm.Vacancies) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Save(vacancies)
}

func DeleteVacancies(id uint) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Delete(&sm.Vacancies{Id: id})
}

package database

import sm "inkass/inkassback/models/site-models"

func GetNews() []sm.News {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	var news []sm.News
	connection.Preload("SingleNew").Find(&news)
	return news
}

func CreateNews(news sm.News) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&news)
}

func EditNews(news sm.News) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Save(&news)
}

func DeleteNews(id uint) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Delete(&sm.News{Id: id})
}

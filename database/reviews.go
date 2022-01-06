package database

import "inkass/inkassback/models"

func GetReviews() []models.Reviews {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	var reviews []models.Reviews
	connection.Find(&reviews)
	return reviews
}

func CreateReviews(reviews models.Reviews) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Create(&reviews)
}

func EditReviews(reviews models.Reviews) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Save(&reviews)
}

func DeleteReviews(id uint) {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.Delete(&models.Reviews{Id: id})
}

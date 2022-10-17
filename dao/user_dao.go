package dao

import (
	"gateway-test/models"
)

func Init() {
	DB.AutoMigrate(&models.User{})
}

func CreateUser(user *models.User) (err error) {
	err = DB.Create(user).Error

	return
}

func GetUsers() (users []*models.User, err error) {
	if err = DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return
}

package auth

import (
	"errors"
	"example/v3/db"
	"example/v3/models"
	"example/v3/utils"
)

func AuthenticateUser(login, password string) (*models.User, error) {
	db, err := db.OpenDBConnection()
	if err != nil {
		return nil, err
	}

	var user *models.User

	if err := db.Where("login = ?", login).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if user != nil && !utils.CheckPasswordHash(password, user.PassHash) {
		return nil, errors.New("wrong password")
	}

	return user, nil
}

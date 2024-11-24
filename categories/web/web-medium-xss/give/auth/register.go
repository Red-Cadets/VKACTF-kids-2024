package auth

import (
	"errors"
	"example/v3/db"
	"example/v3/models"
	"example/v3/utils"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func SignUpUser(user *models.User) error {
	db, err := db.OpenDBConnection()
	if err != nil {
		return err
	}

	fmt.Println("DB connection established...")

	var existingUser *models.User
	if err := db.Where("login = ?", user.Login).First(&existingUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return fmt.Errorf("error checking existing user: %v", err)
		}
	} else {
		return errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.PassHash = hashedPassword
	user.Password = ""

	if err := db.Create(&user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return errors.New("user with this login already exists")
		}
		return fmt.Errorf("error for registration: %v", err)
	}

	return nil
}

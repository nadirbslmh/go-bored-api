package services

import (
	"go-bored-api/database"
	"go-bored-api/middlewares"
	"go-bored-api/models"

	"golang.org/x/crypto/bcrypt"
)

func Register(input models.User) (models.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(password),
	}

	result := database.DB.Create(&user)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	err = result.Last(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func Login(input models.User, jwtAuth *middlewares.JWTConfig) (string, error) {
	var user models.User

	err := database.DB.First(&user, "email = ?", input.Email).Error

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return "", err
	}

	token, err := jwtAuth.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetAll() ([]models.User, error) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		return []models.User{}, err
	}

	return users, nil
}

func Delete(id string) error {
	var deletedUser models.User

	if err := database.DB.Find(&deletedUser, "id = ?", id).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&deletedUser).Error; err != nil {
		return err
	}

	return nil
}

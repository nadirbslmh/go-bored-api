package controllers

import (
	"go-bored-api/middlewares"
	"go-bored-api/models"
	"go-bored-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var userInput models.User

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failed",
			"message": "invalid request",
		})
	}

	user, err := services.Register(userInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failed",
			"message": "invalid request",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status":  "success",
		"message": "user registered",
		"data":    user,
	})
}

func Login(c echo.Context) error {
	var userInput models.User

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failed",
			"message": "invalid request",
		})
	}

	cfg := middlewares.GetDefaultConfig()

	token, err := services.Login(userInput, &cfg)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failed",
			"message": "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func GetAll(c echo.Context) error {
	users, err := services.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failed",
			"message": "fetch data failed",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "success",
		"message": "all users",
		"data":    users,
	})
}

func Delete(c echo.Context) error {
	id := c.Param("id")

	err := services.Delete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failed",
			"message": "delete data failed",
		})
	}

	return c.JSON(http.StatusNoContent, echo.Map{})
}

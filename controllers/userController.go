package controllers

import (
	"Cobain/config"
	"Cobain/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	userID := c.Param("id")
	var users models.User
	// SELECT * FROM users WHERE id LIKE userID LIMIT BY 1
	result := config.DB.First(&users, userID)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   result,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	save := config.DB.Save(&user)

	if save.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, save.Error)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	userID := c.Param("id")
	var user models.User
	result := config.DB.Delete(&user, userID)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete",
		"user":    nil,
	})

}

// update user by id
func UpdateUserController(c echo.Context) error {
	// masang body ke dalam struktur data User
	var payload models.User
	c.Bind(&payload)

	// nyari data berdasarkan id
	userID := c.Param("id")
	edit := models.User{}
	// fetching ke DB (nyari ke db dan masukin ke dalam result)
	config.DB.First(&edit, userID)

	// masukin body ke variabel
	edit.Name = payload.Name
	edit.Email = payload.Email
	edit.Password = payload.Password
	// save ke db
	config.DB.Save(&edit)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success edit",
		"user":    edit,
	})
}

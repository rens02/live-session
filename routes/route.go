package routes

import (
	"Cobain/controllers"
	"github.com/labstack/echo/v4"
)

func Routing() *echo.Echo {
	e := echo.New()

	// Route / to handler function
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	return e
}

package main

import (
	"Prioritas2/controllers"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	//Login
	e.POST("/login", controllers.LoginController)

	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte("1234")))
	eAuth.GET("/blogs", controllers.GetAllBlog)
	eAuth.GET("/blogs/:id", controllers.GetBlogByID)
	eAuth.POST("/blogs", controllers.CreateBlog)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

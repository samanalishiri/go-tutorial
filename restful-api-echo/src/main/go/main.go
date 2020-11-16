package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"restful-api-echo/src/main/go/user"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", user.Save)
	e.GET("/users/:id", user.GetOne)
	e.PUT("/users/:id", user.UpdateOne)
	e.DELETE("/users/:id", user.DeleteOne)

	e.Logger.Fatal(e.Start(":8081"))
}

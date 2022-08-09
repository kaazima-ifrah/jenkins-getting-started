package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"

	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello! Nice to meet you :)")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/welcome", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Welcome to Jenkins Getting Started session! Hope you have a great learning experience!")
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

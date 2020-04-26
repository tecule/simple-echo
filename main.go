package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(context echo.Context) error {
	return context.String(http.StatusOK, "hello, world")
}

package main

import (
	"net/http"
	"github.com/labstack/echo"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Yura!")
  }

func address(c echo.Context) error {
	return c.String(http.StatusOK, "kazan")
  }

func main() {
	e := echo.New()

	// Routes
	e.GET("/", hello)
	e.GET("kazan", address)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
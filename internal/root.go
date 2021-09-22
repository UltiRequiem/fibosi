// API Logic
package internal

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Initialize a new server with configuration
func NewServer(port int, banner bool) error {
	e := echo.New()

	e.HideBanner = banner

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Paths
	e.GET("/:number", FibonacciHandler)
	e.GET("/sequence/:number", FibonacciSequenceHandler)

	return e.Start(fmt.Sprintf(":%d", port))
}

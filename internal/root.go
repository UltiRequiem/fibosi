package internal

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(port int, banner bool) {
	e := echo.New()

        // Customization
	e.HideBanner = banner

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/:number", FibonacciHandler)
	e.GET("/sequence/:number", FibonacciSequenceHandler)

	e.Start(fmt.Sprintf(":%d", port))
}

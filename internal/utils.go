package internal

import (
	"github.com/labstack/echo/v4"
	"log"
)

// Log on console and send a JSON Response with the error
func logManageableError(message error, code int, c echo.Context) error {
	log.Println(message)
	return c.JSON(code, &EchoError{message, code})
}

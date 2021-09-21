package internal

import (
	"github.com/labstack/echo/v4"
	"log"
)

func logManageableError(message string, code int, c echo.Context) error {
	log.Println(message)
	return c.String(code, message)
}

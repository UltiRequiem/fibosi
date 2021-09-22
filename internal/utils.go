package internal

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func logManageableError(message error, code int, c echo.Context) error {
	log.Println(message)
	return c.JSON(code, &EchoError{fmt.Sprintf(`%v`, message), code})
}

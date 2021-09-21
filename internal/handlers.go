package internal

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	. "github.com/UltiRequiem/fibonnaci/pkg"
)

func FibonacciHandler(c echo.Context) error {
	numberParam := c.Param("number")

	number, err := strconv.Atoi(numberParam)

	if err != nil {
		logManageableError(fmt.Sprintf(`Got "%s", but a number was expected.`, numberParam), http.StatusUnprocessableEntity, c)
	}

	return c.String(http.StatusOK, strconv.Itoa(Fibonacci(number)))
}

func FibonacciSequenceHandler(c echo.Context) error {
	numberParam := c.Param("number")

	number, err := strconv.Atoi(numberParam)

	if err != nil {
		logManageableError(fmt.Sprintf(`Got "%s", but a number was expected.`, numberParam), http.StatusUnprocessableEntity, c)
	}

	return c.JSON(http.StatusOK, FibonacciSequence(number))
}

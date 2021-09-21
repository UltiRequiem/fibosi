package internal

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	. "github.com/UltiRequiem/fibonnaci/pkg"
)

func FibonacciHandler(c echo.Context) error {
	numberParam := c.Param("number")

	number, err := strconv.Atoi(numberParam)

	if err != nil {
		log.Errorf(`Got "%s", but number expected`, numberParam)
	}

	return c.String(http.StatusOK, strconv.Itoa(Fibonacci(number)))
}

func FibonacciSequenceHandler(c echo.Context) error {
	numberParam := c.Param("number")

	number, err := strconv.Atoi(numberParam)

	if err != nil {
		log.Errorf(`Got "%s", but number expected`, numberParam)
	}

	return c.JSON(http.StatusOK, FibonacciSequence(number))
}

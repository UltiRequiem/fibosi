package internal

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	. "github.com/UltiRequiem/fibonnaci/pkg"
)

// Root path handler
func FibonacciHandler(c echo.Context) error {
	numberParam := c.Param("number")

	number, err := strconv.Atoi(numberParam)

	if err != nil {
		return logManageableError(fmt.Errorf(`Got '%s', but a number was expected.`, numberParam), http.StatusUnprocessableEntity, c)
	}

	fiboNum, err := Fibonacci(number)

	if err != nil {
		return logManageableError(err, http.StatusUnprocessableEntity, c)
	}

	return c.JSON(http.StatusOK, &FibonacciNumber{fiboNum})
}

// Sequence Path
func FibonacciSequenceHandler(c echo.Context) error {
	numberParam := c.Param("number")

	number, err := strconv.Atoi(numberParam)

	if err != nil {
		return logManageableError(fmt.Errorf(`Got '%s', but a number was expected.`, numberParam), http.StatusUnprocessableEntity, c)
	}

	fiboNums, err := FibonacciSequence(number)

	if err != nil {
		return logManageableError(err, http.StatusUnprocessableEntity, c)
	}

	return c.JSON(http.StatusOK, &FibonacciNumberSequence{fiboNums})
}

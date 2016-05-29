package echomware

import (
	"log"

	"os"

	"github.com/labstack/echo"
)

var errl = log.New(os.Stderr, "error: ", log.LstdFlags)

// ErrorLogger logs all errors to stderr.
func ErrorLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if err = next(c); err != nil {
				c.Error(err)
				errl.Println(err)
			}
			return
		}
	}
}

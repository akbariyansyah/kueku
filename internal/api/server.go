package api

import (
	"kueku/config"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BuildEcho(conf *config.Config) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.HTTPErrorHandler = customHTTPErrorHandler
	return e

}
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	if err := c.JSON(code, map[string]interface{}{
		"message": err.Error(),
	}); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

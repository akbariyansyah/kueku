package cake

import (
	"kueku/internal/usecase"

	"github.com/labstack/echo/v4"
)

func CakeRoute(
	e *echo.Echo,
	cakeUsecase usecase.CakeUsecase) {
	handler := &Handler{Router: e, CakeUsecase: cakeUsecase}
	v1 := handler.Router.Group("/v1")

	v1.GET("/cakes", handler.List)
	v1.GET("/cakes/:id", handler.Detail)
	v1.POST("/cakes", handler.Create)
	v1.PUT("/cakes/:id", handler.Update)
	v1.DELETE("/cakes/:id", handler.Detail)

}

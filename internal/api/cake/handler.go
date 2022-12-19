package cake

import (
	"kueku/internal/api/response"
	"kueku/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Router      *echo.Echo
	CakeUsecase usecase.CakeUsecase
}

func (h *Handler) List(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.CakeUsecase.List(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

func (h *Handler) Detail(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	res, err := h.CakeUsecase.Detail(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

func (h *Handler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(CreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := h.CakeUsecase.Create(ctx, req.ToCommand())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

func (h *Handler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(UpdateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	req.ID = c.Param("id")
	res, err := h.CakeUsecase.Update(ctx, req.ToCommand())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

func (h *Handler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	err := h.CakeUsecase.Delete(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.NewSuccessResponse(nil))
}

package cake_test

import (
	"context"
	"errors"
	"kueku/internal/api/cake"
	cakedomain "kueku/internal/domain/cake"
	umocks "kueku/internal/usecase/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var mockErr = errors.New("unexpected")

type binderError struct{}

func (b binderError) Bind(i interface{}, c echo.Context) error {
	return mockErr
}

type handlerTestSuite struct {
	suite.Suite
	ctx         context.Context
	mockCtrl    *gomock.Controller
	echo        *echo.Echo
	mockUsecase *umocks.MockCakeUsecase
	handler     *cake.Handler
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}

func (t *handlerTestSuite) SetupTest() {
	t.ctx = context.Background()
	t.mockCtrl = gomock.NewController(t.T())
	t.echo = echo.New()
	t.mockUsecase = umocks.NewMockCakeUsecase(t.mockCtrl)
	t.handler = &cake.Handler{
		Router:      t.echo,
		CakeUsecase: t.mockUsecase,
	}
	cake.CakeRoute(t.echo, t.mockUsecase)
}

func (t *handlerTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
}

func (t *handlerTestSuite) TestList_Error() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)

	t.mockUsecase.EXPECT().List(t.ctx).Return(nil, mockErr)

	err := t.handler.List(c)

	t.Error(err)
}

func (t *handlerTestSuite) TestList_NoError() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)
	mockRes := make(cakedomain.Cakes, 5)

	t.mockUsecase.EXPECT().List(t.ctx).Return(mockRes, nil)

	err := t.handler.List(c)

	t.Nil(err)
}

func (t *handlerTestSuite) TestDetail_Error() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("4")

	t.mockUsecase.EXPECT().Detail(t.ctx, 4).Return(nil, mockErr)

	err := t.handler.Detail(c)

	t.Error(err)
}

func (t *handlerTestSuite) TestDetail_NoError() {
	req := httptest.NewRequest(http.MethodGet, "/4", nil)
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("4")

	mockRes := new(cakedomain.Cake)

	t.mockUsecase.EXPECT().Detail(t.ctx, 4).Return(mockRes, nil)

	err := t.handler.Detail(c)

	t.Nil(err)
}

func (t *handlerTestSuite) TestCreate_BindError() {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)
	c.Echo().Binder = new(binderError)

	err := t.handler.Create(c)

	t.Equal(mockErr, err)
}

func (t *handlerTestSuite) TestCreate_UsecaseError() {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)

	t.mockUsecase.EXPECT().Create(t.ctx, gomock.Any()).Return(nil, mockErr)

	err := t.handler.Create(c)

	t.Equal(mockErr, err)
}

func (t *handlerTestSuite) TestCreate_NoError() {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)

	mockRes := new(cakedomain.Cake)

	t.mockUsecase.EXPECT().Create(t.ctx, gomock.Any()).Return(mockRes, nil)

	err := t.handler.Create(c)

	t.Nil(err)
}


func (t *handlerTestSuite) TestUpdate_BindError() {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)
	c.Echo().Binder = new(binderError)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := t.handler.Update(c)

	t.Equal(mockErr, err)
}

func (t *handlerTestSuite) TestUpdate_UsecaseError() {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	t.mockUsecase.EXPECT().Update(t.ctx, gomock.Any()).Return(nil, mockErr)

	err := t.handler.Update(c)

	t.Equal(mockErr, err)
}

func (t *handlerTestSuite) TestUpdate_NoError() {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")
	mockRes := new(cakedomain.Cake)

	t.mockUsecase.EXPECT().Update(t.ctx, gomock.Any()).Return(mockRes, nil)

	err := t.handler.Update(c)

	t.Nil(err)
}

func (t *handlerTestSuite) TestDelete_Error() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("4")

	t.mockUsecase.EXPECT().Delete(t.ctx, 4).Return(mockErr)

	err := t.handler.Delete(c)

	t.Error(err)
}

func (t *handlerTestSuite) TestDelete_NoError() {
	req := httptest.NewRequest(http.MethodGet, "/4", nil)
	rec := httptest.NewRecorder()
	c := t.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("4")

	t.mockUsecase.EXPECT().Delete(t.ctx, 4).Return(nil)

	err := t.handler.Delete(c)

	t.Nil(err)
}

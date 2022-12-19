package usecase_test

import (
	"context"
	"errors"
	"kueku/internal/usecase"
	"kueku/internal/usecase/model"
	"kueku/internal/domain/cake"
	cmocks "kueku/internal/domain/cake/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type cakeUsecaseTestSuite struct {
	suite.Suite
	ctx          context.Context
	mockCtrl     *gomock.Controller
	mockCakeRepo *cmocks.MockRepository
	cakeUsecase  usecase.CakeUsecase
}

func TestCakeUsecaseSuite(t *testing.T) {
	suite.Run(t, new(cakeUsecaseTestSuite))
}

func (t *cakeUsecaseTestSuite) SetupTest() {
	t.ctx = context.Background()
	t.mockCtrl = gomock.NewController(t.T())
	t.mockCakeRepo = cmocks.NewMockRepository(t.mockCtrl)

	t.cakeUsecase = usecase.NewCake(t.mockCakeRepo)
}

func (t *cakeUsecaseTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
}

func (t *cakeUsecaseTestSuite) TestList_FetchError() {
	mockErr := errors.New("unexpected")

	t.mockCakeRepo.EXPECT().Fetch(t.ctx).Return(nil, mockErr)

	res, err := t.cakeUsecase.List(t.ctx)

	t.Equal(mockErr, err)
	t.Nil(res)
}

func (t *cakeUsecaseTestSuite) TestList_NoError() {
	mockRes := make(cake.Cakes, 5)

	t.mockCakeRepo.EXPECT().Fetch(t.ctx).Return(mockRes, nil)

	res, err := t.cakeUsecase.List(t.ctx)

	t.Equal(mockRes, res)
	t.Nil(err)
}

func (t *cakeUsecaseTestSuite) TestDetail_GetByIDError() {
	mockErr := errors.New("unexpected")
	mockID := 5

	t.mockCakeRepo.EXPECT().GetByID(t.ctx, 5).Return(nil, mockErr)

	res, err := t.cakeUsecase.Detail(t.ctx, mockID)

	t.Equal(mockErr, err)
	t.Nil(res)
}

func (t *cakeUsecaseTestSuite) TestDetail_NoError() {
	mockRes := new(cake.Cake)
	mockID := 5

	t.mockCakeRepo.EXPECT().GetByID(t.ctx, mockID).Return(mockRes, nil)

	res, err := t.cakeUsecase.Detail(t.ctx, mockID)

	t.Equal(mockRes, res)
	t.Nil(err)
}

func (t *cakeUsecaseTestSuite) TestCreate_CreateError() {
	mockErr := errors.New("unexpected")
	mockCommand := new(model.CreateCakeCommand)

	t.mockCakeRepo.EXPECT().Create(t.ctx, gomock.Any()).Return(mockErr)

	res, err := t.cakeUsecase.Create(t.ctx, mockCommand)

	t.Equal(mockErr, err)
	t.Nil(res)
}

func (t *cakeUsecaseTestSuite) TestCreate_GetByIDError() {
	mockErr := errors.New("unexpected")
	mockCommand := new(model.CreateCakeCommand)

	t.mockCakeRepo.EXPECT().Create(t.ctx, gomock.Any()).Return(nil)
	t.mockCakeRepo.EXPECT().GetByID(t.ctx, gomock.Any()).Return(nil, mockErr)

	res, err := t.cakeUsecase.Create(t.ctx, mockCommand)

	t.Equal(mockErr, err)
	t.Nil(res)
}

func (t *cakeUsecaseTestSuite) TestCreate_NoError() {
	mockRes := new(cake.Cake)
	mockCommand := new(model.CreateCakeCommand)

	t.mockCakeRepo.EXPECT().Create(t.ctx, gomock.Any()).Return(nil)
	t.mockCakeRepo.EXPECT().GetByID(t.ctx, gomock.Any()).Return(mockRes, nil)

	res, err := t.cakeUsecase.Create(t.ctx, mockCommand)

	t.Equal(mockRes, res)
	t.Nil(err)
}

func (t *cakeUsecaseTestSuite) TestUpdate_CreateError() {
	mockErr := errors.New("unexpected")
	mockCommand := new(model.UpdateCakeCommand)

	t.mockCakeRepo.EXPECT().Update(t.ctx, gomock.Any()).Return(mockErr)

	res, err := t.cakeUsecase.Update(t.ctx, mockCommand)

	t.Equal(mockErr, err)
	t.Nil(res)
}

func (t *cakeUsecaseTestSuite) TestUpdate_GetByIDError() {
	mockErr := errors.New("unexpected")
	mockCommand := new(model.UpdateCakeCommand)

	t.mockCakeRepo.EXPECT().Update(t.ctx, gomock.Any()).Return(nil)
	t.mockCakeRepo.EXPECT().GetByID(t.ctx, gomock.Any()).Return(nil, mockErr)

	res, err := t.cakeUsecase.Update(t.ctx, mockCommand)

	t.Equal(mockErr, err)
	t.Nil(res)
}

func (t *cakeUsecaseTestSuite) TestUpdate_NoError() {
	mockRes := new(cake.Cake)
	mockCommand := new(model.UpdateCakeCommand)

	t.mockCakeRepo.EXPECT().Update(t.ctx, gomock.Any()).Return(nil)
	t.mockCakeRepo.EXPECT().GetByID(t.ctx, gomock.Any()).Return(mockRes, nil)

	res, err := t.cakeUsecase.Update(t.ctx, mockCommand)

	t.Equal(mockRes, res)
	t.Nil(err)
}

func (t *cakeUsecaseTestSuite) TestDelete_DeleteError() {
	mockErr := errors.New("unexpected")
	mockID := 5

	t.mockCakeRepo.EXPECT().Delete(t.ctx, 5).Return(mockErr)

	err := t.cakeUsecase.Delete(t.ctx, mockID)

	t.Equal(mockErr, err)
}

func (t *cakeUsecaseTestSuite) TestDelete_NoError() {
	mockID := 5

	t.mockCakeRepo.EXPECT().Delete(t.ctx, mockID).Return(nil)

	err := t.cakeUsecase.Delete(t.ctx, mockID)

	t.Nil(err)
}

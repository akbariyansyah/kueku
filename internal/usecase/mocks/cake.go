// Code generated by MockGen. DO NOT EDIT.
// Source: cake.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	cake "kueku/internal/domain/cake"
	model "kueku/internal/usecase/model"
	reflect "reflect"
)

// MockCakeUsecase is a mock of CakeUsecase interface
type MockCakeUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCakeUsecaseMockRecorder
}

// MockCakeUsecaseMockRecorder is the mock recorder for MockCakeUsecase
type MockCakeUsecaseMockRecorder struct {
	mock *MockCakeUsecase
}

// NewMockCakeUsecase creates a new mock instance
func NewMockCakeUsecase(ctrl *gomock.Controller) *MockCakeUsecase {
	mock := &MockCakeUsecase{ctrl: ctrl}
	mock.recorder = &MockCakeUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCakeUsecase) EXPECT() *MockCakeUsecaseMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockCakeUsecase) List(ctx context.Context) (cake.Cakes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].(cake.Cakes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockCakeUsecaseMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCakeUsecase)(nil).List), ctx)
}

// Detail mocks base method
func (m *MockCakeUsecase) Detail(ctx context.Context, id int) (*cake.Cake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Detail", ctx, id)
	ret0, _ := ret[0].(*cake.Cake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Detail indicates an expected call of Detail
func (mr *MockCakeUsecaseMockRecorder) Detail(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detail", reflect.TypeOf((*MockCakeUsecase)(nil).Detail), ctx, id)
}

// Create mocks base method
func (m *MockCakeUsecase) Create(ctx context.Context, command *model.CreateCakeCommand) (*cake.Cake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, command)
	ret0, _ := ret[0].(*cake.Cake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCakeUsecaseMockRecorder) Create(ctx, command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCakeUsecase)(nil).Create), ctx, command)
}

// Update mocks base method
func (m *MockCakeUsecase) Update(ctx context.Context, command *model.UpdateCakeCommand) (*cake.Cake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, command)
	ret0, _ := ret[0].(*cake.Cake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockCakeUsecaseMockRecorder) Update(ctx, command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCakeUsecase)(nil).Update), ctx, command)
}

// Delete mocks base method
func (m *MockCakeUsecase) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCakeUsecaseMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCakeUsecase)(nil).Delete), ctx, id)
}
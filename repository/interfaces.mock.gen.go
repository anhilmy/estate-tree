// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetAllTree mocks base method.
func (m *MockRepositoryInterface) GetAllTree(ctx context.Context, input UuidInput) ([]TreeModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTree", ctx, input)
	ret0, _ := ret[0].([]TreeModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTree indicates an expected call of GetAllTree.
func (mr *MockRepositoryInterfaceMockRecorder) GetAllTree(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTree", reflect.TypeOf((*MockRepositoryInterface)(nil).GetAllTree), ctx, input)
}

// GetEstate mocks base method.
func (m *MockRepositoryInterface) GetEstate(ctx context.Context, input UuidInput) (EstateModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEstate", ctx, input)
	ret0, _ := ret[0].(EstateModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEstate indicates an expected call of GetEstate.
func (mr *MockRepositoryInterfaceMockRecorder) GetEstate(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEstate", reflect.TypeOf((*MockRepositoryInterface)(nil).GetEstate), ctx, input)
}

// GetEstateStats mocks base method.
func (m *MockRepositoryInterface) GetEstateStats(ctx context.Context, input UuidInput) (EstateStatsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEstateStats", ctx, input)
	ret0, _ := ret[0].(EstateStatsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEstateStats indicates an expected call of GetEstateStats.
func (mr *MockRepositoryInterfaceMockRecorder) GetEstateStats(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEstateStats", reflect.TypeOf((*MockRepositoryInterface)(nil).GetEstateStats), ctx, input)
}

// GetTree mocks base method.
func (m *MockRepositoryInterface) GetTree(ctx context.Context, input GetTreeByCoordinateInput) (TreeModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTree", ctx, input)
	ret0, _ := ret[0].(TreeModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTree indicates an expected call of GetTree.
func (mr *MockRepositoryInterfaceMockRecorder) GetTree(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTree", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTree), ctx, input)
}

// InsertEstate mocks base method.
func (m *MockRepositoryInterface) InsertEstate(ctx context.Context, input CreateEstateInput) (UuidOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertEstate", ctx, input)
	ret0, _ := ret[0].(UuidOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertEstate indicates an expected call of InsertEstate.
func (mr *MockRepositoryInterfaceMockRecorder) InsertEstate(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertEstate", reflect.TypeOf((*MockRepositoryInterface)(nil).InsertEstate), ctx, input)
}

// InsertTree mocks base method.
func (m *MockRepositoryInterface) InsertTree(ctx context.Context, input CreateTreeInput) (UuidOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTree", ctx, input)
	ret0, _ := ret[0].(UuidOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertTree indicates an expected call of InsertTree.
func (mr *MockRepositoryInterfaceMockRecorder) InsertTree(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTree", reflect.TypeOf((*MockRepositoryInterface)(nil).InsertTree), ctx, input)
}

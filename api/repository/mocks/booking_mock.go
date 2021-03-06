// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nmcalinden/footpal/repository (interfaces: BookingRepositoryI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/nmcalinden/footpal/models"
)

// MockBookingRepositoryI is a mock of BookingRepositoryI interface.
type MockBookingRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockBookingRepositoryIMockRecorder
}

// MockBookingRepositoryIMockRecorder is the mock recorder for MockBookingRepositoryI.
type MockBookingRepositoryIMockRecorder struct {
	mock *MockBookingRepositoryI
}

// NewMockBookingRepositoryI creates a new mock instance.
func NewMockBookingRepositoryI(ctrl *gomock.Controller) *MockBookingRepositoryI {
	mock := &MockBookingRepositoryI{ctrl: ctrl}
	mock.recorder = &MockBookingRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookingRepositoryI) EXPECT() *MockBookingRepositoryIMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockBookingRepositoryI) FindAll() (*[]models.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].(*[]models.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockBookingRepositoryIMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockBookingRepositoryI)(nil).FindAll))
}

// FindAvailableVenues mocks base method.
func (m *MockBookingRepositoryI) FindAvailableVenues(arg0 *int, arg1 string, arg2 *string, arg3 *int) (*[]models.Venue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAvailableVenues", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*[]models.Venue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAvailableVenues indicates an expected call of FindAvailableVenues.
func (mr *MockBookingRepositoryIMockRecorder) FindAvailableVenues(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAvailableVenues", reflect.TypeOf((*MockBookingRepositoryI)(nil).FindAvailableVenues), arg0, arg1, arg2, arg3)
}

// FindById mocks base method.
func (m *MockBookingRepositoryI) FindById(arg0 *int) (*models.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*models.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockBookingRepositoryIMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockBookingRepositoryI)(nil).FindById), arg0)
}

// FindMatchesByBookingId mocks base method.
func (m *MockBookingRepositoryI) FindMatchesByBookingId(arg0 *int) (*[]models.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMatchesByBookingId", arg0)
	ret0, _ := ret[0].(*[]models.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMatchesByBookingId indicates an expected call of FindMatchesByBookingId.
func (mr *MockBookingRepositoryIMockRecorder) FindMatchesByBookingId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMatchesByBookingId", reflect.TypeOf((*MockBookingRepositoryI)(nil).FindMatchesByBookingId), arg0)
}

// IsExistingMatchPresent mocks base method.
func (m *MockBookingRepositoryI) IsExistingMatchPresent(arg0 *string, arg1 *int) (*bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExistingMatchPresent", arg0, arg1)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExistingMatchPresent indicates an expected call of IsExistingMatchPresent.
func (mr *MockBookingRepositoryIMockRecorder) IsExistingMatchPresent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExistingMatchPresent", reflect.TypeOf((*MockBookingRepositoryI)(nil).IsExistingMatchPresent), arg0, arg1)
}

// Save mocks base method.
func (m *MockBookingRepositoryI) Save(arg0 *models.Booking, arg1 *[]models.Match, arg2 *[]models.PitchSlot) (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1, arg2)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockBookingRepositoryIMockRecorder) Save(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockBookingRepositoryI)(nil).Save), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockBookingRepositoryI) Update(arg0 *models.Booking) (*models.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*models.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBookingRepositoryIMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookingRepositoryI)(nil).Update), arg0)
}

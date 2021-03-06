// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nmcalinden/footpal/repository (interfaces: MatchPlayerRepositoryI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/nmcalinden/footpal/models"
)

// MockMatchPlayerRepositoryI is a mock of MatchPlayerRepositoryI interface.
type MockMatchPlayerRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockMatchPlayerRepositoryIMockRecorder
}

// MockMatchPlayerRepositoryIMockRecorder is the mock recorder for MockMatchPlayerRepositoryI.
type MockMatchPlayerRepositoryIMockRecorder struct {
	mock *MockMatchPlayerRepositoryI
}

// NewMockMatchPlayerRepositoryI creates a new mock instance.
func NewMockMatchPlayerRepositoryI(ctrl *gomock.Controller) *MockMatchPlayerRepositoryI {
	mock := &MockMatchPlayerRepositoryI{ctrl: ctrl}
	mock.recorder = &MockMatchPlayerRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatchPlayerRepositoryI) EXPECT() *MockMatchPlayerRepositoryIMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMatchPlayerRepositoryI) Delete(arg0, arg1 *int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMatchPlayerRepositoryIMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMatchPlayerRepositoryI)(nil).Delete), arg0, arg1)
}

// FindByMatchId mocks base method.
func (m *MockMatchPlayerRepositoryI) FindByMatchId(arg0 *int) (*[]models.MatchPlayer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMatchId", arg0)
	ret0, _ := ret[0].(*[]models.MatchPlayer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMatchId indicates an expected call of FindByMatchId.
func (mr *MockMatchPlayerRepositoryIMockRecorder) FindByMatchId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMatchId", reflect.TypeOf((*MockMatchPlayerRepositoryI)(nil).FindByMatchId), arg0)
}

// FindByMatchIdAndPlayerId mocks base method.
func (m *MockMatchPlayerRepositoryI) FindByMatchIdAndPlayerId(arg0, arg1 *int) (*models.MatchPlayer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMatchIdAndPlayerId", arg0, arg1)
	ret0, _ := ret[0].(*models.MatchPlayer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMatchIdAndPlayerId indicates an expected call of FindByMatchIdAndPlayerId.
func (mr *MockMatchPlayerRepositoryIMockRecorder) FindByMatchIdAndPlayerId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMatchIdAndPlayerId", reflect.TypeOf((*MockMatchPlayerRepositoryI)(nil).FindByMatchIdAndPlayerId), arg0, arg1)
}

// FindMatchesByPlayer mocks base method.
func (m *MockMatchPlayerRepositoryI) FindMatchesByPlayer(arg0 *int) (*[]models.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMatchesByPlayer", arg0)
	ret0, _ := ret[0].(*[]models.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMatchesByPlayer indicates an expected call of FindMatchesByPlayer.
func (mr *MockMatchPlayerRepositoryIMockRecorder) FindMatchesByPlayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMatchesByPlayer", reflect.TypeOf((*MockMatchPlayerRepositoryI)(nil).FindMatchesByPlayer), arg0)
}

// Save mocks base method.
func (m *MockMatchPlayerRepositoryI) Save(arg0 models.MatchPlayer) (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockMatchPlayerRepositoryIMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockMatchPlayerRepositoryI)(nil).Save), arg0)
}

// Update mocks base method.
func (m *MockMatchPlayerRepositoryI) Update(arg0 *models.MatchPlayer) (*models.MatchPlayer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*models.MatchPlayer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMatchPlayerRepositoryIMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMatchPlayerRepositoryI)(nil).Update), arg0)
}

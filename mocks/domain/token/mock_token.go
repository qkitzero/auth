// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/token/token.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/token/token.go -destination=mocks/domain/token/mock_token.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockToken is a mock of Token interface.
type MockToken struct {
	ctrl     *gomock.Controller
	recorder *MockTokenMockRecorder
	isgomock struct{}
}

// MockTokenMockRecorder is the mock recorder for MockToken.
type MockTokenMockRecorder struct {
	mock *MockToken
}

// NewMockToken creates a new mock instance.
func NewMockToken(ctrl *gomock.Controller) *MockToken {
	mock := &MockToken{ctrl: ctrl}
	mock.recorder = &MockTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToken) EXPECT() *MockTokenMockRecorder {
	return m.recorder
}

// AccessToken mocks base method.
func (m *MockToken) AccessToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// AccessToken indicates an expected call of AccessToken.
func (mr *MockTokenMockRecorder) AccessToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessToken", reflect.TypeOf((*MockToken)(nil).AccessToken))
}

// RefreshToken mocks base method.
func (m *MockToken) RefreshToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockTokenMockRecorder) RefreshToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockToken)(nil).RefreshToken))
}

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/Xe/pawd/cmd/pawd/database (interfaces: Tokens)

package mock_database

import (
	context "context"
	database "github.com/Xe/pawd/cmd/pawd/database"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Tokens interface
type MockTokens struct {
	ctrl     *gomock.Controller
	recorder *_MockTokensRecorder
}

// Recorder for MockTokens (not exported)
type _MockTokensRecorder struct {
	mock *MockTokens
}

func NewMockTokens(ctrl *gomock.Controller) *MockTokens {
	mock := &MockTokens{ctrl: ctrl}
	mock.recorder = &_MockTokensRecorder{mock}
	return mock
}

func (_m *MockTokens) EXPECT() *_MockTokensRecorder {
	return _m.recorder
}

func (_m *MockTokens) Check(_param0 context.Context, _param1 string) (*database.Token, error) {
	ret := _m.ctrl.Call(_m, "Check", _param0, _param1)
	ret0, _ := ret[0].(*database.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTokensRecorder) Check(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Check", arg0, arg1)
}

func (_m *MockTokens) Create(_param0 context.Context, _param1 string) (*database.Token, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0, _param1)
	ret0, _ := ret[0].(*database.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTokensRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0, arg1)
}

func (_m *MockTokens) UpdateLastSeen(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "UpdateLastSeen", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokensRecorder) UpdateLastSeen(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateLastSeen", arg0, arg1)
}
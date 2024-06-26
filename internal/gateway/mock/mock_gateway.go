// Code generated by MockGen. DO NOT EDIT.
// Source: internal/gateway/gateway.go

// Package mock_gateway is a generated GoMock package.
package mock_gateway

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ykkssyaa/Bash_Service/internal/models"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockStorage) Get(id int) context.CancelFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(context.CancelFunc)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockStorageMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStorage)(nil).Get), id)
}

// Remove mocks base method.
func (m *MockStorage) Remove(id int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", id)
}

// Remove indicates an expected call of Remove.
func (mr *MockStorageMockRecorder) Remove(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockStorage)(nil).Remove), id)
}

// Set mocks base method.
func (m *MockStorage) Set(id int, ctxFunc context.CancelFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", id, ctxFunc)
}

// Set indicates an expected call of Set.
func (mr *MockStorageMockRecorder) Set(id, ctxFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStorage)(nil).Set), id, ctxFunc)
}

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCache) Get(id int) (models.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(models.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), id)
}

// Remove mocks base method.
func (m *MockCache) Remove(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockCacheMockRecorder) Remove(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCache)(nil).Remove), id)
}

// Set mocks base method.
func (m *MockCache) Set(id int, cmd models.Command) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", id, cmd)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheMockRecorder) Set(id, cmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCache)(nil).Set), id, cmd)
}

// MockCommand is a mock of Command interface.
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand.
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance.
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// CreateCommand mocks base method.
func (m *MockCommand) CreateCommand(command models.Command) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommand", command)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCommand indicates an expected call of CreateCommand.
func (mr *MockCommandMockRecorder) CreateCommand(command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommand", reflect.TypeOf((*MockCommand)(nil).CreateCommand), command)
}

// GetAllCommands mocks base method.
func (m *MockCommand) GetAllCommands(limit, offset int) ([]models.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCommands", limit, offset)
	ret0, _ := ret[0].([]models.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCommands indicates an expected call of GetAllCommands.
func (mr *MockCommandMockRecorder) GetAllCommands(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCommands", reflect.TypeOf((*MockCommand)(nil).GetAllCommands), limit, offset)
}

// GetCommand mocks base method.
func (m *MockCommand) GetCommand(commandId int) (models.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommand", commandId)
	ret0, _ := ret[0].(models.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommand indicates an expected call of GetCommand.
func (mr *MockCommandMockRecorder) GetCommand(commandId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommand", reflect.TypeOf((*MockCommand)(nil).GetCommand), commandId)
}

// UpdateCommand mocks base method.
func (m *MockCommand) UpdateCommand(command models.Command) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCommand", command)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCommand indicates an expected call of UpdateCommand.
func (mr *MockCommandMockRecorder) UpdateCommand(command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCommand", reflect.TypeOf((*MockCommand)(nil).UpdateCommand), command)
}

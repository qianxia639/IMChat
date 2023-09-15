// Code generated by MockGen. DO NOT EDIT.
// Source: IMChat/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	db "IMChat/db/sqlc"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddFriend mocks base method.
func (m *MockStore) AddFriend(arg0 context.Context, arg1 *db.AddFriendParams) (db.Friend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFriend", arg0, arg1)
	ret0, _ := ret[0].(db.Friend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFriend indicates an expected call of AddFriend.
func (mr *MockStoreMockRecorder) AddFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFriend", reflect.TypeOf((*MockStore)(nil).AddFriend), arg0, arg1)
}

// AddFriendTx mocks base method.
func (m *MockStore) AddFriendTx(arg0 context.Context, arg1 *db.AddFriendTxParams) (db.Friend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFriendTx", arg0, arg1)
	ret0, _ := ret[0].(db.Friend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFriendTx indicates an expected call of AddFriendTx.
func (mr *MockStoreMockRecorder) AddFriendTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFriendTx", reflect.TypeOf((*MockStore)(nil).AddFriendTx), arg0, arg1)
}

// AddMessage mocks base method.
func (m *MockStore) AddMessage(arg0 context.Context, arg1 *db.AddMessageParams) (db.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMessage", arg0, arg1)
	ret0, _ := ret[0].(db.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMessage indicates an expected call of AddMessage.
func (mr *MockStoreMockRecorder) AddMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMessage", reflect.TypeOf((*MockStore)(nil).AddMessage), arg0, arg1)
}

// AddUserLoginLog mocks base method.
func (m *MockStore) AddUserLoginLog(arg0 context.Context, arg1 *db.AddUserLoginLogParams) (db.UserLoginLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserLoginLog", arg0, arg1)
	ret0, _ := ret[0].(db.UserLoginLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUserLoginLog indicates an expected call of AddUserLoginLog.
func (mr *MockStoreMockRecorder) AddUserLoginLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserLoginLog", reflect.TypeOf((*MockStore)(nil).AddUserLoginLog), arg0, arg1)
}

// CreateFriendClsuterApply mocks base method.
func (m *MockStore) CreateFriendClsuterApply(arg0 context.Context, arg1 *db.CreateFriendClsuterApplyParams) (db.FriendClusterApply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFriendClsuterApply", arg0, arg1)
	ret0, _ := ret[0].(db.FriendClusterApply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFriendClsuterApply indicates an expected call of CreateFriendClsuterApply.
func (mr *MockStoreMockRecorder) CreateFriendClsuterApply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFriendClsuterApply", reflect.TypeOf((*MockStore)(nil).CreateFriendClsuterApply), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 *db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteFriend mocks base method.
func (m *MockStore) DeleteFriend(arg0 context.Context, arg1 *db.DeleteFriendParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFriend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFriend indicates an expected call of DeleteFriend.
func (mr *MockStoreMockRecorder) DeleteFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFriend", reflect.TypeOf((*MockStore)(nil).DeleteFriend), arg0, arg1)
}

// DeleteFriendTx mocks base method.
func (m *MockStore) DeleteFriendTx(arg0 context.Context, arg1 *db.DeleteFriendTxParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFriendTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFriendTx indicates an expected call of DeleteFriendTx.
func (mr *MockStoreMockRecorder) DeleteFriendTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFriendTx", reflect.TypeOf((*MockStore)(nil).DeleteFriendTx), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// ExistsFriendClusterApply mocks base method.
func (m *MockStore) ExistsFriendClusterApply(arg0 context.Context, arg1 *db.ExistsFriendClusterApplyParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsFriendClusterApply", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsFriendClusterApply indicates an expected call of ExistsFriendClusterApply.
func (mr *MockStoreMockRecorder) ExistsFriendClusterApply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsFriendClusterApply", reflect.TypeOf((*MockStore)(nil).ExistsFriendClusterApply), arg0, arg1)
}

// GetFriend mocks base method.
func (m *MockStore) GetFriend(arg0 context.Context, arg1 *db.GetFriendParams) (db.Friend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriend", arg0, arg1)
	ret0, _ := ret[0].(db.Friend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriend indicates an expected call of GetFriend.
func (mr *MockStoreMockRecorder) GetFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriend", reflect.TypeOf((*MockStore)(nil).GetFriend), arg0, arg1)
}

// GetLastUserLoginLog mocks base method.
func (m *MockStore) GetLastUserLoginLog(arg0 context.Context, arg1 int32) (db.UserLoginLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastUserLoginLog", arg0, arg1)
	ret0, _ := ret[0].(db.UserLoginLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastUserLoginLog indicates an expected call of GetLastUserLoginLog.
func (mr *MockStoreMockRecorder) GetLastUserLoginLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastUserLoginLog", reflect.TypeOf((*MockStore)(nil).GetLastUserLoginLog), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListFriendClusterApply mocks base method.
func (m *MockStore) ListFriendClusterApply(arg0 context.Context, arg1 int32) ([]db.ListFriendClusterApplyRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFriendClusterApply", arg0, arg1)
	ret0, _ := ret[0].([]db.ListFriendClusterApplyRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFriendClusterApply indicates an expected call of ListFriendClusterApply.
func (mr *MockStoreMockRecorder) ListFriendClusterApply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFriendClusterApply", reflect.TypeOf((*MockStore)(nil).ListFriendClusterApply), arg0, arg1)
}

// ListFriends mocks base method.
func (m *MockStore) ListFriends(arg0 context.Context, arg1 int32) ([]db.ListFriendsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFriends", arg0, arg1)
	ret0, _ := ret[0].([]db.ListFriendsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFriends indicates an expected call of ListFriends.
func (mr *MockStoreMockRecorder) ListFriends(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFriends", reflect.TypeOf((*MockStore)(nil).ListFriends), arg0, arg1)
}

// UpdateFriendClusterApply mocks base method.
func (m *MockStore) UpdateFriendClusterApply(arg0 context.Context, arg1 *db.UpdateFriendClusterApplyParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFriendClusterApply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFriendClusterApply indicates an expected call of UpdateFriendClusterApply.
func (mr *MockStoreMockRecorder) UpdateFriendClusterApply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFriendClusterApply", reflect.TypeOf((*MockStore)(nil).UpdateFriendClusterApply), arg0, arg1)
}

// UpdateFriendNote mocks base method.
func (m *MockStore) UpdateFriendNote(arg0 context.Context, arg1 *db.UpdateFriendNoteParams) (db.Friend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFriendNote", arg0, arg1)
	ret0, _ := ret[0].(db.Friend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFriendNote indicates an expected call of UpdateFriendNote.
func (mr *MockStoreMockRecorder) UpdateFriendNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFriendNote", reflect.TypeOf((*MockStore)(nil).UpdateFriendNote), arg0, arg1)
}

// UpdateLastUserLoginLog mocks base method.
func (m *MockStore) UpdateLastUserLoginLog(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastUserLoginLog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastUserLoginLog indicates an expected call of UpdateLastUserLoginLog.
func (mr *MockStoreMockRecorder) UpdateLastUserLoginLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastUserLoginLog", reflect.TypeOf((*MockStore)(nil).UpdateLastUserLoginLog), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 *db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

// UpdateUserPassword mocks base method.
func (m *MockStore) UpdateUserPassword(arg0 context.Context, arg1 *db.UpdateUserPasswordParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockStoreMockRecorder) UpdateUserPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockStore)(nil).UpdateUserPassword), arg0, arg1)
}

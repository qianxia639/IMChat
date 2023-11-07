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
func (m *MockStore) AddFriend(arg0 context.Context, arg1 *db.AddFriendParams) (db.Friendship, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFriend", arg0, arg1)
	ret0, _ := ret[0].(db.Friendship)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFriend indicates an expected call of AddFriend.
func (mr *MockStoreMockRecorder) AddFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFriend", reflect.TypeOf((*MockStore)(nil).AddFriend), arg0, arg1)
}

// AddGroupMember mocks base method.
func (m *MockStore) AddGroupMember(arg0 context.Context, arg1 *db.AddGroupMemberParams) (db.GroupMemberShip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGroupMember", arg0, arg1)
	ret0, _ := ret[0].(db.GroupMemberShip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddGroupMember indicates an expected call of AddGroupMember.
func (mr *MockStoreMockRecorder) AddGroupMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGroupMember", reflect.TypeOf((*MockStore)(nil).AddGroupMember), arg0, arg1)
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

// AddMessageTx mocks base method.
func (m *MockStore) AddMessageTx(arg0 context.Context, arg1 *db.AddMessageTxParams) (db.AddMessageTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMessageTx", arg0, arg1)
	ret0, _ := ret[0].(db.AddMessageTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMessageTx indicates an expected call of AddMessageTx.
func (mr *MockStoreMockRecorder) AddMessageTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMessageTx", reflect.TypeOf((*MockStore)(nil).AddMessageTx), arg0, arg1)
}

// CreateFriendGroupApply mocks base method.
func (m *MockStore) CreateFriendGroupApply(arg0 context.Context, arg1 *db.CreateFriendGroupApplyParams) (db.FriendGroupApply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFriendGroupApply", arg0, arg1)
	ret0, _ := ret[0].(db.FriendGroupApply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFriendGroupApply indicates an expected call of CreateFriendGroupApply.
func (mr *MockStoreMockRecorder) CreateFriendGroupApply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFriendGroupApply", reflect.TypeOf((*MockStore)(nil).CreateFriendGroupApply), arg0, arg1)
}

// CreateGroup mocks base method.
func (m *MockStore) CreateGroup(arg0 context.Context, arg1 *db.CreateGroupParams) (db.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", arg0, arg1)
	ret0, _ := ret[0].(db.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockStoreMockRecorder) CreateGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockStore)(nil).CreateGroup), arg0, arg1)
}

// CreateGroupTx mocks base method.
func (m *MockStore) CreateGroupTx(arg0 context.Context, arg1 *db.CreateGroupTxParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroupTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGroupTx indicates an expected call of CreateGroupTx.
func (mr *MockStoreMockRecorder) CreateGroupTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroupTx", reflect.TypeOf((*MockStore)(nil).CreateGroupTx), arg0, arg1)
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

// DeleteFriendGroupApply mocks base method.
func (m *MockStore) DeleteFriendGroupApply(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFriendGroupApply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFriendGroupApply indicates an expected call of DeleteFriendGroupApply.
func (mr *MockStoreMockRecorder) DeleteFriendGroupApply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFriendGroupApply", reflect.TypeOf((*MockStore)(nil).DeleteFriendGroupApply), arg0, arg1)
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

// DeleteUserTx mocks base method.
func (m *MockStore) DeleteUserTx(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserTx indicates an expected call of DeleteUserTx.
func (mr *MockStoreMockRecorder) DeleteUserTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserTx", reflect.TypeOf((*MockStore)(nil).DeleteUserTx), arg0, arg1)
}

// ExistEmail mocks base method.
func (m *MockStore) ExistEmail(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistEmail", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistEmail indicates an expected call of ExistEmail.
func (mr *MockStoreMockRecorder) ExistEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistEmail", reflect.TypeOf((*MockStore)(nil).ExistEmail), arg0, arg1)
}

// ExistNickname mocks base method.
func (m *MockStore) ExistNickname(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistNickname", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistNickname indicates an expected call of ExistNickname.
func (mr *MockStoreMockRecorder) ExistNickname(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistNickname", reflect.TypeOf((*MockStore)(nil).ExistNickname), arg0, arg1)
}

// ExistsFriendGroupApply mocks base method.
func (m *MockStore) ExistsFriendGroupApply(arg0 context.Context, arg1 *db.ExistsFriendGroupApplyParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsFriendGroupApply", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsFriendGroupApply indicates an expected call of ExistsFriendGroupApply.
func (mr *MockStoreMockRecorder) ExistsFriendGroupApply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsFriendGroupApply", reflect.TypeOf((*MockStore)(nil).ExistsFriendGroupApply), arg0, arg1)
}

// GetFriend mocks base method.
func (m *MockStore) GetFriend(arg0 context.Context, arg1 *db.GetFriendParams) (db.Friendship, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriend", arg0, arg1)
	ret0, _ := ret[0].(db.Friendship)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriend indicates an expected call of GetFriend.
func (mr *MockStoreMockRecorder) GetFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriend", reflect.TypeOf((*MockStore)(nil).GetFriend), arg0, arg1)
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

// GetUserById mocks base method.
func (m *MockStore) GetUserById(arg0 context.Context, arg1 int32) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockStoreMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStore)(nil).GetUserById), arg0, arg1)
}

// ListFriendGroupApply mocks base method.
func (m *MockStore) ListFriendGroupApply(arg0 context.Context, arg1 int32) ([]db.ListFriendGroupApplyRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFriendGroupApply", arg0, arg1)
	ret0, _ := ret[0].([]db.ListFriendGroupApplyRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFriendGroupApply indicates an expected call of ListFriendGroupApply.
func (mr *MockStoreMockRecorder) ListFriendGroupApply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFriendGroupApply", reflect.TypeOf((*MockStore)(nil).ListFriendGroupApply), arg0, arg1)
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

// ReplyFriendGroupApplyTx mocks base method.
func (m *MockStore) ReplyFriendGroupApplyTx(arg0 context.Context, arg1 *db.ReplyFriendGroupApplyTxParams) (db.Friendship, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplyFriendGroupApplyTx", arg0, arg1)
	ret0, _ := ret[0].(db.Friendship)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplyFriendGroupApplyTx indicates an expected call of ReplyFriendGroupApplyTx.
func (mr *MockStoreMockRecorder) ReplyFriendGroupApplyTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyFriendGroupApplyTx", reflect.TypeOf((*MockStore)(nil).ReplyFriendGroupApplyTx), arg0, arg1)
}

// UpdateFriendComment mocks base method.
func (m *MockStore) UpdateFriendComment(arg0 context.Context, arg1 *db.UpdateFriendCommentParams) (db.Friendship, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFriendComment", arg0, arg1)
	ret0, _ := ret[0].(db.Friendship)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFriendComment indicates an expected call of UpdateFriendComment.
func (mr *MockStoreMockRecorder) UpdateFriendComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFriendComment", reflect.TypeOf((*MockStore)(nil).UpdateFriendComment), arg0, arg1)
}

// UpdateFriendGroupApply mocks base method.
func (m *MockStore) UpdateFriendGroupApply(arg0 context.Context, arg1 *db.UpdateFriendGroupApplyParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFriendGroupApply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFriendGroupApply indicates an expected call of UpdateFriendGroupApply.
func (mr *MockStoreMockRecorder) UpdateFriendGroupApply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFriendGroupApply", reflect.TypeOf((*MockStore)(nil).UpdateFriendGroupApply), arg0, arg1)
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

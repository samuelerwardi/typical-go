// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/typical-go/typical-rest-server/internal/app/repo (interfaces: BookRepo)

// Package repo_mock is a generated GoMock package.
package repo_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repo "github.com/typical-go/typical-rest-server/internal/app/repo"
	sqkit "github.com/typical-go/typical-rest-server/pkg/sqkit"
	reflect "reflect"
)

// MockBookRepo is a mock of BookRepo interface
type MockBookRepo struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepoMockRecorder
}

// MockBookRepoMockRecorder is the mock recorder for MockBookRepo
type MockBookRepoMockRecorder struct {
	mock *MockBookRepo
}

// NewMockBookRepo creates a new mock instance
func NewMockBookRepo(ctrl *gomock.Controller) *MockBookRepo {
	mock := &MockBookRepo{ctrl: ctrl}
	mock.recorder = &MockBookRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookRepo) EXPECT() *MockBookRepoMockRecorder {
	return m.recorder
}

// BulkInsert mocks base method
func (m *MockBookRepo) BulkInsert(arg0 context.Context, arg1 ...*repo.Book) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BulkInsert", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkInsert indicates an expected call of BulkInsert
func (mr *MockBookRepoMockRecorder) BulkInsert(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkInsert", reflect.TypeOf((*MockBookRepo)(nil).BulkInsert), varargs...)
}

// Count mocks base method
func (m *MockBookRepo) Count(arg0 context.Context, arg1 ...sqkit.SelectOption) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockBookRepoMockRecorder) Count(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockBookRepo)(nil).Count), varargs...)
}

// Delete mocks base method
func (m *MockBookRepo) Delete(arg0 context.Context, arg1 ...sqkit.DeleteOption) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockBookRepoMockRecorder) Delete(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBookRepo)(nil).Delete), varargs...)
}

// Find mocks base method
func (m *MockBookRepo) Find(arg0 context.Context, arg1 ...sqkit.SelectOption) ([]*repo.Book, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].([]*repo.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockBookRepoMockRecorder) Find(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockBookRepo)(nil).Find), varargs...)
}

// Insert mocks base method
func (m *MockBookRepo) Insert(arg0 context.Context, arg1 *repo.Book) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockBookRepoMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockBookRepo)(nil).Insert), arg0, arg1)
}

// Patch mocks base method
func (m *MockBookRepo) Patch(arg0 context.Context, arg1 *repo.Book, arg2 ...sqkit.UpdateOption) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockBookRepoMockRecorder) Patch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockBookRepo)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockBookRepo) Update(arg0 context.Context, arg1 *repo.Book, arg2 ...sqkit.UpdateOption) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockBookRepoMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookRepo)(nil).Update), varargs...)
}

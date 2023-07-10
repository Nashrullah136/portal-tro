// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	filesystem "nashrul-be/crm/utils/filesystem"

	mock "github.com/stretchr/testify/mock"
)

// Folder is an autogenerated mock type for the Folder type
type Folder struct {
	mock.Mock
}

type Folder_Expecter struct {
	mock *mock.Mock
}

func (_m *Folder) EXPECT() *Folder_Expecter {
	return &Folder_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: filename
func (_m *Folder) Create(filename string) (filesystem.File, error) {
	ret := _m.Called(filename)

	var r0 filesystem.File
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (filesystem.File, error)); ok {
		return rf(filename)
	}
	if rf, ok := ret.Get(0).(func(string) filesystem.File); ok {
		r0 = rf(filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filesystem.File)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Folder_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Folder_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - filename string
func (_e *Folder_Expecter) Create(filename interface{}) *Folder_Create_Call {
	return &Folder_Create_Call{Call: _e.mock.On("Create", filename)}
}

func (_c *Folder_Create_Call) Run(run func(filename string)) *Folder_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Folder_Create_Call) Return(_a0 filesystem.File, _a1 error) *Folder_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Folder_Create_Call) RunAndReturn(run func(string) (filesystem.File, error)) *Folder_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetPath provides a mock function with given fields: filename
func (_m *Folder) GetPath(filename string) string {
	ret := _m.Called(filename)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(filename)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Folder_GetPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPath'
type Folder_GetPath_Call struct {
	*mock.Call
}

// GetPath is a helper method to define mock.On call
//   - filename string
func (_e *Folder_Expecter) GetPath(filename interface{}) *Folder_GetPath_Call {
	return &Folder_GetPath_Call{Call: _e.mock.On("GetPath", filename)}
}

func (_c *Folder_GetPath_Call) Run(run func(filename string)) *Folder_GetPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Folder_GetPath_Call) Return(_a0 string) *Folder_GetPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Folder_GetPath_Call) RunAndReturn(run func(string) string) *Folder_GetPath_Call {
	_c.Call.Return(run)
	return _c
}

// IsExist provides a mock function with given fields: filename
func (_m *Folder) IsExist(filename string) bool {
	ret := _m.Called(filename)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(filename)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Folder_IsExist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsExist'
type Folder_IsExist_Call struct {
	*mock.Call
}

// IsExist is a helper method to define mock.On call
//   - filename string
func (_e *Folder_Expecter) IsExist(filename interface{}) *Folder_IsExist_Call {
	return &Folder_IsExist_Call{Call: _e.mock.On("IsExist", filename)}
}

func (_c *Folder_IsExist_Call) Run(run func(filename string)) *Folder_IsExist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Folder_IsExist_Call) Return(_a0 bool) *Folder_IsExist_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Folder_IsExist_Call) RunAndReturn(run func(string) bool) *Folder_IsExist_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: filename
func (_m *Folder) Remove(filename string) error {
	ret := _m.Called(filename)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Folder_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type Folder_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - filename string
func (_e *Folder_Expecter) Remove(filename interface{}) *Folder_Remove_Call {
	return &Folder_Remove_Call{Call: _e.mock.On("Remove", filename)}
}

func (_c *Folder_Remove_Call) Run(run func(filename string)) *Folder_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Folder_Remove_Call) Return(_a0 error) *Folder_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Folder_Remove_Call) RunAndReturn(run func(string) error) *Folder_Remove_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewFolder interface {
	mock.TestingT
	Cleanup(func())
}

// NewFolder creates a new instance of Folder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFolder(t mockConstructorTestingTNewFolder) *Folder {
	mock := &Folder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
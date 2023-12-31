// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	os "os"

	mock "github.com/stretchr/testify/mock"
)

// File is an autogenerated mock type for the File type
type File struct {
	mock.Mock
}

type File_Expecter struct {
	mock *mock.Mock
}

func (_m *File) EXPECT() *File_Expecter {
	return &File_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *File) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// File_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type File_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *File_Expecter) Close() *File_Close_Call {
	return &File_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *File_Close_Call) Run(run func()) *File_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *File_Close_Call) Return(_a0 error) *File_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *File_Close_Call) RunAndReturn(run func() error) *File_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Filename provides a mock function with given fields:
func (_m *File) Filename() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// File_Filename_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Filename'
type File_Filename_Call struct {
	*mock.Call
}

// Filename is a helper method to define mock.On call
func (_e *File_Expecter) Filename() *File_Filename_Call {
	return &File_Filename_Call{Call: _e.mock.On("Filename")}
}

func (_c *File_Filename_Call) Run(run func()) *File_Filename_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *File_Filename_Call) Return(_a0 string) *File_Filename_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *File_Filename_Call) RunAndReturn(run func() string) *File_Filename_Call {
	_c.Call.Return(run)
	return _c
}

// Open provides a mock function with given fields:
func (_m *File) Open() (*os.File, error) {
	ret := _m.Called()

	var r0 *os.File
	var r1 error
	if rf, ok := ret.Get(0).(func() (*os.File, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *os.File); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// File_Open_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Open'
type File_Open_Call struct {
	*mock.Call
}

// Open is a helper method to define mock.On call
func (_e *File_Expecter) Open() *File_Open_Call {
	return &File_Open_Call{Call: _e.mock.On("Open")}
}

func (_c *File_Open_Call) Run(run func()) *File_Open_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *File_Open_Call) Return(_a0 *os.File, _a1 error) *File_Open_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *File_Open_Call) RunAndReturn(run func() (*os.File, error)) *File_Open_Call {
	_c.Call.Return(run)
	return _c
}

// Path provides a mock function with given fields:
func (_m *File) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// File_Path_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Path'
type File_Path_Call struct {
	*mock.Call
}

// Path is a helper method to define mock.On call
func (_e *File_Expecter) Path() *File_Path_Call {
	return &File_Path_Call{Call: _e.mock.On("Path")}
}

func (_c *File_Path_Call) Run(run func()) *File_Path_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *File_Path_Call) Return(_a0 string) *File_Path_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *File_Path_Call) RunAndReturn(run func() string) *File_Path_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields:
func (_m *File) Remove() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// File_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type File_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
func (_e *File_Expecter) Remove() *File_Remove_Call {
	return &File_Remove_Call{Call: _e.mock.On("Remove")}
}

func (_c *File_Remove_Call) Run(run func()) *File_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *File_Remove_Call) Return(_a0 error) *File_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *File_Remove_Call) RunAndReturn(run func() error) *File_Remove_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewFile interface {
	mock.TestingT
	Cleanup(func())
}

// NewFile creates a new instance of File. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFile(t mockConstructorTestingTNewFile) *File {
	mock := &File{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

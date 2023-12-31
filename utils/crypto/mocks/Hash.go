// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Hash is an autogenerated mock type for the Hash type
type Hash struct {
	mock.Mock
}

type Hash_Expecter struct {
	mock *mock.Mock
}

func (_m *Hash) EXPECT() *Hash_Expecter {
	return &Hash_Expecter{mock: &_m.Mock}
}

// Compare provides a mock function with given fields: plainPassword, hashedPassword
func (_m *Hash) Compare(plainPassword string, hashedPassword string) error {
	ret := _m.Called(plainPassword, hashedPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(plainPassword, hashedPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Hash_Compare_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Compare'
type Hash_Compare_Call struct {
	*mock.Call
}

// Compare is a helper method to define mock.On call
//   - plainPassword string
//   - hashedPassword string
func (_e *Hash_Expecter) Compare(plainPassword interface{}, hashedPassword interface{}) *Hash_Compare_Call {
	return &Hash_Compare_Call{Call: _e.mock.On("Compare", plainPassword, hashedPassword)}
}

func (_c *Hash_Compare_Call) Run(run func(plainPassword string, hashedPassword string)) *Hash_Compare_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Hash_Compare_Call) Return(_a0 error) *Hash_Compare_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Hash_Compare_Call) RunAndReturn(run func(string, string) error) *Hash_Compare_Call {
	_c.Call.Return(run)
	return _c
}

// Hash provides a mock function with given fields: pwd
func (_m *Hash) Hash(pwd string) (string, error) {
	ret := _m.Called(pwd)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(pwd)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(pwd)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pwd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Hash_Hash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Hash'
type Hash_Hash_Call struct {
	*mock.Call
}

// Hash is a helper method to define mock.On call
//   - pwd string
func (_e *Hash_Expecter) Hash(pwd interface{}) *Hash_Hash_Call {
	return &Hash_Hash_Call{Call: _e.mock.On("Hash", pwd)}
}

func (_c *Hash_Hash_Call) Run(run func(pwd string)) *Hash_Hash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Hash_Hash_Call) Return(_a0 string, _a1 error) *Hash_Hash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Hash_Hash_Call) RunAndReturn(run func(string) (string, error)) *Hash_Hash_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewHash interface {
	mock.TestingT
	Cleanup(func())
}

// NewHash creates a new instance of Hash. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHash(t mockConstructorTestingTNewHash) *Hash {
	mock := &Hash{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

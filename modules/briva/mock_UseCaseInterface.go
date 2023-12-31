// Code generated by mockery v2.20.0. DO NOT EDIT.

package briva

import (
	context "context"
	entities "nashrul-be/crm/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockUseCaseInterface is an autogenerated mock type for the UseCaseInterface type
type MockUseCaseInterface struct {
	mock.Mock
}

type MockUseCaseInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUseCaseInterface) EXPECT() *MockUseCaseInterface_Expecter {
	return &MockUseCaseInterface_Expecter{mock: &_m.Mock}
}

// GetByBrivaNo provides a mock function with given fields: ctx, brivano
func (_m *MockUseCaseInterface) GetByBrivaNo(ctx context.Context, brivano string) (entities.Briva, error) {
	ret := _m.Called(ctx, brivano)

	var r0 entities.Briva
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entities.Briva, error)); ok {
		return rf(ctx, brivano)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entities.Briva); ok {
		r0 = rf(ctx, brivano)
	} else {
		r0 = ret.Get(0).(entities.Briva)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, brivano)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUseCaseInterface_GetByBrivaNo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByBrivaNo'
type MockUseCaseInterface_GetByBrivaNo_Call struct {
	*mock.Call
}

// GetByBrivaNo is a helper method to define mock.On call
//   - ctx context.Context
//   - brivano string
func (_e *MockUseCaseInterface_Expecter) GetByBrivaNo(ctx interface{}, brivano interface{}) *MockUseCaseInterface_GetByBrivaNo_Call {
	return &MockUseCaseInterface_GetByBrivaNo_Call{Call: _e.mock.On("GetByBrivaNo", ctx, brivano)}
}

func (_c *MockUseCaseInterface_GetByBrivaNo_Call) Run(run func(ctx context.Context, brivano string)) *MockUseCaseInterface_GetByBrivaNo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUseCaseInterface_GetByBrivaNo_Call) Return(_a0 entities.Briva, _a1 error) *MockUseCaseInterface_GetByBrivaNo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUseCaseInterface_GetByBrivaNo_Call) RunAndReturn(run func(context.Context, string) (entities.Briva, error)) *MockUseCaseInterface_GetByBrivaNo_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, briva
func (_m *MockUseCaseInterface) Update(ctx context.Context, briva entities.Briva) error {
	ret := _m.Called(ctx, briva)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Briva) error); ok {
		r0 = rf(ctx, briva)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUseCaseInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockUseCaseInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - briva entities.Briva
func (_e *MockUseCaseInterface_Expecter) Update(ctx interface{}, briva interface{}) *MockUseCaseInterface_Update_Call {
	return &MockUseCaseInterface_Update_Call{Call: _e.mock.On("Update", ctx, briva)}
}

func (_c *MockUseCaseInterface_Update_Call) Run(run func(ctx context.Context, briva entities.Briva)) *MockUseCaseInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.Briva))
	})
	return _c
}

func (_c *MockUseCaseInterface_Update_Call) Return(_a0 error) *MockUseCaseInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUseCaseInterface_Update_Call) RunAndReturn(run func(context.Context, entities.Briva) error) *MockUseCaseInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateBriva provides a mock function with given fields: briva, validations
func (_m *MockUseCaseInterface) ValidateBriva(briva entities.Briva, validations ...validateFunc) (error, error) {
	_va := make([]interface{}, len(validations))
	for _i := range validations {
		_va[_i] = validations[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, briva)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.Briva, ...validateFunc) (error, error)); ok {
		return rf(briva, validations...)
	}
	if rf, ok := ret.Get(0).(func(entities.Briva, ...validateFunc) error); ok {
		r0 = rf(briva, validations...)
	} else {
		r0 = ret.Error(0)
	}

	if rf, ok := ret.Get(1).(func(entities.Briva, ...validateFunc) error); ok {
		r1 = rf(briva, validations...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUseCaseInterface_ValidateBriva_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateBriva'
type MockUseCaseInterface_ValidateBriva_Call struct {
	*mock.Call
}

// ValidateBriva is a helper method to define mock.On call
//   - briva entities.Briva
//   - validations ...validateFunc
func (_e *MockUseCaseInterface_Expecter) ValidateBriva(briva interface{}, validations ...interface{}) *MockUseCaseInterface_ValidateBriva_Call {
	return &MockUseCaseInterface_ValidateBriva_Call{Call: _e.mock.On("ValidateBriva",
		append([]interface{}{briva}, validations...)...)}
}

func (_c *MockUseCaseInterface_ValidateBriva_Call) Run(run func(briva entities.Briva, validations ...validateFunc)) *MockUseCaseInterface_ValidateBriva_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]validateFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(validateFunc)
			}
		}
		run(args[0].(entities.Briva), variadicArgs...)
	})
	return _c
}

func (_c *MockUseCaseInterface_ValidateBriva_Call) Return(_a0 error, _a1 error) *MockUseCaseInterface_ValidateBriva_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUseCaseInterface_ValidateBriva_Call) RunAndReturn(run func(entities.Briva, ...validateFunc) (error, error)) *MockUseCaseInterface_ValidateBriva_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockUseCaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUseCaseInterface creates a new instance of MockUseCaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUseCaseInterface(t mockConstructorTestingTNewMockUseCaseInterface) *MockUseCaseInterface {
	mock := &MockUseCaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

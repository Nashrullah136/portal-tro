// Code generated by mockery v2.20.0. DO NOT EDIT.

package span

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

// GetByDocumentNumberPatchBankRiau provides a mock function with given fields: ctx, documentNumber
func (_m *MockUseCaseInterface) GetByDocumentNumberPatchBankRiau(ctx context.Context, documentNumber string) (entities.SPAN, error) {
	ret := _m.Called(ctx, documentNumber)

	var r0 entities.SPAN
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entities.SPAN, error)); ok {
		return rf(ctx, documentNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entities.SPAN); ok {
		r0 = rf(ctx, documentNumber)
	} else {
		r0 = ret.Get(0).(entities.SPAN)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, documentNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUseCaseInterface_GetByDocumentNumberPatchBankRiau_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByDocumentNumberPatchBankRiau'
type MockUseCaseInterface_GetByDocumentNumberPatchBankRiau_Call struct {
	*mock.Call
}

// GetByDocumentNumberPatchBankRiau is a helper method to define mock.On call
//   - ctx context.Context
//   - documentNumber string
func (_e *MockUseCaseInterface_Expecter) GetByDocumentNumberPatchBankRiau(ctx interface{}, documentNumber interface{}) *MockUseCaseInterface_GetByDocumentNumberPatchBankRiau_Call {
	return &MockUseCaseInterface_GetByDocumentNumberPatchBankRiau_Call{Call: _e.mock.On("GetByDocumentNumberPatchBankRiau", ctx, documentNumber)}
}

func (_c *MockUseCaseInterface_GetByDocumentNumberPatchBankRiau_Call) Run(run func(ctx context.Context, documentNumber string)) *MockUseCaseInterface_GetByDocumentNumberPatchBankRiau_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUseCaseInterface_GetByDocumentNumberPatchBankRiau_Call) Return(_a0 entities.SPAN, _a1 error) *MockUseCaseInterface_GetByDocumentNumberPatchBankRiau_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUseCaseInterface_GetByDocumentNumberPatchBankRiau_Call) RunAndReturn(run func(context.Context, string) (entities.SPAN, error)) *MockUseCaseInterface_GetByDocumentNumberPatchBankRiau_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePatchBankRiau provides a mock function with given fields: ctx, span
func (_m *MockUseCaseInterface) UpdatePatchBankRiau(ctx context.Context, span entities.SPAN) error {
	ret := _m.Called(ctx, span)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.SPAN) error); ok {
		r0 = rf(ctx, span)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUseCaseInterface_UpdatePatchBankRiau_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePatchBankRiau'
type MockUseCaseInterface_UpdatePatchBankRiau_Call struct {
	*mock.Call
}

// UpdatePatchBankRiau is a helper method to define mock.On call
//   - ctx context.Context
//   - span entities.SPAN
func (_e *MockUseCaseInterface_Expecter) UpdatePatchBankRiau(ctx interface{}, span interface{}) *MockUseCaseInterface_UpdatePatchBankRiau_Call {
	return &MockUseCaseInterface_UpdatePatchBankRiau_Call{Call: _e.mock.On("UpdatePatchBankRiau", ctx, span)}
}

func (_c *MockUseCaseInterface_UpdatePatchBankRiau_Call) Run(run func(ctx context.Context, span entities.SPAN)) *MockUseCaseInterface_UpdatePatchBankRiau_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.SPAN))
	})
	return _c
}

func (_c *MockUseCaseInterface_UpdatePatchBankRiau_Call) Return(_a0 error) *MockUseCaseInterface_UpdatePatchBankRiau_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUseCaseInterface_UpdatePatchBankRiau_Call) RunAndReturn(run func(context.Context, entities.SPAN) error) *MockUseCaseInterface_UpdatePatchBankRiau_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateSpan provides a mock function with given fields: span, validations
func (_m *MockUseCaseInterface) ValidateSpan(span entities.SPAN, validations ...validateFunc) (error, error) {
	_va := make([]interface{}, len(validations))
	for _i := range validations {
		_va[_i] = validations[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, span)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.SPAN, ...validateFunc) (error, error)); ok {
		return rf(span, validations...)
	}
	if rf, ok := ret.Get(0).(func(entities.SPAN, ...validateFunc) error); ok {
		r0 = rf(span, validations...)
	} else {
		r0 = ret.Error(0)
	}

	if rf, ok := ret.Get(1).(func(entities.SPAN, ...validateFunc) error); ok {
		r1 = rf(span, validations...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUseCaseInterface_ValidateSpan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateSpan'
type MockUseCaseInterface_ValidateSpan_Call struct {
	*mock.Call
}

// ValidateSpan is a helper method to define mock.On call
//   - span entities.SPAN
//   - validations ...validateFunc
func (_e *MockUseCaseInterface_Expecter) ValidateSpan(span interface{}, validations ...interface{}) *MockUseCaseInterface_ValidateSpan_Call {
	return &MockUseCaseInterface_ValidateSpan_Call{Call: _e.mock.On("ValidateSpan",
		append([]interface{}{span}, validations...)...)}
}

func (_c *MockUseCaseInterface_ValidateSpan_Call) Run(run func(span entities.SPAN, validations ...validateFunc)) *MockUseCaseInterface_ValidateSpan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]validateFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(validateFunc)
			}
		}
		run(args[0].(entities.SPAN), variadicArgs...)
	})
	return _c
}

func (_c *MockUseCaseInterface_ValidateSpan_Call) Return(_a0 error, _a1 error) *MockUseCaseInterface_ValidateSpan_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUseCaseInterface_ValidateSpan_Call) RunAndReturn(run func(entities.SPAN, ...validateFunc) (error, error)) *MockUseCaseInterface_ValidateSpan_Call {
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

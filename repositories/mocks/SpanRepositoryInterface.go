// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "nashrul-be/crm/entities"
	db "nashrul-be/crm/utils/db"

	mock "github.com/stretchr/testify/mock"

	repositories "nashrul-be/crm/repositories"
)

// SpanRepositoryInterface is an autogenerated mock type for the SpanRepositoryInterface type
type SpanRepositoryInterface struct {
	mock.Mock
}

type SpanRepositoryInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *SpanRepositoryInterface) EXPECT() *SpanRepositoryInterface_Expecter {
	return &SpanRepositoryInterface_Expecter{mock: &_m.Mock}
}

// Begin provides a mock function with given fields:
func (_m *SpanRepositoryInterface) Begin() db.Transactor {
	ret := _m.Called()

	var r0 db.Transactor
	if rf, ok := ret.Get(0).(func() db.Transactor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Transactor)
		}
	}

	return r0
}

// SpanRepositoryInterface_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type SpanRepositoryInterface_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
func (_e *SpanRepositoryInterface_Expecter) Begin() *SpanRepositoryInterface_Begin_Call {
	return &SpanRepositoryInterface_Begin_Call{Call: _e.mock.On("Begin")}
}

func (_c *SpanRepositoryInterface_Begin_Call) Run(run func()) *SpanRepositoryInterface_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SpanRepositoryInterface_Begin_Call) Return(_a0 db.Transactor) *SpanRepositoryInterface_Begin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SpanRepositoryInterface_Begin_Call) RunAndReturn(run func() db.Transactor) *SpanRepositoryInterface_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// GetBySpanDocumentNumber provides a mock function with given fields: ctx, documentNumber
func (_m *SpanRepositoryInterface) GetBySpanDocumentNumber(ctx context.Context, documentNumber string) (entities.SPAN, error) {
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

// SpanRepositoryInterface_GetBySpanDocumentNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBySpanDocumentNumber'
type SpanRepositoryInterface_GetBySpanDocumentNumber_Call struct {
	*mock.Call
}

// GetBySpanDocumentNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - documentNumber string
func (_e *SpanRepositoryInterface_Expecter) GetBySpanDocumentNumber(ctx interface{}, documentNumber interface{}) *SpanRepositoryInterface_GetBySpanDocumentNumber_Call {
	return &SpanRepositoryInterface_GetBySpanDocumentNumber_Call{Call: _e.mock.On("GetBySpanDocumentNumber", ctx, documentNumber)}
}

func (_c *SpanRepositoryInterface_GetBySpanDocumentNumber_Call) Run(run func(ctx context.Context, documentNumber string)) *SpanRepositoryInterface_GetBySpanDocumentNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *SpanRepositoryInterface_GetBySpanDocumentNumber_Call) Return(span entities.SPAN, err error) *SpanRepositoryInterface_GetBySpanDocumentNumber_Call {
	_c.Call.Return(span, err)
	return _c
}

func (_c *SpanRepositoryInterface_GetBySpanDocumentNumber_Call) RunAndReturn(run func(context.Context, string) (entities.SPAN, error)) *SpanRepositoryInterface_GetBySpanDocumentNumber_Call {
	_c.Call.Return(run)
	return _c
}

// IsSpanExist provides a mock function with given fields: span
func (_m *SpanRepositoryInterface) IsSpanExist(span entities.SPAN) (bool, error) {
	ret := _m.Called(span)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.SPAN) (bool, error)); ok {
		return rf(span)
	}
	if rf, ok := ret.Get(0).(func(entities.SPAN) bool); ok {
		r0 = rf(span)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(entities.SPAN) error); ok {
		r1 = rf(span)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SpanRepositoryInterface_IsSpanExist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSpanExist'
type SpanRepositoryInterface_IsSpanExist_Call struct {
	*mock.Call
}

// IsSpanExist is a helper method to define mock.On call
//   - span entities.SPAN
func (_e *SpanRepositoryInterface_Expecter) IsSpanExist(span interface{}) *SpanRepositoryInterface_IsSpanExist_Call {
	return &SpanRepositoryInterface_IsSpanExist_Call{Call: _e.mock.On("IsSpanExist", span)}
}

func (_c *SpanRepositoryInterface_IsSpanExist_Call) Run(run func(span entities.SPAN)) *SpanRepositoryInterface_IsSpanExist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.SPAN))
	})
	return _c
}

func (_c *SpanRepositoryInterface_IsSpanExist_Call) Return(exist bool, err error) *SpanRepositoryInterface_IsSpanExist_Call {
	_c.Call.Return(exist, err)
	return _c
}

func (_c *SpanRepositoryInterface_IsSpanExist_Call) RunAndReturn(run func(entities.SPAN) (bool, error)) *SpanRepositoryInterface_IsSpanExist_Call {
	_c.Call.Return(run)
	return _c
}

// MakeAuditUpdate provides a mock function with given fields: ctx, span
func (_m *SpanRepositoryInterface) MakeAuditUpdate(ctx context.Context, span entities.SPAN) (entities.Audit, error) {
	ret := _m.Called(ctx, span)

	var r0 entities.Audit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.SPAN) (entities.Audit, error)); ok {
		return rf(ctx, span)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.SPAN) entities.Audit); ok {
		r0 = rf(ctx, span)
	} else {
		r0 = ret.Get(0).(entities.Audit)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.SPAN) error); ok {
		r1 = rf(ctx, span)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SpanRepositoryInterface_MakeAuditUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MakeAuditUpdate'
type SpanRepositoryInterface_MakeAuditUpdate_Call struct {
	*mock.Call
}

// MakeAuditUpdate is a helper method to define mock.On call
//   - ctx context.Context
//   - span entities.SPAN
func (_e *SpanRepositoryInterface_Expecter) MakeAuditUpdate(ctx interface{}, span interface{}) *SpanRepositoryInterface_MakeAuditUpdate_Call {
	return &SpanRepositoryInterface_MakeAuditUpdate_Call{Call: _e.mock.On("MakeAuditUpdate", ctx, span)}
}

func (_c *SpanRepositoryInterface_MakeAuditUpdate_Call) Run(run func(ctx context.Context, span entities.SPAN)) *SpanRepositoryInterface_MakeAuditUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.SPAN))
	})
	return _c
}

func (_c *SpanRepositoryInterface_MakeAuditUpdate_Call) Return(_a0 entities.Audit, _a1 error) *SpanRepositoryInterface_MakeAuditUpdate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SpanRepositoryInterface_MakeAuditUpdate_Call) RunAndReturn(run func(context.Context, entities.SPAN) (entities.Audit, error)) *SpanRepositoryInterface_MakeAuditUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// MakeAuditUpdateWithOldData provides a mock function with given fields: ctx, oldSpan, newSpan
func (_m *SpanRepositoryInterface) MakeAuditUpdateWithOldData(ctx context.Context, oldSpan entities.SPAN, newSpan entities.SPAN) (entities.Audit, error) {
	ret := _m.Called(ctx, oldSpan, newSpan)

	var r0 entities.Audit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.SPAN, entities.SPAN) (entities.Audit, error)); ok {
		return rf(ctx, oldSpan, newSpan)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.SPAN, entities.SPAN) entities.Audit); ok {
		r0 = rf(ctx, oldSpan, newSpan)
	} else {
		r0 = ret.Get(0).(entities.Audit)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.SPAN, entities.SPAN) error); ok {
		r1 = rf(ctx, oldSpan, newSpan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SpanRepositoryInterface_MakeAuditUpdateWithOldData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MakeAuditUpdateWithOldData'
type SpanRepositoryInterface_MakeAuditUpdateWithOldData_Call struct {
	*mock.Call
}

// MakeAuditUpdateWithOldData is a helper method to define mock.On call
//   - ctx context.Context
//   - oldSpan entities.SPAN
//   - newSpan entities.SPAN
func (_e *SpanRepositoryInterface_Expecter) MakeAuditUpdateWithOldData(ctx interface{}, oldSpan interface{}, newSpan interface{}) *SpanRepositoryInterface_MakeAuditUpdateWithOldData_Call {
	return &SpanRepositoryInterface_MakeAuditUpdateWithOldData_Call{Call: _e.mock.On("MakeAuditUpdateWithOldData", ctx, oldSpan, newSpan)}
}

func (_c *SpanRepositoryInterface_MakeAuditUpdateWithOldData_Call) Run(run func(ctx context.Context, oldSpan entities.SPAN, newSpan entities.SPAN)) *SpanRepositoryInterface_MakeAuditUpdateWithOldData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.SPAN), args[2].(entities.SPAN))
	})
	return _c
}

func (_c *SpanRepositoryInterface_MakeAuditUpdateWithOldData_Call) Return(_a0 entities.Audit, _a1 error) *SpanRepositoryInterface_MakeAuditUpdateWithOldData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SpanRepositoryInterface_MakeAuditUpdateWithOldData_Call) RunAndReturn(run func(context.Context, entities.SPAN, entities.SPAN) (entities.Audit, error)) *SpanRepositoryInterface_MakeAuditUpdateWithOldData_Call {
	_c.Call.Return(run)
	return _c
}

// New provides a mock function with given fields: transact
func (_m *SpanRepositoryInterface) New(transact db.Transactor) repositories.SpanRepositoryInterface {
	ret := _m.Called(transact)

	var r0 repositories.SpanRepositoryInterface
	if rf, ok := ret.Get(0).(func(db.Transactor) repositories.SpanRepositoryInterface); ok {
		r0 = rf(transact)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repositories.SpanRepositoryInterface)
		}
	}

	return r0
}

// SpanRepositoryInterface_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type SpanRepositoryInterface_New_Call struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - transact db.Transactor
func (_e *SpanRepositoryInterface_Expecter) New(transact interface{}) *SpanRepositoryInterface_New_Call {
	return &SpanRepositoryInterface_New_Call{Call: _e.mock.On("New", transact)}
}

func (_c *SpanRepositoryInterface_New_Call) Run(run func(transact db.Transactor)) *SpanRepositoryInterface_New_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(db.Transactor))
	})
	return _c
}

func (_c *SpanRepositoryInterface_New_Call) Return(_a0 repositories.SpanRepositoryInterface) *SpanRepositoryInterface_New_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SpanRepositoryInterface_New_Call) RunAndReturn(run func(db.Transactor) repositories.SpanRepositoryInterface) *SpanRepositoryInterface_New_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, span
func (_m *SpanRepositoryInterface) Update(ctx context.Context, span entities.SPAN) error {
	ret := _m.Called(ctx, span)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.SPAN) error); ok {
		r0 = rf(ctx, span)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SpanRepositoryInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type SpanRepositoryInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - span entities.SPAN
func (_e *SpanRepositoryInterface_Expecter) Update(ctx interface{}, span interface{}) *SpanRepositoryInterface_Update_Call {
	return &SpanRepositoryInterface_Update_Call{Call: _e.mock.On("Update", ctx, span)}
}

func (_c *SpanRepositoryInterface_Update_Call) Run(run func(ctx context.Context, span entities.SPAN)) *SpanRepositoryInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.SPAN))
	})
	return _c
}

func (_c *SpanRepositoryInterface_Update_Call) Return(_a0 error) *SpanRepositoryInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SpanRepositoryInterface_Update_Call) RunAndReturn(run func(context.Context, entities.SPAN) error) *SpanRepositoryInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSpanRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewSpanRepositoryInterface creates a new instance of SpanRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSpanRepositoryInterface(t mockConstructorTestingTNewSpanRepositoryInterface) *SpanRepositoryInterface {
	mock := &SpanRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

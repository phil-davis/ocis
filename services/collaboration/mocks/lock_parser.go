// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LockParser is an autogenerated mock type for the LockParser type
type LockParser struct {
	mock.Mock
}

type LockParser_Expecter struct {
	mock *mock.Mock
}

func (_m *LockParser) EXPECT() *LockParser_Expecter {
	return &LockParser_Expecter{mock: &_m.Mock}
}

// ParseLock provides a mock function with given fields: id
func (_m *LockParser) ParseLock(id string) string {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for ParseLock")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LockParser_ParseLock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseLock'
type LockParser_ParseLock_Call struct {
	*mock.Call
}

// ParseLock is a helper method to define mock.On call
//   - id string
func (_e *LockParser_Expecter) ParseLock(id interface{}) *LockParser_ParseLock_Call {
	return &LockParser_ParseLock_Call{Call: _e.mock.On("ParseLock", id)}
}

func (_c *LockParser_ParseLock_Call) Run(run func(id string)) *LockParser_ParseLock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *LockParser_ParseLock_Call) Return(_a0 string) *LockParser_ParseLock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LockParser_ParseLock_Call) RunAndReturn(run func(string) string) *LockParser_ParseLock_Call {
	_c.Call.Return(run)
	return _c
}

// NewLockParser creates a new instance of LockParser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLockParser(t interface {
	mock.TestingT
	Cleanup(func())
}) *LockParser {
	mock := &LockParser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import (
	context "context"

	providerv1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	mock "github.com/stretchr/testify/mock"

	userv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"

	v0 "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/search/v0"
)

// Searcher is an autogenerated mock type for the Searcher type
type Searcher struct {
	mock.Mock
}

type Searcher_Expecter struct {
	mock *mock.Mock
}

func (_m *Searcher) EXPECT() *Searcher_Expecter {
	return &Searcher_Expecter{mock: &_m.Mock}
}

// IndexSpace provides a mock function with given fields: rID, uID
func (_m *Searcher) IndexSpace(rID *providerv1beta1.StorageSpaceId, uID *userv1beta1.UserId) error {
	ret := _m.Called(rID, uID)

	if len(ret) == 0 {
		panic("no return value specified for IndexSpace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*providerv1beta1.StorageSpaceId, *userv1beta1.UserId) error); ok {
		r0 = rf(rID, uID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Searcher_IndexSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IndexSpace'
type Searcher_IndexSpace_Call struct {
	*mock.Call
}

// IndexSpace is a helper method to define mock.On call
//   - rID *providerv1beta1.StorageSpaceId
//   - uID *userv1beta1.UserId
func (_e *Searcher_Expecter) IndexSpace(rID interface{}, uID interface{}) *Searcher_IndexSpace_Call {
	return &Searcher_IndexSpace_Call{Call: _e.mock.On("IndexSpace", rID, uID)}
}

func (_c *Searcher_IndexSpace_Call) Run(run func(rID *providerv1beta1.StorageSpaceId, uID *userv1beta1.UserId)) *Searcher_IndexSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*providerv1beta1.StorageSpaceId), args[1].(*userv1beta1.UserId))
	})
	return _c
}

func (_c *Searcher_IndexSpace_Call) Return(_a0 error) *Searcher_IndexSpace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Searcher_IndexSpace_Call) RunAndReturn(run func(*providerv1beta1.StorageSpaceId, *userv1beta1.UserId) error) *Searcher_IndexSpace_Call {
	_c.Call.Return(run)
	return _c
}

// MoveItem provides a mock function with given fields: ref, uID
func (_m *Searcher) MoveItem(ref *providerv1beta1.Reference, uID *userv1beta1.UserId) {
	_m.Called(ref, uID)
}

// Searcher_MoveItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MoveItem'
type Searcher_MoveItem_Call struct {
	*mock.Call
}

// MoveItem is a helper method to define mock.On call
//   - ref *providerv1beta1.Reference
//   - uID *userv1beta1.UserId
func (_e *Searcher_Expecter) MoveItem(ref interface{}, uID interface{}) *Searcher_MoveItem_Call {
	return &Searcher_MoveItem_Call{Call: _e.mock.On("MoveItem", ref, uID)}
}

func (_c *Searcher_MoveItem_Call) Run(run func(ref *providerv1beta1.Reference, uID *userv1beta1.UserId)) *Searcher_MoveItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*providerv1beta1.Reference), args[1].(*userv1beta1.UserId))
	})
	return _c
}

func (_c *Searcher_MoveItem_Call) Return() *Searcher_MoveItem_Call {
	_c.Call.Return()
	return _c
}

func (_c *Searcher_MoveItem_Call) RunAndReturn(run func(*providerv1beta1.Reference, *userv1beta1.UserId)) *Searcher_MoveItem_Call {
	_c.Call.Return(run)
	return _c
}

// RestoreItem provides a mock function with given fields: ref, uID
func (_m *Searcher) RestoreItem(ref *providerv1beta1.Reference, uID *userv1beta1.UserId) {
	_m.Called(ref, uID)
}

// Searcher_RestoreItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RestoreItem'
type Searcher_RestoreItem_Call struct {
	*mock.Call
}

// RestoreItem is a helper method to define mock.On call
//   - ref *providerv1beta1.Reference
//   - uID *userv1beta1.UserId
func (_e *Searcher_Expecter) RestoreItem(ref interface{}, uID interface{}) *Searcher_RestoreItem_Call {
	return &Searcher_RestoreItem_Call{Call: _e.mock.On("RestoreItem", ref, uID)}
}

func (_c *Searcher_RestoreItem_Call) Run(run func(ref *providerv1beta1.Reference, uID *userv1beta1.UserId)) *Searcher_RestoreItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*providerv1beta1.Reference), args[1].(*userv1beta1.UserId))
	})
	return _c
}

func (_c *Searcher_RestoreItem_Call) Return() *Searcher_RestoreItem_Call {
	_c.Call.Return()
	return _c
}

func (_c *Searcher_RestoreItem_Call) RunAndReturn(run func(*providerv1beta1.Reference, *userv1beta1.UserId)) *Searcher_RestoreItem_Call {
	_c.Call.Return(run)
	return _c
}

// Search provides a mock function with given fields: ctx, req
func (_m *Searcher) Search(ctx context.Context, req *v0.SearchRequest) (*v0.SearchResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 *v0.SearchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v0.SearchRequest) (*v0.SearchResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v0.SearchRequest) *v0.SearchResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.SearchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v0.SearchRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Searcher_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type Searcher_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - ctx context.Context
//   - req *v0.SearchRequest
func (_e *Searcher_Expecter) Search(ctx interface{}, req interface{}) *Searcher_Search_Call {
	return &Searcher_Search_Call{Call: _e.mock.On("Search", ctx, req)}
}

func (_c *Searcher_Search_Call) Run(run func(ctx context.Context, req *v0.SearchRequest)) *Searcher_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v0.SearchRequest))
	})
	return _c
}

func (_c *Searcher_Search_Call) Return(_a0 *v0.SearchResponse, _a1 error) *Searcher_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Searcher_Search_Call) RunAndReturn(run func(context.Context, *v0.SearchRequest) (*v0.SearchResponse, error)) *Searcher_Search_Call {
	_c.Call.Return(run)
	return _c
}

// TrashItem provides a mock function with given fields: rID
func (_m *Searcher) TrashItem(rID *providerv1beta1.ResourceId) {
	_m.Called(rID)
}

// Searcher_TrashItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TrashItem'
type Searcher_TrashItem_Call struct {
	*mock.Call
}

// TrashItem is a helper method to define mock.On call
//   - rID *providerv1beta1.ResourceId
func (_e *Searcher_Expecter) TrashItem(rID interface{}) *Searcher_TrashItem_Call {
	return &Searcher_TrashItem_Call{Call: _e.mock.On("TrashItem", rID)}
}

func (_c *Searcher_TrashItem_Call) Run(run func(rID *providerv1beta1.ResourceId)) *Searcher_TrashItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*providerv1beta1.ResourceId))
	})
	return _c
}

func (_c *Searcher_TrashItem_Call) Return() *Searcher_TrashItem_Call {
	_c.Call.Return()
	return _c
}

func (_c *Searcher_TrashItem_Call) RunAndReturn(run func(*providerv1beta1.ResourceId)) *Searcher_TrashItem_Call {
	_c.Call.Return(run)
	return _c
}

// UpsertItem provides a mock function with given fields: ref, uID
func (_m *Searcher) UpsertItem(ref *providerv1beta1.Reference, uID *userv1beta1.UserId) {
	_m.Called(ref, uID)
}

// Searcher_UpsertItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertItem'
type Searcher_UpsertItem_Call struct {
	*mock.Call
}

// UpsertItem is a helper method to define mock.On call
//   - ref *providerv1beta1.Reference
//   - uID *userv1beta1.UserId
func (_e *Searcher_Expecter) UpsertItem(ref interface{}, uID interface{}) *Searcher_UpsertItem_Call {
	return &Searcher_UpsertItem_Call{Call: _e.mock.On("UpsertItem", ref, uID)}
}

func (_c *Searcher_UpsertItem_Call) Run(run func(ref *providerv1beta1.Reference, uID *userv1beta1.UserId)) *Searcher_UpsertItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*providerv1beta1.Reference), args[1].(*userv1beta1.UserId))
	})
	return _c
}

func (_c *Searcher_UpsertItem_Call) Return() *Searcher_UpsertItem_Call {
	_c.Call.Return()
	return _c
}

func (_c *Searcher_UpsertItem_Call) RunAndReturn(run func(*providerv1beta1.Reference, *userv1beta1.UserId)) *Searcher_UpsertItem_Call {
	_c.Call.Return(run)
	return _c
}

// NewSearcher creates a new instance of Searcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSearcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *Searcher {
	mock := &Searcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

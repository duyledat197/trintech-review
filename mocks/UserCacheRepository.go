// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "trintech/review/internal/user-management/entity"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// UserCacheRepository is an autogenerated mock type for the UserCacheRepository type
type UserCacheRepository struct {
	mock.Mock
}

// IncrementForgotPassword provides a mock function with given fields: ctx, email, duration
func (_m *UserCacheRepository) IncrementForgotPassword(ctx context.Context, email string, duration time.Duration) (int64, error) {
	ret := _m.Called(ctx, email, duration)

	if len(ret) == 0 {
		panic("no return value specified for IncrementForgotPassword")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) (int64, error)); ok {
		return rf(ctx, email, duration)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) int64); ok {
		r0 = rf(ctx, email, duration)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, time.Duration) error); ok {
		r1 = rf(ctx, email, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveByEmail provides a mock function with given fields: _a0, _a1
func (_m *UserCacheRepository) RemoveByEmail(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RemoveByEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveByResetToken provides a mock function with given fields: _a0, _a1, _a2
func (_m *UserCacheRepository) RemoveByResetToken(_a0 context.Context, _a1 string, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for RemoveByResetToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveByUserName provides a mock function with given fields: _a0, _a1
func (_m *UserCacheRepository) RemoveByUserName(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RemoveByUserName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveByEmail provides a mock function with given fields: _a0, _a1
func (_m *UserCacheRepository) RetrieveByEmail(_a0 context.Context, _a1 string) (*entity.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByEmail")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveByUserName provides a mock function with given fields: _a0, _a1
func (_m *UserCacheRepository) RetrieveByUserName(_a0 context.Context, _a1 string) (*entity.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByUserName")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveResetToken provides a mock function with given fields: ctx, email, resetToken
func (_m *UserCacheRepository) RetrieveResetToken(ctx context.Context, email string, resetToken string) error {
	ret := _m.Called(ctx, email, resetToken)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveResetToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, email, resetToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreByEmail provides a mock function with given fields: _a0, _a1, _a2
func (_m *UserCacheRepository) StoreByEmail(_a0 context.Context, _a1 string, _a2 *entity.User) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for StoreByEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *entity.User) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreByUserName provides a mock function with given fields: _a0, _a1, _a2
func (_m *UserCacheRepository) StoreByUserName(_a0 context.Context, _a1 string, _a2 *entity.User) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for StoreByUserName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *entity.User) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreResetToken provides a mock function with given fields: ctx, email, resetToken, duration
func (_m *UserCacheRepository) StoreResetToken(ctx context.Context, email string, resetToken string, duration time.Duration) error {
	ret := _m.Called(ctx, email, resetToken, duration)

	if len(ret) == 0 {
		panic("no return value specified for StoreResetToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Duration) error); ok {
		r0 = rf(ctx, email, resetToken, duration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserCacheRepository creates a new instance of UserCacheRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserCacheRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserCacheRepository {
	mock := &UserCacheRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

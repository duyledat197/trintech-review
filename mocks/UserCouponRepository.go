// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "trintech/review/internal/coupon-management/entity"
	database "trintech/review/pkg/database"

	mock "github.com/stretchr/testify/mock"
)

// UserCouponRepository is an autogenerated mock type for the UserCouponRepository type
type UserCouponRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, db, data
func (_m *UserCouponRepository) Create(ctx context.Context, db database.Executor, data *entity.UserCoupon) error {
	ret := _m.Called(ctx, db, data)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Executor, *entity.UserCoupon) error); ok {
		r0 = rf(ctx, db, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByCouponID provides a mock function with given fields: ctx, db, id
func (_m *UserCouponRepository) DeleteByCouponID(ctx context.Context, db database.Executor, id int64) error {
	ret := _m.Called(ctx, db, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByCouponID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Executor, int64) error); ok {
		r0 = rf(ctx, db, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveByCouponIDUserID provides a mock function with given fields: ctx, db, couponID, userID
func (_m *UserCouponRepository) RetrieveByCouponIDUserID(ctx context.Context, db database.Executor, couponID int64, userID int64) (*entity.UserCoupon, error) {
	ret := _m.Called(ctx, db, couponID, userID)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByCouponIDUserID")
	}

	var r0 *entity.UserCoupon
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Executor, int64, int64) (*entity.UserCoupon, error)); ok {
		return rf(ctx, db, couponID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.Executor, int64, int64) *entity.UserCoupon); ok {
		r0 = rf(ctx, db, couponID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserCoupon)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.Executor, int64, int64) error); ok {
		r1 = rf(ctx, db, couponID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserCouponRepository creates a new instance of UserCouponRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserCouponRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserCouponRepository {
	mock := &UserCouponRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

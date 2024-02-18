// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"
	coupon "trintech/review/dto/coupon-management/coupon"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// CouponServiceClient is an autogenerated mock type for the CouponServiceClient type
type CouponServiceClient struct {
	mock.Mock
}

// CreateCoupon provides a mock function with given fields: ctx, in, opts
func (_m *CouponServiceClient) CreateCoupon(ctx context.Context, in *coupon.CreateCouponRequest, opts ...grpc.CallOption) (*coupon.CreateCouponResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateCoupon")
	}

	var r0 *coupon.CreateCouponResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coupon.CreateCouponRequest, ...grpc.CallOption) (*coupon.CreateCouponResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coupon.CreateCouponRequest, ...grpc.CallOption) *coupon.CreateCouponResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coupon.CreateCouponResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coupon.CreateCouponRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCouponByID provides a mock function with given fields: ctx, in, opts
func (_m *CouponServiceClient) DeleteCouponByID(ctx context.Context, in *coupon.DeleteCouponByIDRequest, opts ...grpc.CallOption) (*coupon.DeleteCouponByIDResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCouponByID")
	}

	var r0 *coupon.DeleteCouponByIDResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coupon.DeleteCouponByIDRequest, ...grpc.CallOption) (*coupon.DeleteCouponByIDResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coupon.DeleteCouponByIDRequest, ...grpc.CallOption) *coupon.DeleteCouponByIDResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coupon.DeleteCouponByIDResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coupon.DeleteCouponByIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCouponServiceClient creates a new instance of CouponServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCouponServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *CouponServiceClient {
	mock := &CouponServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// number is an autogenerated mock type for the number type
type number struct {
	mock.Mock
}

// newNumber creates a new instance of number. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newNumber(t interface {
	mock.TestingT
	Cleanup(func())
}) *number {
	mock := &number{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

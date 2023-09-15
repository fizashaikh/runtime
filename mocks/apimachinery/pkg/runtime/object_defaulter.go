// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// ObjectDefaulter is an autogenerated mock type for the ObjectDefaulter type
type ObjectDefaulter struct {
	mock.Mock
}

// Default provides a mock function with given fields: in
func (_m *ObjectDefaulter) Default(in runtime.Object) {
	_m.Called(in)
}

// NewObjectDefaulter creates a new instance of ObjectDefaulter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewObjectDefaulter(t interface {
	mock.TestingT
	Cleanup(func())
}) *ObjectDefaulter {
	mock := &ObjectDefaulter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

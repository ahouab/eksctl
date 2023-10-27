// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubeNodeGroup is an autogenerated mock type for the KubeNodeGroup type
type KubeNodeGroup struct {
	mock.Mock
}

// GetAMIFamily provides a mock function with given fields:
func (_m *KubeNodeGroup) GetAMIFamily() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ListOptions provides a mock function with given fields:
func (_m *KubeNodeGroup) ListOptions() v1.ListOptions {
	ret := _m.Called()

	var r0 v1.ListOptions
	if rf, ok := ret.Get(0).(func() v1.ListOptions); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.ListOptions)
	}

	return r0
}

// NameString provides a mock function with given fields:
func (_m *KubeNodeGroup) NameString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Size provides a mock function with given fields:
func (_m *KubeNodeGroup) Size() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// NewKubeNodeGroup creates a new instance of KubeNodeGroup. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKubeNodeGroup(t interface {
	mock.TestingT
	Cleanup(func())
}) *KubeNodeGroup {
	mock := &KubeNodeGroup{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

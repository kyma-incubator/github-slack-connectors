// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import v1beta1 "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"

// Binding is an autogenerated mock type for the Binding type
type Binding struct {
	mock.Mock
}

// Create provides a mock function with given fields: body
func (_m *Binding) Create(body *v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error) {
	ret := _m.Called(body)

	var r0 *v1beta1.ServiceBinding
	if rf, ok := ret.Get(0).(func(*v1beta1.ServiceBinding) *v1beta1.ServiceBinding); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ServiceBinding)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.ServiceBinding) error); ok {
		r1 = rf(body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name, options
func (_m *Binding) Delete(name string, options *v1.DeleteOptions) error {
	ret := _m.Called(name, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *v1.DeleteOptions) error); ok {
		r0 = rf(name, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Prepare provides a mock function with given fields: name, lambdaName
func (_m *Binding) Prepare(name string, lambdaName string) *v1beta1.ServiceBinding {
	ret := _m.Called(name, lambdaName)

	var r0 *v1beta1.ServiceBinding
	if rf, ok := ret.Get(0).(func(string, string) *v1beta1.ServiceBinding); ok {
		r0 = rf(name, lambdaName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ServiceBinding)
		}
	}

	return r0
}

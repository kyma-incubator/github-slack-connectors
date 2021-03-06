// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1beta1 "github.com/google/kf/pkg/client/servicecatalog/clientset/versioned/typed/servicecatalog/v1beta1"

// ServiceCatalogClient is an autogenerated mock type for the ServiceCatalogClient type
type ServiceCatalogClient struct {
	mock.Mock
}

// ServiceBindings provides a mock function with given fields: _a0
func (_m *ServiceCatalogClient) ServiceBindings(_a0 string) v1beta1.ServiceBindingInterface {
	ret := _m.Called(_a0)

	var r0 v1beta1.ServiceBindingInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.ServiceBindingInterface); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.ServiceBindingInterface)
		}
	}

	return r0
}

// ServiceInstances provides a mock function with given fields: _a0
func (_m *ServiceCatalogClient) ServiceInstances(_a0 string) v1beta1.ServiceInstanceInterface {
	ret := _m.Called(_a0)

	var r0 v1beta1.ServiceInstanceInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.ServiceInstanceInterface); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.ServiceInstanceInterface)
		}
	}

	return r0
}

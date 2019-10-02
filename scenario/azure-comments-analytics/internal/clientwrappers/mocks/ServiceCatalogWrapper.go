// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import k8scomponents "github.com/kyma-incubator/github-slack-connectors/scenario/azure-comments-analytics/internal/k8scomponents"
import mock "github.com/stretchr/testify/mock"

// ServiceCatalogWrapper is an autogenerated mock type for the ServiceCatalogWrapper type
type ServiceCatalogWrapper struct {
	mock.Mock
}

// Binding provides a mock function with given fields: namespace
func (_m *ServiceCatalogWrapper) Binding(namespace string) k8scomponents.Binding {
	ret := _m.Called(namespace)

	var r0 k8scomponents.Binding
	if rf, ok := ret.Get(0).(func(string) k8scomponents.Binding); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(k8scomponents.Binding)
		}
	}

	return r0
}

// Instance provides a mock function with given fields: namespace
func (_m *ServiceCatalogWrapper) Instance(namespace string) k8scomponents.ServiceInstance {
	ret := _m.Called(namespace)

	var r0 k8scomponents.ServiceInstance
	if rf, ok := ret.Get(0).(func(string) k8scomponents.ServiceInstance); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(k8scomponents.ServiceInstance)
		}
	}

	return r0
}

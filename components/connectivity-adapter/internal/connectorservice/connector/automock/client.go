// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import apperrors "github.com/kyma-incubator/compass/components/connectivity-adapter/pkg/apperrors"

import externalschema "github.com/kyma-incubator/compass/components/connector/pkg/graphql/externalschema"
import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Configuration provides a mock function with given fields: headers
func (_m *Client) Configuration(headers map[string]string) (externalschema.Configuration, apperrors.AppError) {
	ret := _m.Called(headers)

	var r0 externalschema.Configuration
	if rf, ok := ret.Get(0).(func(map[string]string) externalschema.Configuration); ok {
		r0 = rf(headers)
	} else {
		r0 = ret.Get(0).(externalschema.Configuration)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(map[string]string) apperrors.AppError); ok {
		r1 = rf(headers)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Revoke provides a mock function with given fields: headers
func (_m *Client) Revoke(headers map[string]string) apperrors.AppError {
	ret := _m.Called(headers)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(map[string]string) apperrors.AppError); ok {
		r0 = rf(headers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// SignCSR provides a mock function with given fields: csr, headers
func (_m *Client) SignCSR(csr string, headers map[string]string) (externalschema.CertificationResult, apperrors.AppError) {
	ret := _m.Called(csr, headers)

	var r0 externalschema.CertificationResult
	if rf, ok := ret.Get(0).(func(string, map[string]string) externalschema.CertificationResult); ok {
		r0 = rf(csr, headers)
	} else {
		r0 = ret.Get(0).(externalschema.CertificationResult)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, map[string]string) apperrors.AppError); ok {
		r1 = rf(csr, headers)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

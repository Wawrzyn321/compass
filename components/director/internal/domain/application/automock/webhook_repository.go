// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal/model"

// WebhookRepository is an autogenerated mock type for the WebhookRepository type
type WebhookRepository struct {
	mock.Mock
}

// CreateMany provides a mock function with given fields: items
func (_m *WebhookRepository) CreateMany(items []*model.Webhook) error {
	ret := _m.Called(items)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*model.Webhook) error); ok {
		r0 = rf(items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllByApplicationID provides a mock function with given fields: id
func (_m *WebhookRepository) DeleteAllByApplicationID(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListByApplicationID provides a mock function with given fields: applicationID
func (_m *WebhookRepository) ListByApplicationID(applicationID string) ([]*model.Webhook, error) {
	ret := _m.Called(applicationID)

	var r0 []*model.Webhook
	if rf, ok := ret.Get(0).(func(string) []*model.Webhook); ok {
		r0 = rf(applicationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Webhook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(applicationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/LambdaTest/test-at-scale/pkg/core"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// TestExecutionService is an autogenerated mock type for the TestExecutionService type
type TestExecutionService struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx, testExecutionArgs
func (_m *TestExecutionService) Run(ctx context.Context, testExecutionArgs core.TestExecutionArgs) (*core.ExecutionResults, error) {
	ret := _m.Called(ctx, testExecutionArgs)

	var r0 *core.ExecutionResults
	if rf, ok := ret.Get(0).(func(context.Context, core.TestExecutionArgs) *core.ExecutionResults); ok {
		r0 = rf(ctx, testExecutionArgs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ExecutionResults)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, core.TestExecutionArgs) error); ok {
		r1 = rf(ctx, testExecutionArgs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendResults provides a mock function with given fields: ctx, payload
func (_m *TestExecutionService) SendResults(ctx context.Context, payload *core.ExecutionResults) (*core.TestReportResponsePayload, error) {
	ret := _m.Called(ctx, payload)

	var r0 *core.TestReportResponsePayload
	if rf, ok := ret.Get(0).(func(context.Context, *core.ExecutionResults) *core.TestReportResponsePayload); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TestReportResponsePayload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.ExecutionResults) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTestExecutionService creates a new instance of TestExecutionService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewTestExecutionService(t testing.TB) *TestExecutionService {
	mock := &TestExecutionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

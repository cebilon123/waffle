// Code generated by mockery v2.42.1. DO NOT EDIT.

package rule

import mock "github.com/stretchr/testify/mock"

// MockExpressionTreeBuilder is an autogenerated mock type for the ExpressionTreeBuilder type
type MockExpressionTreeBuilder struct {
	mock.Mock
}

type MockExpressionTreeBuilder_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExpressionTreeBuilder) EXPECT() *MockExpressionTreeBuilder_Expecter {
	return &MockExpressionTreeBuilder_Expecter{mock: &_m.Mock}
}

// BuildExpressionTree provides a mock function with given fields: variable, expression
func (_m *MockExpressionTreeBuilder) BuildExpressionTree(variable string, expression string) (expressionTree, error) {
	ret := _m.Called(variable, expression)

	if len(ret) == 0 {
		panic("no return value specified for BuildExpressionTree")
	}

	var r0 expressionTree
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (expressionTree, error)); ok {
		return rf(variable, expression)
	}
	if rf, ok := ret.Get(0).(func(string, string) expressionTree); ok {
		r0 = rf(variable, expression)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(expressionTree)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(variable, expression)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExpressionTreeBuilder_BuildExpressionTree_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BuildExpressionTree'
type MockExpressionTreeBuilder_BuildExpressionTree_Call struct {
	*mock.Call
}

// BuildExpressionTree is a helper method to define mock.On call
//   - variable string
//   - expression string
func (_e *MockExpressionTreeBuilder_Expecter) BuildExpressionTree(variable interface{}, expression interface{}) *MockExpressionTreeBuilder_BuildExpressionTree_Call {
	return &MockExpressionTreeBuilder_BuildExpressionTree_Call{Call: _e.mock.On("BuildExpressionTree", variable, expression)}
}

func (_c *MockExpressionTreeBuilder_BuildExpressionTree_Call) Run(run func(variable string, expression string)) *MockExpressionTreeBuilder_BuildExpressionTree_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockExpressionTreeBuilder_BuildExpressionTree_Call) Return(_a0 expressionTree, _a1 error) *MockExpressionTreeBuilder_BuildExpressionTree_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExpressionTreeBuilder_BuildExpressionTree_Call) RunAndReturn(run func(string, string) (expressionTree, error)) *MockExpressionTreeBuilder_BuildExpressionTree_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExpressionTreeBuilder creates a new instance of MockExpressionTreeBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExpressionTreeBuilder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExpressionTreeBuilder {
	mock := &MockExpressionTreeBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

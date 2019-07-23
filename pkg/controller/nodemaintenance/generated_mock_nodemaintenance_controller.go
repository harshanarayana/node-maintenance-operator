// Code generated by MockGen. DO NOT EDIT.
// Source: nodemaintenance_controller.go

// Package nodemaintenance is a generated GoMock package.
package nodemaintenance

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockReconcileHandler is a mock of ReconcileHandler interface
type MockReconcileHandler struct {
	ctrl     *gomock.Controller
	recorder *MockReconcileHandlerMockRecorder
}

// MockReconcileHandlerMockRecorder is the mock recorder for MockReconcileHandler
type MockReconcileHandlerMockRecorder struct {
	mock *MockReconcileHandler
}

// NewMockReconcileHandler creates a new mock instance
func NewMockReconcileHandler(ctrl *gomock.Controller) *MockReconcileHandler {
	mock := &MockReconcileHandler{ctrl: ctrl}
	mock.recorder = &MockReconcileHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReconcileHandler) EXPECT() *MockReconcileHandlerMockRecorder {
	return m.recorder
}

// StartPodInformer mocks base method
func (m *MockReconcileHandler) StartPodInformer(node *v1.Node, stop <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartPodInformer", node, stop)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartPodInformer indicates an expected call of StartPodInformer
func (mr *MockReconcileHandlerMockRecorder) StartPodInformer(node, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPodInformer", reflect.TypeOf((*MockReconcileHandler)(nil).StartPodInformer), node, stop)
}

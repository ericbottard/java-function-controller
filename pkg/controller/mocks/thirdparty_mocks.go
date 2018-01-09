// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/projectriff/kubernetes-crds/pkg/client/informers/externalversions/projectriff/v1 (interfaces: TopicInformer,FunctionInformer)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/projectriff/function-controller/vendor/github.com/projectriff/kubernetes-crds/pkg/client/listers/projectriff/v1"
	cache "github.com/projectriff/function-controller/vendor/k8s.io/client-go/tools/cache"
	reflect "reflect"
)

// MockTopicInformer is a mock of TopicInformer interface
type MockTopicInformer struct {
	ctrl     *gomock.Controller
	recorder *MockTopicInformerMockRecorder
}

// MockTopicInformerMockRecorder is the mock recorder for MockTopicInformer
type MockTopicInformerMockRecorder struct {
	mock *MockTopicInformer
}

// NewMockTopicInformer creates a new mock instance
func NewMockTopicInformer(ctrl *gomock.Controller) *MockTopicInformer {
	mock := &MockTopicInformer{ctrl: ctrl}
	mock.recorder = &MockTopicInformerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTopicInformer) EXPECT() *MockTopicInformerMockRecorder {
	return m.recorder
}

// Informer mocks base method
func (m *MockTopicInformer) Informer() cache.SharedIndexInformer {
	ret := m.ctrl.Call(m, "Informer")
	ret0, _ := ret[0].(cache.SharedIndexInformer)
	return ret0
}

// Informer indicates an expected call of Informer
func (mr *MockTopicInformerMockRecorder) Informer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Informer", reflect.TypeOf((*MockTopicInformer)(nil).Informer))
}

// Lister mocks base method
func (m *MockTopicInformer) Lister() v1.TopicLister {
	ret := m.ctrl.Call(m, "Lister")
	ret0, _ := ret[0].(v1.TopicLister)
	return ret0
}

// Lister indicates an expected call of Lister
func (mr *MockTopicInformerMockRecorder) Lister() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lister", reflect.TypeOf((*MockTopicInformer)(nil).Lister))
}

// MockFunctionInformer is a mock of FunctionInformer interface
type MockFunctionInformer struct {
	ctrl     *gomock.Controller
	recorder *MockFunctionInformerMockRecorder
}

// MockFunctionInformerMockRecorder is the mock recorder for MockFunctionInformer
type MockFunctionInformerMockRecorder struct {
	mock *MockFunctionInformer
}

// NewMockFunctionInformer creates a new mock instance
func NewMockFunctionInformer(ctrl *gomock.Controller) *MockFunctionInformer {
	mock := &MockFunctionInformer{ctrl: ctrl}
	mock.recorder = &MockFunctionInformerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFunctionInformer) EXPECT() *MockFunctionInformerMockRecorder {
	return m.recorder
}

// Informer mocks base method
func (m *MockFunctionInformer) Informer() cache.SharedIndexInformer {
	ret := m.ctrl.Call(m, "Informer")
	ret0, _ := ret[0].(cache.SharedIndexInformer)
	return ret0
}

// Informer indicates an expected call of Informer
func (mr *MockFunctionInformerMockRecorder) Informer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Informer", reflect.TypeOf((*MockFunctionInformer)(nil).Informer))
}

// Lister mocks base method
func (m *MockFunctionInformer) Lister() v1.FunctionLister {
	ret := m.ctrl.Call(m, "Lister")
	ret0, _ := ret[0].(v1.FunctionLister)
	return ret0
}

// Lister indicates an expected call of Lister
func (mr *MockFunctionInformerMockRecorder) Lister() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lister", reflect.TypeOf((*MockFunctionInformer)(nil).Lister))
}

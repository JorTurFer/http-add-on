// /*
// Copyright 2023 The KEDA Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: operator/generated/informers/externalversions/http/v1alpha1/httpscalingset.go
//
// Generated by this command:
//
//	mockgen -copyright_file=hack/boilerplate.go.txt -destination=operator/generated/informers/externalversions/http/v1alpha1/mock/httpscalingset.go -package=mock -source=operator/generated/informers/externalversions/http/v1alpha1/httpscalingset.go
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	v1alpha1 "github.com/kedacore/http-add-on/operator/generated/listers/http/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	cache "k8s.io/client-go/tools/cache"
)

// MockHTTPScalingSetInformer is a mock of HTTPScalingSetInformer interface.
type MockHTTPScalingSetInformer struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPScalingSetInformerMockRecorder
}

// MockHTTPScalingSetInformerMockRecorder is the mock recorder for MockHTTPScalingSetInformer.
type MockHTTPScalingSetInformerMockRecorder struct {
	mock *MockHTTPScalingSetInformer
}

// NewMockHTTPScalingSetInformer creates a new mock instance.
func NewMockHTTPScalingSetInformer(ctrl *gomock.Controller) *MockHTTPScalingSetInformer {
	mock := &MockHTTPScalingSetInformer{ctrl: ctrl}
	mock.recorder = &MockHTTPScalingSetInformerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPScalingSetInformer) EXPECT() *MockHTTPScalingSetInformerMockRecorder {
	return m.recorder
}

// Informer mocks base method.
func (m *MockHTTPScalingSetInformer) Informer() cache.SharedIndexInformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Informer")
	ret0, _ := ret[0].(cache.SharedIndexInformer)
	return ret0
}

// Informer indicates an expected call of Informer.
func (mr *MockHTTPScalingSetInformerMockRecorder) Informer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Informer", reflect.TypeOf((*MockHTTPScalingSetInformer)(nil).Informer))
}

// Lister mocks base method.
func (m *MockHTTPScalingSetInformer) Lister() v1alpha1.HTTPScalingSetLister {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lister")
	ret0, _ := ret[0].(v1alpha1.HTTPScalingSetLister)
	return ret0
}

// Lister indicates an expected call of Lister.
func (mr *MockHTTPScalingSetInformerMockRecorder) Lister() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lister", reflect.TypeOf((*MockHTTPScalingSetInformer)(nil).Lister))
}
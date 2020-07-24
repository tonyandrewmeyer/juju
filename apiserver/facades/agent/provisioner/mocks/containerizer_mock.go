// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/network/containerizer (interfaces: LinkLayerDevice)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	network "github.com/juju/juju/core/network"
	containerizer "github.com/juju/juju/network/containerizer"
	state "github.com/juju/juju/state"
)

// MockLinkLayerDevice is a mock of LinkLayerDevice interface
type MockLinkLayerDevice struct {
	ctrl     *gomock.Controller
	recorder *MockLinkLayerDeviceMockRecorder
}

// MockLinkLayerDeviceMockRecorder is the mock recorder for MockLinkLayerDevice
type MockLinkLayerDeviceMockRecorder struct {
	mock *MockLinkLayerDevice
}

// NewMockLinkLayerDevice creates a new mock instance
func NewMockLinkLayerDevice(ctrl *gomock.Controller) *MockLinkLayerDevice {
	mock := &MockLinkLayerDevice{ctrl: ctrl}
	mock.recorder = &MockLinkLayerDeviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLinkLayerDevice) EXPECT() *MockLinkLayerDeviceMockRecorder {
	return m.recorder
}

// Addresses mocks base method
func (m *MockLinkLayerDevice) Addresses() ([]*state.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addresses")
	ret0, _ := ret[0].([]*state.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Addresses indicates an expected call of Addresses
func (mr *MockLinkLayerDeviceMockRecorder) Addresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addresses", reflect.TypeOf((*MockLinkLayerDevice)(nil).Addresses))
}

// EthernetDeviceForBridge mocks base method
func (m *MockLinkLayerDevice) EthernetDeviceForBridge(arg0 string, arg1 bool) (network.InterfaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EthernetDeviceForBridge", arg0, arg1)
	ret0, _ := ret[0].(network.InterfaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EthernetDeviceForBridge indicates an expected call of EthernetDeviceForBridge
func (mr *MockLinkLayerDeviceMockRecorder) EthernetDeviceForBridge(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EthernetDeviceForBridge", reflect.TypeOf((*MockLinkLayerDevice)(nil).EthernetDeviceForBridge), arg0, arg1)
}

// IsAutoStart mocks base method
func (m *MockLinkLayerDevice) IsAutoStart() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAutoStart")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAutoStart indicates an expected call of IsAutoStart
func (mr *MockLinkLayerDeviceMockRecorder) IsAutoStart() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAutoStart", reflect.TypeOf((*MockLinkLayerDevice)(nil).IsAutoStart))
}

// IsUp mocks base method
func (m *MockLinkLayerDevice) IsUp() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUp")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUp indicates an expected call of IsUp
func (mr *MockLinkLayerDeviceMockRecorder) IsUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUp", reflect.TypeOf((*MockLinkLayerDevice)(nil).IsUp))
}

// MACAddress mocks base method
func (m *MockLinkLayerDevice) MACAddress() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MACAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// MACAddress indicates an expected call of MACAddress
func (mr *MockLinkLayerDeviceMockRecorder) MACAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MACAddress", reflect.TypeOf((*MockLinkLayerDevice)(nil).MACAddress))
}

// MTU mocks base method
func (m *MockLinkLayerDevice) MTU() uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MTU")
	ret0, _ := ret[0].(uint)
	return ret0
}

// MTU indicates an expected call of MTU
func (mr *MockLinkLayerDeviceMockRecorder) MTU() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MTU", reflect.TypeOf((*MockLinkLayerDevice)(nil).MTU))
}

// Name mocks base method
func (m *MockLinkLayerDevice) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockLinkLayerDeviceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockLinkLayerDevice)(nil).Name))
}

// ParentDevice mocks base method
func (m *MockLinkLayerDevice) ParentDevice() (containerizer.LinkLayerDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParentDevice")
	ret0, _ := ret[0].(containerizer.LinkLayerDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParentDevice indicates an expected call of ParentDevice
func (mr *MockLinkLayerDeviceMockRecorder) ParentDevice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParentDevice", reflect.TypeOf((*MockLinkLayerDevice)(nil).ParentDevice))
}

// ParentName mocks base method
func (m *MockLinkLayerDevice) ParentName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParentName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ParentName indicates an expected call of ParentName
func (mr *MockLinkLayerDeviceMockRecorder) ParentName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParentName", reflect.TypeOf((*MockLinkLayerDevice)(nil).ParentName))
}

// Type mocks base method
func (m *MockLinkLayerDevice) Type() network.LinkLayerDeviceType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(network.LinkLayerDeviceType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockLinkLayerDeviceMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockLinkLayerDevice)(nil).Type))
}

// VirtualPortType mocks base method
func (m *MockLinkLayerDevice) VirtualPortType() network.VirtualPortType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualPortType")
	ret0, _ := ret[0].(network.VirtualPortType)
	return ret0
}

// VirtualPortType indicates an expected call of VirtualPortType
func (mr *MockLinkLayerDeviceMockRecorder) VirtualPortType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualPortType", reflect.TypeOf((*MockLinkLayerDevice)(nil).VirtualPortType))
}

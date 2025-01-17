/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by MockGen. DO NOT EDIT.
// Source: ./vpc.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	core "github.com/IBM/go-sdk-core/v5/core"
	vpcv1 "github.com/IBM/vpc-go-sdk/vpcv1"
	gomock "github.com/golang/mock/gomock"
)

// MockVpc is a mock of Vpc interface.
type MockVpc struct {
	ctrl     *gomock.Controller
	recorder *MockVpcMockRecorder
}

// MockVpcMockRecorder is the mock recorder for MockVpc.
type MockVpcMockRecorder struct {
	mock *MockVpc
}

// NewMockVpc creates a new mock instance.
func NewMockVpc(ctrl *gomock.Controller) *MockVpc {
	mock := &MockVpc{ctrl: ctrl}
	mock.recorder = &MockVpcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVpc) EXPECT() *MockVpcMockRecorder {
	return m.recorder
}

// AddInstanceNetworkInterfaceFloatingIP mocks base method.
func (m *MockVpc) AddInstanceNetworkInterfaceFloatingIP(options *vpcv1.AddInstanceNetworkInterfaceFloatingIPOptions) (*vpcv1.FloatingIP, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInstanceNetworkInterfaceFloatingIP", options)
	ret0, _ := ret[0].(*vpcv1.FloatingIP)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddInstanceNetworkInterfaceFloatingIP indicates an expected call of AddInstanceNetworkInterfaceFloatingIP.
func (mr *MockVpcMockRecorder) AddInstanceNetworkInterfaceFloatingIP(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInstanceNetworkInterfaceFloatingIP", reflect.TypeOf((*MockVpc)(nil).AddInstanceNetworkInterfaceFloatingIP), options)
}

// CreateFloatingIP mocks base method.
func (m *MockVpc) CreateFloatingIP(options *vpcv1.CreateFloatingIPOptions) (*vpcv1.FloatingIP, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFloatingIP", options)
	ret0, _ := ret[0].(*vpcv1.FloatingIP)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateFloatingIP indicates an expected call of CreateFloatingIP.
func (mr *MockVpcMockRecorder) CreateFloatingIP(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFloatingIP", reflect.TypeOf((*MockVpc)(nil).CreateFloatingIP), options)
}

// CreateInstance mocks base method.
func (m *MockVpc) CreateInstance(options *vpcv1.CreateInstanceOptions) (*vpcv1.Instance, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstance", options)
	ret0, _ := ret[0].(*vpcv1.Instance)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateInstance indicates an expected call of CreateInstance.
func (mr *MockVpcMockRecorder) CreateInstance(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstance", reflect.TypeOf((*MockVpc)(nil).CreateInstance), options)
}

// CreatePublicGateway mocks base method.
func (m *MockVpc) CreatePublicGateway(options *vpcv1.CreatePublicGatewayOptions) (*vpcv1.PublicGateway, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePublicGateway", options)
	ret0, _ := ret[0].(*vpcv1.PublicGateway)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreatePublicGateway indicates an expected call of CreatePublicGateway.
func (mr *MockVpcMockRecorder) CreatePublicGateway(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePublicGateway", reflect.TypeOf((*MockVpc)(nil).CreatePublicGateway), options)
}

// CreateSecurityGroupRule mocks base method.
func (m *MockVpc) CreateSecurityGroupRule(options *vpcv1.CreateSecurityGroupRuleOptions) (vpcv1.SecurityGroupRuleIntf, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecurityGroupRule", options)
	ret0, _ := ret[0].(vpcv1.SecurityGroupRuleIntf)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSecurityGroupRule indicates an expected call of CreateSecurityGroupRule.
func (mr *MockVpcMockRecorder) CreateSecurityGroupRule(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecurityGroupRule", reflect.TypeOf((*MockVpc)(nil).CreateSecurityGroupRule), options)
}

// CreateSubnet mocks base method.
func (m *MockVpc) CreateSubnet(options *vpcv1.CreateSubnetOptions) (*vpcv1.Subnet, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubnet", options)
	ret0, _ := ret[0].(*vpcv1.Subnet)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSubnet indicates an expected call of CreateSubnet.
func (mr *MockVpcMockRecorder) CreateSubnet(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubnet", reflect.TypeOf((*MockVpc)(nil).CreateSubnet), options)
}

// CreateVPC mocks base method.
func (m *MockVpc) CreateVPC(options *vpcv1.CreateVPCOptions) (*vpcv1.VPC, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVPC", options)
	ret0, _ := ret[0].(*vpcv1.VPC)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateVPC indicates an expected call of CreateVPC.
func (mr *MockVpcMockRecorder) CreateVPC(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVPC", reflect.TypeOf((*MockVpc)(nil).CreateVPC), options)
}

// DeleteFloatingIP mocks base method.
func (m *MockVpc) DeleteFloatingIP(options *vpcv1.DeleteFloatingIPOptions) (*core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFloatingIP", options)
	ret0, _ := ret[0].(*core.DetailedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFloatingIP indicates an expected call of DeleteFloatingIP.
func (mr *MockVpcMockRecorder) DeleteFloatingIP(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFloatingIP", reflect.TypeOf((*MockVpc)(nil).DeleteFloatingIP), options)
}

// DeleteInstance mocks base method.
func (m *MockVpc) DeleteInstance(options *vpcv1.DeleteInstanceOptions) (*core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstance", options)
	ret0, _ := ret[0].(*core.DetailedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteInstance indicates an expected call of DeleteInstance.
func (mr *MockVpcMockRecorder) DeleteInstance(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstance", reflect.TypeOf((*MockVpc)(nil).DeleteInstance), options)
}

// DeletePublicGateway mocks base method.
func (m *MockVpc) DeletePublicGateway(options *vpcv1.DeletePublicGatewayOptions) (*core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePublicGateway", options)
	ret0, _ := ret[0].(*core.DetailedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePublicGateway indicates an expected call of DeletePublicGateway.
func (mr *MockVpcMockRecorder) DeletePublicGateway(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePublicGateway", reflect.TypeOf((*MockVpc)(nil).DeletePublicGateway), options)
}

// DeleteSubnet mocks base method.
func (m *MockVpc) DeleteSubnet(options *vpcv1.DeleteSubnetOptions) (*core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubnet", options)
	ret0, _ := ret[0].(*core.DetailedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSubnet indicates an expected call of DeleteSubnet.
func (mr *MockVpcMockRecorder) DeleteSubnet(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubnet", reflect.TypeOf((*MockVpc)(nil).DeleteSubnet), options)
}

// DeleteVPC mocks base method.
func (m *MockVpc) DeleteVPC(options *vpcv1.DeleteVPCOptions) (*core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVPC", options)
	ret0, _ := ret[0].(*core.DetailedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVPC indicates an expected call of DeleteVPC.
func (mr *MockVpcMockRecorder) DeleteVPC(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVPC", reflect.TypeOf((*MockVpc)(nil).DeleteVPC), options)
}

// GetInstance mocks base method.
func (m *MockVpc) GetInstance(options *vpcv1.GetInstanceOptions) (*vpcv1.Instance, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstance", options)
	ret0, _ := ret[0].(*vpcv1.Instance)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetInstance indicates an expected call of GetInstance.
func (mr *MockVpcMockRecorder) GetInstance(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*MockVpc)(nil).GetInstance), options)
}

// GetSubnetPublicGateway mocks base method.
func (m *MockVpc) GetSubnetPublicGateway(options *vpcv1.GetSubnetPublicGatewayOptions) (*vpcv1.PublicGateway, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetPublicGateway", options)
	ret0, _ := ret[0].(*vpcv1.PublicGateway)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSubnetPublicGateway indicates an expected call of GetSubnetPublicGateway.
func (mr *MockVpcMockRecorder) GetSubnetPublicGateway(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetPublicGateway", reflect.TypeOf((*MockVpc)(nil).GetSubnetPublicGateway), options)
}

// ListFloatingIps mocks base method.
func (m *MockVpc) ListFloatingIps(options *vpcv1.ListFloatingIpsOptions) (*vpcv1.FloatingIPCollection, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFloatingIps", options)
	ret0, _ := ret[0].(*vpcv1.FloatingIPCollection)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFloatingIps indicates an expected call of ListFloatingIps.
func (mr *MockVpcMockRecorder) ListFloatingIps(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFloatingIps", reflect.TypeOf((*MockVpc)(nil).ListFloatingIps), options)
}

// ListInstances mocks base method.
func (m *MockVpc) ListInstances(options *vpcv1.ListInstancesOptions) (*vpcv1.InstanceCollection, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInstances", options)
	ret0, _ := ret[0].(*vpcv1.InstanceCollection)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListInstances indicates an expected call of ListInstances.
func (mr *MockVpcMockRecorder) ListInstances(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstances", reflect.TypeOf((*MockVpc)(nil).ListInstances), options)
}

// ListSubnets mocks base method.
func (m *MockVpc) ListSubnets(options *vpcv1.ListSubnetsOptions) (*vpcv1.SubnetCollection, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubnets", options)
	ret0, _ := ret[0].(*vpcv1.SubnetCollection)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSubnets indicates an expected call of ListSubnets.
func (mr *MockVpcMockRecorder) ListSubnets(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubnets", reflect.TypeOf((*MockVpc)(nil).ListSubnets), options)
}

// ListVPCAddressPrefixes mocks base method.
func (m *MockVpc) ListVPCAddressPrefixes(options *vpcv1.ListVPCAddressPrefixesOptions) (*vpcv1.AddressPrefixCollection, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVPCAddressPrefixes", options)
	ret0, _ := ret[0].(*vpcv1.AddressPrefixCollection)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListVPCAddressPrefixes indicates an expected call of ListVPCAddressPrefixes.
func (mr *MockVpcMockRecorder) ListVPCAddressPrefixes(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVPCAddressPrefixes", reflect.TypeOf((*MockVpc)(nil).ListVPCAddressPrefixes), options)
}

// ListVpcs mocks base method.
func (m *MockVpc) ListVpcs(options *vpcv1.ListVpcsOptions) (*vpcv1.VPCCollection, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVpcs", options)
	ret0, _ := ret[0].(*vpcv1.VPCCollection)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListVpcs indicates an expected call of ListVpcs.
func (mr *MockVpcMockRecorder) ListVpcs(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVpcs", reflect.TypeOf((*MockVpc)(nil).ListVpcs), options)
}

// SetSubnetPublicGateway mocks base method.
func (m *MockVpc) SetSubnetPublicGateway(options *vpcv1.SetSubnetPublicGatewayOptions) (*vpcv1.PublicGateway, *core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSubnetPublicGateway", options)
	ret0, _ := ret[0].(*vpcv1.PublicGateway)
	ret1, _ := ret[1].(*core.DetailedResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SetSubnetPublicGateway indicates an expected call of SetSubnetPublicGateway.
func (mr *MockVpcMockRecorder) SetSubnetPublicGateway(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubnetPublicGateway", reflect.TypeOf((*MockVpc)(nil).SetSubnetPublicGateway), options)
}

// UnsetSubnetPublicGateway mocks base method.
func (m *MockVpc) UnsetSubnetPublicGateway(options *vpcv1.UnsetSubnetPublicGatewayOptions) (*core.DetailedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsetSubnetPublicGateway", options)
	ret0, _ := ret[0].(*core.DetailedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnsetSubnetPublicGateway indicates an expected call of UnsetSubnetPublicGateway.
func (mr *MockVpcMockRecorder) UnsetSubnetPublicGateway(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsetSubnetPublicGateway", reflect.TypeOf((*MockVpc)(nil).UnsetSubnetPublicGateway), options)
}

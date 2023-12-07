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
// Source: ../agentpools.go
//
// Generated by this command:
//
//	mockgen -destination agentpools_mock.go -package mock_agentpools -source ../agentpools.go AgentPoolScope
//
// Package mock_agentpools is a generated GoMock package.
package mock_agentpools

import (
	reflect "reflect"

	azcore "github.com/Azure/azure-sdk-for-go/sdk/azcore"
	v1api20231001 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001"
	gomock "go.uber.org/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockAgentPoolScope is a mock of AgentPoolScope interface.
type MockAgentPoolScope struct {
	ctrl     *gomock.Controller
	recorder *MockAgentPoolScopeMockRecorder
}

// MockAgentPoolScopeMockRecorder is the mock recorder for MockAgentPoolScope.
type MockAgentPoolScopeMockRecorder struct {
	mock *MockAgentPoolScope
}

// NewMockAgentPoolScope creates a new mock instance.
func NewMockAgentPoolScope(ctrl *gomock.Controller) *MockAgentPoolScope {
	mock := &MockAgentPoolScope{ctrl: ctrl}
	mock.recorder = &MockAgentPoolScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentPoolScope) EXPECT() *MockAgentPoolScopeMockRecorder {
	return m.recorder
}

// AdditionalTags mocks base method.
func (m *MockAgentPoolScope) AdditionalTags() v1beta1.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1beta1.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockAgentPoolScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockAgentPoolScope)(nil).AdditionalTags))
}

// AgentPoolSpec mocks base method.
func (m *MockAgentPoolScope) AgentPoolSpec() azure.ASOResourceSpecGetter[*v1api20231001.ManagedClustersAgentPool] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentPoolSpec")
	ret0, _ := ret[0].(azure.ASOResourceSpecGetter[*v1api20231001.ManagedClustersAgentPool])
	return ret0
}

// AgentPoolSpec indicates an expected call of AgentPoolSpec.
func (mr *MockAgentPoolScopeMockRecorder) AgentPoolSpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentPoolSpec", reflect.TypeOf((*MockAgentPoolScope)(nil).AgentPoolSpec))
}

// AvailabilitySetEnabled mocks base method.
func (m *MockAgentPoolScope) AvailabilitySetEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilitySetEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AvailabilitySetEnabled indicates an expected call of AvailabilitySetEnabled.
func (mr *MockAgentPoolScopeMockRecorder) AvailabilitySetEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilitySetEnabled", reflect.TypeOf((*MockAgentPoolScope)(nil).AvailabilitySetEnabled))
}

// BaseURI mocks base method.
func (m *MockAgentPoolScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockAgentPoolScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockAgentPoolScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockAgentPoolScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockAgentPoolScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockAgentPoolScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockAgentPoolScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockAgentPoolScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockAgentPoolScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockAgentPoolScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockAgentPoolScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockAgentPoolScope)(nil).CloudEnvironment))
}

// CloudProviderConfigOverrides mocks base method.
func (m *MockAgentPoolScope) CloudProviderConfigOverrides() *v1beta1.CloudProviderConfigOverrides {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderConfigOverrides")
	ret0, _ := ret[0].(*v1beta1.CloudProviderConfigOverrides)
	return ret0
}

// CloudProviderConfigOverrides indicates an expected call of CloudProviderConfigOverrides.
func (mr *MockAgentPoolScopeMockRecorder) CloudProviderConfigOverrides() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderConfigOverrides", reflect.TypeOf((*MockAgentPoolScope)(nil).CloudProviderConfigOverrides))
}

// ClusterName mocks base method.
func (m *MockAgentPoolScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockAgentPoolScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockAgentPoolScope)(nil).ClusterName))
}

// DeleteLongRunningOperationState mocks base method.
func (m *MockAgentPoolScope) DeleteLongRunningOperationState(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteLongRunningOperationState", arg0, arg1, arg2)
}

// DeleteLongRunningOperationState indicates an expected call of DeleteLongRunningOperationState.
func (mr *MockAgentPoolScopeMockRecorder) DeleteLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLongRunningOperationState", reflect.TypeOf((*MockAgentPoolScope)(nil).DeleteLongRunningOperationState), arg0, arg1, arg2)
}

// ExtendedLocation mocks base method.
func (m *MockAgentPoolScope) ExtendedLocation() *v1beta1.ExtendedLocationSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocation")
	ret0, _ := ret[0].(*v1beta1.ExtendedLocationSpec)
	return ret0
}

// ExtendedLocation indicates an expected call of ExtendedLocation.
func (mr *MockAgentPoolScopeMockRecorder) ExtendedLocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocation", reflect.TypeOf((*MockAgentPoolScope)(nil).ExtendedLocation))
}

// ExtendedLocationName mocks base method.
func (m *MockAgentPoolScope) ExtendedLocationName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocationName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtendedLocationName indicates an expected call of ExtendedLocationName.
func (mr *MockAgentPoolScopeMockRecorder) ExtendedLocationName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocationName", reflect.TypeOf((*MockAgentPoolScope)(nil).ExtendedLocationName))
}

// ExtendedLocationType mocks base method.
func (m *MockAgentPoolScope) ExtendedLocationType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocationType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtendedLocationType indicates an expected call of ExtendedLocationType.
func (mr *MockAgentPoolScopeMockRecorder) ExtendedLocationType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocationType", reflect.TypeOf((*MockAgentPoolScope)(nil).ExtendedLocationType))
}

// FailureDomains mocks base method.
func (m *MockAgentPoolScope) FailureDomains() []*string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailureDomains")
	ret0, _ := ret[0].([]*string)
	return ret0
}

// FailureDomains indicates an expected call of FailureDomains.
func (mr *MockAgentPoolScopeMockRecorder) FailureDomains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureDomains", reflect.TypeOf((*MockAgentPoolScope)(nil).FailureDomains))
}

// GetClient mocks base method.
func (m *MockAgentPoolScope) GetClient() client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockAgentPoolScopeMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockAgentPoolScope)(nil).GetClient))
}

// GetLongRunningOperationState mocks base method.
func (m *MockAgentPoolScope) GetLongRunningOperationState(arg0, arg1, arg2 string) *v1beta1.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongRunningOperationState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Future)
	return ret0
}

// GetLongRunningOperationState indicates an expected call of GetLongRunningOperationState.
func (mr *MockAgentPoolScopeMockRecorder) GetLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongRunningOperationState", reflect.TypeOf((*MockAgentPoolScope)(nil).GetLongRunningOperationState), arg0, arg1, arg2)
}

// HashKey mocks base method.
func (m *MockAgentPoolScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockAgentPoolScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockAgentPoolScope)(nil).HashKey))
}

// Location mocks base method.
func (m *MockAgentPoolScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockAgentPoolScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockAgentPoolScope)(nil).Location))
}

// Name mocks base method.
func (m *MockAgentPoolScope) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAgentPoolScopeMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAgentPoolScope)(nil).Name))
}

// NodeResourceGroup mocks base method.
func (m *MockAgentPoolScope) NodeResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// NodeResourceGroup indicates an expected call of NodeResourceGroup.
func (mr *MockAgentPoolScopeMockRecorder) NodeResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeResourceGroup", reflect.TypeOf((*MockAgentPoolScope)(nil).NodeResourceGroup))
}

// RemoveCAPIMachinePoolAnnotation mocks base method.
func (m *MockAgentPoolScope) RemoveCAPIMachinePoolAnnotation(key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveCAPIMachinePoolAnnotation", key)
}

// RemoveCAPIMachinePoolAnnotation indicates an expected call of RemoveCAPIMachinePoolAnnotation.
func (mr *MockAgentPoolScopeMockRecorder) RemoveCAPIMachinePoolAnnotation(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCAPIMachinePoolAnnotation", reflect.TypeOf((*MockAgentPoolScope)(nil).RemoveCAPIMachinePoolAnnotation), key)
}

// ResourceGroup mocks base method.
func (m *MockAgentPoolScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockAgentPoolScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockAgentPoolScope)(nil).ResourceGroup))
}

// SetAgentPoolProviderIDList mocks base method.
func (m *MockAgentPoolScope) SetAgentPoolProviderIDList(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAgentPoolProviderIDList", arg0)
}

// SetAgentPoolProviderIDList indicates an expected call of SetAgentPoolProviderIDList.
func (mr *MockAgentPoolScopeMockRecorder) SetAgentPoolProviderIDList(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAgentPoolProviderIDList", reflect.TypeOf((*MockAgentPoolScope)(nil).SetAgentPoolProviderIDList), arg0)
}

// SetAgentPoolReady mocks base method.
func (m *MockAgentPoolScope) SetAgentPoolReady(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAgentPoolReady", arg0)
}

// SetAgentPoolReady indicates an expected call of SetAgentPoolReady.
func (mr *MockAgentPoolScopeMockRecorder) SetAgentPoolReady(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAgentPoolReady", reflect.TypeOf((*MockAgentPoolScope)(nil).SetAgentPoolReady), arg0)
}

// SetAgentPoolReplicas mocks base method.
func (m *MockAgentPoolScope) SetAgentPoolReplicas(arg0 int32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAgentPoolReplicas", arg0)
}

// SetAgentPoolReplicas indicates an expected call of SetAgentPoolReplicas.
func (mr *MockAgentPoolScopeMockRecorder) SetAgentPoolReplicas(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAgentPoolReplicas", reflect.TypeOf((*MockAgentPoolScope)(nil).SetAgentPoolReplicas), arg0)
}

// SetCAPIMachinePoolAnnotation mocks base method.
func (m *MockAgentPoolScope) SetCAPIMachinePoolAnnotation(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCAPIMachinePoolAnnotation", key, value)
}

// SetCAPIMachinePoolAnnotation indicates an expected call of SetCAPIMachinePoolAnnotation.
func (mr *MockAgentPoolScopeMockRecorder) SetCAPIMachinePoolAnnotation(key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCAPIMachinePoolAnnotation", reflect.TypeOf((*MockAgentPoolScope)(nil).SetCAPIMachinePoolAnnotation), key, value)
}

// SetCAPIMachinePoolReplicas mocks base method.
func (m *MockAgentPoolScope) SetCAPIMachinePoolReplicas(replicas *int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCAPIMachinePoolReplicas", replicas)
}

// SetCAPIMachinePoolReplicas indicates an expected call of SetCAPIMachinePoolReplicas.
func (mr *MockAgentPoolScopeMockRecorder) SetCAPIMachinePoolReplicas(replicas any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCAPIMachinePoolReplicas", reflect.TypeOf((*MockAgentPoolScope)(nil).SetCAPIMachinePoolReplicas), replicas)
}

// SetLongRunningOperationState mocks base method.
func (m *MockAgentPoolScope) SetLongRunningOperationState(arg0 *v1beta1.Future) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLongRunningOperationState", arg0)
}

// SetLongRunningOperationState indicates an expected call of SetLongRunningOperationState.
func (mr *MockAgentPoolScopeMockRecorder) SetLongRunningOperationState(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLongRunningOperationState", reflect.TypeOf((*MockAgentPoolScope)(nil).SetLongRunningOperationState), arg0)
}

// SetSubnetName mocks base method.
func (m *MockAgentPoolScope) SetSubnetName() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSubnetName")
}

// SetSubnetName indicates an expected call of SetSubnetName.
func (mr *MockAgentPoolScopeMockRecorder) SetSubnetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubnetName", reflect.TypeOf((*MockAgentPoolScope)(nil).SetSubnetName))
}

// SubscriptionID mocks base method.
func (m *MockAgentPoolScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockAgentPoolScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockAgentPoolScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockAgentPoolScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockAgentPoolScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockAgentPoolScope)(nil).TenantID))
}

// Token mocks base method.
func (m *MockAgentPoolScope) Token() azcore.TokenCredential {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(azcore.TokenCredential)
	return ret0
}

// Token indicates an expected call of Token.
func (mr *MockAgentPoolScopeMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockAgentPoolScope)(nil).Token))
}

// UpdateDeleteStatus mocks base method.
func (m *MockAgentPoolScope) UpdateDeleteStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDeleteStatus", arg0, arg1, arg2)
}

// UpdateDeleteStatus indicates an expected call of UpdateDeleteStatus.
func (mr *MockAgentPoolScopeMockRecorder) UpdateDeleteStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeleteStatus", reflect.TypeOf((*MockAgentPoolScope)(nil).UpdateDeleteStatus), arg0, arg1, arg2)
}

// UpdatePatchStatus mocks base method.
func (m *MockAgentPoolScope) UpdatePatchStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePatchStatus", arg0, arg1, arg2)
}

// UpdatePatchStatus indicates an expected call of UpdatePatchStatus.
func (mr *MockAgentPoolScopeMockRecorder) UpdatePatchStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePatchStatus", reflect.TypeOf((*MockAgentPoolScope)(nil).UpdatePatchStatus), arg0, arg1, arg2)
}

// UpdatePutStatus mocks base method.
func (m *MockAgentPoolScope) UpdatePutStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePutStatus", arg0, arg1, arg2)
}

// UpdatePutStatus indicates an expected call of UpdatePutStatus.
func (mr *MockAgentPoolScopeMockRecorder) UpdatePutStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePutStatus", reflect.TypeOf((*MockAgentPoolScope)(nil).UpdatePutStatus), arg0, arg1, arg2)
}

// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/flyteorg/flyteadmin/pkg/runtime/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// ClusterConfiguration is an autogenerated mock type for the ClusterConfiguration type
type ClusterConfiguration struct {
	mock.Mock
}

type ClusterConfiguration_GetClusterConfigs struct {
	*mock.Call
}

func (_m ClusterConfiguration_GetClusterConfigs) Return(_a0 []interfaces.ClusterConfig) *ClusterConfiguration_GetClusterConfigs {
	return &ClusterConfiguration_GetClusterConfigs{Call: _m.Call.Return(_a0)}
}

func (_m *ClusterConfiguration) OnGetClusterConfigs() *ClusterConfiguration_GetClusterConfigs {
	c := _m.On("GetClusterConfigs")
	return &ClusterConfiguration_GetClusterConfigs{Call: c}
}

func (_m *ClusterConfiguration) OnGetClusterConfigsMatch(matchers ...interface{}) *ClusterConfiguration_GetClusterConfigs {
	c := _m.On("GetClusterConfigs", matchers...)
	return &ClusterConfiguration_GetClusterConfigs{Call: c}
}

// GetClusterConfigs provides a mock function with given fields:
func (_m *ClusterConfiguration) GetClusterConfigs() []interfaces.ClusterConfig {
	ret := _m.Called()

	var r0 []interfaces.ClusterConfig
	if rf, ok := ret.Get(0).(func() []interfaces.ClusterConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interfaces.ClusterConfig)
		}
	}

	return r0
}

type ClusterConfiguration_GetLabelClusterMap struct {
	*mock.Call
}

func (_m ClusterConfiguration_GetLabelClusterMap) Return(_a0 map[string][]interfaces.ClusterEntity) *ClusterConfiguration_GetLabelClusterMap {
	return &ClusterConfiguration_GetLabelClusterMap{Call: _m.Call.Return(_a0)}
}

func (_m *ClusterConfiguration) OnGetLabelClusterMap() *ClusterConfiguration_GetLabelClusterMap {
	c := _m.On("GetLabelClusterMap")
	return &ClusterConfiguration_GetLabelClusterMap{Call: c}
}

func (_m *ClusterConfiguration) OnGetLabelClusterMapMatch(matchers ...interface{}) *ClusterConfiguration_GetLabelClusterMap {
	c := _m.On("GetLabelClusterMap", matchers...)
	return &ClusterConfiguration_GetLabelClusterMap{Call: c}
}

// GetLabelClusterMap provides a mock function with given fields:
func (_m *ClusterConfiguration) GetLabelClusterMap() map[string][]interfaces.ClusterEntity {
	ret := _m.Called()

	var r0 map[string][]interfaces.ClusterEntity
	if rf, ok := ret.Get(0).(func() map[string][]interfaces.ClusterEntity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]interfaces.ClusterEntity)
		}
	}

	return r0
}

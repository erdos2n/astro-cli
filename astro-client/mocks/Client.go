// Code generated by mockery v2.10.0. DO NOT EDIT.

package astro_mocks

import (
	astro "github.com/astronomer/astro-cli/astro-client"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateDeployment provides a mock function with given fields: input
func (_m *Client) CreateDeployment(input *astro.DeploymentCreateInput) (astro.Deployment, error) {
	ret := _m.Called(input)

	var r0 astro.Deployment
	if rf, ok := ret.Get(0).(func(*astro.DeploymentCreateInput) astro.Deployment); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(astro.Deployment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*astro.DeploymentCreateInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateImage provides a mock function with given fields: input
func (_m *Client) CreateImage(input astro.ImageCreateInput) (*astro.Image, error) {
	ret := _m.Called(input)

	var r0 *astro.Image
	if rf, ok := ret.Get(0).(func(astro.ImageCreateInput) *astro.Image); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*astro.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(astro.ImageCreateInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDeployment provides a mock function with given fields: input
func (_m *Client) DeleteDeployment(input astro.DeploymentDeleteInput) (astro.Deployment, error) {
	ret := _m.Called(input)

	var r0 astro.Deployment
	if rf, ok := ret.Get(0).(func(astro.DeploymentDeleteInput) astro.Deployment); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(astro.Deployment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(astro.DeploymentDeleteInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeployImage provides a mock function with given fields: input
func (_m *Client) DeployImage(input astro.ImageDeployInput) (*astro.Image, error) {
	ret := _m.Called(input)

	var r0 *astro.Image
	if rf, ok := ret.Get(0).(func(astro.ImageDeployInput) *astro.Image); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*astro.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(astro.ImageDeployInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeploymentConfig provides a mock function with given fields:
func (_m *Client) GetDeploymentConfig() (astro.DeploymentConfig, error) {
	ret := _m.Called()

	var r0 astro.DeploymentConfig
	if rf, ok := ret.Get(0).(func() astro.DeploymentConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(astro.DeploymentConfig)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeploymentHistory provides a mock function with given fields: vars
func (_m *Client) GetDeploymentHistory(vars map[string]interface{}) (astro.DeploymentHistory, error) {
	ret := _m.Called(vars)

	var r0 astro.DeploymentHistory
	if rf, ok := ret.Get(0).(func(map[string]interface{}) astro.DeploymentHistory); ok {
		r0 = rf(vars)
	} else {
		r0 = ret.Get(0).(astro.DeploymentHistory)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(vars)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusters provides a mock function with given fields: organizationId
func (_m *Client) ListClusters(organizationId string) ([]astro.Cluster, error) {
	ret := _m.Called(organizationId)

	var r0 []astro.Cluster
	if rf, ok := ret.Get(0).(func(string) []astro.Cluster); ok {
		r0 = rf(organizationId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]astro.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(organizationId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeployments provides a mock function with given fields: input
func (_m *Client) ListDeployments(input astro.DeploymentsInput) ([]astro.Deployment, error) {
	ret := _m.Called(input)

	var r0 []astro.Deployment
	if rf, ok := ret.Get(0).(func(astro.DeploymentsInput) []astro.Deployment); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]astro.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(astro.DeploymentsInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInternalRuntimeReleases provides a mock function with given fields:
func (_m *Client) ListInternalRuntimeReleases() ([]astro.RuntimeRelease, error) {
	ret := _m.Called()

	var r0 []astro.RuntimeRelease
	if rf, ok := ret.Get(0).(func() []astro.RuntimeRelease); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]astro.RuntimeRelease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPublicRuntimeReleases provides a mock function with given fields:
func (_m *Client) ListPublicRuntimeReleases() ([]astro.RuntimeRelease, error) {
	ret := _m.Called()

	var r0 []astro.RuntimeRelease
	if rf, ok := ret.Get(0).(func() []astro.RuntimeRelease); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]astro.RuntimeRelease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUserRoleBindings provides a mock function with given fields:
func (_m *Client) ListUserRoleBindings() ([]astro.RoleBinding, error) {
	ret := _m.Called()

	var r0 []astro.RoleBinding
	if rf, ok := ret.Get(0).(func() []astro.RoleBinding); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]astro.RoleBinding)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkspaces provides a mock function with given fields:
func (_m *Client) ListWorkspaces() ([]astro.Workspace, error) {
	ret := _m.Called()

	var r0 []astro.Workspace
	if rf, ok := ret.Get(0).(func() []astro.Workspace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]astro.Workspace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyDeploymentVariable provides a mock function with given fields: input
func (_m *Client) ModifyDeploymentVariable(input astro.EnvironmentVariablesInput) ([]astro.EnvironmentVariablesObject, error) {
	ret := _m.Called(input)

	var r0 []astro.EnvironmentVariablesObject
	if rf, ok := ret.Get(0).(func(astro.EnvironmentVariablesInput) []astro.EnvironmentVariablesObject); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]astro.EnvironmentVariablesObject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(astro.EnvironmentVariablesInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeployment provides a mock function with given fields: input
func (_m *Client) UpdateDeployment(input *astro.DeploymentUpdateInput) (astro.Deployment, error) {
	ret := _m.Called(input)

	var r0 astro.Deployment
	if rf, ok := ret.Get(0).(func(*astro.DeploymentUpdateInput) astro.Deployment); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(astro.Deployment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*astro.DeploymentUpdateInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

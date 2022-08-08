// Code generated by mockery v2.10.0. DO NOT EDIT.

package houston_mocks

import (
	config "github.com/astronomer/astro-cli/config"
	houston "github.com/astronomer/astro-cli/houston"

	mock "github.com/stretchr/testify/mock"
)

// ClientInterface is an autogenerated mock type for the ClientInterface type
type ClientInterface struct {
	mock.Mock
}

// AddDeploymentTeam provides a mock function with given fields: deploymentID, teamID, role
func (_m *ClientInterface) AddDeploymentTeam(deploymentID string, teamID string, role string) (*houston.RoleBinding, error) {
	ret := _m.Called(deploymentID, teamID, role)

	var r0 *houston.RoleBinding
	if rf, ok := ret.Get(0).(func(string, string, string) *houston.RoleBinding); ok {
		r0 = rf(deploymentID, teamID, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.RoleBinding)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(deploymentID, teamID, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddDeploymentUser provides a mock function with given fields: variables
func (_m *ClientInterface) AddDeploymentUser(variables houston.UpdateDeploymentUserRequest) (*houston.RoleBinding, error) {
	ret := _m.Called(variables)

	var r0 *houston.RoleBinding
	if rf, ok := ret.Get(0).(func(houston.UpdateDeploymentUserRequest) *houston.RoleBinding); ok {
		r0 = rf(variables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.RoleBinding)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(houston.UpdateDeploymentUserRequest) error); ok {
		r1 = rf(variables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddWorkspaceTeam provides a mock function with given fields: workspaceID, teamID, role
func (_m *ClientInterface) AddWorkspaceTeam(workspaceID string, teamID string, role string) (*houston.Workspace, error) {
	ret := _m.Called(workspaceID, teamID, role)

	var r0 *houston.Workspace
	if rf, ok := ret.Get(0).(func(string, string, string) *houston.Workspace); ok {
		r0 = rf(workspaceID, teamID, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Workspace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(workspaceID, teamID, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddWorkspaceUser provides a mock function with given fields: workspaceID, email, role
func (_m *ClientInterface) AddWorkspaceUser(workspaceID string, email string, role string) (*houston.Workspace, error) {
	ret := _m.Called(workspaceID, email, role)

	var r0 *houston.Workspace
	if rf, ok := ret.Get(0).(func(string, string, string) *houston.Workspace); ok {
		r0 = rf(workspaceID, email, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Workspace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(workspaceID, email, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticateWithBasicAuth provides a mock function with given fields: username, password, ctx
func (_m *ClientInterface) AuthenticateWithBasicAuth(username string, password string, ctx *config.Context) (string, error) {
	ret := _m.Called(username, password, ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, *config.Context) string); ok {
		r0 = rf(username, password, ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *config.Context) error); ok {
		r1 = rf(username, password, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelUpdateDeploymentRuntime provides a mock function with given fields: variables
func (_m *ClientInterface) CancelUpdateDeploymentRuntime(variables map[string]interface{}) (*houston.Deployment, error) {
	ret := _m.Called(variables)

	var r0 *houston.Deployment
	if rf, ok := ret.Get(0).(func(map[string]interface{}) *houston.Deployment); ok {
		r0 = rf(variables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(variables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDeployment provides a mock function with given fields: vars
func (_m *ClientInterface) CreateDeployment(vars map[string]interface{}) (*houston.Deployment, error) {
	ret := _m.Called(vars)

	var r0 *houston.Deployment
	if rf, ok := ret.Get(0).(func(map[string]interface{}) *houston.Deployment); ok {
		r0 = rf(vars)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(vars)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDeploymentServiceAccount provides a mock function with given fields: variables
func (_m *ClientInterface) CreateDeploymentServiceAccount(variables *houston.CreateServiceAccountRequest) (*houston.DeploymentServiceAccount, error) {
	ret := _m.Called(variables)

	var r0 *houston.DeploymentServiceAccount
	if rf, ok := ret.Get(0).(func(*houston.CreateServiceAccountRequest) *houston.DeploymentServiceAccount); ok {
		r0 = rf(variables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.DeploymentServiceAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*houston.CreateServiceAccountRequest) error); ok {
		r1 = rf(variables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTeamSystemRoleBinding provides a mock function with given fields: teamID, role
func (_m *ClientInterface) CreateTeamSystemRoleBinding(teamID string, role string) (string, error) {
	ret := _m.Called(teamID, role)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(teamID, role)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(teamID, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: email, password
func (_m *ClientInterface) CreateUser(email string, password string) (*houston.AuthUser, error) {
	ret := _m.Called(email, password)

	var r0 *houston.AuthUser
	if rf, ok := ret.Get(0).(func(string, string) *houston.AuthUser); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.AuthUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateWorkspace provides a mock function with given fields: label, description
func (_m *ClientInterface) CreateWorkspace(label string, description string) (*houston.Workspace, error) {
	ret := _m.Called(label, description)

	var r0 *houston.Workspace
	if rf, ok := ret.Get(0).(func(string, string) *houston.Workspace); ok {
		r0 = rf(label, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Workspace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(label, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateWorkspaceServiceAccount provides a mock function with given fields: variables
func (_m *ClientInterface) CreateWorkspaceServiceAccount(variables *houston.CreateServiceAccountRequest) (*houston.WorkspaceServiceAccount, error) {
	ret := _m.Called(variables)

	var r0 *houston.WorkspaceServiceAccount
	if rf, ok := ret.Get(0).(func(*houston.CreateServiceAccountRequest) *houston.WorkspaceServiceAccount); ok {
		r0 = rf(variables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.WorkspaceServiceAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*houston.CreateServiceAccountRequest) error); ok {
		r1 = rf(variables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDeployment provides a mock function with given fields: deploymentID, doHardDelete
func (_m *ClientInterface) DeleteDeployment(deploymentID string, doHardDelete bool) (*houston.Deployment, error) {
	ret := _m.Called(deploymentID, doHardDelete)

	var r0 *houston.Deployment
	if rf, ok := ret.Get(0).(func(string, bool) *houston.Deployment); ok {
		r0 = rf(deploymentID, doHardDelete)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(deploymentID, doHardDelete)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDeploymentServiceAccount provides a mock function with given fields: deploymentID, serviceAccountID
func (_m *ClientInterface) DeleteDeploymentServiceAccount(deploymentID string, serviceAccountID string) (*houston.ServiceAccount, error) {
	ret := _m.Called(deploymentID, serviceAccountID)

	var r0 *houston.ServiceAccount
	if rf, ok := ret.Get(0).(func(string, string) *houston.ServiceAccount); ok {
		r0 = rf(deploymentID, serviceAccountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.ServiceAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(deploymentID, serviceAccountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDeploymentUser provides a mock function with given fields: deploymentID, email
func (_m *ClientInterface) DeleteDeploymentUser(deploymentID string, email string) (*houston.RoleBinding, error) {
	ret := _m.Called(deploymentID, email)

	var r0 *houston.RoleBinding
	if rf, ok := ret.Get(0).(func(string, string) *houston.RoleBinding); ok {
		r0 = rf(deploymentID, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.RoleBinding)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(deploymentID, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTeamSystemRoleBinding provides a mock function with given fields: teamID, role
func (_m *ClientInterface) DeleteTeamSystemRoleBinding(teamID string, role string) (string, error) {
	ret := _m.Called(teamID, role)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(teamID, role)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(teamID, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkspace provides a mock function with given fields: workspaceID
func (_m *ClientInterface) DeleteWorkspace(workspaceID string) (*houston.Workspace, error) {
	ret := _m.Called(workspaceID)

	var r0 *houston.Workspace
	if rf, ok := ret.Get(0).(func(string) *houston.Workspace); ok {
		r0 = rf(workspaceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Workspace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(workspaceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkspaceServiceAccount provides a mock function with given fields: workspaceID, serviceAccountID
func (_m *ClientInterface) DeleteWorkspaceServiceAccount(workspaceID string, serviceAccountID string) (*houston.ServiceAccount, error) {
	ret := _m.Called(workspaceID, serviceAccountID)

	var r0 *houston.ServiceAccount
	if rf, ok := ret.Get(0).(func(string, string) *houston.ServiceAccount); ok {
		r0 = rf(workspaceID, serviceAccountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.ServiceAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(workspaceID, serviceAccountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkspaceTeam provides a mock function with given fields: workspaceID, teamID
func (_m *ClientInterface) DeleteWorkspaceTeam(workspaceID string, teamID string) (*houston.Workspace, error) {
	ret := _m.Called(workspaceID, teamID)

	var r0 *houston.Workspace
	if rf, ok := ret.Get(0).(func(string, string) *houston.Workspace); ok {
		r0 = rf(workspaceID, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Workspace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(workspaceID, teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkspaceUser provides a mock function with given fields: workspaceID, userID
func (_m *ClientInterface) DeleteWorkspaceUser(workspaceID string, userID string) (*houston.Workspace, error) {
	ret := _m.Called(workspaceID, userID)

	var r0 *houston.Workspace
	if rf, ok := ret.Get(0).(func(string, string) *houston.Workspace); ok {
		r0 = rf(workspaceID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Workspace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(workspaceID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppConfig provides a mock function with given fields:
func (_m *ClientInterface) GetAppConfig() (*houston.AppConfig, error) {
	ret := _m.Called()

	var r0 *houston.AppConfig
	if rf, ok := ret.Get(0).(func() *houston.AppConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.AppConfig)
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

// GetAuthConfig provides a mock function with given fields: ctx
func (_m *ClientInterface) GetAuthConfig(ctx *config.Context) (*houston.AuthConfig, error) {
	ret := _m.Called(ctx)

	var r0 *houston.AuthConfig
	if rf, ok := ret.Get(0).(func(*config.Context) *houston.AuthConfig); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.AuthConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*config.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAvailableNamespaces provides a mock function with given fields:
func (_m *ClientInterface) GetAvailableNamespaces() ([]houston.Namespace, error) {
	ret := _m.Called()

	var r0 []houston.Namespace
	if rf, ok := ret.Get(0).(func() []houston.Namespace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]houston.Namespace)
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

// GetDeployment provides a mock function with given fields: deploymentID
func (_m *ClientInterface) GetDeployment(deploymentID string) (*houston.Deployment, error) {
	ret := _m.Called(deploymentID)

	var r0 *houston.Deployment
	if rf, ok := ret.Get(0).(func(string) *houston.Deployment); ok {
		r0 = rf(deploymentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeploymentConfig provides a mock function with given fields:
func (_m *ClientInterface) GetDeploymentConfig() (*houston.DeploymentConfig, error) {
	ret := _m.Called()

	var r0 *houston.DeploymentConfig
	if rf, ok := ret.Get(0).(func() *houston.DeploymentConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.DeploymentConfig)
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

// GetRuntimeReleases provides a mock function with given fields: airflowVersion
func (_m *ClientInterface) GetRuntimeReleases(airflowVersion string) (houston.RuntimeReleases, error) {
	ret := _m.Called(airflowVersion)

	var r0 houston.RuntimeReleases
	if rf, ok := ret.Get(0).(func(string) houston.RuntimeReleases); ok {
		r0 = rf(airflowVersion)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(houston.RuntimeReleases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(airflowVersion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeam provides a mock function with given fields: teamID
func (_m *ClientInterface) GetTeam(teamID string) (*houston.Team, error) {
	ret := _m.Called(teamID)

	var r0 *houston.Team
	if rf, ok := ret.Get(0).(func(string) *houston.Team); ok {
		r0 = rf(teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamUsers provides a mock function with given fields: teamID
func (_m *ClientInterface) GetTeamUsers(teamID string) ([]houston.User, error) {
	ret := _m.Called(teamID)

	var r0 []houston.User
	if rf, ok := ret.Get(0).(func(string) []houston.User); ok {
		r0 = rf(teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]houston.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkspace provides a mock function with given fields: workspaceID
func (_m *ClientInterface) GetWorkspace(workspaceID string) (*houston.Workspace, error) {
	ret := _m.Called(workspaceID)

	var r0 *houston.Workspace
	if rf, ok := ret.Get(0).(func(string) *houston.Workspace); ok {
		r0 = rf(workspaceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Workspace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(workspaceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkspaceTeamRole provides a mock function with given fields: workspaceID, teamID
func (_m *ClientInterface) GetWorkspaceTeamRole(workspaceID string, teamID string) (*houston.Team, error) {
	ret := _m.Called(workspaceID, teamID)

	var r0 *houston.Team
	if rf, ok := ret.Get(0).(func(string, string) *houston.Team); ok {
		r0 = rf(workspaceID, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(workspaceID, teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkspaceUserRole provides a mock function with given fields: workspaceID, email
func (_m *ClientInterface) GetWorkspaceUserRole(workspaceID string, email string) (houston.WorkspaceUserRoleBindings, error) {
	ret := _m.Called(workspaceID, email)

	var r0 houston.WorkspaceUserRoleBindings
	if rf, ok := ret.Get(0).(func(string, string) houston.WorkspaceUserRoleBindings); ok {
		r0 = rf(workspaceID, email)
	} else {
		r0 = ret.Get(0).(houston.WorkspaceUserRoleBindings)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(workspaceID, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeploymentLogs provides a mock function with given fields: filters
func (_m *ClientInterface) ListDeploymentLogs(filters houston.ListDeploymentLogsRequest) ([]houston.DeploymentLog, error) {
	ret := _m.Called(filters)

	var r0 []houston.DeploymentLog
	if rf, ok := ret.Get(0).(func(houston.ListDeploymentLogsRequest) []houston.DeploymentLog); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]houston.DeploymentLog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(houston.ListDeploymentLogsRequest) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeploymentServiceAccounts provides a mock function with given fields: deploymentID
func (_m *ClientInterface) ListDeploymentServiceAccounts(deploymentID string) ([]houston.ServiceAccount, error) {
	ret := _m.Called(deploymentID)

	var r0 []houston.ServiceAccount
	if rf, ok := ret.Get(0).(func(string) []houston.ServiceAccount); ok {
		r0 = rf(deploymentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]houston.ServiceAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeploymentTeamsAndRoles provides a mock function with given fields: deploymentID
func (_m *ClientInterface) ListDeploymentTeamsAndRoles(deploymentID string) ([]houston.Team, error) {
	ret := _m.Called(deploymentID)

	var r0 []houston.Team
	if rf, ok := ret.Get(0).(func(string) []houston.Team); ok {
		r0 = rf(deploymentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]houston.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeploymentUsers provides a mock function with given fields: filters
func (_m *ClientInterface) ListDeploymentUsers(filters houston.ListDeploymentUsersRequest) ([]houston.DeploymentUser, error) {
	ret := _m.Called(filters)

	var r0 []houston.DeploymentUser
	if rf, ok := ret.Get(0).(func(houston.ListDeploymentUsersRequest) []houston.DeploymentUser); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]houston.DeploymentUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(houston.ListDeploymentUsersRequest) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeployments provides a mock function with given fields: filters
func (_m *ClientInterface) ListDeployments(filters houston.ListDeploymentsRequest) ([]houston.Deployment, error) {
	ret := _m.Called(filters)

	var r0 []houston.Deployment
	if rf, ok := ret.Get(0).(func(houston.ListDeploymentsRequest) []houston.Deployment); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]houston.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(houston.ListDeploymentsRequest) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTeams provides a mock function with given fields: cursor, take
func (_m *ClientInterface) ListTeams(cursor string, take int) (houston.ListTeamsResp, error) {
	ret := _m.Called(cursor, take)

	var r0 houston.ListTeamsResp
	if rf, ok := ret.Get(0).(func(string, int) houston.ListTeamsResp); ok {
		r0 = rf(cursor, take)
	} else {
		r0 = ret.Get(0).(houston.ListTeamsResp)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(cursor, take)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkspaceServiceAccounts provides a mock function with given fields: workspaceID
func (_m *ClientInterface) ListWorkspaceServiceAccounts(workspaceID string) ([]houston.ServiceAccount, error) {
	ret := _m.Called(workspaceID)

	var r0 []houston.ServiceAccount
	if rf, ok := ret.Get(0).(func(string) []houston.ServiceAccount); ok {
		r0 = rf(workspaceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]houston.ServiceAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(workspaceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkspaceTeamsAndRoles provides a mock function with given fields: workspaceID
func (_m *ClientInterface) ListWorkspaceTeamsAndRoles(workspaceID string) ([]houston.Team, error) {
	ret := _m.Called(workspaceID)

	var r0 []houston.Team
	if rf, ok := ret.Get(0).(func(string) []houston.Team); ok {
		r0 = rf(workspaceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]houston.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(workspaceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkspaceUserAndRoles provides a mock function with given fields: workspaceID
func (_m *ClientInterface) ListWorkspaceUserAndRoles(workspaceID string) ([]houston.WorkspaceUserRoleBindings, error) {
	ret := _m.Called(workspaceID)

	var r0 []houston.WorkspaceUserRoleBindings
	if rf, ok := ret.Get(0).(func(string) []houston.WorkspaceUserRoleBindings); ok {
		r0 = rf(workspaceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]houston.WorkspaceUserRoleBindings)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(workspaceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkspaces provides a mock function with given fields:
func (_m *ClientInterface) ListWorkspaces() ([]houston.Workspace, error) {
	ret := _m.Called()

	var r0 []houston.Workspace
	if rf, ok := ret.Get(0).(func() []houston.Workspace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]houston.Workspace)
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

// RemoveDeploymentTeam provides a mock function with given fields: deploymentID, teamID
func (_m *ClientInterface) RemoveDeploymentTeam(deploymentID string, teamID string) (*houston.RoleBinding, error) {
	ret := _m.Called(deploymentID, teamID)

	var r0 *houston.RoleBinding
	if rf, ok := ret.Get(0).(func(string, string) *houston.RoleBinding); ok {
		r0 = rf(deploymentID, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.RoleBinding)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(deploymentID, teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeployment provides a mock function with given fields: variables
func (_m *ClientInterface) UpdateDeployment(variables map[string]interface{}) (*houston.Deployment, error) {
	ret := _m.Called(variables)

	var r0 *houston.Deployment
	if rf, ok := ret.Get(0).(func(map[string]interface{}) *houston.Deployment); ok {
		r0 = rf(variables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(variables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeploymentAirflow provides a mock function with given fields: variables
func (_m *ClientInterface) UpdateDeploymentAirflow(variables map[string]interface{}) (*houston.Deployment, error) {
	ret := _m.Called(variables)

	var r0 *houston.Deployment
	if rf, ok := ret.Get(0).(func(map[string]interface{}) *houston.Deployment); ok {
		r0 = rf(variables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(variables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeploymentImage provides a mock function with given fields: req
func (_m *ClientInterface) UpdateDeploymentImage(req houston.UpdateDeploymentImageRequest) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(houston.UpdateDeploymentImageRequest) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeploymentRuntime provides a mock function with given fields: variables
func (_m *ClientInterface) UpdateDeploymentRuntime(variables map[string]interface{}) (*houston.Deployment, error) {
	ret := _m.Called(variables)

	var r0 *houston.Deployment
	if rf, ok := ret.Get(0).(func(map[string]interface{}) *houston.Deployment); ok {
		r0 = rf(variables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(variables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeploymentTeamRole provides a mock function with given fields: deploymentID, teamID, role
func (_m *ClientInterface) UpdateDeploymentTeamRole(deploymentID string, teamID string, role string) (*houston.RoleBinding, error) {
	ret := _m.Called(deploymentID, teamID, role)

	var r0 *houston.RoleBinding
	if rf, ok := ret.Get(0).(func(string, string, string) *houston.RoleBinding); ok {
		r0 = rf(deploymentID, teamID, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.RoleBinding)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(deploymentID, teamID, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeploymentUser provides a mock function with given fields: variables
func (_m *ClientInterface) UpdateDeploymentUser(variables houston.UpdateDeploymentUserRequest) (*houston.RoleBinding, error) {
	ret := _m.Called(variables)

	var r0 *houston.RoleBinding
	if rf, ok := ret.Get(0).(func(houston.UpdateDeploymentUserRequest) *houston.RoleBinding); ok {
		r0 = rf(variables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.RoleBinding)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(houston.UpdateDeploymentUserRequest) error); ok {
		r1 = rf(variables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWorkspace provides a mock function with given fields: workspaceID, args
func (_m *ClientInterface) UpdateWorkspace(workspaceID string, args map[string]string) (*houston.Workspace, error) {
	ret := _m.Called(workspaceID, args)

	var r0 *houston.Workspace
	if rf, ok := ret.Get(0).(func(string, map[string]string) *houston.Workspace); ok {
		r0 = rf(workspaceID, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*houston.Workspace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]string) error); ok {
		r1 = rf(workspaceID, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWorkspaceTeamRole provides a mock function with given fields: workspaceID, teamID, role
func (_m *ClientInterface) UpdateWorkspaceTeamRole(workspaceID string, teamID string, role string) (string, error) {
	ret := _m.Called(workspaceID, teamID, role)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(workspaceID, teamID, role)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(workspaceID, teamID, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWorkspaceUserRole provides a mock function with given fields: workspaceID, email, role
func (_m *ClientInterface) UpdateWorkspaceUserRole(workspaceID string, email string, role string) (string, error) {
	ret := _m.Called(workspaceID, email, role)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(workspaceID, email, role)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(workspaceID, email, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

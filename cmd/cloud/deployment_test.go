package cloud

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	airflowversions "github.com/astronomer/astro-cli/airflow_versions"
	astro "github.com/astronomer/astro-cli/astro-client"
	astrocore "github.com/astronomer/astro-cli/astro-client-core"
	astrocore_mocks "github.com/astronomer/astro-cli/astro-client-core/mocks"
	astro_mocks "github.com/astronomer/astro-cli/astro-client/mocks"
	"github.com/astronomer/astro-cli/cloud/deployment"
	"github.com/astronomer/astro-cli/config"
	"github.com/astronomer/astro-cli/context"
	"github.com/astronomer/astro-cli/pkg/fileutil"
	testUtil "github.com/astronomer/astro-cli/pkg/testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/mock"
)

func execDeploymentCmd(args ...string) (string, error) {
	buf := new(bytes.Buffer)
	cmd := newDeploymentRootCmd(buf)
	cmd.SetOut(buf)
	cmd.SetArgs(args)
	testUtil.SetupOSArgsForGinkgo()
	_, err := cmd.ExecuteC()
	return buf.String(), err
}

func (s *Suite) TestDeploymentRootCommand() {
	testUtil.InitTestConfig(testUtil.CloudPlatform)
	buf := new(bytes.Buffer)
	deplyCmd := newDeploymentRootCmd(os.Stdout)
	deplyCmd.SetOut(buf)
	testUtil.SetupOSArgsForGinkgo()
	_, err := deplyCmd.ExecuteC()
	s.NoError(err)
	s.Contains(buf.String(), "deployment")
}

func (s *Suite) TestDeploymentList() {
	testUtil.InitTestConfig(testUtil.CloudPlatform)

	mockClient := new(astro_mocks.Client)
	mockClient.On("ListDeployments", mock.Anything, "").Return([]astro.Deployment{{ID: "test-id-1"}, {ID: "test-id-2"}}, nil).Once()
	astroClient = mockClient

	cmdArgs := []string{"list", "-a"}
	resp, err := execDeploymentCmd(cmdArgs...)
	s.NoError(err)
	s.Contains(resp, "test-id-1")
	s.Contains(resp, "test-id-2")
	mockClient.AssertExpectations(s.T())
}

func (s *Suite) TestDeploymentLogs() {
	testUtil.InitTestConfig(testUtil.CloudPlatform)

	deploymentID := "test-id"
	logLevels := []string{"WARN", "ERROR", "INFO"}
	mockInput := map[string]interface{}{
		"deploymentId":  deploymentID,
		"logCountLimit": logCount,
		"start":         "-24hrs",
		"logLevels":     logLevels,
	}

	mockClient := new(astro_mocks.Client)
	mockClient.On("ListDeployments", mock.Anything, mock.Anything).Return([]astro.Deployment{{ID: "test-id"}, {ID: "test-id-2"}}, nil).Once()
	mockClient.On("GetDeploymentHistory", mockInput).Return(astro.DeploymentHistory{DeploymentID: deploymentID, SchedulerLogs: []astro.SchedulerLog{{Raw: "test log line"}}}, nil).Once()
	astroClient = mockClient

	cmdArgs := []string{"logs", "test-id", "-w", "-e", "-i"}
	_, err := execDeploymentCmd(cmdArgs...)
	s.NoError(err)
	mockClient.AssertExpectations(s.T())
}

func (s *Suite) TestDeploymentCreate() {
	testUtil.InitTestConfig(testUtil.CloudPlatform)

	ws := "test-ws-id"
	csID := "test-cluster-id"

	deploymentCreateInput := astro.CreateDeploymentInput{
		WorkspaceID:           ws,
		ClusterID:             csID,
		Label:                 "test-name",
		Description:           "",
		RuntimeReleaseVersion: "4.2.5",
		DagDeployEnabled:      true,
		DeploymentSpec: astro.DeploymentCreateSpec{
			Executor: "CeleryExecutor",
			Scheduler: astro.Scheduler{
				AU:       5,
				Replicas: 1,
			},
		},
	}

	deploymentCreateInput1 := astro.CreateDeploymentInput{
		WorkspaceID:           ws,
		ClusterID:             csID,
		Label:                 "test-name",
		Description:           "",
		RuntimeReleaseVersion: "4.2.5",
		DagDeployEnabled:      false,
		DeploymentSpec: astro.DeploymentCreateSpec{
			Executor: "CeleryExecutor",
			Scheduler: astro.Scheduler{
				AU:       5,
				Replicas: 1,
			},
		},
	}

	mockClient := new(astro_mocks.Client)
	mockClient.On("GetDeploymentConfig").Return(astro.DeploymentConfig{
		Components: astro.Components{
			Scheduler: astro.SchedulerConfig{
				AU: astro.AuConfig{
					Default: 5,
					Limit:   24,
				},
				Replicas: astro.ReplicasConfig{
					Default: 1,
					Minimum: 1,
					Limit:   4,
				},
			},
		},
		RuntimeReleases: []astro.RuntimeRelease{
			{
				Version: "4.2.5",
			},
		},
	}, nil).Times(10)
	mockClient.On("ListWorkspaces", "test-org-id").Return([]astro.Workspace{{ID: ws, OrganizationID: "test-org-id"}}, nil).Times(5)
	mockClient.On("ListClusters", "test-org-id").Return([]astro.Cluster{{ID: csID}}, nil).Times(4)
	mockClient.On("CreateDeployment", &deploymentCreateInput).Return(astro.Deployment{ID: "test-id"}, nil).Twice()
	mockClient.On("CreateDeployment", &deploymentCreateInput1).Return(astro.Deployment{ID: "test-id"}, nil).Times(6)
	deploymentCreateInput2 := astro.CreateDeploymentInput{
		WorkspaceID:           ws,
		ClusterID:             csID,
		Label:                 "test-name",
		Description:           "",
		RuntimeReleaseVersion: "4.2.5",
		DagDeployEnabled:      false,
		SchedulerSize:         "small",
		IsHighAvailability:    false,
		DeploymentSpec: astro.DeploymentCreateSpec{
			Executor: "CeleryExecutor",
			Scheduler: astro.Scheduler{
				AU:       5,
				Replicas: 1,
			},
		},
	}
	mockClient.On("CreateDeployment", &deploymentCreateInput2).Return(astro.Deployment{ID: "test-id"}, nil).Once()
	astroClient = mockClient

	mockResponse := &airflowversions.Response{
		RuntimeVersions: map[string]airflowversions.RuntimeVersion{
			"4.2.5": {Metadata: airflowversions.RuntimeVersionMetadata{AirflowVersion: "2.2.5", Channel: "stable"}, Migrations: airflowversions.RuntimeVersionMigrations{}},
		},
	}
	jsonResponse, err := json.Marshal(mockResponse)
	s.NoError(err)

	httpClient = testUtil.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBuffer(jsonResponse)),
			Header:     make(http.Header),
		}
	})

	s.Run("creates a deployment when dag-deploy is disabled", func() {
		cmdArgs := []string{"create", "--name", "test-name", "--workspace-id", ws, "--cluster-id", csID, "--dag-deploy", "disable"}
		_, err = execDeploymentCmd(cmdArgs...)
		s.NoError(err)
	})
	s.Run("creates a deployment when dag deploy is enabled", func() {
		cmdArgs := []string{"create", "--name", "test-name", "--workspace-id", ws, "--cluster-id", csID, "--dag-deploy", "enable"}
		_, err = execDeploymentCmd(cmdArgs...)
		s.NoError(err)
	})
	s.Run("creates a deployment when executor is specified", func() {
		deploymentCreateInput1.DeploymentSpec.Executor = "KubernetesExecutor"
		defer func() { deploymentCreateInput1.DeploymentSpec.Executor = "CeleryExecutor" }()
		cmdArgs := []string{"create", "--name", "test-name", "--workspace-id", ws, "--cluster-id", csID, "--dag-deploy", "disable", "--executor", "KubernetesExecutor"}
		_, err = execDeploymentCmd(cmdArgs...)
		s.NoError(err)
	})
	s.Run("creates a deployment with default executor", func() {
		cmdArgs := []string{"create", "--name", "test-name", "--workspace-id", ws, "--cluster-id", csID, "--dag-deploy", "disable"}
		_, err = execDeploymentCmd(cmdArgs...)
		s.NoError(err)
	})
	s.Run("returns an error if dag-deploy flag has an incorrect value", func() {
		cmdArgs := []string{"create", "--name", "test-name", "--workspace-id", ws, "--cluster-id", csID, "--dag-deploy", "some-value"}
		_, err = execDeploymentCmd(cmdArgs...)
		s.Error(err)
	})
	s.Run("returns an error if executor has an incorrect value", func() {
		cmdArgs := []string{"create", "--name", "test-name", "--workspace-id", ws, "--cluster-id", csID, "--dag-deploy", "disable", "--executor", "KubeExecutor"}
		_, err = execDeploymentCmd(cmdArgs...)
		s.ErrorContains(err, "KubeExecutor is not a valid executor")
	})
	s.Run("creates a deployment from file", func() {
		orgID := "test-org-id"
		filePath := "./test-deployment.yaml"
		data := `
deployment:
  environment_variables:
    - is_secret: false
      key: foo
      updated_at: NOW
      value: bar
    - is_secret: true
      key: bar
      updated_at: NOW+1
      value: baz
  configuration:
    name: test-deployment-label
    description: description
    runtime_version: 6.0.0
    dag_deploy_enabled: true
    executor: CeleryExecutor
    scheduler_au: 5
    scheduler_count: 3
    cluster_name: test-cluster
    workspace_name: test-workspace
  worker_queues:
    - name: default
      is_default: true
      max_worker_count: 130
      min_worker_count: 12
      worker_concurrency: 180
      worker_type: test-worker-1
    - name: test-queue-1
      is_default: false
      max_worker_count: 175
      min_worker_count: 8
      worker_concurrency: 176
      worker_type: test-worker-2
  metadata:
    deployment_id: test-deployment-id
    workspace_id: test-ws-id
    cluster_id: cluster-id
    release_name: great-release-name
    airflow_version: 2.4.0
    status: UNHEALTHY
    created_at: 2022-11-17T13:25:55.275697-08:00
    updated_at: 2022-11-17T13:25:55.275697-08:00
    deployment_url: cloud.astronomer.io/test-ws-id/deployments/test-deployment-id/analytics
    webserver_url: some-url
  alert_emails:
    - test1@test.com
    - test2@test.com
`
		clusters := []astro.Cluster{
			{
				ID:   "test-cluster-id",
				Name: "test-cluster",
				NodePools: []astro.NodePool{
					{
						ID:               "test-pool-id",
						IsDefault:        false,
						NodeInstanceType: "test-worker-1",
					},
					{
						ID:               "test-pool-id-2",
						IsDefault:        false,
						NodeInstanceType: "test-worker-2",
					},
				},
			},
		}
		createdDeployment := astro.Deployment{
			ID:    "test-deployment-id",
			Label: "test-deployment-label",
		}
		mockWorkerQueueDefaultOptions := astro.WorkerQueueDefaultOptions{
			MinWorkerCount: astro.WorkerQueueOption{
				Floor:   1,
				Ceiling: 20,
				Default: 5,
			},
			MaxWorkerCount: astro.WorkerQueueOption{
				Floor:   16,
				Ceiling: 200,
				Default: 125,
			},
			WorkerConcurrency: astro.WorkerQueueOption{
				Floor:   175,
				Ceiling: 275,
				Default: 180,
			},
		}
		mockClient = new(astro_mocks.Client)
		mockClient.On("ListWorkspaces", orgID).Return([]astro.Workspace{{ID: ws, OrganizationID: orgID, Label: "test-workspace"}}, nil)
		mockClient.On("ListClusters", orgID).Return(clusters, nil)
		mockClient.On("ListDeployments", orgID, "").Return([]astro.Deployment{}, nil).Once()
		mockClient.On("GetWorkerQueueOptions").Return(mockWorkerQueueDefaultOptions, nil).Once()
		mockClient.On("CreateDeployment", mock.Anything).Return(createdDeployment, nil)
		mockClient.On("ModifyDeploymentVariable", mock.Anything).Return([]astro.EnvironmentVariablesObject{}, nil)
		mockClient.On("UpdateAlertEmails", mock.Anything).Return(astro.DeploymentAlerts{}, nil)
		mockClient.On("ListDeployments", orgID, ws).Return([]astro.Deployment{createdDeployment}, nil)
		origClient := astroClient
		astroClient = mockClient
		fileutil.WriteStringToFile(filePath, data)
		defer func() {
			astroClient = origClient
			afero.NewOsFs().Remove(filePath)
		}()
		cmdArgs := []string{"create", "--deployment-file", "test-deployment.yaml"}
		_, err = execDeploymentCmd(cmdArgs...)
		s.NoError(err)
		mockClient.AssertExpectations(s.T())
	})
	s.Run("returns an error if creating a deployment from file fails", func() {
		cmdArgs := []string{"create", "--deployment-file", "test-file-name.json"}
		_, err = execDeploymentCmd(cmdArgs...)
		s.ErrorContains(err, "open test-file-name.json: no such file or directory")
	})
	s.Run("returns an error if from-file is specified with any other flags", func() {
		cmdArgs := []string{"create", "--deployment-file", "test-deployment.yaml", "--description", "fail"}
		_, err = execDeploymentCmd(cmdArgs...)
		s.ErrorIs(err, errFlag)
	})
	s.Run("creates a deployment with cloud provider and region", func() {
		ctx, err := context.GetCurrentContext()
		s.NoError(err)
		ctx.SetContextKey("organization_product", "HOSTED")
		ctx.SetContextKey("organization", "test-org-id")
		ctx.SetContextKey("workspace", ws)
		mockOKResponse := &astrocore.GetSharedClusterResponse{
			HTTPResponse: &http.Response{
				StatusCode: 200,
			},
			JSON200: &astrocore.SharedCluster{Id: csID},
		}
		mockCoreClient := new(astrocore_mocks.ClientWithResponsesInterface)
		astroCoreClient = mockCoreClient
		mockCoreClient.On("GetSharedClusterWithResponse", mock.Anything, mock.Anything).Return(mockOKResponse, nil).Once()
		cmdArgs := []string{
			"create", "--name", "test-name", "--workspace-id", ws, "--dag-deploy", "disable",
			"--cloud-provider", "gcp", "--region", "us-central1",
		}
		_, err = execDeploymentCmd(cmdArgs...)
		s.NoError(err)
		mockCoreClient.AssertExpectations(s.T())
	})
	s.Run("returns an error with incorrect high-availability value", func() {
		ctx, err := context.GetCurrentContext()
		s.NoError(err)
		ctx.SetContextKey("organization_product", "HOSTED")
		ctx.SetContextKey("organization", "test-org-id")
		ctx.SetContextKey("workspace", ws)
		cmdArgs := []string{
			"create", "--name", "test-name", "--workspace-id", ws, "--dag-deploy", "disable",
			"--executor", "KubernetesExecutor", "--cloud-provider", "gcp", "--region", "us-east1", "--high-availability", "some-value",
		}
		_, err = execDeploymentCmd(cmdArgs...)
		s.ErrorContains(err, "Invalid --high-availability value")
	})
	s.Run("returns an error if cloud provider is not valid", func() {
		ctx, err := context.GetCurrentContext()
		s.NoError(err)
		ctx.SetContextKey("organization_product", "HOSTED")
		ctx.SetContextKey("organization", "test-org-id")
		ctx.SetContextKey("workspace", ws)
		cmdArgs := []string{
			"create", "--name", "test-name", "--workspace-id", ws, "--dag-deploy", "disable",
			"--executor", "KubernetesExecutor", "--cloud-provider", "azure",
		}
		_, err = execDeploymentCmd(cmdArgs...)
		s.ErrorContains(err, "azure is not a valid cloud provider. It can only be gcp")
	})
	mockClient.AssertExpectations(s.T())
}

func (s *Suite) TestDeploymentUpdate() {
	testUtil.InitTestConfig(testUtil.CloudPlatform)

	ws := "test-ws-id"
	deploymentResp := astro.Deployment{
		ID:             "test-id",
		Label:          "test-name",
		RuntimeRelease: astro.RuntimeRelease{Version: "4.2.5"},
		DeploymentSpec: astro.DeploymentSpec{Executor: "CeleryExecutor", Scheduler: astro.Scheduler{AU: 5, Replicas: 3}},
	}
	deploymentUpdateInput := astro.UpdateDeploymentInput{
		ID:          "test-id",
		ClusterID:   "",
		Label:       "test-name",
		Description: "",
		DeploymentSpec: astro.DeploymentCreateSpec{
			Executor:  "CeleryExecutor",
			Scheduler: astro.Scheduler{AU: 5, Replicas: 3},
		},
		WorkerQueues: nil,
	}

	mockClient := new(astro_mocks.Client)
	mockClient.On("GetDeploymentConfig").Return(astro.DeploymentConfig{
		Components: astro.Components{
			Scheduler: astro.SchedulerConfig{
				AU: astro.AuConfig{
					Default: 5,
					Limit:   24,
				},
				Replicas: astro.ReplicasConfig{
					Default: 1,
					Minimum: 1,
					Limit:   4,
				},
			},
		},
		RuntimeReleases: []astro.RuntimeRelease{
			{
				Version: "4.2.5",
			},
		},
	}, nil).Once()
	mockClient.On("ListDeployments", mock.Anything, ws).Return([]astro.Deployment{deploymentResp}, nil).Once()
	mockClient.On("UpdateDeployment", &deploymentUpdateInput).Return(astro.Deployment{ID: "test-id"}, nil).Once()
	astroClient = mockClient

	s.Run("updates the deployment successfully", func() {
		cmdArgs := []string{"update", "test-id", "--name", "test-name", "--workspace-id", ws, "--force"}
		_, err := execDeploymentCmd(cmdArgs...)
		s.NoError(err)
	})
	s.Run("returns an error if dag-deploy has an incorrect value", func() {
		cmdArgs := []string{"update", "test-id", "--name", "test-name", "--workspace-id", ws, "--force", "--dag-deploy", "some-value"}
		_, err := execDeploymentCmd(cmdArgs...)
		s.Error(err)
	})
	s.Run("returns an error if executor has an incorrect value", func() {
		cmdArgs := []string{"update", "test-id", "--name", "test-name", "--workspace-id", ws, "--force", "--executor", "KubeExecutor"}
		_, err := execDeploymentCmd(cmdArgs...)
		s.ErrorContains(err, "KubeExecutor is not a valid executor")
	})
	s.Run("returns an error when getting workspace fails", func() {
		testUtil.InitTestConfig(testUtil.CloudPlatform)
		ctx, err := config.GetCurrentContext()
		s.NoError(err)
		ctx.Workspace = ""
		err = ctx.SetContext()
		s.NoError(err)
		defer testUtil.InitTestConfig(testUtil.CloudPlatform)
		expectedOut := "Usage:\n"
		cmdArgs := []string{"update", "-n", "doesnotexist"}
		resp, err := execDeploymentCmd(cmdArgs...)
		s.ErrorContains(err, "failed to find a valid workspace")
		s.Contains(resp, expectedOut)
	})
	s.Run("updates a deployment from file", func() {
		orgID := "test-org-id"
		filePath := "./test-deployment.yaml"
		data := `
deployment:
  environment_variables:
    - is_secret: false
      key: foo
      updated_at: NOW
      value: bar
    - is_secret: true
      key: bar
      updated_at: NOW+1
      value: baz
  configuration:
    name: test-deployment-label
    description: description
    runtime_version: 6.0.0
    dag_deploy_enabled: true
    executor: CeleryExecutor
    scheduler_au: 5
    scheduler_count: 3
    cluster_name: test-cluster
    workspace_name: test-workspace
  worker_queues:
    - name: default
      is_default: true
      max_worker_count: 130
      min_worker_count: 12
      worker_concurrency: 180
      worker_type: test-worker-1
    - name: test-queue-1
      is_default: false
      max_worker_count: 175
      min_worker_count: 8
      worker_concurrency: 176
      worker_type: test-worker-2
  metadata:
    deployment_id: test-deployment-id
    workspace_id: test-ws-id
    cluster_id: cluster-id
    release_name: great-release-name
    airflow_version: 2.4.0
    status: UNHEALTHY
    created_at: 2022-11-17T13:25:55.275697-08:00
    updated_at: 2022-11-17T13:25:55.275697-08:00
    deployment_url: cloud.astronomer.io/test-ws-id/deployments/test-deployment-id/analytics
    webserver_url: some-url
  alert_emails:
    - test1@test.com
    - test2@test.com
`
		clusters := []astro.Cluster{
			{
				ID:   "test-cluster-id",
				Name: "test-cluster",
				NodePools: []astro.NodePool{
					{
						ID:               "test-pool-id",
						IsDefault:        false,
						NodeInstanceType: "test-worker-1",
					},
					{
						ID:               "test-pool-id-2",
						IsDefault:        false,
						NodeInstanceType: "test-worker-2",
					},
				},
			},
		}
		updatedDeployment := astro.Deployment{
			ID:      "test-deployment-id",
			Label:   "test-deployment-label",
			Cluster: astro.Cluster{ID: "test-cluster-id", Name: "test-cluster"},
		}
		mockWorkerQueueDefaultOptions := astro.WorkerQueueDefaultOptions{
			MinWorkerCount: astro.WorkerQueueOption{
				Floor:   1,
				Ceiling: 20,
				Default: 5,
			},
			MaxWorkerCount: astro.WorkerQueueOption{
				Floor:   16,
				Ceiling: 200,
				Default: 125,
			},
			WorkerConcurrency: astro.WorkerQueueOption{
				Floor:   175,
				Ceiling: 275,
				Default: 180,
			},
		}
		mockClient.On("ListClusters", orgID).Return(clusters, nil)
		mockClient.On("ListDeployments", orgID, "").Return([]astro.Deployment{updatedDeployment}, nil).Once()
		mockClient.On("GetWorkerQueueOptions").Return(mockWorkerQueueDefaultOptions, nil).Once()
		mockClient.On("UpdateDeployment", mock.Anything).Return(updatedDeployment, nil)
		mockClient.On("ModifyDeploymentVariable", mock.Anything).Return([]astro.EnvironmentVariablesObject{}, nil)
		mockClient.On("UpdateAlertEmails", mock.Anything).Return(astro.DeploymentAlerts{}, nil)
		mockClient.On("ListDeployments", orgID, "").Return([]astro.Deployment{updatedDeployment}, nil)
		origClient := astroClient
		astroClient = mockClient
		fileutil.WriteStringToFile(filePath, data)
		defer func() {
			astroClient = origClient
			afero.NewOsFs().Remove(filePath)
		}()
		cmdArgs := []string{"update", "--deployment-file", "test-deployment.yaml"}
		_, err := execDeploymentCmd(cmdArgs...)
		s.NoError(err)
		mockClient.AssertExpectations(s.T())
	})
	s.Run("returns an error if updating a deployment from file fails", func() {
		cmdArgs := []string{"update", "--deployment-file", "test-file-name.json"}
		_, err := execDeploymentCmd(cmdArgs...)
		s.ErrorContains(err, "open test-file-name.json: no such file or directory")
	})
	s.Run("returns an error if from-file is specified with any other flags", func() {
		cmdArgs := []string{"update", "--deployment-file", "test-deployment.yaml", "--description", "fail"}
		_, err := execDeploymentCmd(cmdArgs...)
		s.ErrorIs(err, errFlag)
	})
	s.Run("updates a deployment with small scheduler size", func() {
		ctx, err := context.GetCurrentContext()
		s.NoError(err)
		ctx.SetContextKey("organization_product", "HOSTED")
		ctx.SetContextKey("organization", "test-org-id")
		ctx.SetContextKey("workspace", ws)
		deploymentResp := astro.Deployment{
			ID:             "test-id",
			Label:          "test-name",
			RuntimeRelease: astro.RuntimeRelease{Version: "4.2.5"},
			DeploymentSpec: astro.DeploymentSpec{Executor: "CeleryExecutor", Scheduler: astro.Scheduler{AU: 5, Replicas: 3}},
		}

		mockClient.On("GetDeploymentConfig").Return(astro.DeploymentConfig{
			Components: astro.Components{
				Scheduler: astro.SchedulerConfig{
					AU: astro.AuConfig{
						Default: 5,
						Limit:   24,
					},
					Replicas: astro.ReplicasConfig{
						Default: 1,
						Minimum: 1,
						Limit:   4,
					},
				},
			},
			RuntimeReleases: []astro.RuntimeRelease{
				{
					Version: "4.2.5",
				},
			},
		}, nil).Once()
		mockClient.On("ListDeployments", mock.Anything, ws).Return([]astro.Deployment{deploymentResp}, nil).Once()
		cmdArgs := []string{"update", "test-id", "--name", "test-name", "--workspace-id", ws, "--scheduler-size", "small", "--force"}
		_, err = execDeploymentCmd(cmdArgs...)
		s.NoError(err)
		mockClient.AssertExpectations(s.T())
	})
	s.Run("returns an error with incorrect high-availability value", func() {
		ctx, err := context.GetCurrentContext()
		s.NoError(err)
		ctx.SetContextKey("organization_product", "HOSTED")
		ctx.SetContextKey("organization", "test-org-id")
		ctx.SetContextKey("workspace", ws)
		cmdArgs := []string{"update", "test-id", "--name", "test-name", "--workspace-id", ws, "--high-availability", "some-value", "--force"}
		_, err = execDeploymentCmd(cmdArgs...)
		s.ErrorContains(err, "Invalid --high-availability value")
	})
	mockClient.AssertExpectations(s.T())
}

func (s *Suite) TestDeploymentDelete() {
	testUtil.InitTestConfig(testUtil.CloudPlatform)

	deploymentResp := astro.Deployment{
		ID:             "test-id",
		RuntimeRelease: astro.RuntimeRelease{Version: "4.2.5"},
		DeploymentSpec: astro.DeploymentSpec{Scheduler: astro.Scheduler{AU: 5, Replicas: 3}},
	}

	mockClient := new(astro_mocks.Client)
	mockClient.On("ListDeployments", mock.Anything, mock.Anything).Return([]astro.Deployment{deploymentResp}, nil).Once()
	mockClient.On("DeleteDeployment", mock.Anything).Return(astro.Deployment{ID: "test-id"}, nil).Once()
	astroClient = mockClient

	cmdArgs := []string{"delete", "test-id", "--force"}
	_, err := execDeploymentCmd(cmdArgs...)
	s.NoError(err)
	mockClient.AssertExpectations(s.T())
}

func (s *Suite) TestDeploymentVariableList() {
	testUtil.InitTestConfig(testUtil.CloudPlatform)

	mockResponse := []astro.Deployment{
		{
			ID: "test-id-1",
			DeploymentSpec: astro.DeploymentSpec{
				EnvironmentVariablesObjects: []astro.EnvironmentVariablesObject{{Key: "test-key-1", Value: "test-value-1"}},
			},
		},
		{
			ID: "test-id-2",
			DeploymentSpec: astro.DeploymentSpec{
				EnvironmentVariablesObjects: []astro.EnvironmentVariablesObject{{Key: "test-key-2", Value: "test-value-2"}},
			},
		},
	}

	mockClient := new(astro_mocks.Client)
	mockClient.On("ListDeployments", mock.Anything, mock.Anything).Return(mockResponse, nil).Once()
	astroClient = mockClient

	cmdArgs := []string{"variable", "list", "--deployment-id", "test-id-1"}
	resp, err := execDeploymentCmd(cmdArgs...)
	s.NoError(err)
	s.Contains(resp, "test-key-1")
	s.Contains(resp, "test-value-1")
	mockClient.AssertExpectations(s.T())
}

func (s *Suite) TestDeploymentVariableModify() {
	testUtil.InitTestConfig(testUtil.CloudPlatform)

	mockListResponse := []astro.Deployment{
		{
			ID: "test-id-1",
			DeploymentSpec: astro.DeploymentSpec{
				EnvironmentVariablesObjects: []astro.EnvironmentVariablesObject{{Key: "test-key-1", Value: "test-value-1"}},
			},
		},
		{
			ID: "test-id-2",
			DeploymentSpec: astro.DeploymentSpec{
				EnvironmentVariablesObjects: []astro.EnvironmentVariablesObject{},
			},
		},
	}

	mockCreateResponse := []astro.EnvironmentVariablesObject{
		{
			Key:   "test-key-1",
			Value: "test-value-1",
		},
		{
			Key:   "test-key-2",
			Value: "test-value-2",
		},
		{
			Key:   "test-key-3",
			Value: "test-value-3",
		},
	}

	mockClient := new(astro_mocks.Client)
	mockClient.On("ListDeployments", mock.Anything, mock.Anything).Return(mockListResponse, nil).Once()
	mockClient.On("ModifyDeploymentVariable", mock.Anything).Return(mockCreateResponse, nil).Once()
	astroClient = mockClient

	cmdArgs := []string{"variable", "create", "test-key-3=test-value-3", "--deployment-id", "test-id-1", "--key", "test-key-2", "--value", "test-value-2"}
	resp, err := execDeploymentCmd(cmdArgs...)
	s.NoError(err)
	s.Contains(resp, "test-key-1")
	s.Contains(resp, "test-value-1")
	s.Contains(resp, "test-key-2")
	s.Contains(resp, "test-value-2")
	s.Contains(resp, "test-key-3")
	s.Contains(resp, "test-value-3")
	mockClient.AssertExpectations(s.T())
}

func (s *Suite) TestDeploymentVariableUpdate() {
	testUtil.InitTestConfig(testUtil.CloudPlatform)

	mockListResponse := []astro.Deployment{
		{
			ID: "test-id-1",
			DeploymentSpec: astro.DeploymentSpec{
				EnvironmentVariablesObjects: []astro.EnvironmentVariablesObject{{Key: "test-key-1", Value: "test-value-1"}},
			},
		},
		{
			ID: "test-id-2",
			DeploymentSpec: astro.DeploymentSpec{
				EnvironmentVariablesObjects: []astro.EnvironmentVariablesObject{},
			},
		},
	}

	mockUpdateResponse := []astro.EnvironmentVariablesObject{
		{
			Key:   "test-key-1",
			Value: "test-value-update",
		},
		{
			Key:   "test-key-2",
			Value: "test-value-2-update",
		},
	}

	mockClient := new(astro_mocks.Client)
	mockClient.On("ListDeployments", mock.Anything, mock.Anything).Return(mockListResponse, nil).Once()
	mockClient.On("ModifyDeploymentVariable", mock.Anything).Return(mockUpdateResponse, nil).Once()
	astroClient = mockClient

	cmdArgs := []string{"variable", "update", "test-key-2=test-value-2-update", "--deployment-id", "test-id-1", "--key", "test-key-1", "--value", "test-value-update"}
	resp, err := execDeploymentCmd(cmdArgs...)
	s.NoError(err)
	s.Contains(resp, "test-key-1")
	s.Contains(resp, "test-value-update")
	s.Contains(resp, "test-key-2")
	s.Contains(resp, "test-value-2-update")
	mockClient.AssertExpectations(s.T())
}

func (s *Suite) TestIsValidExecutor() {
	s.Run("returns true for Kubernetes Executor", func() {
		actual := isValidExecutor(deployment.KubeExecutor)
		s.True(actual)
	})
	s.Run("returns true for Celery Executor", func() {
		actual := isValidExecutor(deployment.CeleryExecutor)
		s.True(actual)
	})
	s.Run("returns true when no Executor is requested", func() {
		actual := isValidExecutor("")
		s.True(actual)
	})
	s.Run("returns false for any invalid executor", func() {
		actual := isValidExecutor("KubeExecutor")
		s.False(actual)
	})
}

func (s *Suite) TestIsValidCloudProvider() {
	s.Run("returns true if cloudProvider is GCP", func() {
		actual := isValidCloudProvider("gcp")
		s.True(actual)
	})
	s.Run("returns false if cloudProvider is not GCP", func() {
		actual := isValidCloudProvider("azure")
		s.False(actual)
	})
}

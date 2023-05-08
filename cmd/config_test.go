package cmd

import (
	testUtil "github.com/astronomer/astro-cli/pkg/testing"
)

func (s *Suite) TestConfigRootCommand() {
	testUtil.InitTestConfig(testUtil.LocalPlatform)
	output, err := executeCommand("config")
	s.NoError(err)
	s.Contains(output, "astro config")
}

func (s *Suite) TestConfigGetCommandSuccess() {
	testUtil.InitTestConfig(testUtil.LocalPlatform)
	_, err := executeCommand("config", "get", "project.name", "-g")
	s.NoError(err)
}

func (s *Suite) TestConfigGetCommandFailure() {
	testUtil.InitTestConfig(testUtil.LocalPlatform)
	_, err := executeCommand("config", "get", "-g", "test")
	s.Error(err)
	s.EqualError(err, errInvalidConfigPath.Error())

	_, err = executeCommand("config", "get", "test")
	s.Error(err)
	s.Contains(err.Error(), "You are attempting to get [setting-name] a project config outside of a project directory")
}

func (s *Suite) TestConfigSetCommandFailure() {
	testUtil.InitTestConfig(testUtil.LocalPlatform)
	_, err := executeCommand("config", "set", "test", "testing", "-g")
	s.Error(err)

	_, err = executeCommand("config", "set", "test", "-g")
	s.ErrorIs(err, errInvalidSetArgs)
}

func (s *Suite) TestConfigSetCommandSuccess() {
	testUtil.InitTestConfig(testUtil.LocalPlatform)
	_, err := executeCommand("config", "set", "-g", "project.name", "testing")
	s.NoError(err)
}

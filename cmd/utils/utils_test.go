package utils

import (
	"testing"

	"github.com/astronomer/astro-cli/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func TestCmdUtilsSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestEnsureProjectDir() {
	currentWorkingPath := config.WorkingPath
	fileName := config.ConfigFileNameWithExt
	dirName := config.ConfigDir
	defer func() {
		config.WorkingPath = currentWorkingPath
		config.ConfigFileNameWithExt = fileName
		config.ConfigDir = dirName
	}()
	// error case when file path is not resolvable
	config.WorkingPath = "./\000x"
	err := EnsureProjectDir(&cobra.Command{}, []string{})
	s.Error(err)
	s.Contains(err.Error(), "failed to verify that your working directory is an Astro project.\nTry running astro dev init to turn your working directory into an Astro project")

	// error case when no such file or dir
	config.WorkingPath = "./test"
	err = EnsureProjectDir(&cobra.Command{}, []string{})
	s.Error(err)
	s.Contains(err.Error(), "this is not an Astro project directory.\nChange to another directory or run astro dev init to turn your working directory into an Astro project")

	// success case
	config.WorkingPath = currentWorkingPath
	config.ConfigFileNameWithExt = "utils_test.go"
	config.ConfigDir = ""
	err = EnsureProjectDir(&cobra.Command{}, []string{})
	s.NoError(err)
}

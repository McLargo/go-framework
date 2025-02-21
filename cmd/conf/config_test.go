package conf_test

import (
	"os"
	"testing"

	"github.com/mclargo/go-framework/cmd/conf"
	"github.com/spf13/viper"

	"github.com/stretchr/testify/suite"
)

type testConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(testConfigSuite))
}

func (s *testConfigSuite) createEnvFile(cont []byte) {
	f, err := os.Create(".env")
	s.Require().NoError(err)

	_, err = f.Write(cont)
	s.Require().NoError(err)
}

func (s *testConfigSuite) TearDownTest() {
	f, _ := os.Stat(".env")
	if f != nil {
		err := os.Remove(".env")
		s.Require().NoError(err)
	}
	// Clear the environment variables
	os.Clearenv()
	// viper is a global variable, so we need to reset it after each test
	viper.Reset()
}

func (s *testConfigSuite) TestInitConfig() {
	// Arrange
	contEnv := []byte("FRAMEWORK_CONFIG_PATH=.")
	s.createEnvFile(contEnv)

	// Act
	cfg, err := conf.InitConfig()

	// Assert
	s.Require().NoError(err)
	s.Require().NotEmpty(cfg)
}

func (s *testConfigSuite) TestInitConfig_NoEnvFile() {
	// Act
	cfg, err := conf.InitConfig()

	// Assert
	s.Require().Error(err)
	s.Require().Empty(cfg)
}

func (s *testConfigSuite) TestInitConfig_ConfigNotFound() {
	// Arrange
	contEnv := []byte("FRAMEWORK_CONFIG_PATH=../missing")
	s.createEnvFile(contEnv)

	// Act
	cfg, err := conf.InitConfig()

	// Assert
	s.Require().Error(err)
	s.Require().Empty(cfg)
}

func (s *testConfigSuite) TestInitConfig_ConfigEmpty() {
	// Arrange
	contEnv := []byte("FRAMEWORK_CONFIG_PATH=../")
	s.createEnvFile(contEnv)

	_, err := os.Create("../config.yml")
	s.Require().NoError(err)

	// Act
	cfg, err := conf.InitConfig()

	// Assert
	s.Require().Error(err)
	s.Require().Empty(cfg)

	// Cleanup
	err = os.Remove("../config.yml")
	s.Require().NoError(err)
}

func (s *testConfigSuite) TestInitConfig_CannotConvertToStruct() {
	// Arrange
	contEnv := []byte("FRAMEWORK_CONFIG_PATH=../")
	s.createEnvFile(contEnv)

	f, err := os.Create("../config.yml")
	s.Require().NoError(err)

	contYaml := []byte(`app:
  verbose: "invalid"
`)
	_, err = f.Write(contYaml)
	s.Require().NoError(err)

	// Act
	cfg, err := conf.InitConfig()

	// Assert
	s.Require().Error(err)
	s.Require().Empty(cfg)

	// Cleanup
	err = os.Remove("../config.yml")
	s.Require().NoError(err)
}

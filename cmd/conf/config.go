package conf

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type Config struct {
	App AppConfig `mapstructure:"app"`
	Log LogConfig `mapstructure:"log"`
}

type AppConfig struct {
	Verbose *bool  `mapstructure:"verbose" validate:"required"`
	Debug   *bool  `mapstructure:"debug"   validate:"required"`
	Port    string `mapstructure:"port"    validate:"required"`
}

type LogConfig struct {
	Path     string `mapstructure:"path"     validate:"required"`
	Filename string `mapstructure:"filename" validate:"required"`
	Debug    *bool  `mapstructure:"debug"    validate:"required"`
}

func InitConfig() (*Config, error) {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("cannot load .env file: %w", err)
	}

	// Set the default values for the config
	viper.SetEnvPrefix("FRAMEWORK")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// Set and load the config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(viper.GetString("CONFIG_PATH"))

	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshall to config struct: %w", err)
	}

	err = cfg.Validate()
	if err != nil {
		return nil, fmt.Errorf("cannot validate config: %w", err)
	}

	return &cfg, nil
}

// Print prints the configuration.
func (cfg *Config) Print(log *zap.Logger) {
	log.Info("Using viper config:", zap.String("file", viper.ConfigFileUsed()))
	set := viper.AllSettings()

	bs, err := yaml.Marshal(set)
	if err != nil {
		log.Error("Error marshalling config to YAML", zap.Error(err))
		return
	}

	// Print the values as YAML
	log.Info("YAML config:")
	log.Info("\n\n" + string(bs) + "\n")
}

// Validate validates the configuration.
func (cfg *Config) Validate() error {
	v := validator.New()

	return v.Struct(cfg)
}

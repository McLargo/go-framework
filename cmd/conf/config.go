package conf

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Config struct {
	App AppConfig `mapstructure:"app"`
	Log LogConfig `mapstructure:"log"`
}

type AppConfig struct {
	Verbose *bool  `mapstructure:"verbose" validate:"required"`
	Debug   *bool  `mapstructure:"debug" validate:"required"`
	Port    string `mapstructure:"port" validate:"required"`
}

type LogConfig struct {
	Path string `mapstructure:"path" validate:"required"`
}

func InitConfig() (*Config, error) {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}

	// Set the default values for the config
	viper.SetEnvPrefix("FRAMEWORK")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// Set and load the config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./cmd/conf")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file")
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("Error unmarshalling config")
		return nil, err
	}

	err = cfg.Validate()
	if err != nil {
		fmt.Println("Error validating config")
		return nil, err
	}

	return &cfg, nil
}

// Print prints the configuration.
func (cfg *Config) Print() {
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	set := viper.AllSettings()
	bs, err := yaml.Marshal(set)
	if err != nil {
		fmt.Println("Error marshalling config to YAML:", err)
		return
	}

	// Print the values as YAML
	fmt.Println(string(bs))
}

// Validate validates the configuration.
func (cfg *Config) Validate() error {
	v := validator.New()
	return v.Struct(cfg)
}

// Watch watches the configuration file for changes.
func (cfg *Config) Watch() {
	// on config change
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.Unmarshal(&cfg)
		if err != nil {
			fmt.Println("unable to unmarshal config:", err)
		}

		err = cfg.Validate()
		if err != nil {
			fmt.Println("fatal error validating config file:", err)
		}
		if *cfg.App.Verbose {
			cfg.Print()
		}
	})
	viper.WatchConfig()
}

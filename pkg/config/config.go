// config/config.go
package config

import (
	"fmt"
	"strings"

	"github.com/mmcloud/solder/pkg/interfaces"
	"github.com/mmcloud/solder/pkg/logging"
	"github.com/mmcloud/solder/pkg/service"
	"github.com/spf13/viper"
)

// DefaultConfigPath is the default path to the configuration file.
const defaultConfigPath = "." // Make sure the path is correct

type ConfigPath string

func (c ConfigPath) String() string {
	return string(c)
}

func ProvideDefaultConfig() ConfigPath {
	return ConfigPath(defaultConfigPath)
}

// Config is the configuration for the config package.
type Config struct {
	LoggerConfig  logging.LoggingConfig `mapstructure:"logger"`
	ServiceConfig service.ServiceConfig `mapstructure:"service"`
}

// Config implements the IConfig interface.
var _ interfaces.Config = (*Config)(nil)

// GetLoggerConfig returns the logger configuration.
func (c *Config) GetLoggingConfig() interfaces.LoggingConfig {
	return &c.LoggerConfig
}

// GetServiceConfig returns the service configuration.
func (c *Config) GetServiceConfig() interfaces.ServiceConfig {
	return &c.ServiceConfig
}

// NewConfig creates a new Viper instance and loads the configuration from the given path.
func NewViper(configPath ConfigPath) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(configPath.String()) // use the configPath parameter
	v.AutomaticEnv()
	v.SetEnvPrefix("SD")
	v.SetEnvKeyReplacer(strings.NewReplacer(
		".", "_",
	))

	return v, nil
}

// LoadConfig unmarshals the configuration from Viper into the Config struct.
func LoadConfig(v *viper.Viper) (interfaces.Config, error) {
	var config Config

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err) // Wrap error for context
	}

	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err) // Wrap error for context
	}
	return &config, nil
}

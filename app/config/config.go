package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type PluginType string
type TriggerType string
type ResultType string

const (
	PreBuilt     PluginType = "prebuilt"
	ShellCommand PluginType = "command"
	ShellScript  PluginType = "script"

	OnLivenessFailure  TriggerType = "on_liveness_failure"
	OnReadinessFailure TriggerType = "on_readiness_failure"
	OnStartupFailure   TriggerType = "on_startup_failure"
	Periodic           TriggerType = "periodic"

	Http2xx        ResultType = "http_2xx"
	NotHttp2xx     ResultType = "not_http_2xx"
	SpecificString ResultType = "specific_string"
	ContainsString ResultType = "contains_string"
)

// Config defines the format of the config we expect
type Config struct {
	Plugins            []Plugin `yaml:"plugins"`
	PushReportEndpoint string   `yaml:"push_report_endpoint"` // this is where the dependency container pushes its report
}

type Plugin struct {
	Name           string     `yaml:"name"`
	Type           PluginType `yaml:"type"`
	DependsOn      []Plugin   `yaml:"depends_on"`
	Image          string     `yaml:"image,omitempty"`
	ShellScript    string     `yaml:"shell_script,omitempty"`
	ShellFile      string     `yaml:"shell_file,omitempty"`
	ExpectedResult ResultType `yaml:"expected_result,omitempty"`
	Trigger        Trigger    `yaml:"trigger"`
}

type Trigger struct {
	TriggerType     []TriggerType `yaml:"trigger_type"`
	TriggerInterval int           `yaml:"trigger_interval,omitempty"`
	TriggerTimeout  int           `yaml:"trigger_timeout,omitempty"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	replaced := os.ExpandEnv(string(file))
	cfg := &Config{}
	err = yaml.Unmarshal([]byte(replaced), cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

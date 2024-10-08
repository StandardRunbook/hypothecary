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
	EventDriven        TriggerType = "event_driven"

	Http2xx        ResultType = "http_2xx"
	NotHttp2xx     ResultType = "not_http_2xx"
	SpecificString ResultType = "specific_string"
	ContainsString ResultType = "contains_string"

	SkipToEnd = "skip-to-end"
)

// Config defines the format of the config we expect
type Config struct {
	Plugins            []Plugin `yaml:"plugins"`
	PushReportEndpoint string   `yaml:"push_report_endpoint"` // this is the endpoint of the log/metric storage
}

type Plugin struct {
	Name         string        `yaml:"name"`
	Type         PluginType    `yaml:"type"`
	Module       string        `yaml:"module,omitempty"`
	ShellScript  string        `yaml:"shell_script,omitempty"`
	ShellFile    string        `yaml:"shell_file,omitempty"`
	Args         []string      `yaml:"args,omitempty"`
	Conditionals []Conditional `yaml:"conditionals,omitempty"`
	Trigger      Trigger       `yaml:"trigger"`
}

type Conditional struct {
	ExpectedResult ResultType `yaml:"expected_result,omitempty"`
	ChildName      string     `yaml:"child_name,omitempty"`
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
	return LoadConfigFromEnv(string(file))
}

func LoadConfigFromEnv(config string) (*Config, error) {
	replaced := os.ExpandEnv(config)
	cfg := &Config{}
	err := yaml.Unmarshal([]byte(replaced), cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

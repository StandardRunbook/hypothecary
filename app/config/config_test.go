package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfigPrebuiltPlugin(t *testing.T) {
	t.Parallel()
	cfg, err := LoadConfigFromEnv(testPreBuiltPluginWithPeriodicTrigger)
	require.NoError(t, err)
	require.NotNil(t, cfg)

	expectedConfig := &Config{
		PushReportEndpoint: "http://example.com/report",
		Plugins: []Plugin{
			{
				Name: "AWS-STS",
				Type: PreBuilt,
				Conditionals: []Conditional{
					{
						ExpectedResult: Http2xx,
						ChildName:      SkipToEnd,
					},
				},
				Trigger: Trigger{
					TriggerType: []TriggerType{
						Periodic,
					},
					TriggerInterval: 5,
					TriggerTimeout:  60,
				},
			},
		},
	}
	require.Equal(t, *expectedConfig, *cfg)
}

const testPreBuiltPluginWithPeriodicTrigger = `
plugins:
  - name: AWS-STS
    type: prebuilt
    conditionals:
      - expected_result: http_2xx
        child_name: skip-to-end
    trigger:
      trigger_type:
        - periodic
      trigger_interval: 5
      trigger_timeout: 60
push_report_endpoint: "http://example.com/report"
`

func TestHealthCheckOnLiveness(t *testing.T) {
	t.Parallel()
	cfg, err := LoadConfigFromEnv(testHealthCheckOnLivenessFailure)
	require.NoError(t, err)
	require.NotNil(t, cfg)

	expectedConfig := &Config{
		PushReportEndpoint: "http://example.com/report",
		Plugins: []Plugin{
			{
				Name: "HealthCheck",
				Type: ShellCommand,
				Args: []string{"/usr/local/bin/check_health"},
				Conditionals: []Conditional{
					{
						ExpectedResult: Http2xx,
						ChildName:      SkipToEnd,
					},
				},
				Trigger: Trigger{
					TriggerType: []TriggerType{
						OnLivenessFailure,
					},
					TriggerInterval: 10,
					TriggerTimeout:  120,
				},
			},
		},
	}
	require.Equal(t, *expectedConfig, *cfg)
}

const testHealthCheckOnLivenessFailure = `
plugins:
  - name: HealthCheck
    type: command
    args:
      - "/usr/local/bin/check_health"
    conditionals:
      - expected_result: http_2xx
        child_name: skip-to-end
    trigger:
      trigger_type:
        - on_liveness_failure
      trigger_interval: 10
      trigger_timeout: 120
push_report_endpoint: "http://example.com/report"
`

func TestScript(t *testing.T) {
	t.Parallel()
	cfg, err := LoadConfigFromEnv(testScript)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	expectedConfig := &Config{
		PushReportEndpoint: "http://example.com/report",
		Plugins: []Plugin{
			{
				Name:        "DataBackup",
				Type:        ShellScript,
				ShellScript: "backup.sh",
				Conditionals: []Conditional{
					{
						ExpectedResult: Http2xx,
						ChildName:      "AWS-STS",
					},
				},
				Trigger: Trigger{
					TriggerType: []TriggerType{
						Periodic,
						OnStartupFailure,
					},
					TriggerInterval: 15,
					TriggerTimeout:  300,
				},
			},
		},
	}
	require.Equal(t, *expectedConfig, *cfg)
}

const testScript = `
plugins:
  - name: DataBackup
    type: script
    shell_script: "backup.sh"
    conditionals:
      - expected_result: http_2xx
        child_name: AWS-STS
    trigger:
      trigger_type:
        - periodic
        - on_startup_failure
      trigger_interval: 15
      trigger_timeout: 300
push_report_endpoint: "http://example.com/report"
`

func TestConditional(t *testing.T) {
	t.Parallel()
	cfg, err := LoadConfigFromEnv(testConditionals)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	expectedConfig := &Config{
		PushReportEndpoint: "http://example.com/report",
		Plugins: []Plugin{
			{
				Name: "MainCheck",
				Type: PreBuilt,
				Conditionals: []Conditional{
					{
						ExpectedResult: NotHttp2xx,
						ChildName:      "FallbackCheck",
					},
				},
				Trigger: Trigger{
					TriggerType: []TriggerType{
						Periodic,
					},
					TriggerInterval: 10,
					TriggerTimeout:  60,
				},
			},
			{
				Name: "FallbackCheck",
				Type: ShellCommand,
				Args: []string{"/usr/local/bin/fallback_check"},
				Conditionals: []Conditional{
					{
						ExpectedResult: NotHttp2xx,
						ChildName:      SkipToEnd,
					},
				},
				Trigger: Trigger{
					TriggerType: []TriggerType{
						EventDriven,
					},
					TriggerInterval: 20,
					TriggerTimeout:  120,
				},
			},
		},
	}
	require.Equal(t, *expectedConfig, *cfg)
}

const testConditionals = `
plugins:
  - name: MainCheck
    type: prebuilt
    trigger:
      trigger_type:
        - periodic
      trigger_interval: 10
      trigger_timeout: 60
    conditionals:
      - expected_result: not_http_2xx
        child_name: FallbackCheck
  - name: FallbackCheck
    type: command
    args:
      - "/usr/local/bin/fallback_check"
    conditionals:
      - expected_result: not_http_2xx
        child_name: skip-to-end
    trigger:
      trigger_type:
        - event_driven
      trigger_interval: 20
      trigger_timeout: 120
push_report_endpoint: "http://example.com/report"
`

func TestSecurityScanning(t *testing.T) {
	t.Parallel()
	cfg, err := LoadConfigFromEnv(testSecurityScanning)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	expectedConfig := &Config{
		PushReportEndpoint: "http://example.com/report",
		Plugins: []Plugin{
			{
				Name:  "SecurityScan",
				Type:  ShellCommand,
				Image: "security_scan:latest",
				Args:  []string{"--scan-all"},
				Trigger: Trigger{
					TriggerType: []TriggerType{
						Periodic,
					},
					TriggerInterval: 60,
					TriggerTimeout:  600,
				},
			},
			{
				Name:      "StartupCheck",
				Type:      ShellScript,
				ShellFile: "/scripts/startup_check.sh",
				Trigger: Trigger{
					TriggerType: []TriggerType{
						OnStartupFailure,
					},
					TriggerInterval: 5,
					TriggerTimeout:  30,
				},
			},
			{
				Name: "PeriodicHealthCheck",
				Type: PreBuilt,
				Trigger: Trigger{
					TriggerType: []TriggerType{
						Periodic,
						OnLivenessFailure,
					},
					TriggerInterval: 10,
					TriggerTimeout:  120,
				},
			},
		},
	}
	require.Equal(t, expectedConfig, cfg)
}

const testSecurityScanning = `
plugins:
  - name: SecurityScan
    type: command
    image: "security_scan:latest"
    args:
      - "--scan-all"
    trigger:
      trigger_type:
        - periodic
      trigger_interval: 60
      trigger_timeout: 600
  - name: StartupCheck
    type: script
    shell_file: "/scripts/startup_check.sh"
    trigger:
      trigger_type:
        - on_startup_failure
      trigger_interval: 5
      trigger_timeout: 30
  - name: PeriodicHealthCheck
    type: prebuilt
    trigger:
      trigger_type:
        - periodic
        - on_liveness_failure
      trigger_interval: 10
      trigger_timeout: 120
push_report_endpoint: "http://example.com/report"
`

package settings

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
)

const SettingsLoggingEndpoint = "api/v1/settings/logging"

type LoggingSettings struct {
	SysLog                  SyslogSpec `json:"syslog,omitempty"`
	StdOut                  StdOutSpec `json:"stdout,omitempty"`
	EnableMetricsCollection bool       `json:"enableMetricsCollection,omitempty"`
	IncludeRuntimeLink      bool       `json:"includeRuntimeLink,omitempty"`
	ConsoleAddress          string     `json:"consoleAddress,omitempty"`
}

type SyslogSpec struct {
	Enabled       bool   `json:"enabled,omitempty"`
	VerboseScan   bool   `json:"verboseScan,omitempty"`
	AllProcEvents bool   `json:"allProcEvents,omitempty"`
	Address       string `json:"addr,omitempty"`
	ID            string `json:"id,omitempty"`
}

type StdOutSpec struct {
	Enabled       bool `json:"enabled,omitempty"`
	VerboseScan   bool `json:"verboseScan,omitempty"`
	AllProcEvents bool `json:"allProcEvents,omitempty"`
}

// Get the current logging settings.
func GetLoggingSettings(c api.Client) (LoggingSettings, error) {
	var ans LoggingSettings
	if err := c.Request(http.MethodGet, SettingsLoggingEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting logging settings: %s", err)
	}
	return ans, nil
}

// Update the current logging settings.
func UpdateLoggingSettings(c api.Client, settings LoggingSettings) error {
	return c.Request(http.MethodPost, SettingsLoggingEndpoint, nil, settings, nil)
}

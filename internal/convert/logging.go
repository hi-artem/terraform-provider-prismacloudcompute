package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/settings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaToLoggingSettings(d *schema.ResourceData) (settings.LoggingSettings, error) {
	loggingSettings := settings.LoggingSettings{}
	loggingSettings.ConsoleAddress = d.Get("console_address").(string)
	loggingSettings.IncludeRuntimeLink = d.Get("include_runtime_link").(bool)
	loggingSettings.EnableMetricsCollection = d.Get("enable_metrics_collection").(bool)
	sysLogElements := d.Get("syslog").([]interface{})
	loggingSettings.SysLog = schemaToSysLogSpec(sysLogElements[0].(map[string]interface{}))
	stdOutLogElements := d.Get("stdout").([]interface{})
	loggingSettings.StdOut = schemaToStdOutSpec(stdOutLogElements[0].(map[string]interface{}))

	return loggingSettings, nil
}

func schemaToSysLogSpec(d map[string]interface{}) settings.SyslogSpec {
	logSpecToSchema := settings.SyslogSpec{}
	logSpecToSchema.Address = d["address"].(string)
	logSpecToSchema.AllProcEvents = d["all_proc_events"].(bool)
	logSpecToSchema.Enabled = d["enabled"].(bool)
	logSpecToSchema.VerboseScan = d["verbose_scan"].(bool)
	logSpecToSchema.ID = d["identifier"].(string)

	return logSpecToSchema
}

func schemaToStdOutSpec(d map[string]interface{}) settings.StdOutSpec {
	logSpecToSchema := settings.StdOutSpec{}
	logSpecToSchema.AllProcEvents = d["all_proc_events"].(bool)
	logSpecToSchema.Enabled = d["enabled"].(bool)
	logSpecToSchema.VerboseScan = d["verbose_scan"].(bool)

	return logSpecToSchema
}

func SysLogSpecToSchema(in settings.SyslogSpec) []interface{} {
	m := make(map[string]interface{})
	m["enabled"] = in.Enabled
	m["verbose_scan"] = in.VerboseScan
	m["all_proc_events"] = in.AllProcEvents
	m["address"] = in.Address
	m["identifier"] = in.ID

	s := make([]interface{}, 1)
	s[0] = m
	return s
}

func StdOutSpecToSchema(in settings.StdOutSpec) []interface{} {
	m := make(map[string]interface{})
	m["enabled"] = in.Enabled
	m["verbose_scan"] = in.VerboseScan
	m["all_proc_events"] = in.AllProcEvents

	s := make([]interface{}, 1)
	s[0] = m
	return s
}

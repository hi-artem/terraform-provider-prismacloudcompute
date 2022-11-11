package provider

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/settings"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingSettings() *schema.Resource {
	return &schema.Resource{
		Create: createLoggingSettings,
		Read:   readLoggingSettings,
		Update: updateLoggingSettings,
		Delete: deleteLoggingSettings,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the logging settings.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"enable_metrics_collection": {
				Description: "Enable prometheus instrumentation.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"include_runtime_link": {
				Description: "Include link to runtime events.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"console_address": {
				Description: "Prisma Cloud Compute console url.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"syslog": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Configuration for the syslog daemons of the underlying hosts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Enable syslog logging.",
						},
						"verbose_scan": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Detailed output for vulnerabilities and compliance.",
						},
						"all_proc_events": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Detailed output of all process activity (not recommended).",
						},
						"address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Send syslog messages to a network endpoint.",
							Default:     "",
						},
						"identifier": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Custom identifier to all syslog messages.",
							Default:     "",
						},
					},
				},
			},
			"stdout": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Configuration for the stdout logging.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Enable syslog logging.",
						},
						"verbose_scan": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Detailed output for vulnerabilities and compliance.",
						},
						"all_proc_events": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Detailed output of all process activity (not recommended).",
						},
					},
				},
			},
		},
	}
}

func createLoggingSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedLoggingSettings, err := convert.SchemaToLoggingSettings(d)

	if err != nil {
		return fmt.Errorf("error creating logging settings feed '%+v': %s", parsedLoggingSettings, err)
	}

	if err := settings.UpdateLoggingSettings(*client, parsedLoggingSettings); err != nil {
		return fmt.Errorf("error creating %s logging settings feed: %s", settingsTypeLogging, err)
	}

	d.SetId(settingsTypeLogging)
	return readLoggingSettings(d, meta)
}

func readLoggingSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	retrievedLoggingSettings, err := settings.GetLoggingSettings(*client)
	if err != nil {
		return fmt.Errorf("error reading %s logging settings: %s", settingsTypeLogging, err)
	}

	d.Set("enable_metrics_collection", retrievedLoggingSettings.EnableMetricsCollection)
	d.Set("include_runtime_link", retrievedLoggingSettings.IncludeRuntimeLink)
	d.Set("console_address", retrievedLoggingSettings.ConsoleAddress)

	if err := d.Set("stdout", convert.StdOutSpecToSchema(retrievedLoggingSettings.StdOut)); err != nil {
		return fmt.Errorf("error reading %s logging settings: %s", settingsTypeLogging, err)
	}

	if err := d.Set("syslog", convert.SysLogSpecToSchema(retrievedLoggingSettings.SysLog)); err != nil {
		return fmt.Errorf("error reading %s logging settings: %s", settingsTypeLogging, err)
	}

	return nil
}

func updateLoggingSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedLoggingSettings, err := convert.SchemaToLoggingSettings(d)
	if err != nil {
		return fmt.Errorf("error updating logging settings feed '%+v': %s", parsedLoggingSettings, err)
	}

	if err := settings.UpdateLoggingSettings(*client, parsedLoggingSettings); err != nil {
		return fmt.Errorf("error creating %s logging settings feed: %s", settingsTypeLogging, err)
	}

	d.SetId(settingsTypeLogging)
	return readLoggingSettings(d, meta)
}

func deleteLoggingSettings(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default policy
	return nil
}

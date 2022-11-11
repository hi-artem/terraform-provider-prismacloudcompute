package provider

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/settings"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
)

type stdOutSpec struct {
	enabled         bool
	verbose_scan    bool
	all_proc_events bool
}

type sysLogSpec struct {
	enabled         bool
	verbose_scan    bool
	all_proc_events bool
	address         string
	identifier      string
}

func TestFlattenStdOut(t *testing.T) {
	cases := []struct {
		stdOut   settings.StdOutSpec
		expected stdOutSpec
	}{
		{
			stdOut: settings.StdOutSpec{
				Enabled:       true,
				VerboseScan:   true,
				AllProcEvents: true,
			},
			expected: stdOutSpec{
				enabled:         true,
				verbose_scan:    true,
				all_proc_events: true,
			},
		},
	}

	for _, c := range cases {
		out := convert.StdOutSpecToSchema(c.stdOut)
		stdOut := out[0].(map[string]interface{})

		if !reflect.DeepEqual(stdOut["enabled"], c.expected.enabled) {
			t.Fatalf("Error matching output and expected for stdout enabled: %#v vs %#v", out, c.expected)
		}

		if !reflect.DeepEqual(stdOut["verbose_scan"], c.expected.verbose_scan) {
			t.Fatalf("Error matching output and expected for stdout verbose_scan: %#v vs %#v", out, c.expected)
		}

		if !reflect.DeepEqual(stdOut["all_proc_events"], c.expected.all_proc_events) {
			t.Fatalf("Error matching output and expected for stdout all_proc_events: %#v vs %#v", out, c.expected)
		}
	}
}

func TestFlattenSysLog(t *testing.T) {
	cases := []struct {
		sysLog   settings.SyslogSpec
		expected sysLogSpec
	}{
		{
			sysLog: settings.SyslogSpec{
				Enabled:       true,
				VerboseScan:   true,
				AllProcEvents: true,
				ID:            "testing",
				Address:       "http://example.com",
			},
			expected: sysLogSpec{
				enabled:         true,
				verbose_scan:    true,
				all_proc_events: true,
				identifier:      "testing",
				address:         "http://example.com",
			},
		},
	}

	for _, c := range cases {
		out := convert.SysLogSpecToSchema(c.sysLog)
		sysLog := out[0].(map[string]interface{})

		if !reflect.DeepEqual(sysLog["enabled"], c.expected.enabled) {
			t.Fatalf("Error matching output and expected for stdout enabled: %#v vs %#v", out, c.expected)
		}

		if !reflect.DeepEqual(sysLog["verbose_scan"], c.expected.verbose_scan) {
			t.Fatalf("Error matching output and expected for stdout verbose_scan: %#v vs %#v", out, c.expected)
		}

		if !reflect.DeepEqual(sysLog["all_proc_events"], c.expected.all_proc_events) {
			t.Fatalf("Error matching output and expected for stdout all_proc_events: %#v vs %#v", out, c.expected)
		}

		if !reflect.DeepEqual(sysLog["address"], c.expected.address) {
			t.Fatalf("Error matching output and expected for stdout address: %#v vs %#v", out, c.expected)
		}

		if !reflect.DeepEqual(sysLog["identifier"], c.expected.identifier) {
			t.Fatalf("Error matching output and expected for stdout identifier: %#v vs %#v", out, c.expected)
		}
	}
}

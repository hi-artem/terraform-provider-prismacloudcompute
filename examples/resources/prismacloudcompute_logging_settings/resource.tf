resource "prismacloudcompute_logging_settings" "enable_all" {
  include_runtime_link = true
  enable_metrics_collection = true
  stdout {
    enabled = true
    verbose_scan = true
    all_proc_events = true
  }
  syslog {
    enabled = true
    verbose_scan = true
    all_proc_events = true
    address = "https://api.datadoghq.com"
    identifier = "prisma-syslog"
  }
}

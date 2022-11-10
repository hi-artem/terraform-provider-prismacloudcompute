resource "prismacloudcompute_custom_rule" "test" {
  name        = "basic-rule"
  description = "this is basic rule"
  message     = "%proc.name doing stuff"
  type        = "processes"
  script      = "proc.name = \"cat\""
}

resource "prismacloudcompute_custom_rule" "test_heredoc" {
  name        = "less-basic-rule"
  description = "this is less basic rule"
  message     = "%proc.name wrote to path"
  type        = "filesystem"
  script      = <<EOT
                  // Example:
                  // user modifies a sensitive file under /etc or its subfolders
                  // proc.user != "root" and file.path startswith "/etc"

                  proc.user != "crond" and file.path startswith "/var/spool"
                EOT
}

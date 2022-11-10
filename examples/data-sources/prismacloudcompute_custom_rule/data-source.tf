# Some default rules coming with the Console
data "prismacloudcompute_custom_rule" "rule" {
  name = "Twistlock Labs - Suspicious networking tool"
}

data "prismacloudcompute_custom_rule" "another_rule" {
  name = "Twistlock Labs - Running privileged process within container"
}

data "prismacloudcompute_custom_rule" "yet_another_rule" {
  name = "Twistlock Labs - Running cron app"
}

# Add container policy using these custom rules
resource "prismacloudcompute_container_runtime_policy" "ruleset" {
  learning_disabled = false

  rule {
    advanced_protection        = true
    cloud_metadata_enforcement = true
    collections = [
      "All",
    ]
    disabled               = false
    kubernetes_enforcement = true
    name                   = "Demo runtime container policy"
    wildfire_analysis      = "block"

    custom_rule {
      action = "audit"
      effect = "block"
      id     = data.prismacloudcompute_custom_rule.rule.prisma_id
    }
    custom_rule {
      action = "audit"
      effect = "block"
      id     = data.prismacloudcompute_custom_rule.another_rule.prisma_id
    }

    custom_rule {
      action = "audit"
      effect = "block"
      id     = data.prismacloudcompute_custom_rule.yet_another_rule.prisma_id
    }

    dns {
      allowed = [
        "amplitutude.com",
      ]
      denied = [
        "ru.com",
        "cn.com",
        "ir.com",
      ]
      deny_effect = "block"
    }

    filesystem {
      allowed = [
        "/etc",
        "/usr/bin/",
        "/var/app",
      ]
      backdoor_files          = true
      check_new_files         = true
      denied                  = []
      deny_effect             = "prevent"
      skip_encrypted_binaries = false
      suspicious_elf_headers  = true
    }

    network {
      allowed_outbound_ips = []
      denied_outbound_ips     = []
      deny_effect             = "alert"
      detect_port_scan        = true
      skip_modified_processes = false
      skip_raw_sockets        = false

      allowed_listening_port {
        deny  = false
        end   = 443
        start = 443
      }

      allowed_outbound_port {
        deny  = false
        end   = 80
        start = 80
      }
      allowed_outbound_port {
        deny  = false
        end   = 443
        start = 443
      }
    }

    processes {
      allowed = [
        "aws-cni",
      ]
      check_crypto_miners    = true
      check_lateral_movement = true
      check_parent_child     = false
      check_suid_binaries    = false
      denied                 = []
      deny_effect            = "block"
      skip_modified          = false
      skip_reverse_shell     = false
    }
  }
}

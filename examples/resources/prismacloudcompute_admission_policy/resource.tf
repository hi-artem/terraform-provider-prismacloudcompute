resource "prismacloudcompute_admission_policy" "ruleset" {
  rule {
    name        = "Block host PID and IPC sharing"
    disabled    = true
    effect      = "block"
    description = "This rule is important, but disabled"
    script      = <<-EOT
            match[{"msg": msg}] {
            	input.request.operation == "CREATE"
            	input.request.kind.kind == "Pod"
            	input.request.resource.resource == "pods"
            	input_share_hostnamespace(input.request.object)
            	msg := sprintf("Sharing the host namespace is not allowed, pod: %v", [input.request.object.metadata.name])
            }

            input_share_hostnamespace(o) {
                o.spec.hostPID
            }

            input_share_hostnamespace(o) {
                o.spec.hostIPC
            }
        EOT
  }
  rule {
    name        = "Allow containers with non read only filesystem"
    disabled    = false
    effect      = "allow"
    description = "This rule is important and enabled, but allowed"
    script      = <<-EOT
            match[{"msg": msg}] {
                operations := {"CREATE"}
                operations[input.request.operation]
                input.request.kind.kind == "Pod"

                containers := input.request.object.spec.containers[_]

                not containers.securityContext.readOnlyRootFilesystem
                msg := sprintf("container '%v' does not have a read only root filesystem", [containers.name])
            }
        EOT
  }
}

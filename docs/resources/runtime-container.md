---
page_title: "Prisma Cloud: prismacloudcompute_policiesruntimecontainer"
---

# prismacloudcompute_policiesruntimecontainer

Manage cloud compute runtime container policies.

## Example Usage

```hcl
resource "prismacloudcompute_policiesruntimecontainer" "example2" {
    learningdisabled = true
    rules {
        name = "my-rule"
	collections = [{
		name = "All"
	}]
        processes = {
            effect = "alert"
        }
        network = {
            effect = "alert"
        }
        dns = {
            effect = "alert"
        }
        filesystem = {
            effect = "alert"
        }
    }
}
```

## Argument Reference

* `filters` - Filter policy results.
* `_id` - ID of the policy set.
* `learningdisabled` - If set to `true`, automatic behavioural learning is disabled.
* [`rules`](#rules) - List of rules in the policies.

### Rules

* `name` - (Required) Name of the rule.
* [`collections`](#collections) - List of collections. Used to scope the rule.
* [`customrules`](#custom-rules) - List of custom rules.
* `disabled` - If set to `true`, the rule is currently disabled.
* [`dns`](#dns) - The DNS runtime rules.
* [`filesystem`](#filesystem) - Represents restrictions or suppression for filesystem changes.
* `kubernetesenforcement` - Detects containers that attempt to compromise the orchestrator.
* `modified` - Date/time when the rule was last modified.
* [`network`](#network) - Represents the restrictions and suppression for networking.
* `notes` - A free-form text description of the collection.
* `owner` - User who created or last modified the rule.
* `previousname` - Previous name of the rule. Required for rule renaming.
* [`processes`](#processes) - Represents restrictions or suppression for running processes.
* `wildfireanalysis` - The effect that will be used in the runtime rule. Can be set to `block`, `prevent`, `alert`, or `disable`.

#### Collections

* `accountids` - List of account IDs.
* `appids` - List of application IDs.
* `clusters` - List of Kubernetes cluster names.
* `coderepos` - List of code repositories.
* `color` - Hex color code for a collection.
* `containers` - List of containers.
* `description` - A free-form text description of the collection.
* `functions` - List of functions.
* `hosts` - List of hosts.
* `images` - List of images.
* `labels` - List of labels.
* `modified` - Date/time when the collection was last modified.
* `name` - Unique collection name.
* `namespaces` - List of Kubernetes namespaces.
* `owner` - User who created or last modified the collection.
* `prisma` - If set to `true`, this collection originates from Prisma Cloud.
* `system` - If set to `true`, this collection was created by the system (i.e., a non-user). Otherwise it was created by a real user.

#### Custom Rules

* `_id` - Custom rule ID.
* `action` - The action to perform if the custom rule applies. Can be set to `audit` or `incident`.
* `effect` - The effect to be used for the custom rule. Can be set to `block`, `prevent`, `alert`, `allow`, `ban`, or `disable`.

#### DNS

* `blacklist` - Deny-listed domain names (e.g., www.bad-url.com, *.bad-url.com).
* `effect` - The effect to be used in the runtime rule. Can be set to `block`, `prevent`, `alert`, or `disable`.
* `whitelist` - Allow-listed domain names (e.g., *.gmail.com, *.s3.amazon.com).

#### Filesystem

* `backdoorfiles` - If set to `true`, monitors files that can create or persist backdoors (SSH or admin account config files).
* `blacklist` - List of denied file system paths.
* `checknewfiles` - If set to `true`, Detects changes to binaries and certificates.
* `effect` - The effect that will be used in the runtime rule. Can be set to `block`, `prevent`, `alert`, or `disable`.
* `skipencryptedbinaries` - If set to `true`, the encrypted binaries check will be skipped.
* `suspiciouselfheaders` - If set to `true`, enables malware detection based on suspicious ELF headers.
* `whitelist` - List of allowed filesystem paths.

#### Network

* `blacklistips` - Deny-listed IP addresses.
* [`blacklistlisteningports`](#blacklist-listening-ports) - Deny-listed listening ports.
* [`blacklistoutboundports`](#blacklist-outbound-ports) - Deny-listed outbound ports.
* `detectportscan` - If set to `true`, port scanning detection is enabled.
* `effect` - Effect used in the runtime rule. Can be set to: `block`, `prevent`, `alert`, or `disable`.
* `skipmodifiedproc` - If set to `true`, Prisma Cloud can detect malicious networking activity from modified processes.
* `skiprawsockets` - If set to `true`, raw socket detection will be skipped.
* `whitelistips` - Allow-listed IP addresses.
* [`whitelistlisteningports`](#whitelist-listening-ports) - Allow-listed listening ports.
* [`whitelistoutboundports`](#whitelist-outbound-ports) - Allow-listed outbound ports.

##### Blacklist Listening Ports

* `deny` - If set to `true`, the connection is denied.
* `end` - End.
* `start` - Start.

##### Blacklist Outbound Ports

* `deny` - If set to `true`, the connection is denied.
* `end` - End.
* `start` - Start.

##### Whitelist Listening Ports

* `deny` - If set to `true`, the connection is denied.
* `end` - End.
* `start` - Start.

##### Whitelist Outbound Ports

* `deny` - If set to `true`, the connection is denied.
* `end` - End.
* `start` - Start.

#### Processes

* `blacklist` - List of processes to deny.
* `blockallbinaries` - If set to `true`, blocks all processes except for the main process.
* `checkcryptominers` - If set to `true`, detect crypto miners.
* `checklateralmovement` - If set to `true`, enables detection of processes that can be used for lateral movement exploits.
* `checknewbinaries` - If set to `true`, binaries which don't belong to the original image are allowed to run.
* `checkparentchild` - If set to `true`, enables check for parent-child relationship when comparing spawned processes in the model.
* `checksuidbinaries` - If set to `true`, enables check for process elevanting privileges (SUID bit).
* `effect` - The effect to be used in the runtime rule. Can be set to `block`, `prevent`, `alert`, or `disable`.
* `skipmodified` - If set to `true`, trigger audits/incidents when a modified proc is spawned.
* `skipreverseshell` - If set to `true`, reverse shell detection is disabled.
* `whitelist` - Allow-list of processes.
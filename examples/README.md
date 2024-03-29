# Examples

This directory contains examples that are mostly used for documentation, but can also be run/tested manually via the Terraform CLI.

The document generation tool looks for files in the following locations by default. All other *.tf files besides the ones mentioned below are ignored by the documentation tool. This is useful for creating examples that can run and/or are testable even if some parts are not relevant for the documentation.

* `provider/provider.tf` example file for the provider index page
* `data-sources/data_source_name/data-source.tf` example file for the named data source page
* `resources/resource_name/resource.tf` example file for the named data source page

The files in the `defaults/` directory are for Prisma Cloud Compute default configurations. This configuration is a subject to change, depending on the Compute version.

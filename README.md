# Terraform Provider for Prisma Cloud Compute

This is a community-driven fork of the official Prisma Cloud Compute provider with some extra sauce.

You can find the Prisma Cloud Compute provider in the [Terraform Registry](https://registry.terraform.io/providers/hi-artem/prismacloudcompute/latest).

## Basic setup
```terraform
terraform {
  required_providers {
    prismacloudcompute = {
      source  = "hi-artem/prismacloudcompute"
      version = "0.8.1"
    }
  }
}

provider "prismacloudcompute" {
  # Configure provider inline
  #
  console_url = "https://foo.bar.com"
  username = "myUsername"
  password = "myPassword"

  # Or you can use file
  #
  # config_file = "creds.json"
}
```
Complete documentation can be found in the [marketplace listing](https://registry.terraform.io/providers/hi-artem/prismacloudcompute/latest/docs).

## Contributing
Contributions are welcome!
Please read the [contributing guide](CONTRIBUTING.md) for more information.

## Support
Please read our [support document](SUPPORT.md) for details on how to get support for this project.

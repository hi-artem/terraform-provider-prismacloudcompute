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

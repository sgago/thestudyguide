
# Terraform configuration file for this project.
# Old versions of Terraform may not support all features we need!
# https://www.terraform.io/docs/language/settings/backends/s3.html
terraform {

  # Required providers are the plugins that Terraform uses to interact with cloud providers.
  # In this case, we are using the AWS provider.
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.8"
    }
  }

  # Backend specifies where Terraform stores its state
  # By default, this is a local file, but we want to use a remote backend.
  # This lets us share the state file with other team members.
  # In this case, we are using a remote S3 bucket.
  # This S3 bucket needs to be created first before running terraform.
  backend "s3" {
    bucket = "tsg-terraform-state-bucket" # The S3 bucket name
    key    = "terraform/state/tsg"            # The path to the state file in the bucket
    region = "us-east-1"                  # The region where the bucket is located
  }

  # Required Terraform version
  # This ensures that we are using a compatible version of Terraform.
  required_version = ">= 1.3.0"
}
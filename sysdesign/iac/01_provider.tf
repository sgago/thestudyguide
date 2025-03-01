
# AWS Provider Configuration
# We use this block to configure the AWS provider.
# https://www.terraform.io/docs/providers/aws/index.html
provider "aws" {
  region = "us-east-1"
  
  # OR
  # You can use these instead of environment variables
  #access_key = "<MY_AWS_ACCESS_KEY_ID>"
  #secret_key = "<AWS_SECRET_ACCESS_KEY>"
  
  # OR
  # You can use config and credentials files in ~/.aws/
  #shared_config_files = ["~/.aws/config"]
  #shared_credentials_file = "~/.aws/credentials"
  #profile = "my_profile" # If you have multiple profiles in your credentials file

  # OR
  # You can assume_role when using terraform
  #assume_role {
  #  role_arn = "arn:aws:iam::123456789012:role/role_name"
  #  session_name = "terraform" # A name for the session for logging
  #  external_id = "external_id" # A secure external ID for the role
  #}
}
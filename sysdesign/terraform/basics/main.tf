// The terraform block contains metadata for terraform.
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws" // Where to get the provider from, this is default Terraform registry.

      // The ~> arrow is a pessimistic constraint operator.
      // It constrains 5 >= version < 6, version 6 might break things!
      version = "~> 5.0" // Constrains the AWS provider version, previous versions might not work!
    }
  }

  required_version = ">= 1.2.0" // Constrains the Terraform version, previous versions might not work!
}

// Configures the provider, in this case AWS.
// It can also be GCP, Azure, etc. Multiple providers
// can be configured in the same file.
provider "aws" {
  region = "us-west-2"
}

// A resource defines a component of your infrastructure.
// These can be logical or physical components.
// The two strings before the block are the resource type and name.
// The prefix of the type, in this case "aws" in "aws_instance",
// maps to the aws provider above.
resource "aws_instance" "app_server" {
  // The following are arguments for the resource.
  // They can include VPC Ids, machine sizes, disk image names, etc.
  
  // AMI is Amazon Machine Image, this AMI Id is for Ubuntu 20.04
  ami           = "ami-830c94e3"
  instance_type = "t2.micro" // This instance type qualifies for free tier on AWS

  // Tags are key-value pairs that can be attached to resources.
  tags = {
    Name = "BasicAppServerInstance"
  }
}
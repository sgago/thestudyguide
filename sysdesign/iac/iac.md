# Terraform

## Why infrastructure as code anyway?
Imagine all the steps it takes for a human to setup a single http webserver. You need to install/setup the http server software, download dependencies, expose ports 80 and 443, install a certificate for secure communications, manage DNS settings, and setup monitoring. That's a lot of work.

Now, imagine, you need to setup 100 more servers. Don't make any mistakes! Also, later this evening, we need you to remove 30 of those servers to save money, because we get less web requests at night. And hopefully, nobody else messes with one of those 100 servers in the meantime and breaks something. Yikes. This is slow, error prone, and difficult to scale.

Infrastructure as code (IaC) via Terraform, OpenTofu, or Pulumi, come to the rescue. These tools will build and configure the same thing every single time. Using cloud APIs, IaC tools can create, destroy, and change resources quickly. Beyond this, our infrastructure is no longer tribal knowledge; the IaC files become our source of truth. As a bonus, we can save our infrastructure knowledge to version control with all the advantages that brings: history, code reviews, and so on. Pretty sweet, right?

This guide will focus on Terraform.

## State
Terraform doesn't _actually_ know what resources you've provisioned in the cloud. It stores all sorts of metadata about your provisioned services in this state file called `terraform.tfstate`. Without the state file, terraform doesn't know which services it should care about, which services depend on which, what unique identifiers are out there, and so on. The cloud providers themselves don't offer a consistent way to track all this themselves via tags or similar. Therefore, enter state files.

The state file is intended to be your actual source of truth of your infrastructure. However, this makes it like a cache, with all the pros and cons of such. For example, what happens if someone manually changes our infrastructure without using terraform? Now our state file and actual infrastructure will be in disagreement. Or what if we want to pull in some resources that were create manually? Fret not, terraform has tools to help us fix these issues.

## Remote State


## Lock file
The dependency lock file, naed `.terraform.lock.hcl` ensures that providers (dependencies) are the same each time. Without it, running `terraform init` could grab new provider versions. You, your team mates, and your CI/CD pipe could all be running different versions of providers. This would make it hard to find bugs and can lead to inconsistencies.

Upgrading dependencies should be intentional, not accidental. To upgrade providers, run `terraform init -upgrade` and then commit this to version control.
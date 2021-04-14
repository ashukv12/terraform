# Change Infrastructure

Here we will modify the resource, and learn how to apply changes to the Terraform projects.

## Prerequisites
- The Terraform CLI installed.
- The AWS CLI installed.
- Applied the configuration as done [here](https://github.com/ashukv12/terraform/tree/main/Build%20Infrastructure) 

## Steps to follow:

- Configuration

1. Update the ami of your instance. Change the `aws_instance.app_server` resource under the provider block in `main.tf`, by replacing the current AMI ID with a new one.

- Apply Changes

1. To apply the change to the existing resources, use:

`terraform apply`

2. Respond with a `yes` to execute the planned steps.

As indicated by the execution plan, Terraform first destroyed the existing instance and then created a new one in its place.

- Inspect state

1. Inspect the current state using:

    `terraform show`

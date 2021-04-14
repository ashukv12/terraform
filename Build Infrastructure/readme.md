# Build Infrastructure

## Prerequisites
- The Terraform CLI installed.
- The AWS CLI installed.
- An AWS account.
- Your AWS credentials. 

## Steps to follow:

- Configure the AWS CLI from your terminal

`aws configure`

- Write configuration

1. `mkdir learn-terraform-aws-instance`
2. `cd learn-terraform-aws-instance`
3. `touch main.tf`
4. Edit the  configuration in main.tf file

- Initialize the directory

1. `terraform init`

This will download the aws provider and installs it in a hidden subdirectory of your current working directory.

- Format and validate the configuration

1. `terraform fmt`
2. `terraform validate`

- Create infrastructure

1. `terraform apply`

Now, infrastructure using Terraform has been created. Just visit the EC2 console and find your new EC2 instance.

- Inspect state
1. Inspect the current state using:
2. 
  `terraform show`


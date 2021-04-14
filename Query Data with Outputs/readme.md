# Query Data with Outputs

Here we will use output values to organize data to be easily queried and to be shown back to the Terraform user.

## Prerequisites
- The Terraform CLI installed.
- The AWS CLI installed.
- Applied the configuration as done [here](https://github.com/ashukv12/terraform/tree/main/Build%20Infrastructure) 

## Steps to follow:

- Output EC2 instance configuration

  Create a file called outputs.tf in your learn-terraform-aws-instance directory. Add outputs to the new file for your EC2 instance's ID and IP address.
  
  ``` output "instance_id" {
        description = "ID of the EC2 instance"
        value       = aws_instance.app_server.id
      }

      output "instance_public_ip" {
        description = "Public IP address of the EC2 instance"
        value       = aws_instance.app_server.public_ip
      }

- Inspect output values

 1. To apply the change to the existing resources, use:

    `terraform apply`

 2. Query the outputs using this command:

    `terraform output`

  As indicated by the execution plan, Terraform first destroyed the existing instance and then created a new one in its place.

- Destroy infrastructure

    `terraform destroy`

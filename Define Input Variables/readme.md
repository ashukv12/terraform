# Define Input Variables

Here we will learn how to set variables via the command-line.

## Prerequisites
- The Terraform CLI installed.
- The AWS CLI installed.
- Applied the configuration as done [here](https://github.com/ashukv12/terraform/tree/main/Change%20Infrastructure) 

## Steps to follow:

- Create a new file called variables.tf with a block defining a new instance_name variable.

    ```variable "instance_name" {
    
    description = "Value of the Name tag for the EC2 instance"
    
    type        = string
    
    default     = "ExampleAppServerInstance"

  }

- In main.tf, update the aws_instance resource block to use the new variable.

  `Name = var.instance_name`

- Apply the configuration:

  `terraform apply`

  Now apply the configuration again. Terraform will update the instance with the new  Name tag with the value of the instance_type variable using the -var flag:

  `terraform apply -var 'instance_name=YetAnotherName'`

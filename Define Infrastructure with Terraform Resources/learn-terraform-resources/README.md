# Learn Terraform Resources

This repo is a companion repo to the [Define Infrastructure with Terraform Resources Learn guide](https://learn.hashicorp.com/tutorials/terraform/resource), containing Terraform configuration files to provision two publicly EC2 instances and an ELB.


## Prerequisites
- an AWS account
- a GitHub account


## Set up and explore your Terraform workspace
- git clone https://github.com/hashicorp/learn-terraform-resources.git
- cd learn-terraform-resources

## Explore the random pet name resource
`resource "random_pet" "name" {}`

## Explore the EC2 instance resource

```
resource "aws_instance" "web" {
  ami                    = "ami-a0cfeed8"
  instance_type          = "t2.micro"
  user_data              = file("init-script.sh")

  tags = {
    Name = random_pet.name.id
  }
}
```

## Initialize and apply Terraform configuration
- terraform init
- terraform apply

## Add security group to instance

## Clean up your infrastructure

## Head over to this to see the EC2 instance deploying the PHP application:

http://ec2-52-39-100-11.us-west-2.compute.amazonaws.com/index.php

# Learn Terraform Resources

This repo is a companion repo to the [Define Infrastructure with Terraform Resources Learn guide](https://learn.hashicorp.com/tutorials/terraform/resource), containing Terraform configuration files to provision two publicly EC2 instances and an ELB.


## Prerequisites
- an AWS account
- a GitHub account


## Set up and explore your Terraform workspace
- `git clone https://github.com/hashicorp/learn-terraform-resources.git`
- `cd learn-terraform-resources`

## Explore the random pet name resource

`resource "random_pet" "name" {}`

The random pet name resource block defines a `random_pet` resource named `name` to generate a random pet name. You will use the name that the `random_pet` resource generates to create a unique name for your EC2 instance.

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

The aws_instance.web resource block defines an aws_instance resource named web to create an AWS EC2 instance.

## Initialize and apply Terraform configuration
- Initialize the directory, using `terraform init`
- Apply your configuration, use `terraform apply` and respond `yes`.

## Add security group to instance

Open the AWS Provider documentation page. Search for security_group and select the aws_security_group resource.

- Add security group resource as given below:

```
resource "aws_security_group" "web-sg" {
  name = "${random_pet.name.id}-sg"
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
```
- Add the vpc_security_group_ids argument to the aws_instance.web resource as a list by placing the aws_security_group.web-sg.id attribute inside square brackets.

```
resource "aws_instance" "web" {
  ami                    = "ami-a0cfeed8"
  instance_type          = "t2.micro"
  user_data              = file("init-script.sh")
+ vpc_security_group_ids = [aws_security_group.web-sg.id]

  tags = {
    Name = random_pet.name.id
  }
}
```
- Apply your configuration. Remember to confirm your apply with a yes

`terraform apply`

- Verify that your EC2 instance is now publicly accessible.

`terraform output application-url`

## Clean up your infrastructure

## Head over to this to see the EC2 instance deploying the PHP application:

http://ec2-52-39-100-11.us-west-2.compute.amazonaws.com/index.php

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }
}

provider "aws" {
  profile = "default"
  region  = var.region1
}

resource "aws_instance" "example" {
  ami           = "ami-0742b4e673072066f"
  instance_type = "t2.micro"

  tags =  {
      Name = var.region2
  }
}

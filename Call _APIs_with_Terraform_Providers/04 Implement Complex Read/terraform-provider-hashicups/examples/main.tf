terraform {
  required_providers {
    hashicups = {
      version = "0.2"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {
  username = "education"
  password = "test123"
}

module "psl" {
  source = "./coffee"

  coffee_name = "Packer Spiced Latte"
}

data "hashicups_order" "order" {
  id = 1
}

output "psl" {
  value = module.psl.coffee
}

output "order" {
  value = data.hashicups_order.order
}

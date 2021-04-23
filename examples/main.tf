terraform {
  required_providers {
    zoom = {
      version = "0.2"
      source  = "hashicorp.com/edu/zoom"
    }
  }
}

provider "zoom" {
  
}

data "zoom_users" "user1"{

}

resource "zoom_orders" "user2" {
  first_name="abc"
  last_name="xyz"
  email="ashutosh99october@gmail.com"
  type=1
}

terraform {
  required_providers {
    zoom = {
      version = "2.0.0"
      source  = "zoom.us/edu/zoom"
    }
  }
}

provider "zoom" {
  
}

/*

resource "zoom_orders" "user2" {
  first_name="abc"
  last_name="xyz123"
  email="ashutoshkverma12@gmail.com"
  type=1
}


output "user123" {
  value=zoom_orders.user2 
}
*/


data "zoom_users" "user3" {
  email = "ekansh0786@gmail.com"
}


output "user3" {
  value = data.zoom_users.user3
}


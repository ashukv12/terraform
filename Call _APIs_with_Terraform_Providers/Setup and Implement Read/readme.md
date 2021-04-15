# Setup and Implement Read

## Prerequisites

- Golang 1.13+ installed and configured
- Terraform 0.14+ CLI 
- Docker and Docker Compose to run an instance of HashiCups locally

## Steps:

1. Set up your development environment

    - `git clone --branch boilerplate https://github.com/hashicorp/terraform-provider-hashicups`
    - `cd terraform-provider-hashicups`
    - `cd docker_compose && docker-compose up`
    - Verify that HashiCups is running by sending a request to its health check endpoint:
        
          `curl localhost:19090/health`

2. Build provider

    - `go mod init terraform-provider-hashicups`
    - `go mod vendor`
    - `make build`
    - To verify things are working correctly, execute:
    
        `./terraform-provider-hashicups`


3. Create new HashiCups user

    - `curl -X POST localhost:19090/signup -d '{"username":"education", "password":"test123"}'`
    - `curl -X POST localhost:19090/signin -d '{"username":"education", "password":"test123"}'`
    - `export HASHICUPS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTEwNzgwODUsInVzZXJfaWQiOjIsInVzZXJuYW1lIjoiZWR1Y2F0aW9uIn0.CguceCNILKdjOQ7Gx0u4UAMlOTaH3Dw-fsll2iXDrYU`

    Now that the HashiCups app is running, you're ready to interact with it using the Terraform provider.
  
4. Initialize workspace

    - Add the following to your main.tf file. This is required for Terraform 0.13+.

        ```terraform {
              required_providers {
                hashicups = {
                  version = "~> 0.3.1"
                  source  = "hashicorp.com/edu/hashicups"
                }
              }
            }

    - `terraform init`

5. Create order

    - Add the following to your main.tf file.

        This authenticate the HashiCups provider, create an order and return the order's values in your output. The order contains total of 4 coffees: 2 of each coffee_id 3 and 2.

        ```provider "hashicups" {
              username = "education"
              password = "test123"
            }

            resource "hashicups_order" "edu" {
              items {
                coffee {
                  id = 3
                }
                quantity = 2
              }
              items {
                coffee {
                  id = 2
                }
                quantity = 2
              }
            }

            output "edu_order" {
              value = hashicups_order.edu
            }
        
    - `terraform apply`
    - `terraform state show hashicups_order.edu`
    

Then, at last, verify if order is created.

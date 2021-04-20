# Implement Import

## Prerequisites

- Golang 1.13+ installed and configured
- Terraform 0.14+ CLI 
- Docker and Docker Compose to run an instance of HashiCups locally

## Steps:

1. Set up your development environment

    - `git clone --branch implement-create https://github.com/hashicorp/terraform-provider-hashicups`
    - `cd terraform-provider-hashicups`
    - `go mod vendor`

2. Build provider binary

    - `go build -o terraform-provider-hashicups`
    - `export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"`
    - `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`


3. Run HashiCups API

    - `cd docker_compose && docker-compose up`
    - Leave this terminal running. In another terminal, verify that HashiCups is running using:
        
        `curl localhost:19090/health`
    - `cd ..`
       
4. Create a HashiCups user

    - `curl -X POST localhost:19090/signin -d '{"username":"education", "password":"test123"}'`
    - `export HASHICUPS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTEwNzgwODUsInVzZXJfaWQiOjIsInVzZXJuYW1lIjoiZWR1Y2F0aW9uIn0.CguceCNILKdjOQ7Gx0u4UAMlOTaH3Dw-fsll2iXDrYU`
    
    
5. Apply example configuration

    - `cd examples`
    - `terraform init`
    - `terraform apply --auto-approve`
    -  `cd ..`
    
6. Implement import
        
    - In hashicups/resource_order.go, add the following Importer attribute to the end of the resourceOrder() schema:
      ```
      Importer: &schema.ResourceImporter{
        StateContext: schema.ImportStatePassthroughContext,
      },

  - `go fmt ./...`

   
8. Test the provider

    - `go build -o terraform-provider-hashicups`
    - `export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"`
    - `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `cd examples`
    - `mkdir import && cd import`
    - Create a new file named main.tf with the following configuration. You will import an existing HashiCups order into the hashicups_order.sample resource:
      ```
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

      resource "hashicups_order" "sample" {}

      output "sample_order" {
        value = hashicups_order.sample
      }

    - `terraform init`
    - `curl -X POST -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders -d '[{"coffee": { "id":1 }, "quantity":4}, {"coffee": { "id":3 }, "quantity":3}]'`
    - `terraform import hashicups_order.sample 3`
    - `terraform state show hashicups_order.sample`
    
Finally, we have added import capabilities to the order resource.

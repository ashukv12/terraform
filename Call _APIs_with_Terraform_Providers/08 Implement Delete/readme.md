# Implement Delete

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
    - `cd ..`
    
    
5. Apply example configuration

    - `cd examples`
    - `terraform init`
    - `terraform apply --auto-approve`
    -  `cd ..`
    
6. Implement delete
        
    - Replace the resourceOrderDelete function in hashicups/resource_order.go with the code snippet below. This function will delete the HashiCups order and Terraform resource:
      ```
      func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
        c := m.(*hc.Client)

        // Warning or errors can be collected in a slice type
        var diags diag.Diagnostics

        orderID := d.Id()

        err := c.DeleteOrder(orderID)
        if err != nil {
          return diag.FromErr(err)
        }

        // d.SetId("") is automatically called assuming delete returns no errors, but
        // it is added here for explicitness.
        d.SetId("")

        return diags
      }

  - `go fmt ./...`

   
8. Test the provider

    - `go build -o terraform-provider-hashicups`
    - `export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"`
    - `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `cd examples`
    - `terraform apply --auto-approve`
    - `curl -X GET -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders/1`

Finally, we have added update capabilities to the order resource.

# Debug a Terraform Provider

## Prerequisites

- Golang 1.13+ installed and configured
- Terraform 0.14+ CLI 
- Docker and Docker Compose to run an instance of HashiCups locally

## Steps:

1. Set up your development environment

    - `git clone --branch implement-complex-read  https://github.com/hashicorp/terraform-provider-hashicups`
    - `cd terraform-provider-hashicups`
    - `go mod vendor`

2. Build provider binary

    - `go build -o terraform-provider-hashicups`
    - `export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"`
    - `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `make build`
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`


3. Run HashiCups API

    - `cd docker_compose && docker-compose up`
    - Leave this terminal running. In another terminal, verify that HashiCups is running using:
        
        `curl localhost:19090/health`
    - `cd ..`
       
4. Update error messages

    - Replace the return nil, diag.FromError(err) line in your providerConfigure function with the following code snippet
      ```
      diags = append(diags, diag.Diagnostic{
        Severity: diag.Error,
        Summary:  "Unable to create HashiCups client",
        Detail:   "Unable to auth user for authenticated HashiCups client",
      })

      return nil, diags

      
    - `go fmt ./...`
    
5. Test error message

    - `pwd`
    - `go build -o terraform-provider-hashicups`
    - `export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"`
    - `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `make build`
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `cd examples`
    - Update your HashiCups provider configuration so it uses the wrong credentials:
        `password = "test1234"`
        
    - `terraform init && terraform apply --auto-approve`

6. Add warning message
    - Add the following snippet after diags is declared in providerConfigure function:
      ```
      diags = append(diags, diag.Diagnostic{
        Severity: diag.Warning,
        Summary:  "Warning Message Summary",
        Detail:   "This is the detailed warning message from providerConfigure",
      })

    - `go fmt ./...`
    
7. Test error message

    - `pwd`
    - `go build -o terraform-provider-hashicups`
    - `export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"`
    - `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `make build`
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `cd examples`
    -     
6. Implement complex read
        
    - Add the dataSourceOrderRead function to hashicups/data_source_order.go
        ```
        func dataSourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
          c := m.(*hc.Client)

          // Warning or errors can be collected in a slice type
          var diags diag.Diagnostics

          orderID := strconv.Itoa(d.Get("id").(int))

          order, err := c.GetOrder(orderID)
          if err != nil {
            return diag.FromErr(err)
          }

          orderItems := flattenOrderItemsData(&order.Items)
          if err := d.Set("items", orderItems); err != nil {
            return diag.FromErr(err)
          }

          d.SetId(orderID)

          return diags
        }

    - `go fmt ./...`
    - Add the flattenOrderItemsData function to your hashicups/data_source_order.go file:
        ```
        func flattenOrderItemsData(orderItems *[]hc.OrderItem) []interface{} {
          if orderItems != nil {
            ois := make([]interface{}, len(*orderItems), len(*orderItems))

            for i, orderItem := range *orderItems {
              oi := make(map[string]interface{})

              oi["coffee_id"] = orderItem.Coffee.ID
              oi["coffee_name"] = orderItem.Coffee.Name
              oi["coffee_teaser"] = orderItem.Coffee.Teaser
              oi["coffee_description"] = orderItem.Coffee.Description
              oi["coffee_price"] = orderItem.Coffee.Price
              oi["coffee_image"] = orderItem.Coffee.Image
              oi["quantity"] = orderItem.Quantity

              ois[i] = oi
            }

            return ois
          }

          return make([]interface{}, 0)
        }

    - `go fmt ./...`
    - In your hashicups/provider.go file, add the order data source to the DataSourcesMap of your Provider() function
        `"hashicups_order":       dataSourceOrder(),`
  

2. Build provider binary

    - `pwd`
    - `go build -o terraform-provider-hashicups`
    - `export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"`
    - `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `make build`
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `cd examples`
    - Add the following Terraform configuration to main.tf:
        ```
        data "hashicups_order" "order" {
          id = 1
        }

        output "order" {
          value = data.hashicups_order.order
        }

        
    - `terraform init && terraform apply --auto-approve`

    
The provider should have invoked a request to the signin endpoint.   

Finally, we have implemented a nested read function. This will be useful when we will create a resource using the HashiCups provider in the Implement Create tutorial.

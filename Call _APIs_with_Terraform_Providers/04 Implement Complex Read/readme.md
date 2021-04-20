# Implement Complex Read

## Prerequisites

- Golang 1.13+ installed and configured
- Terraform 0.14+ CLI 
- Docker and Docker Compose to run an instance of HashiCups locally

## Steps:

1. Set up your development environment

    - `git clone --branch auth-configuration  https://github.com/hashicorp/terraform-provider-hashicups`
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
       
4. Create an order via API

    - `curl -X POST -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders -d '[{"coffee": { "id":1 }, "quantity":4}, {"coffee": { "id":3 }, "quantity":3}]''`
    - `curl -X GET -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders/1`
    
    
5. Define order data resource

    - Now, create a file named hashicups/data_source_order.go in your hashicups directory and add the following snippet:
    
        ```package hashicups

            import (
              "context"
              "strconv"

              hc "github.com/hashicorp-demoapp/hashicups-client-go"
              "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
              "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
            )

            func dataSourceOrder() *schema.Resource {
              return &schema.Resource{
                ReadContext: dataSourceOrderRead,
                Schema: map[string]*schema.Schema{
                  "id": &schema.Schema{
                    Type:     schema.TypeInt,
                    Required: true,
                  },
                  "items": &schema.Schema{
                    Type:     schema.TypeList,
                    Computed: true,
                    Elem: &schema.Resource{
                      Schema: map[string]*schema.Schema{
                        "coffee_id": &schema.Schema{
                          Type:     schema.TypeInt,
                          Computed: true,
                        },
                        "coffee_name": &schema.Schema{
                          Type:     schema.TypeString,
                          Computed: true,
                        },
                        "coffee_teaser": &schema.Schema{
                          Type:     schema.TypeString,
                          Computed: true,
                        },
                        "coffee_description": &schema.Schema{
                          Type:     schema.TypeString,
                          Computed: true,
                        },
                        "coffee_price": &schema.Schema{
                          Type:     schema.TypeInt,
                          Computed: true,
                        },
                        "coffee_image": &schema.Schema{
                          Type:     schema.TypeString,
                          Computed: true,
                        },
                        "quantity": &schema.Schema{
                          Type:     schema.TypeInt,
                          Computed: true,
                        },
                      },
                    },
                  },
                },
              }
            }


    
    - `go fmt ./...`
    
    
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
        
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `cd examples`
    - Set HASHICUPS_USERNAME and HASHICUPS_PASSWORD to education and test123 respectively.
        - $ `export HASHICUPS_USERNAME=education`
        - $ `export HASHICUPS_PASSWORD=test123`
        
    - `terraform init && terraform apply --auto-approve`
   
Check the terminal containing your HashiCups logs for the recorded operations invoked by the HashiCups provider.

    api_1  | 2020-12-10T09:26:23.349Z [INFO]  Handle User | signin
    api_1  | 2020-12-10T09:26:23.357Z [INFO]  Handle Coffee
    api_1  | 2020-12-10T09:26:23.488Z [INFO]  Handle User | signin
    api_1  | 2020-12-10T09:26:23.606Z [INFO]  Handle User | signin
    
The provider should have invoked a request to the signin endpoint.   

Finally, we have added authentication to our HashiCups provider.

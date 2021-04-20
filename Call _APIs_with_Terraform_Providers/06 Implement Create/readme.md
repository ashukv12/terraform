# Implement Create

## Prerequisites

- Golang 1.13+ installed and configured
- Terraform 0.14+ CLI 
- Docker and Docker Compose to run an instance of HashiCups locally

## Steps:

1. Set up your development environment

    - `git clone --branch debug-tf-provider https://github.com/hashicorp/terraform-provider-hashicups`
    - `cd terraform-provider-hashicups`
    - `go mod vendor`

2. Build provider binary

    - `go build -o terraform-provider-hashicups`
    - `export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"`
    - `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`


3. Create a HashiCups user

    - `cd docker_compose && docker-compose up`
    - Leave this terminal running. In another terminal, verify that HashiCups is running using:
        
        `curl localhost:19090/health`
    - `cd ..`
       
4. Create an order via API

    - `curl -X POST localhost:19090/signin -d '{"username":"education", "password":"test123"}'`
    - `export HASHICUPS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTEwNzgwODUsInVzZXJfaWQiOjIsInVzZXJuYW1lIjoiZWR1Y2F0aW9uIn0.CguceCNILKdjOQ7Gx0u4UAMlOTaH3Dw-fsll2iXDrYU`
    - `cd ..`
    
    
5. Define order data resource

    - Add the following code snippet to a new file named hashicups/resource_order.go in the hashicups directory:
    
        ```
        package hashicups

        import (
          "context"

          "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
          "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
        )

        func resourceOrder() *schema.Resource {
          return &schema.Resource{
            CreateContext: resourceOrderCreate,
            ReadContext:   resourceOrderRead,
            UpdateContext: resourceOrderUpdate,
            DeleteContext: resourceOrderDelete,
            Schema: map[string]*schema.Schema{},
          }
        }

        func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
          // Warning or errors can be collected in a slice type
          var diags diag.Diagnostics

          return diags
        }

        func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
          // Warning or errors can be collected in a slice type
          var diags diag.Diagnostics

          return diags
        }

        func resourceOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
          return resourceOrderRead(ctx, d, m)
        }

        func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
          // Warning or errors can be collected in a slice type
          var diags diag.Diagnostics

          return diags
        }

    
    - `go fmt ./...`
    
    
6. Define order schema
        
    - `curl -X POST -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders -d '[{"coffee": { "id":1 }, "quantity":4}, {"coffee": { "id":3 }, "quantity":3}]'`
    - Replace the line Schema: map[string]*schema.Schema{}, in your resourceOrder function with the following schema:
        ```
        Schema: map[string]*schema.Schema{
          "items": &schema.Schema{
            Type:     schema.TypeList,
            Required: true,
            Elem: &schema.Resource{
              Schema: map[string]*schema.Schema{
                "coffee": &schema.Schema{
                  Type:     schema.TypeList,
                  MaxItems: 1,
                  Required: true,
                  Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                      "id": &schema.Schema{
                        Type:     schema.TypeInt,
                        Required: true,
                      },
                      "name": &schema.Schema{
                        Type:     schema.TypeString,
                        Computed: true,
                      },
                      "teaser": &schema.Schema{
                        Type:     schema.TypeString,
                        Computed: true,
                      },
                      "description": &schema.Schema{
                        Type:     schema.TypeString,
                        Computed: true,
                      },
                      "price": &schema.Schema{
                        Type:     schema.TypeInt,
                        Computed: true,
                      },
                      "image": &schema.Schema{
                        Type:     schema.TypeString,
                        Computed: true,
                      },
                    },
                  },
                },
                "quantity": &schema.Schema{
                  Type:     schema.TypeInt,
                  Required: true,
                },
              },
            },
          },
        },


    - `go fmt ./...`

7. Implement create

    - Now that you have defined the order resource schema, replace the resourceOrderCreate function in hashicups/resource_order.go with the following code snippet:
      ```
      func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
        c := m.(*hc.Client)

        // Warning or errors can be collected in a slice type
        var diags diag.Diagnostics

        items := d.Get("items").([]interface{})
        ois := []hc.OrderItem{}

        for _, item := range items {
          i := item.(map[string]interface{})

          co := i["coffee"].([]interface{})[0]
          coffee := co.(map[string]interface{})

          oi := hc.OrderItem{
            Coffee: hc.Coffee{
              ID: coffee["id"].(int),
            },
            Quantity: i["quantity"].(int),
          }

          ois = append(ois, oi)
        }

        o, err := c.CreateOrder(ois)
        if err != nil {
          return diag.FromErr(err)
        }

        d.SetId(strconv.Itoa(o.ID))

        return diags
      }

    - `"strconv"
    - `hc "github.com/hashicorp-demoapp/hashicups-client-go"`
    - `go fmt ./...`

8. Implement read

    - 
        ```
        func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
          c := m.(*hc.Client)

          // Warning or errors can be collected in a slice type
          var diags diag.Diagnostics

          orderID := d.Id()

          order, err := c.GetOrder(orderID)
          if err != nil {
            return diag.FromErr(err)
          }

          orderItems := flattenOrderItems(&order.Items)
          if err := d.Set("items", orderItems); err != nil {
            return diag.FromErr(err)
          }

          return diags
        }
        
    -`go fmt ./...`
    
    
9. Add flattening functions

    - 
        ```
        func flattenOrderItems(orderItems *[]hc.OrderItem) []interface{} {
          if orderItems != nil {
            ois := make([]interface{}, len(*orderItems), len(*orderItems))

            for i, orderItem := range *orderItems {
              oi := make(map[string]interface{})

              oi["coffee"] = flattenCoffee(orderItem.Coffee)
              oi["quantity"] = orderItem.Quantity
              ois[i] = oi
            }

            return ois
          }

          return make([]interface{}, 0)
        }
        
    - 
        ```
        func flattenCoffee(coffee hc.Coffee) []interface{} {
          c := make(map[string]interface{})
          c["id"] = coffee.ID
          c["name"] = coffee.Name
          c["teaser"] = coffee.Teaser
          c["description"] = coffee.Description
          c["price"] = coffee.Price
          c["image"] = coffee.Image

          return []interface{}{c}
        }
      
    -`go fmt ./...`
    
10. Add read function to create function

    - `resourceOrderRead(ctx, d, m)`
    - `go fmt ./...`

11. Add order resource to provider

    - `"hashicups_order": resourceOrder(),`
    - `go fmt ./...`

12. Test the provider

    - `go build -o terraform-provider-hashicups`
    - `export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"`
    - `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `cd examples`
    - 
        ```
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
        
    - `terraform init && terraform apply --auto-approve` 
     

Finally, we have created the order resource with create and read capabilities.


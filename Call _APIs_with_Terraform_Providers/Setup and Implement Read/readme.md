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


3. Define coffees schema

    - `curl localhost:19090/coffees`
    - Update your coffees data source's schema with the following code snippet: 
        ```
        Schema: map[string]*schema.Schema{
          "coffees": &schema.Schema{
            Type:     schema.TypeList,
            Computed: true,
            Elem: &schema.Resource{
              Schema: map[string]*schema.Schema{
                "id": &schema.Schema{
                  Type:     schema.TypeInt,
                  Computed: true,
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
                "ingredients": &schema.Schema{
                  Type:     schema.TypeList,
                  Computed: true,
                  Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                      "ingredient_id": &schema.Schema{
                        Type:     schema.TypeInt,
                        Computed: true,
                      },
                    },
                  },
                },
              },
            },
          },
        },

        
    - `go fmt ./...`
    
    Note: The coffees schema is a schema.TypeList of coffee (schema.Resource).
       
4. Implement read

    - Add the following read function to your hashicups/data_source_coffee.go file.

        ```func dataSourceCoffeesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
              client := &http.Client{Timeout: 10 * time.Second}

              // Warning or errors can be collected in a slice type
              var diags diag.Diagnostics

              req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees", "http://localhost:19090"), nil)
              if err != nil {
                return diag.FromErr(err)
              }

              r, err := client.Do(req)
              if err != nil {
                return diag.FromErr(err)
              }
              defer r.Body.Close()

              coffees := make([]map[string]interface{}, 0)
              err = json.NewDecoder(r.Body).Decode(&coffees)
              if err != nil {
                return diag.FromErr(err)
              }

              if err := d.Set("coffees", coffees); err != nil {
                return diag.FromErr(err)
              }

              // always run
              d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

              return diags
            }


    - `go fmt ./...`

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

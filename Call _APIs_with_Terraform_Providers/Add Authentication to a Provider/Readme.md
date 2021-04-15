# Add Authentication to a Provider

## Prerequisites

- Golang 1.13+ installed and configured
- Terraform 0.14+ CLI 
- Docker and Docker Compose to run an instance of HashiCups locally

## Steps:

1. Set up your development environment

    - `git clone --branch implement-read https://github.com/hashicorp/terraform-provider-hashicups`
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
    - `curl localhost:19090/health`
    - `cd ..`
       
4. Update provider schema

    - `curl -X POST localhost:19090/signup -d '{"username":"education", "password":"test123"}'`
    - `export HASHICUPS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTEwNzgwODUsInVzZXJfaWQiOjIsInVzZXJuYW1lIjoiZWR1Y2F0aW9uIn0.CguceCNILKdjOQ7Gx0u4UAMlOTaH3Dw-fsll2iXDrYU`
    - In your hashicups/provider.go file, replace the Provider() function with the code snippet below.:

        ```// Provider -
          func Provider() *schema.Provider {
            return &schema.Provider{
              Schema: map[string]*schema.Schema{
                "username": &schema.Schema{
                  Type:        schema.TypeString,
                  Optional:    true,
                  DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_USERNAME", nil),
                },
                "password": &schema.Schema{
                  Type:        schema.TypeString,
                  Optional:    true,
                  Sensitive:   true,
                  DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_PASSWORD", nil),
                },
              },
              ResourcesMap: map[string]*schema.Resource{},
              DataSourcesMap: map[string]*schema.Resource{
                "hashicups_coffees":     dataSourceCoffees(),
              },
              ConfigureContextFunc: providerConfigure,
            }
          }


    - `go fmt ./...`

5. Define providerConfigure

    - Import the context, API client and diag libraries into the provider.go file. The providerConfigure function will use these libraries:
    
        ```"context"
           "github.com/hashicorp-demoapp/hashicups-client-go"
           "github.com/hashicorp/terraform-plugin-sdk/v2/diag"DataSourcesMap: map[string]*schema.Resource{
                "hashicups_coffees":     dataSourceCoffees(),
           },
    
    - Add the providerConfigure function below your Provider() function. This function retrieves the username and password from the provider schema to authenticate and configure your provider:
    
        ```func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
              username := d.Get("username").(string)
              password := d.Get("password").(string)

              // Warning or errors can be collected in a slice type
              var diags diag.Diagnostics

              if (username != "") && (password != "") {
                c, err := hashicups.NewClient(nil, &username, &password)
                if err != nil {
                  return nil, diag.FromErr(err)
                }

                return c, diags
              }

              c, err := hashicups.NewClient(nil, nil, nil)
              if err != nil {
                return nil, diag.FromErr(err)
              }

              return c, diags
            }

    
    - `go fmt ./...`
    - Save your hashicups/provider.go file, then run go mod vendor to download the API client library into your /vendor directory.
         `go mod vendor`
    
6. Test the provider
        
    - `pwd`
    - `go build -o terraform-provider-hashicups`
    - `export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"`
    - ` mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `cd examples`
    - Set HASHICUPS_USERNAME and HASHICUPS_PASSWORD to education and test123 respectively.
        - $ export HASHICUPS_USERNAME=education
        - $ export HASHICUPS_PASSWORD=test123
    - `terraform init && terraform apply --auto-approve`
   
   
Now, a Terraform provider and data resource to reference information from an API in the Terraform configuration has been created.    

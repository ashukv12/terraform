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

5. Add data source to provider

    - In your hashicups/provider.go file, add the coffees data source to the DataSourcesMap:
    
        ```DataSourcesMap: map[string]*schema.Resource{
                "hashicups_coffees":     dataSourceCoffees(),
           },
        
    - `go fmt ./...`
    - `terraform state show hashicups_order.edu`
    
5. »Test the provider

    - In your hashicups/provider.go file, add the coffees data source to the DataSourcesMap:
    
        ```DataSourcesMap: map[string]*schema.Resource{
                "hashicups_coffees":     dataSourceCoffees(),
           },
        
    - `pwd`
    - `go build -o terraform-provider-hashicups`
    - `export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"`
    - ` mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `cd examples`
    - `terraform init && terraform apply --auto-approve`
   
   
Now, a Terraform provider and data resource to reference information from an API in the Terraform configuration has been created.    

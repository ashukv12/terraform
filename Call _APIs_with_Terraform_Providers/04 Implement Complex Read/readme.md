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

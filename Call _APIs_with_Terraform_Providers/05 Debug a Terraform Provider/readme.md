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
    - `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/$OS_ARCH`
    - `cd examples`
    -   `terraform init && terraform apply --auto-approve`
      
8. Fix provider

    - Comment out or remove the warning message, then recompile your Terraform provider:
        ```
        - diags = append(diags, diag.Diagnostic{
        -   Severity: diag.Warning,
        -   Summary:  "Warning Message Summary",
        -   Detail:   "This is the detailed warning message from providerConfigure",
        - })
        
    - Fix your Terraform configuration to use the correct credentials.
        `password = "test123"`
    - terraform init && terraform apply --auto-approve
    
Finally, we have added error and warning messages to our HashiCups provider and implemented a nested read function.

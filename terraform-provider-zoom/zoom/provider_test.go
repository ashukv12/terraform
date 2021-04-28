package zoom
import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    // "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
    "os"
    "testing"
)
var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
func init() {
    testAccProvider = Provider()
    testAccProviders = map[string]*schema.Provider{
        "zoom": testAccProvider,
    }
}
func TestProvider(t *testing.T) {
    if err := Provider().InternalValidate(); err != nil {
        t.Fatalf("err: %s", err)
    }
}
func testAccPreCheck(t *testing.T) {
    os.Setenv("AUTHORIZATION_TOKEN", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTk0NDMxMjgsImlhdCI6MTYxODgzODM0MH0.YG6Qr5Ce12uPRCG396zKl7myb4Co9cVmo8uokjD0NUA")
    if v := os.Getenv("AUTHORIZATION_TOKEN"); v == "" {
        t.Fatal("jwt must be authenticated")
    }
}
package zoom
import (
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
    "os"
    "testing"
	"time"
	"net/http"
	"io/ioutil"
)

func TestAccItem_Basic(t *testing.T) {
    os.Setenv("TF_ACC", "1")
    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        CheckDestroy: testAccCheckItemDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccCheckItemBasic(),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckExampleItemExists("zoom_orders.test_orders"),
					resource.TestCheckResourceAttr(
                        "zoom_orders.test_orders", "email", "ashutoshkverma12@gmail.com"),
                    resource.TestCheckResourceAttr(
                        "zoom_orders.test_orders", "type", "1"),
                    resource.TestCheckResourceAttr(
                        "zoom_orders.test_orders", "last_name", "xyz"),
                    resource.TestCheckResourceAttr(
                        "zoom_orders.test_orders", "first_name", "abc"),
                ),
				//ExpectNonEmptyPlan: true,
            },
        },
    })
}

func testAccCheckItemBasic() string {
    return fmt.Sprintf(`
        resource "zoom_orders" "test_orders" {
            first_name="abc"
            last_name="xyz"
            email="ashutoshkverma12@gmail.com"
            type=1
        }
    `)
}

// func testAccCheckItemDestroy(s *terraform.State) error {
//     // apiClient := testAccProvider.Meta().(*client.Client)
//     for _, rs := range s.RootModule().Resources {
//         if rs.Type != "example_item" {
//             continue
//         }       
//     }
//     return nil
// }

func testAccCheckItemDestroy(s *terraform.State) error{
    auth := testAccProvider.Meta().(tok)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "zoom_orders" {
            continue
        }
        time.Sleep(10 * time.Second)
        url := fmt.Sprintf("https://api.zoom.us/v2/users/%v?action=%v", rs.Primary.ID, "delete")
        req, err := http.NewRequest("DELETE", url, nil)
        if err != nil {
            return err
        }
        req.Header.Add("authorization", auth.token)
        res, err := http.DefaultClient.Do(req)
        if err != nil {
            return err
        }
        defer res.Body.Close()
        
		_, err = ioutil.ReadAll(res.Body)
        if err != nil {
            return err
        }
        // if res.StatusCode!=404{
        //     return fmt.Errorf("During Test\nStatus Code - %v\n%v", res.StatusCode, string(body))
        // }
    }
    return nil
}

func TestAccItem_Update(t *testing.T) {
	os.Setenv("TF_ACC", "1")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExampleItemExists("zoom_orders.test_orders"),
					resource.TestCheckResourceAttr(
						"zoom_orders.test_orders", "email", "ashutoshkverma12@gmail.com"),
					resource.TestCheckResourceAttr(
						"zoom_orders.test_orders", "type", "1"),
					resource.TestCheckResourceAttr(
						"zoom_orders.test_orders", "last_name", "xyz"),
					resource.TestCheckResourceAttr(
						"zoom_orders.test_orders", "first_name", "abc"),
				),
				//ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckItemUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExampleItemExists("zoom_orders.test_orders"),
					resource.TestCheckResourceAttr(
						"zoom_orders.test_orders", "email", "ashutoshkverma12@gmail.com"),
					resource.TestCheckResourceAttr(
						"zoom_orders.test_orders", "type", "1"),
					resource.TestCheckResourceAttr(
						"zoom_orders.test_orders", "last_name", "xyz123"),
					resource.TestCheckResourceAttr(
						"zoom_orders.test_orders", "first_name", "abc"),
				),
				//ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckItemUpdatePre() string {
	return fmt.Sprintf(`
	resource "zoom_orders" "test_orders" {
		email = "ashutoshkverma12@gmail.com"
		first_name = "abc"
		last_name = "xyz"
		type = 1
	}`)
}

func testAccCheckItemUpdatePost() string {
	return fmt.Sprintf(`
	resource "zoom_orders" "test_orders" { 
		email = "ashutoshkverma12@gmail.com"
		first_name = "abc"
		last_name = "xyz123"
		type = 1
	}`)
}

func testAccCheckExampleItemExists(resource string) resource.TestCheckFunc {
    return func(state *terraform.State) error {
        rs, ok := state.RootModule().Resources[resource]
        if !ok {
            return fmt.Errorf("Not found: %s", resource)
        }
        if rs.Primary.ID == "" {
            return fmt.Errorf("No Record ID is set")
        }
        return nil
        // name := rs.Primary.ID
    }
}


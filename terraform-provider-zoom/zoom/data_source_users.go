package zoom

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	//"os"
	//"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Users struct {
	Id         string `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}

type whole_body struct {
	Users []Users `json:"users"`
}

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUsersRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceUsersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//var bearer = os.Getenv("bearer")

	email := d.Get("email").(string)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s", "https://api.zoom.us/v2", email), nil)
	
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTk0NDMxMjgsImlhdCI6MTYxODgzODM0MH0.YG6Qr5Ce12uPRCG396zKl7myb4Co9cVmo8uokjD0NUA")
	
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)

	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	users := Users{}
	err = json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("email", users.Email)
	d.Set("first_name", users.First_name)
	d.Set("last_name", users.Last_name)
	
	//decode the response into a []interface{}
	// decode_response := make([]interface{}, len(Users))

	// //sets the response body to Terraform users data source,
	// //assigning each value to its respective schema position.
	// if err := d.Set("users", decode_response); err != nil {
	// 	return diag.FromErr(err)
	// }
	// // always run, to set the resource ID.
	d.SetId(users.Id)

	return diags
}

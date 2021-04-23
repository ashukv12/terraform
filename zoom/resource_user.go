package zoom

import (
	"context"
	//"fmt"
	"encoding/json"
	"net/http"
	"strings"
	//"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Users_res struct {
	Type int `json:"type"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}

type whole_body_res struct {
	Action string `json:"action"`
	Users_List Users_res `json:"users"`
}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrderCreate,
		ReadContext:   resourceOrderRead,
		UpdateContext: resourceOrderUpdate,
		DeleteContext: resourceOrderDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	email := d.Get("email").(string)
	first := d.Get("first_name").(string)
	last := d.Get("last_name").(string)
	Type := d.Get("type").(int)

	values := whole_body_res{
		Action: "create",
		Users_List: Users_res{
			Email:      email,
			Type:       Type,
			First_name: first,
			Last_name:  last,
		},
	}

	url := "https://api.zoom.us/v2/users"

	payload, _ := json.Marshal(values)

	//var bearer = os.Getenv("bearer")

	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payload)))

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTk0NDMxMjgsImlhdCI6MTYxODgzODM0MH0.YG6Qr5Ce12uPRCG396zKl7myb4Co9cVmo8uokjD0NUA")

	res, _ := http.DefaultClient.Do(req)

	d.SetId(values.Users_List.Email)

	defer res.Body.Close()

	resourceOrderRead(ctx, d, m)

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

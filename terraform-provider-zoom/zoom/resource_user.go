package zoom

import (
	"context"
	//"fmt"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	// "strconv"
	"os"
	//"io/ioutil"
	"fmt"

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
	Users_List Users_res `json:"user_info"`
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
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAyMTgzMTksImlhdCI6MTYxOTYxMzUyMX0.oU-kb6pSAvFU_SJ-KM7gnNeknRh3dM5eLq1TMSp9vuM")

	res, _ := http.DefaultClient.Do(req)

	d.SetId(values.Users_List.Email)

	defer res.Body.Close()

	//resourceOrderRead(ctx, d, m)

	fmt.Println("User created")

	return diags
}

func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	
	client := &http.Client{Timeout: 10 * time.Second}

	UserID := d.Id()
	url := "https://api.zoom.us/v2/users/" + UserID

	req, _ := http.NewRequest("GET", url, nil)
	jwt := os.Getenv("bearer")
	req.Header.Add("authorization", jwt)

	_, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	Email:=      d.Get("email").(string)
	First_name:= d.Get("first_name").(string)
	Last_name:=  d.Get("last_name").(string)
	Type:=       d.Get("type").(int)
	
	if err:= d.Set("email", Email); err!= nil{
		return diag.FromErr(err)
	}
	if err:= d.Set("first_name", First_name); err != nil{
		return diag.FromErr(err)
	}
	if err:= d.Set("last_name", Last_name); err != nil{
		return diag.FromErr(err)
	}
	if err:= d.Set("type", Type); err != nil{
		return diag.FromErr(err)
	}
	
	// json.NewDecoder(r.Body).Decode(&User_i)
	
	// defer r.Body.Close()

	d.SetId(d.Id())

	fmt.Println("User read")
	

	return diags

/*	
	var diags diag.Diagnostics
	
	client := &http.Client{Timeout: 10 * time.Second}

	UserID := d.Id()
	url := "https://api.zoom.us/v2/users/" + UserID

	req, _ := http.NewRequest("GET", url, nil)
	jwt := os.Getenv("bearer")
	req.Header.Add("authorization", jwt)

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	User_i := Users_res{}
	json.NewDecoder(r.Body).Decode(&User_i)
	// d.Set("email", User_i.Email)
	// d.Set("first_name", User_i.First_name)
	// d.Set("last_name", User_i.Last_name)
	// d.Set("type", User_i.Type)

	defer r.Body.Close()
	
	// d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	
	return diags
*/
}

func resourceOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  
	var diags diag.Diagnostics

  orderID := d.Id()

  if d.HasChange("type") || d.HasChange("email") || d.HasChange("first_name") || d.HasChange("last_name") {
    email := d.Get("email").(string)
	first := d.Get("first_name").(string)
	last := d.Get("last_name").(string)
	Type := d.Get("type").(int)


	Users_List := Users_res{
		Email:      email,
		Type:       Type,
		First_name: first,
		Last_name:  last,
	}

	url := "https://api.zoom.us/v2/users/" + orderID

	payload, err := json.Marshal(Users_List)

	if err != nil {
		return diag.FromErr(err)
	}
	//var bearer = os.Getenv("bearer")

	req, err := http.NewRequest("PATCH", url, strings.NewReader(string(payload)))

	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAyMTgzMTksImlhdCI6MTYxOTYxMzUyMX0.oU-kb6pSAvFU_SJ-KM7gnNeknRh3dM5eLq1TMSp9vuM")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return diag.FromErr(err)
	}

	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)

	// if res.StatusCode!=204{
	// 	return diag.Errorf("Status Code - %v\n%v", res.StatusCode, string(body))
	// }

	d.Set("last_updated", time.Now().Format(time.RFC850))

  }
 //resourceOrderRead(ctx, d, m)

  fmt.Println("User updated")
	
  return diags
}

func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	orderID := d.Id()

	url := "https://api.zoom.us/v2/users/" + orderID +"?action=disassociate" 

	//var bearer = os.Getenv("bearer")

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return diag.FromErr(err)
	}

	// req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAyMTgzMTksImlhdCI6MTYxOTYxMzUyMX0.oU-kb6pSAvFU_SJ-KM7gnNeknRh3dM5eLq1TMSp9vuM")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return diag.FromErr(err)
	}

	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)

	// if res.StatusCode!=204{
	// 	return diag.Errorf("Status Code - %v\n%v", res.StatusCode, string(body))
	// }

	fmt.Println("Delete function is working")

	return diags
}

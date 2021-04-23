package zoom

import (
	"context"
	"os"
	

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"jwt": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZOOM_JWT", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"zoom_orders": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"zoom_users": dataSourceUsers(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

type tok struct {
	token string
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	jwtoken := "Bearer " + d.Get("jwt").(string)
	
	ct := tok{
		token: jwtoken,
	}
	
	os.Setenv("bearer", jwtoken)

	var diags diag.Diagnostics
	
	return ct, diags
}

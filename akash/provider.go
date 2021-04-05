package akash

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"network": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AKASH_NET", "mainnet"),
			},
			"akash_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AKASH_VERSION", "0.10.2-rc1"),
			},
			"chain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CHAIN_ID", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"akash_key":     dataSourceKey(),
			"akash_account": dataSourceAccount(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	network := d.Get("network").(string)
	version := d.Get("akash_version").(string)
	chain_id := d.Get("chain_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c, err := NewAkashClient(&network, nil, &version, &chain_id, nil)

	if err != nil {
		log.Println(err)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Akash client",
			Detail:   "Unable to setup Akash client",
		})
		return nil, diags
	}

	return c, diags
}

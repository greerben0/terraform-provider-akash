package akash

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKeyRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pubkey": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*TerraformAkashClient)

	name := d.Get("name").(string)

	key, err := c.GetKey(name)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("address", key.Address)
	d.Set("pubkey", key.Pubkey)
	d.Set("type", key.Type)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

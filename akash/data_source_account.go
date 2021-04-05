package akash

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAccountRead,
		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"balance": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"denom": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*TerraformAkashClient)

	name := d.Get("address").(string)

	account, err := c.GetAccount(name)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("balance", account.Amount)
	d.Set("denom", account.Denom)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

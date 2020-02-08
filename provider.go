package main

import (
	"context"

	salt "github.com/finarfin/go-salt-netapi-client/cherrypy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backend": {
				Type:     schema.TypeString,
				Required: true,
			},
			"skip_verify": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"salt_minion": resourceSaltMinion(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c := salt.NewClient(
		d.Get("address").(string),
		d.Get("username").(string),
		d.Get("password").(string),
		d.Get("backend").(string),
		d.Get("skip_verify").(bool),
	)

	if err := c.Login(context.Background()); err != nil {
		return nil, err
	}

	return c, nil
}

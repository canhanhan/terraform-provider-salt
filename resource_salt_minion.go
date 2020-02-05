package main

import (
	"context"
	"fmt"

	salt "github.com/finarfin/go-salt-netapi-client/cherrypy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSaltMinion() *schema.Resource {
	return &schema.Resource{
		Create: resourceSaltMinionCreate,
		Read:   resourceSaltMinionRead,
		Delete: resourceSaltMinionDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  2048,
			},
			"private_key": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSaltMinionCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	keySize := d.Get("key_size").(int)

	c := m.(*salt.Client)
	res, err := c.GenerateKeyPair(context.Background(), name, keySize, false)
	if err != nil {
		return err
	}

	if err = d.Set("private_key", res.Private); err != nil {
		return err
	}

	if err = d.Set("public_key", res.Public); err != nil {
		return err
	}

	d.SetId(res.ID)

	return resourceSaltMinionRead(d, m)
}

func resourceSaltMinionRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*salt.Client)
	_, err := c.Key(context.Background(), d.Id())
	if err != nil {
		return err
	}

	return nil
}

func resourceSaltMinionDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*salt.Client)

	t := make(map[string]interface{})
	t["match"] = d.Id()

	cmd := salt.Command{
		Client:    salt.WheelClient,
		Function:  "key.delete",
		Arguments: t,
	}

	res, err := c.RunCommand(context.Background(), cmd)
	if err != nil {
		return err
	}

	dict := res.(map[string]interface{})
	data := dict["data"].(map[string]interface{})
	if !data["success"].(bool) {
		if m, ok := data["return"]; ok {
			return fmt.Errorf("command failed: %s", m)
		}

		return fmt.Errorf("command failed on Salt Master")
	}
	return nil
}

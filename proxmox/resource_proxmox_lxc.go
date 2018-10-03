package proxmox

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceProxmoxLxc() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true
			},
			"ctid": &schema.Schema{
				Type: schema.TypeInt,
				Optional: true,
				ForceNew: true
			},
			"name": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				ForceNew: true
			},
			"cpu_cores": &schema.Schema{
				Type: schema.TypeInt,
				Required: true
			},
			"cpu_limit": &schema.Schema{
				Type: schema.TypeInt,
				Optional: true,
				Default: -1
			},
			"cpu_units": &schema.Schema{
				Type: schema.TypeInt,
				Optional: true,
				Default: 1024
			},
			"mem_total": &schema.Schema{
				Type: schema.TypeInt,
				Required: true
			},
			"mem_swap": &schema.Schema{
				Type: schema.TypeInt,
				Required: True
			},
		}
	}
}

package proxmox

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceProxmoxQemu() *schema.Resource {
	return &schema.Resource{
		Read:   resourceProxmoxQemucRead,
		Create: resourceProxmoxQemucCreate,
		Update: resourceProxmoxQemucUpdate,
		Delete: resourceProxmoxQemucDelete,

		Schema: map[string]*schema.Schema{
			"node": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true
			},
			"vmid": &schema.Schema{
				Type: schema.TypeInt,
				Optional: true,
				ForceNew: true
			},
			"name": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				ForceNew: true
			},
			"cpu_sockets": &schema.Schema{
				Type: schema.TypeInt,
				Required: true
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
			"mem_minimum": &schema.Schema{
				Type: schema.TypeInt,
				Optional: True
			},
			"mem_ballooning": &schema.Schema{
				Type: schema.TypeBool,
				Optional: True,
				Default: true
			}
		}
	}
}

func resourceProxmoxQemuRead(d *schema.ResourceData, meta interface{}) error {
	pve := meta.(*gopve.GoPVE)
	node := d.Get("node").(string)
	vmid := d.Get("vmid").(string)
}

func resourceProxmoxQemuCreate(d *schema.ResourceData, meta interface{}) error {
	pve := meta.(*gopve.GoPVE)
	node := d.Get("node").(string)
	vmid := d.Get("vmid").(string)

	return resourceProxmoxLxcRead(d, meta)
}

func resourceProxmoxQemuUpdate(d *schema.ResourceData, meta interface{}) error {
	pve := meta.(*gopve.GoPVE)

	return resourceProxmoxLxcRead(d, meta)
}

func resourceProxmoxQemuDelete(d *schema.ResourceData, meta interface{}) error {
	pve := meta.(*gopve.GoPVE)
}

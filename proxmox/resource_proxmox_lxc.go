package proxmox

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceProxmoxLxc() *schema.Resource {
	return &schema.Resource{
		Read:   resourceProxmoxLxcRead,
		Create: resourceProxmoxLxcCreate,
		Update: resourceProxmoxLxcUpdate,
		Delete: resourceProxmoxLxcDelete,

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

func resourceProxmoxLxcRead(d *schema.ResourceData, meta interface{}) error {
	pve := meta.(*gopve.GoPVE)
	node := d.Get("node").(string)
	vmid := d.Get("vmid").(string)

	res, err := pve.Node.Get(node).LXC().Get(vmid)
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("vmid", res.VMID)
	d.Set("name", res.name)
}

func resourceProxmoxLxcCreate(d *schema.ResourceData, meta interface{}) error {
	pve := meta.(*gopve.GoPVE)
	node := d.Get("node").(string)
	vmid := d.Get("vmid").(string)

	return resourceProxmoxLxcRead(d, meta)
}

func resourceProxmoxLxcUpdate(d *schema.ResourceData, meta interface{}) error {
	pve := meta.(*gopve.GoPVE)

	return resourceProxmoxLxcRead(d, meta)
}

func resourceProxmoxLxcDelete(d *schema.ResourceData, meta interface{}) error {
	pve := meta.(*gopve.GoPVE)
}

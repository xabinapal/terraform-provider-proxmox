package proxmox

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"proxmox_uri": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("PROXMOX_URI", ""),
				Description: "Proxmox complete URI, with schema, hostname and port"
			},
			"proxmox_user": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("PROXMOX_USER", ""),
				Description: "Proxmox username and auth type (pam, pve...)"
			},
			"proxmox_password": {
				Type: schema.TypeString,
				Required: true,
				Sensitive: true,
				DefaultFunc: schema.EnvDefaultFunc("PROXMOX_PASSWORD", ""),
				Description: "Proxmox authentication password"
			},
			"proxmox_invalid_cert": {
				Type: schema.TypeBool,
				Required: true,
				Default: false,
				Description: "Allow validating self signed certificates" 
			}
		},

		ResourcesMap: map[string]*schema.Resource{
			"proxmox_qemu_clone": resourceProxmoxQemu(),
			"proxmox_lxc_clone":  resourceProxmoxLxc()
		},

		ConfigureFunc: providerConfigure
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	cfg := Config{
		proxmoxUri:         d.Get("proxmox_uri").(string),
		proxmoxUser:        d.Get("proxmox_user").(string),
		proxmoxPassword:    d.Get("proxmox_password").(string),
		proxmoxInvalidCert: d.Get("proxmox_invalid_cert").(bool),
	}

	pve, err := proxmox.CreateFromConfig(cfg)
	return  &pve, err
}
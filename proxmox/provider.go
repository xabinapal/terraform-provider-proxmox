package proxmox

import (
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/xabito/gopve"
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

func providerConfigure(d *schema.ResourceData) (*gopve.GoPVE, error) {
	uri, err := url.Parse(d.Get("proxmox_uri").(string))
	if err != nil {
		return nil, err
	}

	scheme := uri.Scheme
	if scheme == "" {
		scheme = "https"
	}

	port := uri.Port()
	if port == "" {
		port = "8006"
	}

	cfg := &gopve.Config{
		Schema:      scheme,
		Host:        uri.Hostname(),
		Port:        port,
		User:        d.Get("proxmox_user").(string),
		Password:    d.Get("proxmox_password").(string),
		InvalidCert: d.Get("proxmox_invalid_cert").(bool),
	}

	pve, err := NewGoPVE(cfg)
	return &pve, err
}
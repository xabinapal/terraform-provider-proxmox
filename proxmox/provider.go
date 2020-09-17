package proxmox

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/xabinapal/gopve/pkg/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PROXMOX_ENDPOINT", nil),
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if value == "" {
						errors = append(errors, fmt.Errorf("Endpoint must not be an empty string"))
					}

					return
				},
			},

			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PROXMOX_USERNAME", nil),
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if value == "" {
						errors = append(errors, fmt.Errorf("Username must not be an empty string"))
					}

					return
				},
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PROXMOX_PASSWORD", nil),
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if value == "" {
						errors = append(errors, fmt.Errorf("Password must not be an empty string"))
					}

					return
				},
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PROXMOX_INSECURE", nil),
			},

			"cacert_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PROXMOX_CACERT", ""),
			},
		},

		ConfigureFunc: providerConfigure,

		ResourcesMap: map[string]*schema.Resource{},

		DataSourcesMap: map[string]*schema.Resource{},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var endpoint = d.Get("endpoint").(string)
	var username = d.Get("username").(string)
	var password = d.Get("password").(string)
	var insecure = d.Get("insecure").(bool)
	var cacertFile = d.Get("cacert_file").(string)

	uri, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(uri.Port())
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: insecure,
	}

	if cacertFile != "" {
		caCert, err := ioutil.ReadFile(cacertFile)
		if err != nil {
			return nil, err
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.RootCAs = caCertPool
	}

	cli, err := client.NewClient(client.Config{
		Host:   uri.Hostname(),
		Port:   uint16(port),
		Path:   uri.Path,
		Secure: uri.Scheme == "https",
		HTTPTransport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	})
	if err != nil {
		return nil, err
	}

	if err := cli.AuthenticateWithCredentials(username, password); err != nil {
		return nil, err
	}

	return cli.API(), nil
}

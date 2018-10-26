package playground

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func ResourceHCLPlayground() *schema.Resource {
	return &schema.Resource {
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:          schema.TypeString,
				Required:      true,
				ForceNew:      true,
			},
			"network_interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							Computed:         true,
						},

						"address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true, // Computed because it is set if network_ip is set.
							Optional: true,
							ForceNew: true,
						},

						"network_ip": &schema.Schema{
							Type:       schema.TypeString,
							Computed:   true, // Computed because it is set if address is set.
							Optional:   true,
							ForceNew:   true,
							Deprecated: "Please use address",
						},

						"subnetwork": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							Computed:         true,
						},

						"subnetwork_project": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},

						"access_config": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nat_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"network_tier": &schema.Schema{
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.StringInSlice([]string{"PREMIUM", "STANDARD"}, false),
									},
									// Instance templates will never have an
									// 'assigned NAT IP', but we need this in
									// the schema to allow us to share flatten
									// code with an instance, which could.
									"assigned_nat_ip": &schema.Schema{
										Type:       schema.TypeString,
										Computed:   true,
										Deprecated: "Use network_interface.access_config.nat_ip instead.",
									},
								},
							},
						},
					},
				},
			},
		},
		Create: resourcePlaygroundCreate,
		Read: resourcePlaygroundRead,
		Update: nil,
		Delete:resourcePlaygroundDelete,
	}
}

func resourcePlaygroundCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	d.SetId(name)
	return nil
}

func resourcePlaygroundRead(d *schema.ResourceData, m interface{}) error {
	networkInterface := d.Get("network_interface").([]interface{})
	for _, raw := range networkInterface {
		data := raw.(map[string]interface{})
		_ = data["subnetwork"].(string)
		_ = data["address"].(string)
		accessConfigs := data["access_config"].([]interface{})
		for _, rawConfigs := range accessConfigs {
			dataconfigs := rawConfigs.(map[string]interface{})
			_ = dataconfigs["nat_ip"].(string)
		}
	}
	return nil
}

func resourcePlaygroundDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
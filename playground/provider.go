package playground

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"playground_hcl_playground": ResourceHCLPlayground(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  configureFunc,
	}
}

func configureFunc(d *schema.ResourceData) (interface{}, error) {
	return nil, nil
}
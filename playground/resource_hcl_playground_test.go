package playground

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestHCLPlayground(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hclUnderTest(),
				Check:  noOpCheck(),
			},
		},
	})
}

func noOpCheck() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return nil
	}
}

func hclUnderTest() string {
	return fmt.Sprintf(`
locals {
	nat-ip = "192.168.0.2"
	cond = "true"
	instance_network_access_config = {
		no-external-ip = "${list()}"
		external-ip = [{
			nat_ip = "${local.nat-ip}"
		}]
	}
}

resource "playground_hcl_playground" "mytest1" {
	name = "mytest"

	network_interface {
		subnetwork = "sub"
		address = "an address"
		access_config = ["${local.instance_network_access_config["${local.cond ? "external-ip" : "no-external-ip"}"]}"]
	}
}
`)
}
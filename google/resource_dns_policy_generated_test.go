// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDNSPolicy_dnsPolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNSPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSPolicy_dnsPolicyBasicExample(context),
			},
			{
				ResourceName:      "google_dns_policy.example-policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDNSPolicy_dnsPolicyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dns_policy" "example-policy" {
  name                      = "tf-test-example-policy%{random_suffix}"
  enable_inbound_forwarding = true

  enable_logging = true

  alternative_name_server_config {
    target_name_servers {
      ipv4_address = "172.16.1.10"
    }
    target_name_servers {
      ipv4_address = "172.16.1.20"
    }
  }

  networks {
    network_url = google_compute_network.network-1.self_link
  }
  networks {
    network_url = google_compute_network.network-2.self_link
  }
}

resource "google_compute_network" "network-1" {
  name                    = "tf-test-network-1%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-2" {
  name                    = "tf-test-network-2%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func testAccCheckDNSPolicyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dns_policy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DNSBasePath}}projects/{{project}}/policies/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("DNSPolicy still exists at %s", url)
			}
		}

		return nil
	}
}
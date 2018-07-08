package newrelic

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNewRelicAlertChannelDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNewRelicAlertChannelDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccNewRelicAlertChannel("data.newrelic_alert_channel.channel"),
				),
			},
		},
	})
}

func testAccNewRelicAlertChannel(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to get an alert channel from New Relic")
		}

		if strings.Contains(strings.ToLower(testAccExpectedAlertChannelName), strings.ToLower(a["name"])) {
			return fmt.Errorf("Expected the alert channel name to be: %s, but got: %s", testAccExpectedAlertChannelName, a["name"])
		}

		return nil
	}
}

// The test newrelic channel name
const testAccNewRelicAlertChannelDataSourceConfig = `
data "newrelic_alert_channel" "channel" {
	name = "tf-test@example.com"
}
`

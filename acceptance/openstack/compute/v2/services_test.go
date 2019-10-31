// +build acceptance compute services

package v2

import (
	"testing"

	"github.com/phouverneyuff/gophercloud/acceptance/clients"
	"github.com/phouverneyuff/gophercloud/acceptance/tools"
	"github.com/phouverneyuff/gophercloud/openstack/compute/v2/extensions/services"
	th "github.com/phouverneyuff/gophercloud/testhelper"
)

func TestServicesList(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	allPages, err := services.List(client).AllPages()
	th.AssertNoErr(t, err)

	allServices, err := services.ExtractServices(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, service := range allServices {
		tools.PrintResource(t, service)

		if service.Binary == "nova-scheduler" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}

package testing

import (
	"testing"

	"github.com/phouverneyuff/gophercloud/openstack/loadbalancer/v2/apiversions"
	th "github.com/phouverneyuff/gophercloud/testhelper"
	"github.com/phouverneyuff/gophercloud/testhelper/client"
)

func TestListVersions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	allVersions, err := apiversions.List(client.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)

	actual, err := apiversions.ExtractAPIVersions(allVersions)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, OctaviaAllAPIVersionResults, actual)
}

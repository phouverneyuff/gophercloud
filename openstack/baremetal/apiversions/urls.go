package apiversions

import (
	"github.com/phouverneyuff/gophercloud"
)

func getURL(c *gophercloud.ServiceClient, version string) string {
	return c.ServiceURL(version)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL()
}

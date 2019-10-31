package buildinfo

import "github.com/phouverneyuff/gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}

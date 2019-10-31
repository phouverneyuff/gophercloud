package diagnostics

import (
	"github.com/phouverneyuff/gophercloud"
)

// Diagnostics
func Get(client *gophercloud.ServiceClient, serverId string) (r serverDiagnosticsResult) {
	_, r.Err = client.Get(serverDiagnosticsURL(client, serverId), &r.Body, nil)
	return
}

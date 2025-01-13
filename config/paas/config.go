package paas

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("gridscale_filesystem", func(r *config.Resource) {
		// TODO: Investigate cross resource reference for the security_zone_uuid

		// Cross Resource Reference for network_uuid to gridscale_network
		r.References["network_uuid"] = config.Reference{
			TerraformName: "gridscale_network",
		}
	})
}

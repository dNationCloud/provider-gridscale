package compute

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("gridscale_k8s", func(r *config.Resource) {
		// TODO: Investigate cross resource reference for the security_zone_uuid
		// r.References["security_zone_uuid"] = config.Reference{
		// 	TerraformName: "gridscale_paas_securityzone",
		// }
	})
	p.AddResourceConfigurator("gridscale_server", func(r *config.Resource) {
		// Cross Resource Reference for storage.object_uuid to gridscale_storage
		r.References["storage.object_uuid"] = config.Reference{
			TerraformName: "gridscale_storage",
		}
		// Cross Resource Reference for network.object_uuid to gridscale_network
		r.References["network.object_uuid"] = config.Reference{
			TerraformName: "gridscale_network",
		}
		// Cross Resource Reference for network.firewall_template_uuid to gridscale_firewall
		r.References["network.firewall_template_uuid"] = config.Reference{
			TerraformName: "gridscale_firewall",
		}

		// TODO: Cross Resource Reference for isoimage to gridscale_isoimage
		// r.References["isoimage"] = config.Reference{
		// 	TerraformName: "gridscale_isoimage",
		// }

		// Cross Resource Reference for ipv4 to gridscale_ipv4
		r.References["ipv4"] = config.Reference{
			TerraformName: "gridscale_ipv4",
		}
		// Cross Resource Reference for ipv6 to gridscale_ipv6
		r.References["ipv6"] = config.Reference{
			TerraformName: "gridscale_ipv6",
		}
	})
}

// TODO: Implement timeout logic for the resources.
// The timeout logic should handle scenarios where a timeout is set, and the resource is not synchronized within the specified time.
// Without this, the reconcile loop will repeatedly attempt to create the resource in gridscale,
// resulting in multiple unwanted resource instances being created every time the reconcile loop runs.
// This can lead to resource duplication and unnecessary load on the gridscale API.

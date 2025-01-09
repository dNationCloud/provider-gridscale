package compute

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("gridscale_k8s", func(r *config.Resource) {
		// TODO: Investigate cross resource reference for the security_zone_uuid
	})
}

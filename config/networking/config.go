package networking

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("gridscale_ipv4", func(r *config.Resource) {})
	p.AddResourceConfigurator("gridscale_ipv6", func(r *config.Resource) {})
	p.AddResourceConfigurator("gridscale_network", func(r *config.Resource) {})
	p.AddResourceConfigurator("gridscale_sshkey", func(r *config.Resource) {})
}

// TODO: Implement timeout logic for the resources.
// The timeout logic should handle scenarios where a timeout is set, and the resource is not synchronized within the specified time.
// Without this, the reconcile loop will repeatedly attempt to create the resource in gridscale,
// resulting in multiple unwanted resource instances being created every time the reconcile loop runs.
// This can lead to resource duplication and unnecessary load on the gridscale API.
// The `addTimeoutSchema` function is a potential starting point for defining timeout behavior by adding a `timeouts` schema
// that allows users to configure timeouts for create, update, and delete operations. Also take a look at this
// https://github.com/gridscale/terraform-provider-gridscale/blob/master/gridscale/resource_gridscale_ipv4.go#L108
// func addTimeoutSchema(r *config.Resource) {
// 	r.TerraformResource.Schema["timeouts"] = &schema.Schema{
// 		Type:        schema.TypeMap,
// 		Optional:    true,
// 		Description: "Timeouts configuration (create, update, delete).",
// 		Elem: &schema.Schema{
// 			Type: schema.TypeString,
// 		},
// 	}
// }

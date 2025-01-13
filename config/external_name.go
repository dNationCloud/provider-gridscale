/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Networking
	"gridscale_ipv4":            config.IdentifierFromProvider,
	"gridscale_ipv6":            config.IdentifierFromProvider,
	"gridscale_network":         config.IdentifierFromProvider,
	"gridscale_sshkey":          config.IdentifierFromProvider,
	"gridscale_loadbalancer":    config.IdentifierFromProvider,
	"gridscale_ssl_certificate": config.IdentifierFromProvider,
	"gridscale_firewall":        config.IdentifierFromProvider,

	// Compute
	"gridscale_k8s":    config.IdentifierFromProvider,
	"gridscale_server": config.IdentifierFromProvider,

	// Storage
	"gridscale_storage":                  config.IdentifierFromProvider,
	"gridscale_template":                 config.IdentifierFromProvider,
	"gridscale_snapshot":                 config.IdentifierFromProvider,
	"gridscale_isoimage":                 config.IdentifierFromProvider,
	"gridscale_object_storage_accesskey": config.IdentifierFromProvider,

	// PaaS
	"gridscale_filesystem": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

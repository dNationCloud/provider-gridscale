/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/dNationCloud/provider-gridscale/config/compute"
	"github.com/dNationCloud/provider-gridscale/config/networking"
	"github.com/dNationCloud/provider-gridscale/config/paas"
	"github.com/dNationCloud/provider-gridscale/config/storage"
)

const (
	resourcePrefix = "gridscale"
	modulePath     = "github.com/dNationCloud/provider-gridscale"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		networking.Configure,
		compute.Configure,
		storage.Configure,
		paas.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

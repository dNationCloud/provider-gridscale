package storage

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("gridscale_snapshot", func(r *config.Resource) {
		// Cross Resource Reference for storage_uuid to gridscale_storage
		r.References["storage_uuid"] = config.Reference{
			TerraformName: "gridscale_storage",
		}
	})
	p.AddResourceConfigurator("gridscale_template", func(r *config.Resource) {
		// Cross Resource Reference for snapshot_uuid to gridscale_snapshot
		r.References["snapshot_uuid"] = config.Reference{
			TerraformName: "gridscale_snapshot",
		}
	})
	p.AddResourceConfigurator("gridscale_storage", func(r *config.Resource) {
		// TODO: Investigate if cross resource reference is needed for the rollback_from_backup_uuid

		// Cross Resource Reference for template.template_uuid to gridscale_template
		r.References["template.template_uuid"] = config.Reference{
			TerraformName: "gridscale_template",
		}
		// Cross Resource Reference for template.sshkeys to gridscale_sshkey
		r.References["template.sshkeys"] = config.Reference{
			TerraformName: "gridscale_sshkey",
		}
	})
	// p.AddResourceConfigurator("gridscale_isoimage", func(r *config.Resource) {})
}

// TODO: Implement timeout logic for the resources.
// The timeout logic should handle scenarios where a timeout is set, and the resource is not synchronized within the specified time.
// Without this, the reconcile loop will repeatedly attempt to create the resource in gridscale,
// resulting in multiple unwanted resource instances being created every time the reconcile loop runs.
// This can lead to resource duplication and unnecessary load on the gridscale API.

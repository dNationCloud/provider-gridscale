package networking

import (
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("gridscale_ipv4", func(r *config.Resource) {
		// Add validation for the "name" field (max length 64 characters)
		r.TerraformResource.Schema["name"].ValidateFunc = func(val interface{}, key string) ([]string, []error) {
			value, ok := val.(string)
			if !ok {
				return nil, []error{fmt.Errorf("expected type string for name, got %T", val)}
			}
			if len(value) > 64 {
				return nil, []error{fmt.Errorf("name must be 64 characters or fewer")}
			}
			return nil, nil
		}

		r.TerraformResource.Schema["timeouts"] = &schema.Schema{
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "Timeouts configuration (create, update, delete).",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		}
	})
}

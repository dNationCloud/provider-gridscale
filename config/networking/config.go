package networking

import (
	"fmt"
	"time"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("gridscale_ipv4", func(r *config.Resource) {
		addNameValidation(r)
		addTimeoutSchema(r)
	})
}

// addNameValidation adds validation for the "name" field.
func addNameValidation(r *config.Resource) {
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
}

// addTimeoutSchema adds the timeouts schema and handles defaults.
func addTimeoutSchema(r *config.Resource) {
	r.TerraformResource.Schema["timeouts"] = &schema.Schema{
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "Timeouts configuration (create, update, delete).",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}

	// Set default timeouts
	r.OperationTimeouts = config.OperationTimeouts{
		Create: 5 * time.Minute,
		Update: 5 * time.Minute,
		Delete: 5 * time.Minute,
	}

	// Check if timeouts are provided in the resource configuration
	if timeouts, ok := r.TerraformResource.Schema["timeouts"].Default.(map[string]interface{}); ok {
		// Parse timeouts for create, update, and delete
		parseTimeout("create", timeouts, &r.OperationTimeouts.Create)
		parseTimeout("update", timeouts, &r.OperationTimeouts.Update)
		parseTimeout("delete", timeouts, &r.OperationTimeouts.Delete)
	}
}

// parseTimeout parses a timeout value from the configuration and sets it to the provided pointer.
func parseTimeout(operation string, timeouts map[string]interface{}, timeoutDuration *time.Duration) {
	if timeoutVal, exists := timeouts[operation]; exists && timeoutVal != "" {
		duration, err := time.ParseDuration(timeoutVal.(string))
		if err == nil {
			*timeoutDuration = duration
		} else {
			// Handle invalid duration format
			fmt.Printf("Invalid duration format for %s timeout: %v\n", operation, err)
		}
	}
}

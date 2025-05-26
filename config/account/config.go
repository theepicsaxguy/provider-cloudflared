package account

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for account group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_account", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "Account"

		// Configure external name using NameAsIdentifier
		r.ExternalName = config.NameAsIdentifier

		// Configure which fields are for spec vs status
		r.References = config.References{
			"name": {
				Type: "string",
			},
		}

		// Enable async operations
		r.UseAsync = true

		// Configure fields
		if s, ok := r.TerraformResource.Schema["id"]; ok {
			s.Required = false
			s.Computed = true
		}

		// Configure name field
		if s, ok := r.TerraformResource.Schema["name"]; ok {
			s.Required = true
		}
	})

	p.AddResourceConfigurator("cloudflare_account_member", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "AccountMember"

		// Configure external name
		r.ExternalName = config.IdentifierFromProvider

		// Reference to parent Account
		r.References = config.References{
			"account_id": {
				Type: "Account",
			},
		}

		// Ensure proper status tracking
		r.UseAsync = true

		// Set which fields belong in status vs spec
		if s, ok := r.TerraformResource.Schema["id"]; ok {
			s.Required = false
			s.Computed = true
		}

		// Configure email field
		if s, ok := r.TerraformResource.Schema["email"]; ok {
			s.Required = true
		}

		// Configure connection details
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if id, ok := attr["id"].(string); ok {
				conn["member_id"] = []byte(id)
			}
			if email, ok := attr["email"].(string); ok {
				conn["email"] = []byte(email)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_api_token", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "APIToken"

		// Configure external name
		r.ExternalName = config.NameAsIdentifier

		// Ensure proper status tracking
		r.UseAsync = true

		// Configure fields
		if s, ok := r.TerraformResource.Schema["id"]; ok {
			s.Required = false
			s.Computed = true
		}

		// Mark sensitive fields
		if s, ok := r.TerraformResource.Schema["value"]; ok {
			s.Sensitive = true
		}

		// Configure name field
		if s, ok := r.TerraformResource.Schema["name"]; ok {
			s.Required = true
		}
	})
}

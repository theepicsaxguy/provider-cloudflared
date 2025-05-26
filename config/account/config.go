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
		r.TerraformResource.Schema["id"].Required = false
		r.TerraformResource.Schema["id"].Computed = true

		// Ensure proper status tracking
		r.UseAsync = true

		// Configure name field
		if s, ok := r.TerraformResource.Schema["name"]; ok {
			s.Required = true
		}

		// Set up base types
		r.References = config.References{
			"name": {
				Type: "string",
			},
		}
	})

	p.AddResourceConfigurator("cloudflare_account_member", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "AccountMember"

		// Configure external name
		r.ExternalName = config.IdentifierFromProvider

		// Reference to parent Account
		r.References["account_id"] = config.Reference{
			Type: "Account",
		}

		// Ensure proper status tracking
		r.UseAsync = true

		// Set which fields belong in status vs spec
		r.TerraformResource.Schema["id"].Required = false
		r.TerraformResource.Schema["id"].Computed = true

		// Configure email field
		if s, ok := r.TerraformResource.Schema["email"]; ok {
			s.Required = true
		}

		// Set up base types
		r.References = config.References{
			"email": {
				Type: "string",
			},
		}

		// Ensure proper managed resource setup
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

		// Mark sensitive fields
		if s, ok := r.TerraformResource.Schema["value"]; ok {
			s.Sensitive = true
		}

		// Set which fields belong in status vs spec
		r.TerraformResource.Schema["id"].Required = false
		r.TerraformResource.Schema["id"].Computed = true

		// Configure name field
		if s, ok := r.TerraformResource.Schema["name"]; ok {
			s.Required = true
		}

		// Ensure proper status tracking
		r.UseAsync = true

		// Set up base types
		r.References = config.References{
			"name": {
				Type: "string",
			},
		}

		// Configure sensitive values
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if val, ok := attr["value"].(string); ok {
				conn["token"] = []byte(val)
			}
			if id, ok := attr["id"].(string); ok {
				conn["token_id"] = []byte(id)
			}
			return conn, nil
		}
	})
}

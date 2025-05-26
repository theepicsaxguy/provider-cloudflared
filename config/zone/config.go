package zone

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	shortGroupName = "Zone"
)

// Configure adds configurations for zone group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_zone", func(r *config.Resource) {
		r.ShortGroup = "zone"
		r.Kind = "Zone"

		// Configure external name using NameAsIdentifier
		r.ExternalName = config.NameAsIdentifier

		r.References["account_id"] = config.Reference{
			Type: "github.com/theepicsaxguy/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zone_settings_override", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "ZoneSettingsOverride"

		// Configure external name
		r.ExternalName = config.NameAsIdentifier

		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_zone_dnssec", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "DNSSEC"

		// Configure external name
		r.ExternalName = config.IdentifierFromProvider

		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}

		// Configure which fields are for spec vs status
		r.TerraformResource.Schema["id"].Required = false
		r.TerraformResource.Schema["id"].Computed = true

		// Configure managed resource fields
		r.UseAsync = true
	})

	p.AddResourceConfigurator("cloudflare_total_tls", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "TotalTLS"

		// Configure external name
		r.ExternalName = config.IdentifierFromProvider

		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}

		// Configure managed resource fields
		r.UseAsync = true
	})

	p.AddResourceConfigurator("cloudflare_tiered_cache", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "TieredCache"

		// Configure external name
		r.ExternalName = config.IdentifierFromProvider

		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}

		// Configure managed resource fields
		r.UseAsync = true
	})

	p.AddResourceConfigurator("cloudflare_logpull_retention", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "LogpullRetention"

		// Configure external name
		r.ExternalName = config.IdentifierFromProvider

		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}

		// Configure managed resource fields
		r.UseAsync = true
	})
}

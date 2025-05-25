/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/theepicsaxguy/provider-cloudflare/config/access"
	"github.com/theepicsaxguy/provider-cloudflare/config/account"
	"github.com/theepicsaxguy/provider-cloudflare/config/apishield"
	"github.com/theepicsaxguy/provider-cloudflare/config/argo"
	"github.com/theepicsaxguy/provider-cloudflare/config/authenticatedoriginpulls"
	"github.com/theepicsaxguy/provider-cloudflare/config/byoip"
	"github.com/theepicsaxguy/provider-cloudflare/config/certificate"
	"github.com/theepicsaxguy/provider-cloudflare/config/custom"
	"github.com/theepicsaxguy/provider-cloudflare/config/customhostname"
	"github.com/theepicsaxguy/provider-cloudflare/config/dlp"
	"github.com/theepicsaxguy/provider-cloudflare/config/dns"
	"github.com/theepicsaxguy/provider-cloudflare/config/emailrouting"
	"github.com/theepicsaxguy/provider-cloudflare/config/filters"
	"github.com/theepicsaxguy/provider-cloudflare/config/firewall"
	"github.com/theepicsaxguy/provider-cloudflare/config/lists"
	"github.com/theepicsaxguy/provider-cloudflare/config/loadbalancer"
	"github.com/theepicsaxguy/provider-cloudflare/config/logpush"
	"github.com/theepicsaxguy/provider-cloudflare/config/magic"
	"github.com/theepicsaxguy/provider-cloudflare/config/notification"
	"github.com/theepicsaxguy/provider-cloudflare/config/originca"
	"github.com/theepicsaxguy/provider-cloudflare/config/page"
	"github.com/theepicsaxguy/provider-cloudflare/config/pages"
	"github.com/theepicsaxguy/provider-cloudflare/config/ruleset"
	"github.com/theepicsaxguy/provider-cloudflare/config/spectrum"
	"github.com/theepicsaxguy/provider-cloudflare/config/teams"
	"github.com/theepicsaxguy/provider-cloudflare/config/waf"
	"github.com/theepicsaxguy/provider-cloudflare/config/waitingroom"
	"github.com/theepicsaxguy/provider-cloudflare/config/warp"
	"github.com/theepicsaxguy/provider-cloudflare/config/web3"
	"github.com/theepicsaxguy/provider-cloudflare/config/worker"
	"github.com/theepicsaxguy/provider-cloudflare/config/zone"
)

const (
	resourcePrefix = "cloudflare"
	modulePath     = "github.com/theepicsaxguy/provider-cloudflare"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		access.Configure,
		account.Configure,
		apishield.Configure,
		argo.Configure,
		authenticatedoriginpulls.Configure,
		byoip.Configure,
		certificate.Configure,
		custom.Configure,
		customhostname.Configure,
		dlp.Configure,
		dns.Configure,
		emailrouting.Configure,
		filters.Configure,
		firewall.Configure,
		lists.Configure,
		loadbalancer.Configure,
		logpush.Configure,
		magic.Configure,
		notification.Configure,
		originca.Configure,
		page.Configure,
		pages.Configure,
		ruleset.Configure,
		spectrum.Configure,
		teams.Configure,
		waf.Configure,
		waitingroom.Configure,
		warp.Configure,
		web3.Configure,
		worker.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

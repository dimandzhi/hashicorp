package provider

import (
	"github.com/dimandzhi/hashicorp/go-discover/provider/dns"
	discover "github.com/hashicorp/go-discover"
)

func AllProvaiders() map[string]discover.Provider {
	return map[string]discover.Provider{
		"dns": &dns.Provider{},
	}
}

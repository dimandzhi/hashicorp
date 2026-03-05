// Package dns provides node discovery via DNS.
package dns

import (
	"fmt"
	"io"
	"log"
	"net"

	discover "github.com/hashicorp/go-discover"
)

// Provider implements the Provider interface.
type Provider struct{}

var _ discover.Provider = (*Provider)(nil)

// Help returns help information for the mDNS package.
func (p *Provider) Help() string {
	return `DNS:

    provider:          "dns"
    hostname:          The hostname to resolve.
`
}

// Addrs returns discovered addresses for the mDNS package.
func (p *Provider) Addrs(args map[string]string, l *log.Logger) (addrs []string, err error) {
	// default to null logger
	if l == nil {
		l = log.New(io.Discard, "", 0)
	}

	// validate and set hostname record
	if args["hostname"] == "" {
		return nil, fmt.Errorf("discover-dns: hostname record not provided." +
			"  Please specify a hostname record for the DNS lookup")
	}

	addrs = make([]string, 0)
	if addrs, err = net.LookupHost(args["hostname"]); err != nil {
		err = fmt.Errorf("discover-dns: hostname lookup error: %w", err)
		return
	}

	return
}

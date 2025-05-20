package main

import (
	plugin "github.com/jsiebens/libdns-plugin"
	"github.com/libdns/route53"
)

func main() {
	plugin.Serve(&route53.Provider{})
}

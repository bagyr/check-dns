package main

import (
	"github.com/miekg/dns"

	"github.com/bagyr/dns-check/internal/log"
)

func main() {
	// cfg := dns.ClientConfig{
	// 	Servers:  []string{"8.8.8.8"},
	// 	Port:     "53",
	// 	Ndots:    1,
	// 	Timeout:  5,
	// 	Attempts: 2,
	// }

	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion("api.github.com.", dns.TypeCNAME)
	m.RecursionAvailable = true

	r, _, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		log.S.Fatalf("error on exchange: %s", err)
	}

	for _, a := range r.Answer {
		if mx, ok := a.(*dns.CNAME); ok {
			log.S.Infof("%v\n", mx)
		}
	}
}

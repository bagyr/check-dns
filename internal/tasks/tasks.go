package tasks

import "github.com/miekg/dns"

type CNameTask struct {
	Url  string
	Name string
}

func (t *CNameTask) Check(dnsServer string) (bool, error) {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(t.Url, dns.TypeCNAME)
	m.RecursionAvailable = true

	r, _, err := c.Exchange(m, dnsServer)
	if err != nil {
		return false, err
	}

	for _, a := range r.Answer {
		if mx, ok := a.(*dns.CNAME); ok {
			return mx.Target == t.Name, nil
		}
	}

	return false, nil
}

type ATask struct {
	Url     string
	Records []string
	rMap    map[string]struct{}
}

func (t *ATask) Check(dnsServer string) (bool, error) {
	if t.rMap == nil {
		t.rMap = make(map[string]struct{})

		for i := range t.Records {
			t.rMap[t.Records[i]] = struct{}{}
		}
	}

	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(t.Url, dns.TypeA)
	m.RecursionAvailable = true

	r, _, err := c.Exchange(m, dnsServer)
	if err != nil {
		return false, err
	}

	for _, a := range r.Answer {
		if mx, ok := a.(*dns.A); ok {
			if _, ok := t.rMap[string(mx.A)]; !ok {
				return false, nil
			}
		}
	}

	return true, nil
}

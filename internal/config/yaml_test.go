package config

import (
	"testing"
	"time"

	"github.com/bagyr/dns-check/internal/tasks"
)

func TestFromYamlString(t *testing.T) {
	config := `
options:
  interval: 1m
dns:
  api.some-api.org:
    cname: alias-api.some-dns.org.
  prod-some-lb.some-dns.org:
    a:
      - 1.1.1.1
      - 2.2.2.2
      - 3.3.3.3
      - 4.4.4.4
      - 5.5.5.5
      - 6.6.6.6
dns_servers:
  - 10.0.0.1
  - 10.1.0.1`

	appConf, err := FromYamlString(config)
	if err != nil {
		t.Error(err)
	}

	if appConf.UpdateInterval != time.Minute {
		t.Errorf("wrong interval: %v", appConf.UpdateInterval)
	}

	cNameUrl := "api.some-api.org"
	if _, ok := appConf.Tasks[cNameUrl]; !ok {
		t.Fatalf("task %s not found", cNameUrl)
	}

	if _, ok := appConf.Tasks[cNameUrl].(*tasks.CNameTask); !ok {
		t.Errorf("wrong task type: %T", appConf.Tasks[cNameUrl])
	}

	aRecordUrl := "prod-some-lb.some-dns.org"
	if _, ok := appConf.Tasks[aRecordUrl]; !ok {
		t.Fatalf("task %s not found", aRecordUrl)
	}

	aTask, ok := appConf.Tasks[aRecordUrl].(*tasks.ATask)
	if !ok {
		t.Errorf("wrong task type: %T", appConf.Tasks[aRecordUrl])
	}

	if len(aTask.Records) != 6 {
		t.Errorf("wrong A task ip list len: %d, %v", len(aTask.Records), aTask.Records)
	}

	if aTask.Records[0] != "1.1.1.1" {
		t.Errorf("wrong A task record: %s", aTask.Records[0])
	}
}

package config

import (
	"fmt"
	"time"

	"github.com/bagyr/dns-check/internal/tasks"
	"gopkg.in/yaml.v3"
)

type task map[string]any

type yamlScheme struct {
	Options struct {
		CheckInteval time.Duration `yaml:"interval"`
	} `yaml:"options"`
	Servers []string        `yaml:"dns_servers"`
	Tasks   map[string]task `yaml:"dns"`
}

func FromYamlString(config string) (*AppConfig, error) {
	cfg := yamlScheme{}

	err := yaml.Unmarshal([]byte(config), &cfg)
	if err != nil {
		return nil, err
	}

	out := AppConfig{
		UpdateInterval: cfg.Options.CheckInteval,
		Tasks:          make(map[string]Task),
	}

	for url, task := range cfg.Tasks {
		for typ, val := range task {
			switch typ {
			case "cname":
				name, ok := val.(string)
				if !ok {
					return nil, fmt.Errorf("error on task parse: unknown %v near %s", val, url)
				}
				out.Tasks[url] = &tasks.CNameTask{
					Name: name,
				}
			case "a":
				ips, ok := val.([]any)
				if !ok {
					return nil, fmt.Errorf("error on task parse: unknown %v(%T) near %s", val, val, url)
				}
				newTask := &tasks.ATask{
					Records: make([]string, len(ips)),
				}
				for idx, ip := range ips {
					i, ok := ip.(string)
					if !ok {
						return nil, fmt.Errorf("cannot parse %+v", ip)
					}
					newTask.Records[idx] = i
				}
				out.Tasks[url] = newTask
			default:
				return nil, fmt.Errorf("unknown task type: %s", typ)
			}
		}
	}

	return &out, nil
}

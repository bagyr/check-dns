package config

import "time"

type Task interface {
	Check() bool
}

type AppConfig struct {
	UpdateInterval time.Duration
	Tasks          map[string]Task
}

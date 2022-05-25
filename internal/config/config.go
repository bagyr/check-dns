package config

import "time"

type Task interface {
	Check(string) (bool, error)
}

type AppConfig struct {
	UpdateInterval time.Duration
	Tasks          map[string]Task
	Servers        []string
}

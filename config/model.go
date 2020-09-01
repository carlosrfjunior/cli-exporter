package config

import (
	"sync"
)

// Services ....
// type Services struct {
// 	Name  string `yaml:"name"`
// 	Value bool   `yaml:"value"`
// }

// Provider ...
type Provider struct {
	Name     string          `yaml:"name"`
	Services map[string]bool `yaml:"services"`
}

// Config ....
type Config struct {
	Provider Provider `yaml:"provider"`
}

var (
	instance *Config
	once     sync.Once
)

//Once Singleton
func Once() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}

//Get ...
func (cfg *Config) Get() *Config {
	return cfg
}

//GetProviderName ...
func (cfg *Config) GetProviderName() string {
	return cfg.Provider.Name
}

//SetProviderName ...
func (cfg *Config) SetProviderName(name string) {
	cfg.Provider.Name = name
}

//GetProviderServices ...
func (cfg *Config) GetProviderServices() map[string]bool {
	return cfg.Provider.Services
}

//SetProviderServices ..
func (cfg *Config) SetProviderServices(services map[string]bool) {
	cfg.Provider.Services = services
}

//GetProviderService ...
func (cfg *Config) GetProviderService(name string) bool {
	return cfg.Provider.Services[name]
}

//SetProviderService ..
func (cfg *Config) SetProviderService(name string, value bool) {
	cfg.Provider.Services[name] = value
}

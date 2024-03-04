package domain

import (
	"fmt"
	"waffle/internal/config"
)

type YamlNameSystemProvider struct {
	cfg *config.YamlConfig
}

func NewYamlNameSystemProvider(cfg *config.YamlConfig) *YamlNameSystemProvider {
	return &YamlNameSystemProvider{
		cfg: cfg,
	}
}

func (y *YamlNameSystemProvider) GetAddress(host string) (string, error) {
	for _, h := range y.cfg.DNS {
		if h.Host == host {
			return h.Address, nil
		}
	}

	return "", fmt.Errorf("%s not found in configuration", host)
}

package domain

import (
	"fmt"
	"net/url"

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

func (y *YamlNameSystemProvider) GetAddress(host string) (*url.URL, error) {
	for _, h := range y.cfg.DNS {
		if h.Host == host {
			urlAddress, err := url.Parse(h.Address)
			if err != nil {
				return nil, fmt.Errorf("parse dns address: %w", err)
			}

			return urlAddress, nil
		}
	}

	return nil, fmt.Errorf("'%s' not found in configuration", host)
}

package config

import (
	"fmt"

	"github.com/goccy/go-yaml"
)

type YamlConfig struct {
	DNS []*Dns `yaml:"dns"`
}

type Dns struct {
	Host    string `yaml:"host"`
	Address string `yaml:"address"`
}

func NewYamlConfig(yamlBytes []byte) (*YamlConfig, error) {
	var yamlCfg YamlConfig
	if err := yaml.Unmarshal(yamlBytes, &yamlCfg); err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}

	return &yamlCfg, nil
}

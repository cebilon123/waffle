package config

import (
	"fmt"

	"github.com/goccy/go-yaml"
)

type YamlConfig struct {
	DNS   []*Dns `yaml:"dns"`
	Rules Rules  `yaml:"rules"`
}

type Dns struct {
	Host    string `yaml:"host"`
	Address string `yaml:"address"`
}

type Rules struct {
	Custom []CustomRule `yaml:"custom"`
}

// CustomRule is a rule that do the string matching on the request body.
// If negated is true, then it will do the !contains validation on the payload.
type CustomRule struct {
	Name      string `yaml:"name"`
	Predicate string `yaml:"predicate"`
}

func NewYamlConfig(yamlBytes []byte) (*YamlConfig, error) {
	var yamlCfg YamlConfig
	if err := yaml.Unmarshal(yamlBytes, &yamlCfg); err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}

	return &yamlCfg, nil
}

package config

type YamlConfig struct {
	DNS []*Dns
}

type Dns struct {
	Domain  string `yaml:"domain"`
	Address string `yaml:"address"`
}

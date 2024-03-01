package domain

type NameSystem interface {
	GetAddress(domain string) (string, error)
}

// YAMLBasedNameSystem is a DNS based on the configuration
// passed using YAML config file.
type YAMLBasedNameSystem struct {
}

func NewYAMLBasedNameSystem() *YAMLBasedNameSystem {
	return &YAMLBasedNameSystem{}
}

func (Y *YAMLBasedNameSystem) GetAddress(domain string) (string, error) {
	//TODO implement me
	panic("implement me")
}

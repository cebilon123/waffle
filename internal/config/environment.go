package config

import (
	"errors"
	"fmt"

	"github.com/Netflix/go-env"
)

type Environment struct {
	Extras env.EnvSet
}

func (e *Environment) Validate() error {
	errs := make([]error, 0)
	
	return errors.Join(errs...)
}

func LoadEnvironmentConfig() (*Environment, error) {
	var environment Environment

	es, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		return nil, fmt.Errorf("load config from environment variables: %w", err)
	}

	environment.Extras = es

	if err := environment.Validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}

	return &environment, nil
}

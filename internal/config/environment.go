package config

import (
	"errors"
	"fmt"

	"github.com/Netflix/go-env"
)

type Environment struct {
	Database database
	Extras   env.EnvSet
}

type database struct {
	URI string `env:"DATABASE_URI"`
}

func (e *Environment) Validate() error {
	errs := make([]error, 0)

	if len(e.Database.URI) == 0 {
		errs = append(errs, fmt.Errorf("DATABASE_URI is empty"))
	}

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

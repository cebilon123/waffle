package rule

import (
	"errors"
	"fmt"
	"strings"
)

const (
	variableExpressionDivider = "=>"
)

var (
	ErrEmptyPredicateName    = errors.New("predicate name is empty")
	ErrEmptyPredicateValue   = errors.New("value to compile predicate is empty")
	ErrMissingPredicateParts = errors.New("predicate parts are missing")
	ErrInvalidPredicate      = errors.New("predicate is invalid")
)

type Compiler interface {
	// Compile compiles rule to the functional predicate.
	// 	Name - is a name of the predicate.
	//  Value - is a value that should be compiled into predicate.
	Compile(name, value string) (*Predicate, error)
}

type CustomCompiler struct {
	builder Builder
}

var _ Compiler = (*CustomCompiler)(nil)

func (c *CustomCompiler) Compile(name, value string) (*Predicate, error) {
	if err := validateInput(name, value); err != nil {
		return nil, fmt.Errorf("validate name and value: %w", err)
	}

	variable, expression, err := getVariableAndLogicalExpression(value)
	if err != nil {
		return nil, fmt.Errorf("get variable and logical expression for predicate: %w", err)
	}

	pred, err := c.builder.Build(name, variable, expression)
	if err != nil {
		return nil, fmt.Errorf("build predicate: %w", err)
	}

	return pred, nil
}

func validateInput(name string, value string) error {
	if len(name) == 0 {
		return ErrEmptyPredicateName
	}

	if len(value) == 0 {
		return ErrEmptyPredicateValue
	}

	return nil
}

func getVariableAndLogicalExpression(value string) (string, string, error) {
	values := strings.Split(value, variableExpressionDivider)

	if len(values) < 2 {
		return "", "", ErrMissingPredicateParts
	}

	if len(values) > 2 {
		return "", "", ErrInvalidPredicate
	}

	return values[0], values[1], nil
}

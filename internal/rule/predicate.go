package rule

import (
	"fmt"
	
	"waffle/internal/request"
)

type ExpressionTreeBuilder interface {
	BuildExpressionTree(variable, expression string) (expressionTree, error)
}

type Builder interface {
	Build(name, variable, expression string) (*Predicate, error)
}

type Predicate struct {
	Name string
	Eval func(r request.Wrapper) bool
}

type predicateBuilder struct {
	treeBuilder ExpressionTreeBuilder
}

func newPredicateBuilder() *predicateBuilder {
	return &predicateBuilder{
		treeBuilder: newExpressionTreeBuilder(),
	}
}

var _ Builder = (*predicateBuilder)(nil)

func (p *predicateBuilder) Build(name, variable, expression string) (*Predicate, error) {
	tree, err := p.treeBuilder.BuildExpressionTree(variable, expression)
	if err != nil {
		return nil, fmt.Errorf("build expression tree: %w", err)
	}

	return &Predicate{
		Name: name,
		Eval: tree.Eval,
	}, nil
}

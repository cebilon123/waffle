package rule

import "fmt"

type expressionTree node

type ExpressionTreeFactory interface {
	CreateExpressionTree(tokens []Token) (expressionTree, error)
}

// Tokenizer should be implemented by the structs that tokenizes input.
type Tokenizer interface {
	// BuildTokens builds tokens based on the variable and expression.
	// Returns error if tokenizer cannot be done with the input.
	BuildTokens(variable string, expression string) ([]Token, error)
}

type expressionTreeBuilder struct {
	tokenizer             Tokenizer
	expressionTreeFactory ExpressionTreeFactory
}

var _ ExpressionTreeBuilder = (*expressionTreeBuilder)(nil)

func newExpressionTreeBuilder() *expressionTreeBuilder {
	return &expressionTreeBuilder{
		tokenizer:             &tokenizer{},
		expressionTreeFactory: &expressionTreeFactory{},
	}
}

func (c *expressionTreeBuilder) BuildExpressionTree(variable, expression string) (expressionTree, error) {
	tokens, err := c.tokenizer.BuildTokens(variable, expression)
	if err != nil {
		return nil, fmt.Errorf("build tokens: %w", err)
	}

	tree, err := c.expressionTreeFactory.CreateExpressionTree(tokens)
	if err != nil {
		return nil, fmt.Errorf("create expression tree: %w", err)
	}

	return tree, nil
}

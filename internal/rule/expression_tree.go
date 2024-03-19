package rule

import "fmt"

type expressionTree node

type ExpressionTreeBuilder interface {
	BuildExpressionTree(variable, expression string) (expressionTree, error)
}

type expressionTreeBuilder struct {
	tokenizer             Tokenizer
	expressionTreeFactory ExpressionTreeFactory
}

var _ ExpressionTreeBuilder = (*expressionTreeBuilder)(nil)

func newCustomRulesTokenizer() *expressionTreeBuilder {
	return &expressionTreeBuilder{
		tokenizer: &tokenizer{},
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

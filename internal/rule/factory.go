package rule

import "github.com/emirpasic/gods/stacks/arraystack"

var (
	// operatorPrecedence is a slice of operators in the precedence order.
	// The most important ones are on the top, the lest important ones are on the bottom.
	// Also if there are two operators: A and B, and if those are in the same row,
	// those are equal in terms of precedence.
	operatorPrecedence = [][]string{
		// Postfix
		{
			tokenLParen,
			tokenRParen,
			tokenDot,
		},
		// Relational
		{
			tokenMoreThan,
			tokenLessThan,
		},
		// Equality
		{
			tokenEqual,
			tokenNotEqual,
		},
		// AND
		{
			tokenDoubleAmpersand,
		},
		// OR
		{
			tokenOr,
		},
		// Dot
		{
			tokenDot,
		},
	}
)

type expressionTreeFactory struct {
}

var _ ExpressionTreeFactory = (*expressionTreeFactory)(nil)

func (e *expressionTreeFactory) CreateExpressionTree(tokens []Token) (expressionTree, error) {

}

func reversePolishNotationSort(tokens []Token) []Token {
	var (
		operatorStack arraystack.Stack
		outputTokens  []Token
	)

	for _, token := range tokens {
		if isOperator(token) {

		}
	}
}

func isOperator(tkn Token) bool {
	for _, tknOperator := range tokensOperators {
		if tkn.Name == tknOperator {
			return true
		}
	}

	return false
}

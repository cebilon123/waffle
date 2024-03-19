package rule

type ExpressionTreeFactory interface {
	CreateExpressionTree(tokens []Token) (expressionTree, error)
}

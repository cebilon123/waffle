package rule

type expressionTreeFactory struct {
}

var _ ExpressionTreeFactory = (*expressionTreeFactory)(nil)

func (e *expressionTreeFactory) CreateExpressionTree(tokens []Token) (expressionTree, error) {
	//TODO implement me
	panic("implement me")
}

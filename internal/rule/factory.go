package rule

import (
	"fmt"
	"reflect"
)

var (
	matches = []match{}

	adjustableNodes = []reflect.Type{
		reflect.TypeOf(gt{}),
	}
)

type NodeAdjuster interface {
	AdjustNode(nd node, tokensMatch []Token) (node, error)
}

type expressionTreeFactory struct {
	nodeAdjuster NodeAdjuster
}

var _ ExpressionTreeFactory = (*expressionTreeFactory)(nil)

func (e *expressionTreeFactory) CreateExpressionTree(tokens []Token) (expressionTree, error) {
	var (
		nodes      []node
		tokenMatch []Token
	)

	for _, token := range tokens {
		tokenMatch = append(tokenMatch, token)

		for _, m := range matches {
			if m.isMatching(tokenMatch) {
				nd, err := e.adaptNode(m.node, tokenMatch)
				if err != nil {
					return nil, fmt.Errorf("adapt node: %w", err)
				}

				nodes = append(nodes, nd)

				tokenMatch = []Token{}
			}
		}
	}

	for _, nd := range nodes {
		
	}

	return nil, nil
}

func (e *expressionTreeFactory) adaptNode(base node, tokenMatch []Token) (node, error) {
	if !isAdjustableNode(base) {
		return base, nil
	}

	nd, err := e.nodeAdjuster.AdjustNode(base, tokenMatch)
	if err != nil {
		return nil, fmt.Errorf("adjust node %s: %w", reflect.TypeOf(base).Name(), err)
	}

	return nd, nil
}

type match struct {
	tokens []Token
	node   node
}

func (m *match) isMatching(tokens []Token) bool {
	if len(m.tokens) != len(tokens) {
		return false
	}

	for i, token := range m.tokens {
		if tokens[i].Name != token.Name {
			return false
		}
	}

	return true
}

func isAdjustableNode(nd node) bool {
	nodeType := reflect.TypeOf(nd)

	for _, adjustableNode := range adjustableNodes {
		if nodeType.Name() == adjustableNode.Name() {
			return true
		}
	}

	return false
}

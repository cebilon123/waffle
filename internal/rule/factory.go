package rule

import (
	"errors"
	"fmt"
	"reflect"
	
	"github.com/emirpasic/gods/stacks/arraystack"
)

var (
	matches = []match{}

	adjustableNodes = []reflect.Type{
		reflect.TypeOf(gt{}),
	}

	higherOrderNodes = []reflect.Type{
		reflect.TypeOf(and{}),
		reflect.TypeOf(or{}),
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

	eTree, err := buildExpressionTree(nodes)
	if err != nil {
		return nil, fmt.Errorf("build expression tree: %w", err)
	}

	return eTree, nil
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

func buildExpressionTree(nodes []node) (expressionTree, error) {
	stack := arraystack.Stack{}

	// todo: we need to add additional check if less than 2 then we need to validate if NEXT node is not a higherOrderNode
	// and then we need to build based on that
	for i := 0; i < len(nodes); i++ {
		switch nodes[i].(type) {
		case and:
			// after nth iteration it could be an and, we want then continue
			// in order to add additional elements to the stack
			if stack.Size() < 2 {
				continue
			}

			if stack.Size() > 2 {
				return nil, errors.New("expression is invalid, and can only evaluate 2 predicates")
			}

			nd := and{}

			var childrenNodes []node

			for j := 0; j < 2; j++ {
				v, _ := stack.Pop()
				vt, _ := v.(node)
				childrenNodes = append(childrenNodes, vt)
			}

			// nodes are popped from the stack, therefore we need to change order of the children
			nd.SetChild(childrenNodes[1], childrenNodes[0])

			stack.Clear()

			stack.Push(nd)
		case or:
			if stack.Size() < 2 {
				continue
			}

			if stack.Size() > 2 {
				return nil, errors.New("expression is invalid, and can only evaluate 2 predicates")
			}

			nd := or{}

			var childrenNodes []node

			for j := 0; j < 2; j++ {
				v, _ := stack.Pop()
				vt, _ := v.(node)
				childrenNodes = append(childrenNodes, vt)
			}

			nd.SetChild(childrenNodes[1], childrenNodes[0])

			stack.Clear()

			stack.Push(nd)
		default:
			stack.Push(nodes[i])
		}
	}

	v, ok := stack.Pop()
	if !ok {
		return nil, fmt.Errorf("empty expression stack")
	}

	vt, _ := v.(node)

	return expressionTree(vt), nil
}

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
	return nil, nil
}

// 1.  While there are tokens to be read:
// 2.        Read a token
// 3.        If it's a number add it to queue
// 4.        If it's an operator
// 5.               While there's an operator on the top of the stack with greater precedence:
// 6.                       Pop operators from the stack onto the output queue
// 7.               Push the current operator onto the stack
// 8.        If it's a left bracket push it onto the stack
// 9.        If it's a right bracket
// 10.            While there's not a left bracket at the top of the stack:
// 11.                     Pop operators from the stack onto the output queue.
// 12.             Pop the left bracket from the stack and discard it
// 13. While there are operators on the stack, pop them to the queue

func reversePolishNotationSort(tokens []Token) []Token {
	var (
		operatorStack arraystack.Stack
		outputTokens  []Token
	)

	for _, token := range tokens {
		if isOperator(token) {
			if v, ok := operatorStack.Peek(); ok {
				vt := v.(Token)
				if isTokenWithGreaterPrecedenceFromThePrecedenceSlice(token, vt) {
					for !operatorStack.Empty() {
						v, ok := operatorStack.Pop()
						if !ok {
							continue
						}

						vt := v.(Token)

						outputTokens = append(outputTokens, vt)
					}

					operatorStack.Push(token)
				}
			}

			continue
		}

		if token.Name == tokenLParen {
			operatorStack.Push(token)

			continue
		}

		if token.Name == tokenRParen {
			var lParen *Token
			for !operatorStack.Empty() || lParen == nil {
				v, _ := operatorStack.Pop()

				vt := v.(Token)

				if vt.Name == tokenLParen {
					lParen = &vt
					continue
				}

				outputTokens = append(outputTokens, vt)
			}

			continue
		}

		tokens = append(tokens, token)
	}

	return tokens
}

func isOperator(tkn Token) bool {
	for _, tknOperator := range tokensOperators {
		if tkn.Name == tknOperator {
			return true
		}
	}

	return false
}

func isTokenWithGreaterPrecedenceFromThePrecedenceSlice(tkn Token, lastTknFromStack Token) bool {
	var (
		tknIdx              int
		lastTknFromStackIdx int
	)

	for i, op := range operatorPrecedence {
		for _, opp := range op {
			if opp == tkn.Name {
				tknIdx = i
			}
			if opp == lastTknFromStack.Name {
				lastTknFromStackIdx = i
			}
		}
	}

	return lastTknFromStackIdx < tknIdx
}

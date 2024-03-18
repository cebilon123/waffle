package rule

import "fmt"

var (
	tokenVariable        = "var"
	tokenFunction        = "func"
	tokenLParen          = "("
	tokenRParen          = ")"
	tokenDot             = "."
	tokenDoubleAmpersand = "&&"
	tokenMoreThan        = ">"
	tokenLessThan        = "<"
	tokenNumber          = "token_number"
	tokenField           = "field"

	methodLen    = "LEN"
	methodFormat = "FORMAT"
	methods      = []string{
		methodLen,
		methodFormat,
	}
)

type Tokenizer interface {
	BuildTokens(variable string, expression string) ([]Token, error)
}

type Token struct {
	Name  string
	Value string
}

type tokenizer struct {
}

func (t *tokenizer) BuildTokens(variable string, expression string) ([]Token, error) {
	var (
		tokens []Token
		match  []rune
	)

	runeExpression := []rune(expression)

	for i, r := range runeExpression {
		match = append(match, r)
		strMatch := string(match)

		v, ok := getFunction(match)
		if ok {
			tokens = append(tokens, Token{
				Name:  tokenFunction,
				Value: v,
			})

			match = []rune{}
			continue
		}

		if isVariable(match, variable) && len(expression) >= i+1 && runeExpression[i+1] == '.' {
			tokens = append(tokens, Token{
				Name:  tokenVariable,
				Value: string(match),
			})

			match = []rune{}
			continue
		}

		if strMatch == "." {
			tokens = append(tokens, Token{
				Name:  tokenDot,
				Value: tokenDot,
			})

			match = []rune{}
			continue
		}

	}

	return tokens, nil
}

func getFunction(match []rune) (string, bool) {
	strMatch := string(match)

	for _, method := range methods {
		if method == strMatch {
			return method, true
		}
	}

	return "", false
}

func isVariable(match []rune, variable string) bool {
	strMatch := string(match)

	return strMatch == fmt.Sprintf("%s", variable)
}

func getSpecialCharacter()

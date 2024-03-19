package rule

import (
	"fmt"
	"unicode"
)

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
	tokenSpace           = "space"
	specialCharacters    = []string{
		tokenLParen,
		tokenRParen,
		tokenDot,
		tokenDoubleAmpersand,
		tokenMoreThan,
		tokenLessThan,
	}

	methodLen    = "LEN"
	methodFormat = "FORMAT"
	methods      = []string{
		methodLen,
		methodFormat,
	}

	fieldPayload = "payload"
	fieldHeaders = "headers"
	fields       = []string{
		fieldPayload,
		fieldHeaders,
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

		if v, ok := getSpecialCharacter(match); ok {
			tokens = append(tokens, Token{
				Name:  v,
				Value: v,
			})

			match = []rune{}
			continue
		}

		if v, ok := getField(match); ok {
			tokens = append(tokens, Token{
				Name:  tokenField,
				Value: v,
			})

			match = []rune{}
			continue
		}

		if unicode.IsSpace(match[0]) {
			tokens = append(tokens, Token{
				Name:  tokenSpace,
				Value: tokenSpace,
			})

			match = []rune{}
			continue
		}

		if unicode.IsNumber(match[0]) {
			tokens = append(tokens, Token{
				Name:  tokenNumber,
				Value: string(match[0]),
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

func getSpecialCharacter(match []rune) (string, bool) {
	strMatch := string(match)

	for _, sch := range specialCharacters {
		if sch == strMatch {
			return sch, true
		}
	}

	return "", false
}

func getField(match []rune) (string, bool) {
	strMatch := string(match)

	for _, field := range fields {
		if field == strMatch {
			return field, true
		}
	}

	return "", false
}

package rule

import (
	"reflect"
	"testing"
)

func Test_isOperator(t *testing.T) {
	type args struct {
		tkn Token
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is an operator, true returned",
			args: args{
				tkn: Token{
					Name: tokenMoreThan,
				},
			},
			want: true,
		},
		{
			name: "isn't an operator, false returned",
			args: args{
				tkn: Token{
					Name: tokenLParen,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOperator(tt.args.tkn); got != tt.want {
				t.Errorf("isOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTokenWithGreaterPrecedenceFromThePrecedenceSlices(t *testing.T) {
	type args struct {
		tkn              Token
		lastTknFromStack Token
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "token with greater precedence, true returned",
			args: args{
				tkn: Token{
					Name:  tokenOr,
					Value: tokenOr,
				},
				lastTknFromStack: Token{
					Name:  tokenLParen,
					Value: tokenLParen,
				},
			},
			want: true,
		},
		{
			name: "token with less precedence, false returned",
			args: args{
				tkn: Token{
					Name:  tokenLParen,
					Value: tokenLParen,
				},
				lastTknFromStack: Token{
					Name:  tokenOr,
					Value: tokenOr,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTokenWithGreaterPrecedenceFromThePrecedenceSlice(tt.args.tkn, tt.args.lastTknFromStack); got != tt.want {
				t.Errorf("isTokenWithGreaterPrecedenceFromThePrecedenceSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reversePolishNotationSort(t *testing.T) {
	type args struct {
		tokens []Token
	}
	tests := []struct {
		name string
		args args
		want []Token
	}{
		{
			name: "tokens should be ordered",
			args: args{
				tokens: []Token{
					{
						Name:  tokenFunction,
						Value: "LEN",
					},
					{
						Name:  tokenLParen,
						Value: tokenLParen,
					},
					{
						Name:  tokenVariable,
						Value: "p",
					},
					{
						Name:  tokenDot,
						Value: ".",
					},
					{
						Name:  tokenField,
						Value: "headers",
					},
					{
						Name:  tokenRParen,
						Value: tokenRParen,
					},
					{
						Name:  tokenSpace,
						Value: tokenSpace,
					},
					{
						Name:  tokenMoreThan,
						Value: tokenMoreThan,
					},
					{
						Name:  tokenSpace,
						Value: tokenSpace,
					},
					{
						Name:  tokenNumber,
						Value: "5",
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePolishNotationSort(tt.args.tokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePolishNotationSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

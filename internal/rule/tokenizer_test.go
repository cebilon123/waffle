package rule

import (
	"reflect"
	"testing"
)

func Test_tokenizer_BuildTokens(t1 *testing.T) {
	type args struct {
		variable   string
		expression string
	}
	tests := []struct {
		name    string
		args    args
		want    []Token
		wantErr bool
	}{
		{
			name: "slice of tokens returned",
			args: args{
				variable:   "p",
				expression: "LEN(p.payload) > 0 && LEN(p.headers) > 0",
			},
			want: []Token{
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
					Value: tokenDot,
				},
				{
					Name:  tokenField,
					Value: "payload",
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
					Value: ">",
				},
				{
					Name:  tokenSpace,
					Value: tokenSpace,
				},
				{
					Name:  tokenNumber,
					Value: "0",
				},
				{
					Name:  tokenSpace,
					Value: tokenSpace,
				},
				{
					Name:  tokenDoubleAmpersand,
					Value: "&&",
				},
				{
					Name:  tokenSpace,
					Value: tokenSpace,
				},
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
					Value: tokenDot,
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
					Value: ">",
				},
				{
					Name:  tokenSpace,
					Value: tokenSpace,
				},
				{
					Name:  tokenNumber,
					Value: "0",
				},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &tokenizer{}
			got, err := t.BuildTokens(tt.args.variable, tt.args.expression)
			if (err != nil) != tt.wantErr {
				t1.Errorf("BuildTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("BuildTokens() got = %v, want %v", got, tt.want)
			}
		})
	}
}

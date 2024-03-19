package rule

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_expressionTreeBuilder_BuildExpressionTree(t *testing.T) {
	type fields struct {
		tokenizerFunc             func() Tokenizer
		expressionTreeFactoryFunc func() ExpressionTreeFactory
	}
	type args struct {
		variable   string
		expression string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    expressionTree
		wantErr bool
	}{
		{
			name: "tokenizer returns error, error returned",
			fields: fields{
				tokenizerFunc: func() Tokenizer {
					tkn := NewMockTokenizer(t)

					tkn.EXPECT().
						BuildTokens(mock.Anything, mock.Anything).
						Return(nil, errors.New("error"))

					return tkn
				},
			},
			wantErr: true,
		},
		{
			name: "expression tree factory returns error, error returned",
			fields: fields{
				tokenizerFunc: func() Tokenizer {
					tkn := NewMockTokenizer(t)

					tkn.EXPECT().
						BuildTokens(mock.Anything, mock.Anything).
						Return(testTokens, nil)

					return tkn
				},
				expressionTreeFactoryFunc: func() ExpressionTreeFactory {
					etf := NewMockExpressionTreeFactory(t)

					etf.EXPECT().
						CreateExpressionTree(mock.Anything).
						Return(nil, errors.New("error"))

					return etf
				},
			},
			wantErr: true,
		},
		{
			name: "expression tree factory returns expression tree, tree returned",
			fields: fields{
				tokenizerFunc: func() Tokenizer {
					tkn := NewMockTokenizer(t)

					tkn.EXPECT().
						BuildTokens(mock.Anything, mock.Anything).
						Return(testTokens, nil)

					return tkn
				},
				expressionTreeFactoryFunc: func() ExpressionTreeFactory {
					etf := NewMockExpressionTreeFactory(t)

					etf.EXPECT().
						CreateExpressionTree(testTokens).
						Return(and{}, nil)

					return etf
				},
			},
			wantErr: false,
			want:    and{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expressionTreeFactory ExpressionTreeFactory
			if tt.fields.expressionTreeFactoryFunc != nil {
				expressionTreeFactory = tt.fields.expressionTreeFactoryFunc()
			}

			var tokenizer Tokenizer
			if tt.fields.tokenizerFunc != nil {
				tokenizer = tt.fields.tokenizerFunc()
			}

			c := &expressionTreeBuilder{
				tokenizer:             tokenizer,
				expressionTreeFactory: expressionTreeFactory,
			}
			got, err := c.BuildExpressionTree(tt.args.variable, tt.args.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildExpressionTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildExpressionTree() got = %v, want %v", got, tt.want)
			}
		})
	}
}

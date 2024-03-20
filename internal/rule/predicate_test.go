package rule

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_predicateBuilder_Build(t *testing.T) {
	andNode := &and{}

	type fields struct {
		treeBuilderFunc func() ExpressionTreeBuilder
	}
	type args struct {
		name       string
		variable   string
		expression string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantPredicate bool
		wantErr       bool
	}{
		{
			name: "builder returns error, error returned",
			fields: fields{
				treeBuilderFunc: func() ExpressionTreeBuilder {
					etb := NewMockExpressionTreeBuilder(t)

					etb.EXPECT().BuildExpressionTree(mock.Anything, mock.Anything).Return(nil, errors.New("error"))

					return etb
				},
			},
			wantErr: true,
		},
		{
			name: "builder returns tree node, predicate returned",
			fields: fields{
				treeBuilderFunc: func() ExpressionTreeBuilder {
					etb := NewMockExpressionTreeBuilder(t)

					etb.EXPECT().
						BuildExpressionTree(mock.Anything, mock.Anything).
						Return(andNode, nil)

					return etb
				},
			},
			args: args{
				name:       "test is test",
				variable:   "test",
				expression: "test => LEN(test.payload) > 0",
			},
			wantPredicate: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var treeBuilder ExpressionTreeBuilder
			if tt.fields.treeBuilderFunc != nil {
				treeBuilder = tt.fields.treeBuilderFunc()
			}

			p := &predicateBuilder{
				treeBuilder: treeBuilder,
			}
			got, err := p.Build(tt.args.name, tt.args.variable, tt.args.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil) != tt.wantPredicate {
				t.Errorf("Build() got = %v, want %v", got, tt.wantPredicate)
			}
		})
	}
}

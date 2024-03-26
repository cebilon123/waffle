package rule

import "testing"

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

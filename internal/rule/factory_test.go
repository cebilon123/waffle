package rule

import "testing"

func Test_isAdjustableNode(t *testing.T) {
	type args struct {
		nd node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "type isn't adjustable, false returned",
			args: args{
				nd: and{},
			},
			want: false,
		},
		{
			name: "type is adjustable, true returned",
			args: args{
				nd: gt{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdjustableNode(tt.args.nd); got != tt.want {
				t.Errorf("isAdjustableNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

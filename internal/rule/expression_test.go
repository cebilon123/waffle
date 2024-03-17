package rule

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"waffle/internal/request"
)

func Test_and_Eval(t *testing.T) {
	type fields struct {
		leftFunc  func() node
		rightFunc func() node
	}
	type args struct {
		r request.Wrapper
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "both nodes true, true returned",
			fields: fields{
				leftFunc: func() node {
					n := NewMocknode(t)

					n.EXPECT().Eval(mock.Anything).Return(true)

					return n
				},
				rightFunc: func() node {
					n := NewMocknode(t)

					n.EXPECT().Eval(mock.Anything).Return(true)

					return n
				},
			},
			want: true,
		},
		{
			name: "one of nodes false, false returned",
			fields: fields{
				leftFunc: func() node {
					n := NewMocknode(t)

					n.EXPECT().Eval(mock.Anything).Return(true)

					return n
				},
				rightFunc: func() node {
					n := NewMocknode(t)

					n.EXPECT().Eval(mock.Anything).Return(false)

					return n
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var left, right node
			if tt.fields.leftFunc != nil {
				left = tt.fields.leftFunc()
			}
			if tt.fields.rightFunc != nil {
				right = tt.fields.rightFunc()
			}

			a := and{
				left:  left,
				right: right,
			}
			if got := a.Eval(tt.args.r); got != tt.want {
				t.Errorf("Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_or_Eval(t *testing.T) {
	type fields struct {
		leftFunc  func() node
		rightFunc func() node
	}
	type args struct {
		r request.Wrapper
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "first node true, returns true",
			fields: fields{
				leftFunc: func() node {
					n := NewMocknode(t)

					n.EXPECT().Eval(mock.Anything).Return(true)

					return n
				},
			},
			want: true,
		},
		{
			name: "both nodes false, returns false",
			fields: fields{
				leftFunc: func() node {
					n := NewMocknode(t)

					n.EXPECT().Eval(mock.Anything).Return(false)

					return n
				},
				rightFunc: func() node {
					n := NewMocknode(t)

					n.EXPECT().Eval(mock.Anything).Return(false)

					return n
				},
			},
			want: false,
		},
		{
			name: "first node false, second true, returns true",
			fields: fields{
				leftFunc: func() node {
					n := NewMocknode(t)

					n.EXPECT().Eval(mock.Anything).Return(false)

					return n
				},
				rightFunc: func() node {
					n := NewMocknode(t)

					n.EXPECT().Eval(mock.Anything).Return(true)

					return n
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var left, right node
			if tt.fields.leftFunc != nil {
				left = tt.fields.leftFunc()
			}
			if tt.fields.rightFunc != nil {
				right = tt.fields.rightFunc()
			}

			o := or{
				left:  left,
				right: right,
			}
			if got := o.Eval(tt.args.r); got != tt.want {
				t.Errorf("Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}

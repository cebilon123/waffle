package rule

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestCustomCompiler_Compile(t *testing.T) {
	type fields struct {
		builderFunc func() Builder
	}
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Predicate
		wantErr bool
	}{
		{
			name: "empty name, error returned",
			args: args{
				name:  "",
				value: "p => p",
			},
			wantErr: true,
		},
		{
			name: "empty value, error returned",
			args: args{
				name:  "name",
				value: "",
			},
			wantErr: true,
		},
		{
			name: "missing => sign, error returned",
			args: args{
				name:  "name",
				value: "p p.length() > 0",
			},
			wantErr: true,
		},
		{
			name: "to many => signs, error returned",
			args: args{
				name:  "name",
				value: "p => p.format() => json",
			},
			wantErr: true,
		},
		{
			name: "builder returns error, error returned",
			fields: fields{
				builderFunc: func() Builder {
					b := NewMockBuilder(t)

					b.EXPECT().Build(mock.Anything, mock.Anything, mock.Anything).
						Return(nil, errors.New("err"))

					return b
				},
			},
			args: args{
				name:  "name",
				value: "p => p.length() > 0",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder Builder
			if tt.fields.builderFunc != nil {
				builder = tt.fields.builderFunc()
			}

			c := &CustomCompiler{
				builder: builder,
			}
			got, err := c.Compile(tt.args.name, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

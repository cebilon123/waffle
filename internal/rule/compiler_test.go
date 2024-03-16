package rule

import (
	"reflect"
	"testing"
)

func TestCustomCompiler_Compile(t *testing.T) {
	type fields struct {
		builder Builder
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CustomCompiler{
				builder: tt.fields.builder,
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

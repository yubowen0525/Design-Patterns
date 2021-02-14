package Builder_test

import (
	"fmt"
	"my_design_pattern/Creational_Patterns/Builder"
	"reflect"
	"testing"
)

func TestNewResourcePoolConfig(t *testing.T) {
	type args struct {
		name string
		opts []Builder.ResourcePoolConfigOptFunc
	}

	tests := []struct {
		name    string
		args    args
		want    *Builder.ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			args: args{
				name: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				name: "test",
				opts: []Builder.ResourcePoolConfigOptFunc{
					func(option *Builder.ResourcePoolConfigOption) {
						option.MinIdle = 2
						option.MaxIdle = 5
					},
				},
			},
			want: &Builder.ResourcePoolConfig{
				Name:     "test",
				MaxTotal: 10,
				MaxIdle:  5,
				MinIdle:  2,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Builder.NewResourcePoolConfig(tt.args.name, tt.args.opts...)
			if err != nil {
				t.Fatal(err.Error())
			}
			fmt.Println("got:", got)
			fmt.Println("tt.want:", tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error")
			}
		})
	}
}

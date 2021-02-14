package Builder_test

import (
	"my_design_pattern/Creational_Patterns/Builder"
	"reflect"
	"testing"
)

func TestResourcePoolConfigBuilder_Build(t *testing.T) {
	tests := []struct {
		name    string
		builder *Builder.ResourcePoolConfigBuilder
		want    *Builder.ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			builder: &Builder.ResourcePoolConfigBuilder{
				Name:     "",
				MaxTotal: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "maxIdle < minIdle",
			builder: &Builder.ResourcePoolConfigBuilder{
				Name:     "test",
				MaxTotal: 0,
				MaxIdle:  10,
				MinIdle:  20,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			builder: &Builder.ResourcePoolConfigBuilder{
				Name: "test",
			},
			want: &Builder.ResourcePoolConfig{
				Name:     "test",
				MaxTotal: Builder.DefaultMaxTotal,
				MaxIdle:  Builder.DefaultMaxIdle,
				MinIdle:  Builder.DefaultMinIdle,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.builder.Build()
			if err != nil {
				t.Fatal(err.Error())
			}
			//fmt.Println("got:",got)
			//fmt.Println("tt.want:",tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error ")
			}

		})
	}
}

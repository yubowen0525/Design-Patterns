package factory_test

import (
	factory "my_design_pattern/Creational_Patterns/Factory"
	"reflect"
	"testing"
)

func TestNewIRuleConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want factory.IRuleConfigParserFactory
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: factory.JsonRuleConfigParserFactory{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: factory.YamlRuleConfigParserFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factory.NewIRuleConfigParserFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

package factory_test

import (
	factory "my_design_pattern/Creational_Patterns/Factory"
	"testing"
)

func TestImplA_Say(t *testing.T) {
	api := factory.NewApi("A")
	if api == nil {
		t.Fatal("error api is nil")
	}
	s := api.Say("Tom")
	if s != "Hello ImplA Tom" {
		t.Fatal("TestA fail")
	}
}

func TestImplB_Say(t *testing.T) {
	api := factory.NewApi("B")
	if api == nil {
		t.Fatal("error api is nil")
	}
	s := api.Say("Tom")
	if s != "Hello ImplB Tom" {
		t.Fatal("TestB fail")
	}
}

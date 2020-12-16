package SimpleFactory_test

import (
	"my_design_pattern/Creational_Patterns/SimpleFactory"
	"testing"
)

func TestImplA_Say(t *testing.T) {
	api := SimpleFactory.NewApi("A")
	if api == nil {
		t.Fatal("error api is nil")
	}
	s := api.Say("Tom")
	if s != "Hello ImplA Tom" {
		t.Fatal("TestA fail")
	}
}

func TestImplB_Say(t *testing.T) {
	api := SimpleFactory.NewApi("B")
	if api == nil {
		t.Fatal("error api is nil")
	}
	s := api.Say("Tom")
	if s != "Hello ImplB Tom" {
		t.Fatal("TestB fail")
	}
}

package SimpleFactory

import "testing"

func TestImplA_Say(t *testing.T) {
	api := NewApi("A")
	if api == nil {
		t.Fatal("error api is nil")
	}
	s := api.Say("Tom")
	if s != "Hello ImplA Tom" {
		t.Fatal("TestA fail")
	}
}

func TestImplB_Say(t *testing.T) {
	api := NewApi("B")
	if api == nil {
		t.Fatal("error api is nil")
	}
	s := api.Say("Tom")
	if s != "Hello ImplB Tom" {
		t.Fatal("TestB fail")
	}
}

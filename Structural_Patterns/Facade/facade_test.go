package Facade_test

import (
	"my_design_pattern/Structural_Patterns/Facade"
	"testing"
)

var expect = "a module test\nb module test"

func TestGetAPI(t *testing.T) {
	content := Facade.GetAPI().Test()
	if content != expect {
		t.Fatalf("expect=%s,but content=%s", expect, content)
	}
}

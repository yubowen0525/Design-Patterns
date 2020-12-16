package Facade

import "testing"

var expect = "a module test\nb module test"

func TestGetAPI(t *testing.T) {
	content := GetAPI().Test()
	if content != expect {
		t.Fatalf("expect=%s,but content=%s", expect, content)
	}
}

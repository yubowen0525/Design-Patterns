package Adapter_test

import (
	"my_design_pattern/Structural_Patterns/Adapter"
	"testing"
)

var test string = "adaptee method"

func TestAdapter_Request(t *testing.T) {
	// 原始客户端
	adaptee := Adapter.CreateAdaptee()
	// 原始接口
	content := adaptee.SpecificRequest()
	if content != test {
		t.Fatal("adaptee class not true")
	}

	// 转接器客户端
	NewAdaptee := Adapter.CreateAdaptee()
	target := Adapter.CreateAdapter(NewAdaptee)
	// 接口已经改变，但底层实现未变，依然调用的是adaptee接口
	content = target.Request()
	if content != test {
		t.Fatal("adapter class not true")
	}

}

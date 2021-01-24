package Singleton_test

import (
	"my_design_pattern/Creational_Patterns/Singleton"
	"testing"
)

func TestGetInstance(t *testing.T) {
	a := Singleton.GetLazyInstance()
	b := Singleton.GetLazyInstance()

	if a != b {
		t.Fatal("Not the same Instance")
	}

}

// 并行测试，一定范围时间下，该语句能执行多少次
// 此方法符合该语句执行环境下的预期
// 并行的goroutine个数是默认等于runtime.GOMAXPROCS(0)。
func BenchmarkGetLazyInstanceInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if Singleton.GetLazyInstance() != Singleton.GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

// 串行测试，一定范围时间下(1s)，该语句能执行的次数
// b.N就是最大执行次数
func BenchmarkGetLazyInstanceInstanceSerial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if Singleton.GetLazyInstance() != Singleton.GetLazyInstance() {
			b.Errorf("test fail")
		}
	}
}

func BenchmarkGetHungryInstanceInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if Singleton.GetHungryInstance() != Singleton.GetHungryInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

func BenchmarkGetHungryInstanceInstanceSerial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if Singleton.GetHungryInstance() != Singleton.GetHungryInstance() {
			b.Errorf("test fail")
		}
	}
}

package Singleton

import "sync"

type Singleton2 struct{}

var singleton2 *Singleton2
var once sync.Once

//GetInstance 获取对象实例
// 懒汉式
// 延迟加载
func GetLazyInstance() *Singleton2 {
	once.Do(func() {
		singleton2 = &Singleton2{}
	})

	return singleton2
}

package Singleton

type Singleton1 struct{}

var singleton1 *Singleton1

func init() {
	singleton1 = &Singleton1{}
}

func GetHungryInstance() *Singleton1 {
	return singleton1
}

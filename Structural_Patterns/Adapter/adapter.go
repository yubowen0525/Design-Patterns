package Adapter

//Target是适配的接口
type Target interface {
	Request() string
}

// Adaptee是被适配的的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// Adaptee实现
type AdapteeImpl struct {
}

func (a *AdapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// Adaptee工厂函数
func CreateAdaptee() Adaptee {
	return &AdapteeImpl{}
}

// Adapter的工厂函数
// Adaptee 为不适配Target接口的类
func CreateAdapter(adaptee Adaptee) Target {
	return &adapter{
		adaptee: adaptee,
	}
}

// Adapter是转换Adaptee适配Target接口的适配器
type adapter struct {
	adaptee Adaptee
}

func (a *adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

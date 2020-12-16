package Facade

import "fmt"

// factory of facade API create
func GetAPI() API {
	return &facadeAPIImpl{
		a: GetAModule(),
		b: GetBModule(),
	}
}

// facade API
type API interface {
	Test() string
}

type facadeAPIImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (f *facadeAPIImpl) Test() string {
	aRet := f.a.TestA()
	bRet := f.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// a module
// a module factory
func GetAModule() AModuleAPI {
	return &aModuleImpl{}
}

// a module interface
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func (a *aModuleImpl) TestA() string {
	return "a module test"
}

// b module
// b module factory
func GetBModule() BModuleAPI {
	return &bModuleImpl{}
}

// b module interface
type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (b *bModuleImpl) TestB() string {
	return "b module test"
}

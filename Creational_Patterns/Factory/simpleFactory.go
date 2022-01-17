package factory

import "fmt"

// API is interface
type API interface {
	Say(name string) string
}

//A implement
type ImplA struct{}

func (*ImplA) Say(name string) string {

	return fmt.Sprintf("Hello ImplA %s", name)
}

//B implement
type ImplB struct{}

func (*ImplB) Say(name string) string {
	return fmt.Sprintf("Hello ImplB %s", name)
}

// factory
// type A or B
func NewApi(t string) API {
	if t == "A" {
		return &ImplA{}
	} else if t == "B" {
		return &ImplB{}
	}
	return nil
}

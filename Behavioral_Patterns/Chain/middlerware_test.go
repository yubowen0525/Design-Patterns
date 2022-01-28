package chain

import (
	"fmt"
	"testing"
)

type engine struct {
	Handlers []HandlerFunc
}

type Context struct{}

type HandlerFunc func(*Context)

func (e *engine) Use(middleware ...HandlerFunc) {
	e.Handlers = append(e.Handlers, middleware...)
}

// 工厂函数，生成 HandlerFunc
func ErrorHandler() HandlerFunc {
	return func(c *Context) {
		fmt.Println("test handler")
	}
}

// 工厂函数，生成 HandlerFunc
func LogHandler() HandlerFunc {
	return func(c *Context) {
		fmt.Println("log handler")
	}
}

func TestMiddleware(t *testing.T) {
	r := engine{}

	r.Use(ErrorHandler(), LogHandler())

	c := Context{}

	for i := 0; i < len(r.Handlers); i++ {
		r.Handlers[i](&c)
	}
}

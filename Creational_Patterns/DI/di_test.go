package di_test

import (
	"fmt"
	di "my_design_pattern/Creational_Patterns/DI"
	"testing"

	"go.uber.org/dig"
)

func TestDig(t *testing.T) {
	container := dig.New()

	container.Provide(di.NewRepo)
	container.Provide(di.NewService)

	container.Invoke(func(s *di.Service) {
		u, err := s.R.Get(1)
		fmt.Printf("user %+v err %+v", u, err)
	})
}

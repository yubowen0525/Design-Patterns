package wi

import (
	"fmt"
	"log"
	"testing"
)

func TestWire(t *testing.T) {
	s, err := InitApp()
	if err != nil {
		log.Fatal(err)
		return
	}
	u, err := s.R.Get(1)
	fmt.Printf("%+v,+%v", u, err)
}

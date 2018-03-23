package http

import (
	"fmt"
	"testing"
)

func TestSpecies(t *testing.T) {
	s, _ := NewService2()
	ret := s.Species("aaa")
	fmt.Println(ret)
}

func TestHomeworld(t *testing.T) {
	s, _ := NewService2()
	ret := s.Homeworld("aaa")
	fmt.Println(ret)
}

func TestAllHomeworld(t *testing.T) {
	s, _ := NewService2()
	ret := s.AllHomeworld()
	fmt.Println(ret)
}

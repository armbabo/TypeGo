package model

import "fmt"

type Cat struct {
	Name string
}

func (cat Cat) AnyObject() {
	panic("implement me")
}

func (cat Cat) Sua() {
	fmt.Println("meo meo")
}

func (cat Cat) Chay() {
	panic("implement me")
}

func (cat Cat)PrintOld() string {
	return cat.Name
}

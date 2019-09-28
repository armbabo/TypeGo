package model

import "fmt"

type Dog struct {
	Old int
}

func (dog Dog) Sua() {
	fmt.Println("gâu gâu")
}

func (dog Dog) Chay() {
	panic("implement me")
}

func (dog *Dog)PrintOld() int {
	return dog.Old
}



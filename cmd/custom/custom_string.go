package custom

import "fmt"

type StringUltimate []string

func (c *StringUltimate) Print() {

}

func (c *StringUltimate) ToArray() *[]string {
	fmt.Println("====================================")
	fmt.Printf("Address Array Origin: %p \n",c)
	a := []string(*c)
	fmt.Printf("%p",a)
	return &a
}

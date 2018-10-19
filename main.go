package main

import "fmt"

type Animal interface {
	Speak() string
}
type Cat struct{}

func (c Cat) Speak() string {
	return "cat"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "dog"
}
func Test(params interface{}) {
	fmt.Println(params)
}
func main() {
	animals := []Animal{Cat{}, Dog{}}
	for _, animal := range animals {
		fmt.Println("speak %s", animal.Speak)
	}
	Test("String")
	Test(123)
	Test(true)
	names := []string{"stanely", "david", "oscar"}
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)

}
func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

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
		fmt.Println(animal.Speak())
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
	fmt.Println(Add(1, 2))

}
func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}
func Add(x, y int) int {
	return x + y
}

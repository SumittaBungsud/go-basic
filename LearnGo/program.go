// package main

import (
	"fmt"
	"gobasic/greeting"
)

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("My name is ", p.Name)
}

func main() {
	num, name := 23, "sumitta"
	fmt.Printf("Hello Go Programming! %v %v \n", num, name)

	numbers := []int{100, 200, 300, 400, 500, 600}
	fmt.Println(len(numbers))

	for i, v := range numbers {
		fmt.Println("index = ", i, " value = ", v)
	}

	a := new(Android)
	a.Name = "Sumitta"
	a.Talk()
	fmt.Println(greeting.Hello("sumitta"))
}

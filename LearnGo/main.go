package main

import (
	"fmt"
)

type Speaker interface{
	Speak() string
}

type Dog struct{
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Walk() string {
	return "Walk!"
}

func makeSound(s Speaker){
	fmt.Println(s.Speak())
}

func main(){
	dog := Dog{Name: "Lilly"}

	fmt.Print(dog.Speak())
	fmt.Print(dog.Walk())
	makeSound(dog)
}

// nodemon --exec go run main.go --signal SIGTERM
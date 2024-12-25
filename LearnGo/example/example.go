package example

import "fmt"

func intro(){
	fmt.Println("Good morning")
}

func SayHi(name string) {
	intro()
	fmt.Println("Hi "+name)
}
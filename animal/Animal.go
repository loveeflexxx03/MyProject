package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Printf("%s издает какой-то звук\n", a.Name)
}

type Dog struct {
	Animal
}

func (d Dog) Bark() {
	fmt.Printf("%s гавкает\n", d.Name)
}

func main() {
	dog := Dog{Animal{Name: "Rex"}}

	dog.Speak() // Rex издает какой-то звук — метод унаследован от Animal
	dog.Bark()  // Rex гавкает
}

package main

import "fmt"

type Pessoa struct {
	Name     string
	LastName string
	Age      int8
}

func main() {
	list := List{}

	junior := Pessoa{"Júnior", "Sanson", 25}
	holices := Pessoa{"Holices", "Sanson", 45}
	grasi := Pessoa{"Grasi", "de Lima", 45}
	joao := Pessoa{"João", "Sanson", 21}
	pedro := Pessoa{"Pedro", "Sanson", 17}

	list.Append(junior)
	list.Append(holices)
	list.Append(grasi)
	list.Append(joao)
	list.Append(pedro)

	list.Display()

	fmt.Println("--------------------")

	pessoa := list.Search("Grasi")
	fmt.Println(pessoa)

	fmt.Println("--------------------")

	list.Delete("Holices")
	list.Display()
}
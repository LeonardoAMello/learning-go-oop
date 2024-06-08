package main

import "./animals"

func main() {
	dog := animals.Animal{"Dog", "Roof", "Neutral"}
	cat := animals.Animal{"Cat", "Meow", "Neutral"}
	someAnimal := animals.Animal{Sound: "???"}

	cat.Pet()

	dog.DescribeAnimal()
	cat.DescribeAnimal()
	someAnimal.DescribeAnimal()
}

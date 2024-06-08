package main

import (
	"fmt"
)

type Animal struct {
	name  string
	sound string
	mood  string
}

func (a *Animal) pet() {
	a.mood = "Happy"
}

func (a *Animal) describeAnimal() {
	fmt.Println(a.name, "makes", a.sound, "sound and its mood is", a.mood)
}

func main() {
	dog := Animal{"Dog", "Roof", "Neutral"}
	cat := Animal{"Cat", "Meow", "Neutral"}
	someAnimal := Animal{sound: "???"}

	cat.pet()

	dog.describeAnimal()
	cat.describeAnimal()
	someAnimal.describeAnimal()
}

package animals

import "fmt"

type Animal struct {
	Name  string
	Sound string
	Mood  string
}

func (a *Animal) Pet() {
	a.Mood = "Happy"
}

func (a *Animal) DescribeAnimal() {
	fmt.Println(a.Name, "makes", a.Sound, "sound and its mood is", a.Mood)
}

package main

import "github.com/LeonardoAMello/learning-go-oop/animals"

type hasFavouriteFood interface {
	Feed(food string)
}

func feed(animal hasFavouriteFood, food string) {
	animal.Feed(food)
}

func main() {
	dog := animals.Animal{"Dog", "Roof", "Neutral"}
	cat := animals.Animal{"Cat", "Meow", "Neutral"}
	someAnimal := animals.Animal{Sound: "???"}

	cat.Pet()

	dog.DescribeAnimal()
	cat.DescribeAnimal()
	someAnimal.DescribeAnimal()

	monkey := animals.Monkey{FavouriteFood: "Banana"}
	parrot := animals.Parrot{FavouriteFood: "Seeds"}

	feed(&monkey, "Banana")
	feed(&parrot, "Banana")
	// feed(&dog, "Banana") // Dog doesnt have function implemented, editor throws error
}

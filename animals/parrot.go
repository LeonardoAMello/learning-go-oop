package animals

import "fmt"

type Parrot struct {
	FavouriteFood string
}

func (a *Parrot) Feed(food string) {
	if a.FavouriteFood == food {
		fmt.Println("Parrot is happy")
	} else {
		fmt.Println("Parrot is angry")
	}
}

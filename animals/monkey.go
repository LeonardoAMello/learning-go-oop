package animals

import "fmt"

type Monkey struct {
	FavouriteFood string
}

func (a *Monkey) Feed(food string) {
	if a.FavouriteFood == food {
		fmt.Println("Monkey is happy")
	} else {
		fmt.Println("Monkey is angry")
	}
}

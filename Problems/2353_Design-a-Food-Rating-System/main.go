package main

import "fmt"

func main() {
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	cuisines := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	ratings := []int{9, 12, 8, 15, 14, 7}

	obj := Constructor(foods, cuisines, ratings)
	highestRating := obj.HighestRated("korean")
	fmt.Println("OUT:", highestRating)
	highestRating = obj.HighestRated("japanese")
	fmt.Println("OUT:", highestRating)

	obj.ChangeRating("sushi", 16)
	highestRating = obj.HighestRated("japanese")
	fmt.Println("OUT:", highestRating)

}

package main

import "fmt"

type FoodRatings struct {
	ratings         map[string]*Rating
	foodCategoryMap map[string]string
}

type Rating struct {
	name  string
	value int
	next  *Rating
}

func isCurrlexicographiclyLarger(s1, s2 string) bool {
	for i := range s1 {
		if i >= len(s2) {
			return false
		}
		if s1[i] != s2[i] {
			return s1[i] < s2[i]
		}
	}
	return true
}

func isCurrRatingHigherThanNext(curr, next *Rating) bool {
	if curr.value == next.value {
		res := isCurrlexicographiclyLarger(curr.name, next.name)
		return res
	}
	res := curr.value > next.value
	return res
}

func insertRating(root *Rating, food string, rating int) *Rating {
	node := Rating{
		name:  food,
		value: rating,
	}
	if root == nil {
		return &node
	}

	// If Higher than head, replace head and then return
	if isCurrRatingHigherThanNext(&node, root) {
		node.next = root
		return &node
	}
	curr := root

	// Else Find where next should be where the rating be inserted
	for curr != nil {
		if curr.next == nil {
			curr.next = &node
			return root
		}
		if curr.value >= node.value && isCurrRatingHigherThanNext(&node, curr.next) {
			nextNode := curr.next
			curr.next = &node
			node.next = nextNode
			return root
		}

		curr = curr.next
	}
	return root
}

func deleteRating(root *Rating, food string) *Rating {
	if root.name == food {
		return root.next
	}
	curr := root
	for curr.next != nil {
		if curr.next.name == food {
			curr.next = curr.next.next
			return root
		}
		curr = curr.next
	}
	return root
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	this := FoodRatings{
		ratings:         make(map[string]*Rating),
		foodCategoryMap: make(map[string]string),
	}
	for i := range foods {
		food := foods[i]
		cuisine := cuisines[i]
		rating := ratings[i]

		fmt.Printf("Category: %10s %10s : %d\n", cuisine, food, rating)

		// Check if
		root := this.ratings[cuisine]
		root = insertRating(root, food, rating)

		this.ratings[cuisine] = root

		this.foodCategoryMap[food] = cuisine

		this.Print()
	}
	return this
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	foodCategory := this.foodCategoryMap[food]
	root := this.ratings[foodCategory]
	root = deleteRating(root, food)

	this.ratings[foodCategory] = insertRating(root, food, newRating)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	if v, ok := this.ratings[cuisine]; ok && v != nil {
		return v.name
	}
	return ""
}

func (this *FoodRatings) Print() {
	fmt.Println()
	for category, rootRating := range this.ratings {
		curr := rootRating
		fmt.Println(category)
		for curr != nil {
			fmt.Printf("| %s : %d |", curr.name, curr.value)
			curr = curr.next
		}
		fmt.Println()
	}
	fmt.Println()
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */

package main

import (
	"container/heap"
	"fmt"
)

type FoodRatings struct {
	ratings         map[string]*RatingHeap
	ratingMap       map[string]*Rating
	foodCategoryMap map[string]string
}

type Rating struct {
	name  string
	value int
}

type RatingHeap []*Rating

func (h RatingHeap) Len() int           { return len(h) }
func (h RatingHeap) Less(i, j int) bool { return isCurrRatingHigherThanNext(h[i], h[j]) }
func (h RatingHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *RatingHeap) Push(x any) {
	*h = append(*h, x.(*Rating))
}

func (h *RatingHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *RatingHeap) ViewTop() any {
	return fmt.Sprintf("%s %v", (*h)[0].name, (*h)[0].value)
}

func (h *RatingHeap) GetTop() string {
	return (*h)[0].name
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

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	this := FoodRatings{
		ratings:         make(map[string]*RatingHeap),
		foodCategoryMap: make(map[string]string),
		ratingMap:       make(map[string]*Rating),
	}
	for i := range foods {
		food := foods[i]
		cuisine := cuisines[i]
		rating := ratings[i]

		node := Rating{
			name:  food,
			value: rating,
		}
		// If category not extist init heap
		h, ok := this.ratings[cuisine]
		if !ok {
			h = &RatingHeap{}
			heap.Init(h)
			this.ratings[cuisine] = h
		}
		heap.Push(h, &node)
		this.ratings[cuisine] = h
		this.ratingMap[food] = &node
		this.foodCategoryMap[food] = cuisine

	}
	return this
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {

	category := this.foodCategoryMap[food]
	fmt.Println("before changes", food, newRating, this.ratings[category])

	fmt.Println(this.ratingMap[food].name, this.ratingMap[food].value)
	(*this.ratingMap[food]).value = newRating

	heap.Init(this.ratings[category])

	fmt.Println("after changes", food, newRating, this.ratings[category])
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return this.ratings[cuisine].GetTop()
}

func (this *FoodRatings) PopAll(cuisine string) {
	for this.ratings[cuisine].Len() > 0 {
		fmt.Println(this.ratings[cuisine].ViewTop())
		fmt.Println(this.ratings[cuisine])
		heap.Pop(this.ratings[cuisine])
	}
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */

package main

import "container/heap"

type Data struct {
	priority int
}

type Heap []Data

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].priority < h[j].priority } // Default is Min Heap, reverse this logic to get Max Heap
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.(Data))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap) GetTop() Data {
	return (*h)[0]
}

func main() {
	// Init using a regular array
	h := &Heap{
		Data{
			priority: 5,
		},
	}
	heap.Init(h)

	heap.Push(h, Data{priority: 10})

	h.GetTop() // Data Prioriy 5
}

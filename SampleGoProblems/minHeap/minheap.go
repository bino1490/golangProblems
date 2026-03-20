package main

import "fmt"

type Minheap struct {
	data []int
}

// Heapify Up (Insert)
func (h *Minheap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2

		if h.data[parent] <= h.data[index] {
			break // IMPORTANT
		}

		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

// Heapify Down (Delete)
func (h *Minheap) heapifyDown(index int) {
	n := len(h.data)

	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}

		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}

// Insert
func (h *Minheap) Insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

// Extract Min
func (h *Minheap) ExtractMin() int {
	if len(h.data) == 0 {
		return -1
	}

	min := h.data[0]
	last := h.data[len(h.data)-1]

	// Move last to root
	h.data[0] = last
	h.data = h.data[:len(h.data)-1]

	// Fix heap
	h.heapifyDown(0)

	return min
}

func main() {
	h := Minheap{}

	arr := []int{2, 5, 3, 6, 0, 9}

	// Insert
	for _, v := range arr {
		h.Insert(v)
	}

	fmt.Println("Heap:", h.data)

	// Extract Min
	fmt.Println("Extract:", h.ExtractMin())
	fmt.Println("After Extract:", h.data)
}

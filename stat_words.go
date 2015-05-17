package main

import "container/heap"

type wordCounter struct {
	word  string
	count int
	index int
}

type wordsHeap []*wordCounter

func (h wordsHeap) Len() int {
	return len(h)
}

func (h wordsHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h wordsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *wordsHeap) Push(x interface{}) {
	n := len(*h)
	wc := x.(*wordCounter)
	wc.index = n
	*h = append(*h, wc)
}

func (h *wordsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	wc := old[n-1]
	wc.index = -1
	*h = old[0 : n-1]
	return wc
}

func (h *wordsHeap) update(wc *wordCounter, count int) {
	wc.count = count
	heap.Fix(h, wc.index)
}

func (h *wordsHeap) copy() *wordsHeap {
	copy := wordsHeap{}
	for _, wc := range *h {
		copy = append(copy, &wordCounter{word: wc.word, count: wc.count, index: wc.index})
	}
	return &copy
}

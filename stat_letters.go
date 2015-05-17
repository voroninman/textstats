package main

type letterCounter struct {
	letter string
	count  int
}

type lettersHeap []*letterCounter // This "heap" don't correlate with data structure called heap

func (h lettersHeap) Len() int           { return len(h) }
func (h lettersHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h lettersHeap) Less(i, j int) bool { return h[i].count > h[j].count }

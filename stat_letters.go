package main

type letterCounter struct {
	letter string
	count  int
}

// This "heap" doesn't correlate with the data structure called heap
// Sorry for being ambiguous
// TODO: Devise a proper variable name
type lettersHeap []*letterCounter

func (h lettersHeap) Len() int           { return len(h) }
func (h lettersHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h lettersHeap) Less(i, j int) bool { return h[i].count > h[j].count }

package main

import (
	"container/heap"
	"sort"
	"strconv"
)

// Stat holding the statictics in memory
// We keep heaps for sorting
type Stat struct {
	count        int
	lettersItems map[string]*letterCounter
	lettersHeap  lettersHeap
	wordsItems   map[string]*wordCounter
	wordsHeap    wordsHeap
}

// NewStat Stat constructor
func NewStat() *Stat {
	return &Stat{
		wordsItems:   map[string]*wordCounter{},
		wordsHeap:    wordsHeap{},
		lettersItems: map[string]*letterCounter{},
		lettersHeap:  lettersHeap{},
		count:        0,
	}
}

// RecordWord eat a word
func (stat *Stat) RecordWord(word string) {

	wc, ok := stat.wordsItems[word]
	if !ok {
		wc = &wordCounter{word: word, count: 0}
		stat.wordsItems[word] = wc
		heap.Push(&stat.wordsHeap, wc)
	}

	for _, char := range word {
		letter := string(char)
		lc, ok := stat.lettersItems[letter]
		if !ok {
			lc = &letterCounter{letter: letter, count: 0}
			stat.lettersItems[letter] = lc
			stat.lettersHeap = append(stat.lettersHeap, lc)
		}
		lc.count++
	}
	stat.wordsHeap.update(wc, wc.count+1)
	stat.count++
}

// Dump the statictics of words count, frequent words and letters
func (stat Stat) Dump(n int) map[string]interface{} {
	dump := map[string]interface{}{}

	dump["count"] = stat.count

	jsonKeyWords := "top_" + strconv.Itoa(n) + "_words"
	dump[jsonKeyWords] = []string{}
	heapCopy := stat.wordsHeap.copy()

	jsonKeyLetters := "top_" + strconv.Itoa(n) + "_letters"
	dump[jsonKeyLetters] = []string{}
	sort.Sort(stat.lettersHeap)

	for i := 0; i < n; i++ {
		if heapCopy.Len() > 0 {
			word := heap.Pop(heapCopy).(*wordCounter).word
			dump[jsonKeyWords] = append(dump[jsonKeyWords].([]string), word)
		}
		if i < len(stat.lettersHeap) {
			dump[jsonKeyLetters] = append(
				dump[jsonKeyLetters].([]string),
				stat.lettersHeap[i].letter)
		}
	}

	return dump
}

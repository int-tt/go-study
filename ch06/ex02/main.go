package main

import "fmt"

type IntSet struct {
	words []uint64
}

func main() {
	intset := &IntSet{
		words: []uint64{1, 2, 3, 4, 5, 6, 7},
	}
	fmt.Println("IntSet.words:", intset.words)
	intset.AddAll(8, 9, 10)
	fmt.Println("Len:", intset.words)
}
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// 要素数を返します
func (s *IntSet) Len() int {
	return len(s.words)
}

// セットからxを取り除きます
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

// セットからすべての要素をとりのぞきます
func (s *IntSet) Clear() {
	s.words = []uint64{}
}

// セットのコピーを返します
func (s *IntSet) Copy() *IntSet {
	var set IntSet
	set.words = make([]uint64, len(s.words))
	copy(set.words, s.words)
	return &set
}
func (s *IntSet) AddAll(vals ...int) {
	list := make([]uint64, 0, len(vals))
	for _, v := range vals {
		list = append(list, uint64(v))
	}
	s.words = append(s.words, list...)
}

package main

import "fmt"

type IntSet struct {
	words []uint64
}

func main() {
	a := &IntSet{
		words: []uint64{1, 2, 3, 4, 5},
	}
	fmt.Println("a.words:", a.words)
	fmt.Println("Enum(int):", a.Enums())
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

// 和集合
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < s.Len() {
			s.words[i] &= tword
		}
	}

	for i := t.Len(); i < s.Len(); i++ {
		s.words[i] = 0
	}
}

// 積集合
func (s *IntSet) DifferencWith(t *IntSet) {
	for i, tword := range t.words {
		if i < s.Len() {
			s.words[i] &^= tword
		}
	}
}

// どちらかの集合にはあるが、両方には無い集合
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < s.Len() {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Enums() []int {
	len := s.Len()
	if len == 0 {
		return []int{}
	}
	enums := make([]int, 0, len)
	for i, sword := range s.words {
		for bit := uint(0); bit < 64; bit++ {
			if sword&(1<<bit) != 0 {
				enums = append(enums, i*64+int(bit))
			}
		}
	}
	return enums
}

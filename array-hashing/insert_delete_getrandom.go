package algoritmosexercicios

import "math/rand"

type RandomizedSet struct {
	m map[int]int
	s []int
}

func constructor() RandomizedSet {
	return RandomizedSet{m: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.m[val]

	if !ok {
		this.m[val] = len(this.s) // index do slices 's'
		this.s = append(this.s, val)
	}

	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.m[val]

	if ok {
		last := this.s[len(this.s)-1]
		this.s[idx] = last // substitui o elemento a ser removido pelo ultimo

		this.m[last] = idx // atualiza o index do ultimo

		this.s = this.s[:len(this.s)-1] // remove o ultimo elemento via slicing O(1)

		delete(this.m, val) // remove do map
	}

	return ok
}

func (this *RandomizedSet) GetRandom() int {

	return this.s[rand.Intn(len(this.s))]
}

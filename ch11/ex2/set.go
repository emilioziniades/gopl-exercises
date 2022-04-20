package main

type Set map[int]bool

func (s Set) Has(x int) bool {
	return s[x]
}

func (s Set) Add(x int) {
	s[x] = true
}

func (s Set) UnionWith(t Set) {
	for i := range t {
		s[i] = true
	}
}

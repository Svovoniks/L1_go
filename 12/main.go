package main

import "fmt"

type Set struct {
	mp map[string]bool
}

func GetSet() Set {
	return Set{mp: make(map[string]bool)}
}

func (s *Set) Add(st string) {
	s.mp[st] = true
}

func (s *Set) Contains(st string) bool {
	return s.mp[st]
}

func (s *Set) GetAll() []string {
	all := make([]string, len(s.mp))

	c := 0
	for val := range s.mp {
		all[c] = val
		c++
	}

	return all
}

func (s *Set) Delete(st string) {
	delete(s.mp, st)
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := GetSet()

	for _, i := range arr {
		set.Add(i)
	}

	fmt.Println(set.GetAll())

}

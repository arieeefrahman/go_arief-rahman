package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name string
	count int
}

type countSorted []pair

func (c countSorted) Len() int {
	return len(c)
}

func (c countSorted) Swap(i, j int)  {
	c[i], c[j] = c[j], c[i]
}

func (c countSorted) Less(i, j int) bool {
	return c[i].count < c[j].count
}

func MostAppearIteam(items []string) []pair{
	dict := make(map[string]int)
	var pairs []pair

	for _, v := range items  {
		dict[v] = dict[v]+1
	}

	for key, value := range dict {
		results := []pair{{name: key, count: value}}

		for _, details := range results {
			pairs = append(pairs, pair{
				name: details.name, count: details.count,
			})
		}
	}

	sort.Sort(countSorted(pairs))

	return pairs
}
func main() {
	fmt.Println(MostAppearIteam([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearIteam([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearIteam([]string{"football", "basketball", "tenis"}))
}
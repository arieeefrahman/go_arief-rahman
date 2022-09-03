package main

import (
	"fmt"
)

type pair struct {
	name string
	count int
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

	return pairs
}
func main() {
	fmt.Println(MostAppearIteam([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearIteam([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearIteam([]string{"football", "basketball", "tenis"}))
}
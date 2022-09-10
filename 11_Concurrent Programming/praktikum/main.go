package main

import (
	"fmt"
	"strings"
	"gopkg.in/loremipsum.v1"
)

func main() {
	loremIpsumGenerator := loremipsum.NewWithSeed(1234)
	loremString := loremIpsumGenerator.Words(20)
	fmt.Printf("%s\n\n", loremString)
	lorems := strings.Split(loremString, " ")
	result := ConcurrentFrequency(lorems)
	
	for key, value := range result {
		character := string(rune(key))
		fmt.Printf("%s : %d\n", character, value)
	}
}

type frequencyMap map[rune]int

func Frequency(word string) frequencyMap {
	m := frequencyMap{}
	
	for _, r := range word {
		m[r]++
	}

	return m
}

func ConcurrentFrequency(words []string) frequencyMap {
	m := frequencyMap{}
	c := make(chan frequencyMap, len(words))
	
	for _, v := range words {
		go func(v string) {
			c <- Frequency(v)
		}(v)
	}

	for range words {
		for r, v := range <- c {
			m[r] += v
		}
	}

	return m
}
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
	"unicode/utf8"
)

type populationItem struct {
	Key   string
	Value float64
}

func main() {
	// Define a random seed
	rand.Seed(time.Now().UnixNano())

	// Define parameters
	sentence := string("This is a genetic algorithm to evaluate, combine, evolve  mutate a string!")
	charmap := []rune(" ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz.,;!?+-*#@^'èéòà€ù=)(&%$£/\\")
	populationNum := 200
	selectionNum := 50
	mutationProb := .1

	// Verify the presence of all char in sentence
	for position, r := range []rune(sentence) {
		find := func() bool {
			for _, n := range charmap {
				if n == r {
					return true
				}
			}
			return false
		}
		if !find() {
			fmt.Println(errors.New("Character not aviable in charmap"), position, "\"", string(r), "\"")
			os.Exit(1)
		}
	}

	// Generate random population
	pop := make([]populationItem, populationNum, populationNum)
	for i := 0; i < populationNum; i++ {
		key := ""
		for x := 0; x < utf8.RuneCountInString(sentence); x++ {
			choice := rand.Intn(len(charmap))
			key += string(charmap[choice])
		}
		pop[i] = populationItem{key, 0}
	}

	for gen, generatedPop := 1, 0; ; gen++ {
		generatedPop += len(pop)

		// Random population created now it's time to evaluate
		for i, item := range pop {
			itemKey, sentenceRune := []rune(item.Key), []rune(sentence)
			for x := 0; x < len(sentence); x++ {
				if itemKey[x] == sentenceRune[x] {
					pop[i].Value++
				}
			}
			pop[i].Value = pop[i].Value / float64(len(sentenceRune))
		}
		// Check if there is a right evolution
		sort.SliceStable(pop, func(i, j int) bool { return pop[i].Value > pop[j].Value })
		if pop[0].Key == sentence {
			fmt.Println("Generation:", strconv.Itoa(gen), "Analyzed:", generatedPop, "Best:", pop[0])
			break
		}
		// Print the best result
		if gen%1000 == 0 {
			fmt.Println("Generation:", strconv.Itoa(gen), "Analyzed:", generatedPop, "Best:", pop[0])
		}
		// Combine, Evolve and Mutate
		var popChildren []populationItem
		for i := 0; i < int(selectionNum); i++ {
			parent1 := pop[i]
			parent2 := pop[i+1]
			split := rand.Intn(utf8.RuneCountInString(sentence))

			// Save Children 1
			child := append([]rune(parent1.Key)[:split], []rune(parent2.Key)[split:]...)
			if rand.Float64() > mutationProb {
				child[rand.Intn(len(child))] = charmap[rand.Intn(len(charmap))]
			}
			popChildren = append(popChildren, populationItem{string(child), 0})

			// Save Children 2
			child = append([]rune(parent2.Key)[:split], []rune(parent1.Key)[split:]...)
			if rand.Float64() > mutationProb {
				child[rand.Intn(len(child))] = charmap[rand.Intn(len(charmap))]
			}
			popChildren = append(popChildren, populationItem{string(child), 0})
		}
		pop = popChildren
	}
}

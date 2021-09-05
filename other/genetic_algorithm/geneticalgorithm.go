/*
Simple multithreaded algorithm to show how the 4 phases of a genetic
algorithm works (Evaluation, Selection, Crossover and Mutation)
https://en.wikipedia.org/wiki/Genetic_algorithm

Link to the same algorithm implemented in python:
https://github.com/TheAlgorithms/Python/blob/master/genetic_algorithm/basic_string.py

Author: D4rkia
*/

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

func geneticString(target string, charmap []rune) (int, int, string) {
	// Define parameters
	// Maximum size of the population.  bigger could be faster but is more memory expensive
	populationNum := 200
	// Number of elements selected in every generation for evolution the selection takes
	// place from the best to the worst of that generation must be smaller than N_POPULATION
	selectionNum := 50
	// Probability that an element of a generation can mutate changing one of its genes this
	// guarantees that all genes will be used during evolution
	mutationProb := .4
	// Just a seed to improve randomness required by the algorithm
	rand.Seed(time.Now().UnixNano())

	// Verify if 'populationNum' s bigger than 'selectionNum'
	if populationNum < selectionNum {
		fmt.Println(errors.New("populationNum must be bigger than selectionNum"))
		os.Exit(1)
	}
	// Verify that the target contains no genes besides the ones inside genes variable.
	for position, r := range []rune(target) {
		find := func() bool {
			for _, n := range charmap {
				if n == r {
					return true
				}
			}
			return false
		}
		if !find() {
			fmt.Println(errors.New("character not aviable in charmap"), position, "\"", string(r), "\"")
			os.Exit(1)
		}
	}

	// Generate random starting population
	pop := make([]populationItem, populationNum)
	for i := 0; i < populationNum; i++ {
		key := ""
		for x := 0; x < utf8.RuneCountInString(target); x++ {
			choice := rand.Intn(len(charmap))
			key += string(charmap[choice])
		}
		pop[i] = populationItem{key, 0}
	}

	// Just some logs to know what the algorithms is doing
	gen, generatedPop := 0, 0

	// This loop will end when we will find a perfect match for our target
	for {
		gen++
		generatedPop += len(pop)

		// Random population created now it's time to evaluate
		for i, item := range pop {
			pop[i].Value = 0
			itemKey, targetRune := []rune(item.Key), []rune(target)
			for x := 0; x < len(target); x++ {
				if itemKey[x] == targetRune[x] {
					pop[i].Value++
				}
			}
			pop[i].Value = pop[i].Value / float64(len(targetRune))
		}
		sort.SliceStable(pop, func(i, j int) bool { return pop[i].Value > pop[j].Value })

		// Check if there is a matching evolution
		if pop[0].Key == target {
			break
		}
		// Print the best resultPrint the Best result every 10 generations
		// just to know that the algorithm is working
		if gen%10 == 0 {
			fmt.Println("Generation:", strconv.Itoa(gen), "Analyzed:", generatedPop, "Best:", pop[0])
		}

		// Generate a new population vector keeping some of the best evolutions
		// Keeping this avoid regression of evolution
		var popChildren []populationItem
		popChildren = append(popChildren, pop[0:int(selectionNum/3)]...)

		// This is Selection
		for i := 0; i < int(selectionNum); i++ {
			parent1 := pop[i]
			// Generate more child proportionally to the fitness score
			nChild := (parent1.Value * 100) + 1
			if nChild >= 10 {
				nChild = 10
			}
			for x := 0.0; x < nChild; x++ {
				parent2 := pop[rand.Intn(selectionNum)]
				// Crossover
				split := rand.Intn(utf8.RuneCountInString(target))
				child1 := append([]rune(parent1.Key)[:split], []rune(parent2.Key)[split:]...)
				child2 := append([]rune(parent2.Key)[:split], []rune(parent1.Key)[split:]...)
				//Clean fitness value
				// Mutate
				if rand.Float64() < mutationProb {
					child1[rand.Intn(len(child1))] = charmap[rand.Intn(len(charmap))]
				}
				if rand.Float64() < mutationProb {
					child2[rand.Intn(len(child2))] = charmap[rand.Intn(len(charmap))]
				}
				// Push into 'popChildren'
				popChildren = append(popChildren, populationItem{string(child1), 0})
				popChildren = append(popChildren, populationItem{string(child2), 0})

				// Check if the population has already reached the maximum value and if so,
				// break the cycle. If this check is disabled the algorithm will take
				// forever to compute large strings but will also calculate small string in
				// a lot fewer generationsù
				if len(popChildren) >= selectionNum {
					break
				}
			}
		}
		pop = popChildren
	}
	return gen, generatedPop, pop[0].Key
}

func main() {
	// Define parameters
	target := string("This is a genetic algorithm to evaluate, combine, evolve and mutate a string!")
	charmap := []rune(" ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz.,;!?+-*#@^'èéòà€ù=)(&%$£/\\")
	gen, generatedPop, best := geneticString(target, charmap)
	fmt.Println("Generation:", strconv.Itoa(gen), "Analyzed:", generatedPop, "Best:", best)
}

// Package genetic provides functions to work with strings
// using genetic algorithm. https://en.wikipedia.org/wiki/Genetic_algorithm
//
// Author: D4rkia
package genetic

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
	"unicode/utf8"
)

// Population item represent a single step in the evolution process.
// One can think of population item as a single species.
// Key stands for the actual data entity of the species, which is a string
// in current implementation. Key can be interpreted as species DNA.
// Value shows how close this species to the desired target, where 1 means,
// that species DNA equals to the targeted one, 0 for no matchings in the DNA.
//
// **Note** In the current implementation species DNA length is suppose to be
// equal to the target length for algorithm to work.
type PopulationItem struct {
	Key   string
	Value float64
}

// Conf stands for configurations set provided to GeneticString function.
type Conf struct {
	// Maximum size of the population.
	// Bigger could be faster but more memory expensive.
	PopulationNum int

	// Number of elements selected in every generation for evolution
	// the selection takes. Place from the best to the worst of that
	// generation must be smaller than PopulationNum.
	SelectionNum int

	// Probability that an element of a generation can mutate changing one of
	// its genes this guarantees that all genes will be used during evolution.
	MutationProb float64

	// Enables debugging output to the console.
	Debug bool
}

// Result structure contains generation process statistics, as well as the
// best resulted population item.
type Result struct {
	// Number of generations steps performed.
	Generation int

	// Number of generated population items.
	Analyzed int

	// Result of generation with the best Value.
	Best PopulationItem
}

// GeneticString generates PopulationItem based on the imputed target
// string, and a set of possible runes to build a string with. In order
// to optimise string generation additional configurations can be provided
// with Conf instance. Empty instance of Conf (&Conf{}) can be provided,
// then default values would be set.
//
// Link to the same algorithm implemented in python:
// https://github.com/TheAlgorithms/Python/blob/master/genetic_algorithm/basic_string.py
func GeneticString(target string, charmap []rune, conf *Conf) (*Result, error) {
	populationNum := conf.PopulationNum
	if populationNum == 0 {
		populationNum = 200
	}

	selectionNum := conf.SelectionNum
	if selectionNum == 0 {
		selectionNum = 50
	}

	// Verify if 'populationNum' s bigger than 'selectionNum'
	if populationNum < selectionNum {
		return nil, errors.New("populationNum must be bigger than selectionNum")
	}

	mutationProb := conf.MutationProb
	if mutationProb == .0 {
		mutationProb = .4
	}

	debug := conf.Debug

	// Just a seed to improve randomness required by the algorithm
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Verify that the target contains no genes besides the ones inside genes variable.
	for position, r := range target {
		invalid := true
		for _, n := range charmap {
			if n == r {
				invalid = false
			}
		}
		if invalid {
			message := fmt.Sprintf("character not available in charmap at position: %v", position)
			return nil, errors.New(message)
		}
	}

	// Generate random starting population
	pop := make([]PopulationItem, populationNum)
	for i := 0; i < populationNum; i++ {
		key := ""
		for x := 0; x < utf8.RuneCountInString(target); x++ {
			choice := rnd.Intn(len(charmap))
			key += string(charmap[choice])
		}
		pop[i] = PopulationItem{key, 0}
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
		if debug && gen%10 == 0 {
			fmt.Println("Generation:", strconv.Itoa(gen), "Analyzed:", generatedPop, "Best:", pop[0])
		}

		// Generate a new population vector keeping some of the best evolutions
		// Keeping this avoid regression of evolution
		var popChildren []PopulationItem
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
				parent2 := pop[rnd.Intn(selectionNum)]
				// Crossover
				split := rnd.Intn(utf8.RuneCountInString(target))
				child1 := append([]rune(parent1.Key)[:split], []rune(parent2.Key)[split:]...)
				child2 := append([]rune(parent2.Key)[:split], []rune(parent1.Key)[split:]...)
				// Clean fitness value
				// Mutate
				if rnd.Float64() < mutationProb {
					child1[rnd.Intn(len(child1))] = charmap[rnd.Intn(len(charmap))]
				}
				if rnd.Float64() < mutationProb {
					child2[rnd.Intn(len(child2))] = charmap[rnd.Intn(len(charmap))]
				}
				// Push into 'popChildren'
				popChildren = append(popChildren, PopulationItem{string(child1), 0})
				popChildren = append(popChildren, PopulationItem{string(child2), 0})

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
	return &Result{gen, generatedPop, pop[0]}, nil
}

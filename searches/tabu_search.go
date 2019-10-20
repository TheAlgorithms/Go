package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func euc_2d(coord1, coord2 City) int {
	diffx := coord1.x - coord2.x
	diffy := coord1.y - coord2.y
	return int(math.Sqrt(float64(diffx*diffx + diffy*diffy)))
}

func cost(permutation []int, cities []City) int {
	distance := 0
	city2 := 0

	for index, city1 := range permutation {
		if index == len(permutation)-1 {
			city2 = permutation[0]
		} else {
			city2 = permutation[index+1]
		}

		distance += euc_2d(cities[city1], cities[city2])
	}

	return distance
}

func randomPermutation(cities []City) []int {
	perm := make([]int, len(cities))

	for i, _ := range perm {
		perm[i] = i
	}

	for i, _ := range perm {
		r := rand.Intn(len(perm)-i) + i
		perm[r], perm[i] = perm[i], perm[r]
	}

	return perm
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func indexOf(s []int, x int) int {
	for i, val := range s {
		if val == x {
			return i
		}
	}

	return -1
}

type Edge struct {
	left  int
	right int
}

func stochasticTwoOpt(parent []int) ([]int, []Edge) {
	//fmt.Println(parent)
	permutation := make([]int, len(parent))
	copy(permutation, parent)

	city1, city2 := rand.Intn(len(permutation)), rand.Intn(len(permutation))

	exclude := []int{city1}

	if city1 == 0 {
		exclude = append(exclude, len(permutation)-1)
	} else {
		exclude = append(exclude, city1-1)
	}

	if city1 == len(permutation)-1 {
		exclude = append(exclude, 0)
	} else {
		exclude = append(exclude, city1+1)
	}

	for {
		if indexOf(exclude, city2) == -1 {
			break
		}

		city2 = rand.Intn(len(permutation))
	}

	if city2 < city1 {
		city1, city2 = city2, city1
	}

	// reverse function will be applied to the permutation slice imidiately no need to assign the reversed
	// portion back to the slice
	reverse(permutation[city1:city2])
	//fmt.Println(city1, city2)

	if city1 == 0 {
		city1++
	}

	if city2 == 0 {
		city2++
	}
	//fmt.Println(parent)
	//fmt.Println(permutation)
	return permutation, []Edge{{parent[city1-1], parent[city1]}, {parent[city2-1], parent[city2]}}
}

func isTabu(permutation []int, tabuList []Edge) bool {
	city2 := 0

	for index, city1 := range permutation {
		if index == len(permutation)-1 {
			city2 = permutation[0]
		} else {
			city2 = permutation[index+1]
		}

		for _, forbiddenEdge := range tabuList {
			edge := Edge{city1, city2}
			//fmt.Println(edge)
			//fmt.Println(forbiddenEdge)
			if forbiddenEdge == edge {
				return true
			}
		}
	}

	return false
}

type Candidate struct {
	vector []int
	cost   int
	edges  []Edge
}

type CandidateVector []Candidate

func (c CandidateVector) Len() int {
	return len(c)
}

func (c CandidateVector) Less(i, j int) bool {
	return c[i].cost < c[j].cost
}

func (c CandidateVector) Swap(i, j int) {
	c[j], c[i] = c[i], c[j]
}

func generateCandidate(best Candidate, tabuList []Edge, cities []City) Candidate {
	permutation := make([]int, len(cities))
	edges := make([]Edge, 2)

	for {
		permutation, edges = stochasticTwoOpt(best.vector)

		if !isTabu(permutation, tabuList) {
			break
		}
	}

	candidate := Candidate{}
	candidate.vector = permutation
	candidate.cost = cost(candidate.vector, cities)
	candidate.edges = edges
	//fmt.Println(permutation)
	return candidate
}

func search(cities []City, tabuSize, maxCandidates, iterations int) Candidate {
	current := Candidate{}
	current.vector = randomPermutation(cities)
	current.cost = cost(current.vector, cities)
	//fmt.Println("current", current)
	best := current

	tabuList := make([]Edge, tabuSize)

	for i := 0; i < iterations; i++ {
		candidates := make([]Candidate, maxCandidates)
		for j, _ := range candidates {
			candidates[j] = generateCandidate(current, tabuList, cities)
		}

		sort.Sort(CandidateVector(candidates))
		bestCandidate := candidates[0]
		bestCandidateEdges := bestCandidate.edges

		if bestCandidate.cost < current.cost {
			current = bestCandidate
			if bestCandidate.cost < best.cost {
				best = bestCandidate

				for _, edge := range bestCandidateEdges {
					tabuList = append(tabuList, edge)
					//fmt.Println("push")
				}

				for {
					//pop
					tabuList = tabuList[1:len(tabuList)]
					//fmt.Println("pop")
					if len(tabuList) <= tabuSize {
						break
					}
				}
			}
		}

		if (i+1)%100 == 0 {
			//fmt.Println("Iteration", i+1, "best=", best, "taboo=", tabuList)
		}
	}

	return best
}

type City struct {
	x int
	y int
}

func main() {
	rand.Seed(time.Now().Unix())

	cities := []City{{565, 575}, {25, 185}, {345, 750}, {945, 685}, {845, 655},
		{880, 660}, {25, 230}, {525, 1000}, {580, 1175}, {650, 1130}, {1605, 620},
		{1220, 580}, {1465, 200}, {1530, 5}, {845, 680}, {725, 370}, {145, 665},
		{415, 635}, {510, 875}, {560, 365}, {300, 465}, {520, 585}, {480, 415},
		{835, 625}, {975, 580}, {1215, 245}, {1320, 315}, {1250, 400}, {660, 180},
		{410, 250}, {420, 555}, {575, 665}, {1150, 1160}, {700, 580}, {685, 595},
		{685, 610}, {770, 610}, {795, 645}, {720, 635}, {760, 650}, {475, 960},
		{95, 260}, {875, 920}, {700, 500}, {555, 815}, {830, 485}, {1170, 65},
		{830, 610}, {605, 625}, {595, 360}, {1340, 725}, {1740, 245}}

	iterations := 1000 //100
	tabuSize := 15     //15
	maxCandidates := 50

	best := search(cities, tabuSize, maxCandidates, iterations)

	//fmt.Println(best)

	fmt.Println("Result is:", best.cost, ", Result should be 7542.")
}

package main

import "fmt"

/**
 *
 * \file
 * \brief [Disjoint Sets Data Structure
 * (Disjoint Sets)](https://en.wikipedia.org/wiki/Disjoint-set_data_structure)
 *
 * \author [leoyang429](https://github.com/leoyang429)
 *
 * \details
 * A disjoint set data structure (also called union find or merge find set)
 * is a data structure that tracks a set of elements partitioned into a number
 * of disjoint (non-overlapping) subsets.
 * Some situations where disjoint sets can be used are-
 * to find connected components of a graph, kruskal's algorithm for finding
 * Minimum Spanning Tree etc.
 * There are two operation which we perform on disjoint sets -
 * 1) Union
 * 2) Find
 * 3) More Operations can be added as a further implementation to this code
 *
 */
type DSU struct {
	parent []int
	rank   []int
	size   []int
}

/**
 *
 * Function to initialize a set
 * @param n number of element
 *
 */
func (d *DSU) init(size int) {
	for i := 0; i < size+1; i++ {
		d.parent = append(d.parent, i)
		d.size = append(d.size, 1)
	}
	d.rank = make([]int, size+1, size+1)
}

/**
 *
 * Find operation takes a number x and returns the set to which this number
 * belongs to.
 * @param x element of some set
 * @return set to which x belongs to
 *
 */
func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

/**
 *
 * A utility function to check if x and y are from same set or not
 * @param x element of some set
 * @param y element of some set
 *
 */
func (d *DSU) belongToSame(x, y int) bool {
	xroot := d.find(x)
	yroot := d.find(y)
	return xroot == yroot
}

/**
 *
 * Union operation combines two disjoint sets to make a single set
 * in this union function we pass two elements and check if they are
 * from different sets then combine those sets
 * @param x element of some set
 * @param y element of some set
 *
 */
func (d *DSU) union(x, y int) bool {
	xroot := d.find(x)
	yroot := d.find(y)
	if xroot == yroot {
		return false
	} else if d.rank[xroot] < d.rank[yroot] {
		d.parent[xroot] = yroot
	} else if d.rank[xroot] > d.rank[yroot] {
		d.parent[yroot] = xroot
	} else {
		d.parent[yroot] = xroot
		d.rank[xroot]++
	}
	return true
}

func main() {

	var size int = 10
	var myDSU DSU
	myDSU.init(size)  // 10 numbers pointed to themselves. Size of individual elements initialized.
	myDSU.union(1, 2) // merging two sets
	myDSU.union(3, 2)
	myDSU.union(9, 3)
	parent := myDSU.find(3)                      // finding the parent
	fmt.Printf("%v\n", parent)                   // checking parent value of 3
	fmt.Printf("%v\n", myDSU.belongToSame(4, 9)) // Test to see if 2 elements belong to same set
	fmt.Printf("%v\n", myDSU.belongToSame(1, 3))
	myDSU.union(4, 9)
	fmt.Printf("%v\n", myDSU.belongToSame(4, 9))
}

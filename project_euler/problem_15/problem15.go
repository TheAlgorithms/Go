/**
* Problem 15 - Lattice paths
* @see {@link https://projecteuler.net/problem=15}
*
* Starting in the top left corner of a 2×2 grid,
* and only being able to move to the right and down,
* there are exactly 6 routes to the bottom right corner.
*
* How many such routes are there through a 20×20 grid?
*
* @author ddaniel27
 */
package problem15

import (
	"github.com/TheAlgorithms/Go/math/factorial"
)

func Problem15(gridSize int) int {
	/**
	Author note:
	We can solve this problem using combinatorics.
	Here is a good blog post that explains the solution:

	[link](https://stemhash.com/counting-lattice-paths/)

	Btw, I'm not related to the author of the blog post.

	After some simplification, we can see that the solution is:
	(2n)! / (n!)^2

	We can use the factorial package to calculate the factorials.
	*/

	n := gridSize

	numerator, _ := factorial.Iterative(2 * n)
	denominator, _ := factorial.Iterative(n)
	denominator *= denominator

	return numerator / denominator
}

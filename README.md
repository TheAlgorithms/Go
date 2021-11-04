# The Algorithms - Go
[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod&style=flat-square)](https://gitpod.io/#https://github.com/TheAlgorithms/Go)&nbsp;
![golangci-lint](https://github.com/TheAlgorithms/Go/workflows/golangci-lint/badge.svg)
![](https://img.shields.io/github/repo-size/TheAlgorithms/Go.svg?label=Repo%20size&style=flat-square)&nbsp;
![update_directory_md](https://github.com/TheAlgorithms/Go/workflows/update_directory_md/badge.svg)
[![Discord chat](https://img.shields.io/discord/808045925556682782.svg?logo=discord&colorB=7289DA&style=flat-square)](https://discord.gg/c7MnfGFGa6)&nbsp;

### Algorithms implemented in Go (for education)

The repository is a collection of open-source implementation of a variety of algorithms implemented in Go and licensed under [MIT License](LICENSE).

Read our [Contribution Guidelines](CONTRIBUTING.md) before you contribute.

## List of Algorithms

<!--- GODOCMD BEGIN --->
# Packages:

<details>
	<summary> <strong> ahocorasick </strong> </summary>	

---

##### Functions:

1. [`AhoCorasick`](./strings/ahocorasick/ahocorasick.go#L15):  AhoCorasick Function performing the Basic Aho-Corasick algorithm. Finds and prints occurrences of each pattern.
2. [`BuildAc`](./strings/ahocorasick/ahocorasick.go#L54):  Functions that builds Aho Corasick automaton.
3. [`ConstructTrie`](./strings/ahocorasick/shared.go#L4):  ConstructTrie Function that constructs Trie as an automaton for a set of reversed & trimmed strings.
4. [`Contains`](./strings/ahocorasick/shared.go#L39):  Contains Returns 'true' if array of int's 's' contains int 'e', 'false' otherwise.
5. [`GetWord`](./strings/ahocorasick/shared.go#L49):  GetWord Function that returns word found in text 't' at position range 'begin' to 'end'.
6. [`ComputeAlphabet`](./strings/ahocorasick/shared.go#L61):  ComputeAlphabet Function that returns string of all the possible characters in given patterns.
7. [`IntArrayCapUp`](./strings/ahocorasick/shared.go#L70):  IntArrayCapUp Dynamically increases an array size of int's by 1.
8. [`BoolArrayCapUp`](./strings/ahocorasick/shared.go#L78):  BoolArrayCapUp Dynamically increases an array size of bool's by 1.
9. [`ArrayUnion`](./strings/ahocorasick/shared.go#L86):  ArrayUnion Concats two arrays of int's into one.
10. [`GetParent`](./strings/ahocorasick/shared.go#L99):  GetParent Function that finds the first previous state of a state and returns it. Used for trie where there is only one parent.
11. [`CreateNewState`](./strings/ahocorasick/shared.go#L111):  CreateNewState Automaton function for creating a new state 'state'.
12. [`CreateTransition`](./strings/ahocorasick/shared.go#L116):  CreateTransition Creates a transition for function σ(state,letter) = end.
13. [`GetTransition`](./strings/ahocorasick/shared.go#L121):  GetTransition Returns ending state for transition σ(fromState,overChar), '-1' if there is none.
14. [`StateExists`](./strings/ahocorasick/shared.go#L133):  StateExists Checks if state 'state' exists. Returns 'true' if it does, 'false' otherwise.
15. [`Advanced`](./strings/ahocorasick/advancedahocorasick.go#L10):  Advanced Function performing the Advanced Aho-Corasick algorithm. Finds and prints occurrences of each pattern.
16. [`BuildExtendedAc`](./strings/ahocorasick/advancedahocorasick.go#L46):  BuildExtendedAc Functions that builds extended Aho Corasick automaton.

---
##### Types

1. [`Result`](./strings/ahocorasick/ahocorasick.go#L9): No description provided.


---
</details><details>
	<summary> <strong> avl </strong> </summary>	

---

#####  Package avl is a Adelson-Velskii and Landis tree implemnation avl is self-balancing tree, i.e for all node in a tree, height difference between its left and right child will not exceed 1 more information : https://en.wikipedia.org/wiki/AVL_tree

---
##### Functions:

1. [`NewTree`](./structure/avl/avl.go#L15):  NewTree create a new AVL tree
2. [`Get`](./structure/avl/avl.go#L20):  Get : return node with given key
3. [`Insert`](./structure/avl/avl.go#L35):  Insert a new item
4. [`Delete`](./structure/avl/avl.go#L72):  Delete : remove given key from the tree

---
##### Types

1. [`Node`](./structure/avl/avl.go#L8): No description provided.


---
</details><details>
	<summary> <strong> binary </strong> </summary>	

---

#####  xorsearch_test.go description: Test for Find a missing number in a sequence author(s) [red_byte](https://github.com/i-redbyte) see xorsearch.go Package binary describes algorithms that use binary operations for different calculations.

---
##### Functions:

1. [`MeanUsingAndXor`](./math/binary/arithmeticmean.go#L11): No description provided.
2. [`MeanUsingRightShift`](./math/binary/arithmeticmean.go#L15): No description provided.
3. [`IsPowerOfTwo`](./math/binary/checkisnumberpoweroftwo.go#L19):  IsPowerOfTwo This function uses the fact that powers of 2 are represented like 10...0 in binary, and numbers one less than the power of 2 are represented like 11...1. Therefore, using the and function:    10...0  & 01...1    00...0 -> 0 This is also true for 0, which is not a power of 2, for which we have to add and extra condition.
4. [`IsPowerOfTwoLeftShift`](./math/binary/checkisnumberpoweroftwo.go#L26):  IsPowerOfTwoLeftShift This function takes advantage of the fact that left shifting a number by 1 is equivalent to multiplying by 2. For example, binary 00000001 when shifted by 3 becomes 00001000, which in decimal system is 8 or = 2 * 2 * 2
5. [`ReverseBits`](./math/binary/reversebits.go#L14):  ReverseBits This function initialized the result by 0 (all bits 0) and process the given number starting from its least significant bit. If the current bit is 1, set the corresponding most significant bit in the result and finally move on to the next bit in the input number. Repeat this till all its bits are processed.
6. [`XorSearchMissingNumber`](./math/binary/xorsearch.go#L10): No description provided.

---
</details><details>
	<summary> <strong> binarytree </strong> </summary>	

---

##### Functions:

1. [`Insert`](./structure/binarysearchtree/bstree.go#L17):  Insert a value in the BSTree
2. [`InOrderSuccessor`](./structure/binarysearchtree/bstree.go#L35):  InOrderSuccessor Goes to the left
3. [`BstDelete`](./structure/binarysearchtree/bstree.go#L44):  BstDelete removes the node
4. [`InOrder`](./structure/binarysearchtree/bstree.go#L79):  Travers the tree in the following order left --> root --> right
5. [`PreOrder`](./structure/binarysearchtree/bstree.go#L96):  Travers the tree in the following order root --> left --> right
6. [`PostOrder`](./structure/binarysearchtree/bstree.go#L113):  Travers the tree in the following order left --> right --> root
7. [`LevelOrder`](./structure/binarysearchtree/bstree.go#L138): No description provided.
8. [`AccessNodesByLayer`](./structure/binarysearchtree/bstree.go#L145):  AccessNodesByLayer Function that access nodes layer by layer instead of printing the results as one line.
9. [`Max`](./structure/binarysearchtree/bstree.go#L174):  Max Function that returns max of two numbers - possibly already declared.
10. [`NewNode`](./structure/binarysearchtree/node.go#L11):  NewNode Returns a new pointer to an empty Node

---
##### Types

1. [`BSTree`](./structure/binarysearchtree/bstree.go#L4): No description provided.

2. [`Node`](./structure/binarysearchtree/node.go#L4): No description provided.


---
</details><details>
	<summary> <strong> caesar </strong> </summary>	

---

#####  Package caesar is the shift cipher ref: https://en.wikipedia.org/wiki/Caesar_cipher

---
##### Functions:

1. [`Encrypt`](./cipher/caesar/caesar.go#L6):  Encrypt encrypts by right shift of "key" each character of "input"
2. [`Decrypt`](./cipher/caesar/caesar.go#L27):  Decrypt decrypts by left shift of "key" each character of "input"

---
</details><details>
	<summary> <strong> coloring </strong> </summary>	

---

#####  Package coloring provides implementation of different graph coloring algorithms, e.g. coloring using BFS, using Backtracking, using greedy approach. Author(s): [Shivam](https://github.com/Shivam010)

---
##### Functions:

1. [`BipartiteCheck`](./graph/coloring/bipartite.go#L40):  basically tries to color the graph in two colors if each edge connects 2 differently colored nodes the graph can be considered bipartite

---
##### Types

1. [`Graph`](./graph/coloring/graph.go#L14): No description provided.


---
</details><details>
	<summary> <strong> combination </strong> </summary>	

---

#####  Package combination ...

---
##### Functions:

1. [`Start`](./strings/combination/combination.go#L13):  Start ...

---
##### Types

1. [`Combinations`](./strings/combination/combination.go#L7): No description provided.


---
</details><details>
	<summary> <strong> conversion </strong> </summary>	

---

#####  Package name. Package name. Package conversion is a package of implementations which converts one data structure to another.

---
##### Functions:

1. [`RomanToInteger`](./conversion/romantointeger.go#L40):  RomanToInteger converts a roman numeral string to an integer. Roman numerals for numbers outside the range 1 to 3,999 will return an error. Nil or empty string return 0 with no error thrown.
2. [`IntToRoman`](./conversion/integertoroman.go#L17):  IntToRoman converts an integer value to a roman numeral string. An error is returned if the integer is not between 1 and 3999.
3. [`HEXToRGB`](./conversion/rgbhex.go#L10):  HEXToRGB splits an RGB input (e.g. a color in hex format; 0x<color-code>) into the individual components: red, green and blue
4. [`RGBToHEX`](./conversion/rgbhex.go#L41):  RGBToHEX does exactly the opposite of HEXToRGB: it combines the three components red, green and blue to an RGB value, which can be converted to e.g. Hex
5. [`BinaryToDecimal`](./conversion/binarytodecimal.go#L26):  BinaryToDecimal() function that will take Binary number as string, and return it's Decimal equivalent as integer.
6. [`Reverse`](./conversion/decimaltobinary.go#L23):  Reverse() function that will take string, and returns the reverse of that string.
7. [`DecimalToBinary`](./conversion/decimaltobinary.go#L33):  DecimalToBinary() function that will take Decimal number as int, and return it's Binary equivalent as string.

---
</details><details>
	<summary> <strong> diffiehellman </strong> </summary>	

---

#####  Package diffiehellman implements Deffie Hellman Key Exchange Algorithm for more information watch : https://www.youtube.com/watch?v=NmM9HA2MQGI

---
##### Functions:

1. [`GenerateShareKey`](./cipher/diffiehellman/diffiehellmankeyexchange.go#L13):  GenerateShareKey : generates a key using client private key , generator and primeNumber this key can be made public shareKey = (g^key)%primeNumber
2. [`GenerateMutualKey`](./cipher/diffiehellman/diffiehellmankeyexchange.go#L19):  GenerateMutualKey : generates a mutual key that can be used by only alice and bob mutualKey = (shareKey^prvKey)%primeNumber

---
</details><details>
	<summary> <strong> dynamic </strong> </summary>	

---

#####  Package dynamic is a package of certain implementations of dynamically run algorithms.

---
##### Functions:

1. [`LongestIncreasingSubsequence`](./dynamic/longestincreasingsubsequence.go#L9):  LongestIncreasingSubsequence returns the longest increasing subsequence where all elements of the subsequence are sorted in increasing order
2. [`CutRodRec`](./dynamic/rodcutting.go#L8):  CutRodRec solve the problem recursively: initial approach
3. [`CutRodDp`](./dynamic/rodcutting.go#L21):  CutRodDp solve the same problem using dynamic programming
4. [`EditDistanceRecursive`](./dynamic/editdistance.go#L10):  EditDistanceRecursive is a naive implementation with exponential time complexity.
5. [`EditDistanceDP`](./dynamic/editdistance.go#L35):  EditDistanceDP is an optimised implementation which builds on the ideas of the recursive implementation. We use dynamic programming to compute the DP table where dp[i][j] denotes the edit distance value of first[0..i-1] and second[0..j-1]. Time complexity is O(m * n) where m and n are lengths of the strings, first and second respectively.
6. [`LongestCommonSubsequence`](./dynamic/longestcommonsubsequence.go#L8):  LongestCommonSubsequence function
7. [`MatrixChainRec`](./dynamic/matrixmultiplication.go#L10):  MatrixChainRec function
8. [`MatrixChainDp`](./dynamic/matrixmultiplication.go#L24):  MatrixChainDp function
9. [`IsSubsetSum`](./dynamic/subsetsum.go#L14): No description provided.
10. [`Max`](./dynamic/knapsack.go#L11):  Max function - possible duplicate
11. [`Knapsack`](./dynamic/knapsack.go#L17):  Knapsack solves knapsack problem return maxProfit
12. [`LpsRec`](./dynamic/longestpalindromicsubsequence.go#L7):  LpsRec function
13. [`LpsDp`](./dynamic/longestpalindromicsubsequence.go#L21):  LpsDp function
14. [`Bin2`](./dynamic/binomialcoefficient.go#L21):  Bin2 function
15. [`NthFibonacci`](./dynamic/fibonacci.go#L6):  NthFibonacci returns the nth Fibonacci Number
16. [`NthCatalanNumber`](./dynamic/catalan.go#L13):  NthCatalan returns the n-th Catalan Number Complexity: O(n²)

---
</details><details>
	<summary> <strong> dynamicarray </strong> </summary>	

---

#####  Package dynamicarray A dynamic array is quite similar to a regular array, but its Size is modifiable during program runtime, very similar to how a slice in Go works. The implementation is for educational purposes and explains how one might go about implementing their own version of slices.  For more details check out those links below here: GeeksForGeeks article : https://www.geeksforgeeks.org/how-do-dynamic-arrays-work/ Go blog: https://blog.golang.org/slices-intro Go blog: https://blog.golang.org/slices authors [Wesllhey Holanda](https://github.com/wesllhey), [Milad](https://github.com/miraddo) see dynamicarray.go, dynamicarray_test.go

---
##### Types

1. [`DynamicArray`](./structure/dynamicarray/dynamicarray.go#L21): No description provided.


---
</details><details>
	<summary> <strong> factorial </strong> </summary>	

---

#####  Package factorial describes algorithms Factorials calculations.

---
##### Functions:

1. [`BruteForceFactorial`](./math/factorial/factorial.go#L11): No description provided.
2. [`RecursiveFactorial`](./math/factorial/factorial.go#L19): No description provided.
3. [`CalculateFactorialUseTree`](./math/factorial/factorial.go#L27): No description provided.

---
</details><details>
	<summary> <strong> gcd </strong> </summary>	

---

##### Functions:

1. [`TemplateTestGCD`](./math/gcd/gcd_test.go#L18): No description provided.
2. [`TemplateBenchmarkGCD`](./math/gcd/gcd_test.go#L37): No description provided.
3. [`Iterative`](./math/gcd/gcditerative.go#L4):  Iterative Faster iterative version of GcdRecursive without holding up too much of the stack
4. [`Extended`](./math/gcd/extended.go#L12):  Extended simple extended gcd
5. [`ExtendedRecursive`](./math/gcd/extendedgcd.go#L4):  ExtendedRecursive finds and returns gcd(a, b), x, y satisfying a*x + b*y = gcd(a, b).
6. [`TemplateTestExtendedGCD`](./math/gcd/extendedgcd_test.go#L7): No description provided.
7. [`TemplateBenchmarkExtendedGCD`](./math/gcd/extendedgcd_test.go#L44): No description provided.
8. [`ExtendedIterative`](./math/gcd/extendedgcditerative.go#L4):  ExtendedIterative finds and returns gcd(a, b), x, y satisfying a*x + b*y = gcd(a, b).
9. [`Recursive`](./math/gcd/gcd.go#L4):  Recursive finds and returns the greatest common divisor of a given integer.

---
</details><details>
	<summary> <strong> generateparentheses </strong> </summary>	

---

##### Functions:

1. [`GenerateParenthesis`](./strings/generateparentheses/generateparentheses.go#L12): No description provided.

---
</details><details>
	<summary> <strong> genetic </strong> </summary>	

---

#####  Package genetic provides functions to work with strings using genetic algorithm. https://en.wikipedia.org/wiki/Genetic_algorithm  Author: D4rkia

---
##### Functions:

1. [`GeneticString`](./strings/genetic/genetic.go#L71):  GeneticString generates PopultaionItem based on the imputed target string, and a set of possible runes to build a string with. In order to optimise string generation additional configurations can be provided with Conf instance. Empty instance of Conf (&Conf{}) can be provided, then default values would be set. Link to the same algorithm implemented in python: https://github.com/TheAlgorithms/Python/blob/master/genetic_algorithm/basic_string.py

---
##### Types

1. [`PopulationItem`](./strings/genetic/genetic.go#L26): No description provided.

2. [`Conf`](./strings/genetic/genetic.go#L32): No description provided.

3. [`Result`](./strings/genetic/genetic.go#L52): No description provided.


---
</details><details>
	<summary> <strong> geometry </strong> </summary>	

---

##### Functions:

1. [`Distance`](./math/geometry/straightlines.go#L17):  Calculates the shortest distance between two points.
2. [`Section`](./math/geometry/straightlines.go#L23):  Calculates the Point that divides a line in specific ratio. DO NOT specify the ratio in the form m:n, specify it as r, where r = m / n.
3. [`Slope`](./math/geometry/straightlines.go#L31):  Calculates the slope (gradient) of a line.
4. [`Intercept`](./math/geometry/straightlines.go#L36):  Calculates the Y-Intercept of a line from a specific Point.
5. [`IsParallel`](./math/geometry/straightlines.go#L41):  Checks if two lines are parallel or not.
6. [`IsPerpendicular`](./math/geometry/straightlines.go#L46):  Checks if two lines are perpendicular or not.
7. [`PointDistance`](./math/geometry/straightlines.go#L52):  Calculates the distance of a given Point from a given line. The slice should contain the coefficiet of x, the coefficient of y and the constant in the respective order.

---
##### Types

1. [`Point`](./math/geometry/straightlines.go#L8): No description provided.

2. [`Line`](./math/geometry/straightlines.go#L12): No description provided.


---
</details><details>
	<summary> <strong> graph </strong> </summary>	

---

#####  The Bellman–Ford algorithm is an algorithm that computes shortest paths from a single source vertex to all of the other vertices in a weighted durected graph. It is slower than Dijkstra but capable of handling negative edge weights. https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm Implementation is based on the book 'Introduction to Algorithms' (CLRS) Package graph demonstrates Graph search algorithms reference: https://en.wikipedia.org/wiki/Tree_traversal

---
##### Functions:

1. [`FloydWarshall`](./graph/floydwarshall.go#L15):  FloydWarshall Returns all pair's shortest path using Floyd Warshall algorithm
2. [`Topological`](./graph/topological.go#L7):  Assumes that graph given is valid and possible to get a topo ordering. constraints are array of []int{a, b}, representing an edge going from a to b
3. [`BreadthFirstSearch`](./graph/breadthfirstsearch.go#L9):  BreadthFirstSearch is an algorithm for traversing and searching graph data structures. It starts at an arbitrary node of a graph, and explores all of the neighbor nodes at the present depth prior to moving on to the nodes at the next depth level. Worst-case performance	 		O(|V|+|E|)=O(b^{d})}O(|V|+|E|)=O(b^{d}) Worst-case space complexity	 	O(|V|)=O(b^{d})}O(|V|)=O(b^{d}) reference: https://en.wikipedia.org/wiki/Breadth-first_search
4. [`GetIdx`](./graph/depthfirstsearch.go#L3): No description provided.
5. [`NotExist`](./graph/depthfirstsearch.go#L12): No description provided.
6. [`DepthFirstSearchHelper`](./graph/depthfirstsearch.go#L21): No description provided.
7. [`DepthFirstSearch`](./graph/depthfirstsearch.go#L53): No description provided.
8. [`New`](./graph/graph.go#L16):  Constructor functions for graphs (undirected by default)
9. [`NewDSU`](./graph/kruskal.go#L34):  NewDSU will return an initialised DSU using the value of n which will be treated as the number of elements out of which the DSU is being made
10. [`KruskalMST`](./graph/kruskal.go#L87):  KruskalMST will return a minimum spanning tree along with its total cost to using Kruskal's algorithm. Time complexity is O(m * log (n)) where m is the number of edges in the graph and n is number of nodes in it.

---
##### Types

1. [`Item`](./graph/dijkstra.go#L5): No description provided.

2. [`Graph`](./graph/graph.go#L9): No description provided.

3. [`Edge`](./graph/kruskal.go#L14): No description provided.

4. [`DisjointSetUnionElement`](./graph/kruskal.go#L21): No description provided.

5. [`DisjointSetUnion`](./graph/kruskal.go#L29): No description provided.

6. [`WeightedGraph`](./graph/floydwarshall.go#L9): No description provided.


---
</details><details>
	<summary> <strong> hashmap </strong> </summary>	

---

##### Functions:

1. [`New`](./structure/hashmap/hashmap.go#L24):  New return new HashMap instance

---
##### Types

1. [`HashMap`](./structure/hashmap/hashmap.go#L17): No description provided.


---
</details><details>
	<summary> <strong> kmp </strong> </summary>	

---

##### Functions:

1. [`Kmp`](./strings/kmp/kmp.go#L70):  Kmp Function kmp performing the Knuth-Morris-Pratt algorithm. Prints whether the word/pattern was found and on what position in the text or not. m - current match in text, i - current character in w, c - amount of comparisons.

---
##### Types

1. [`Result`](./strings/kmp/kmp.go#L15): No description provided.


---
</details><details>
	<summary> <strong> lcm </strong> </summary>	

---

##### Functions:

1. [`Lcm`](./math/lcm/lcm.go#L10):  Lcm returns the lcm of two numbers using the fact that lcm(a,b) * gcd(a,b) = | a * b |

---
</details><details>
	<summary> <strong> levenshtein </strong> </summary>	

---

##### Functions:

1. [`Distance`](./strings/levenshtein/levenshteindistance.go#L10):  Distance Function that gives Levenshtein Distance

---
</details><details>
	<summary> <strong> linkedlist </strong> </summary>	

---

#####  Package linkedlist demonstates different implementations on linkedlists.

---
##### Functions:

1. [`NewDoubly`](./structure/linkedlist/doubly.go#L22): No description provided.
2. [`NewNode`](./structure/linkedlist/shared.go#L12):  Create new node.
3. [`NewSingly`](./structure/linkedlist/singlylinkedlist.go#L19):  NewSingly returns a new instance of a linked list
4. [`NewCyclic`](./structure/linkedlist/cyclic.go#L12):  Create new list.
5. [`JosephusProblem`](./structure/linkedlist/cyclic.go#L120):  https://en.wikipedia.org/wiki/Josephus_problem This is a struct-based solution for Josephus problem.

---
##### Types

1. [`Node`](./structure/linkedlist/shared.go#L5): No description provided.

2. [`Singly`](./structure/linkedlist/singlylinkedlist.go#L10): No description provided.

3. [`Cyclic`](./structure/linkedlist/cyclic.go#L6): No description provided.

4. [`testCase`](./structure/linkedlist/cyclic_test.go#L105): No description provided.

5. [`Doubly`](./structure/linkedlist/doubly.go#L18): No description provided.


---
</details><details>
	<summary> <strong> manacher </strong> </summary>	

---

##### Functions:

1. [`LongestPalindrome`](./strings/manacher/longestpalindrome.go#L37): No description provided.

---
</details><details>
	<summary> <strong> math </strong> </summary>	

---

#####  Package math is a package that contains mathematical algorithms and its different implementations.

---
##### Functions:

1. [`Phi`](./math/eulertotient.go#L5):  Phi is the Euler totient function. This function computes the number of numbers less then n that are coprime with n.
2. [`IsPowOfTwoUseLog`](./math/checkisnumberpoweroftwo.go#L10):  IsPowOfTwoUseLog This function checks if a number is a power of two using the logarithm. The limiting degree can be from 0 to 63. See alternatives in the binary package.

---
</details><details>
	<summary> <strong> max </strong> </summary>	

---

##### Functions:

1. [`Int`](./math/max/max.go#L4):  Int is a function which returns the maximum of all the integers provided as arguments.
2. [`BitwiseMax`](./math/max/bitwisemax.go#L10): No description provided.

---
</details><details>
	<summary> <strong> maxsubarraysum </strong> </summary>	

---

#####  Package maxsubarraysum is a package containing a solution to a common problem of finding max contiguous sum within a array of ints.

---
##### Functions:

1. [`MaxSubarraySum`](./other/maxsubarraysum/maxsubarraysum.go#L13):  MaxSubarraySum returns the maximum subarray sum

---
</details><details>
	<summary> <strong> min </strong> </summary>	

---

##### Functions:

1. [`Int`](./math/min/min.go#L4):  Int is a function which returns the minimum of all the integers provided as arguments.

---
</details><details>
	<summary> <strong> modular </strong> </summary>	

---

##### Functions:

1. [`Exponentiation`](./math/modular/exponentiation.go#L22):  Exponentiation returns base^exponent % mod
2. [`Multiply64BitInt`](./math/modular/exponentiation.go#L51):  Multiply64BitInt Checking if the integer multiplication overflows
3. [`Inverse`](./math/modular/inverse.go#L20):  Inverse Modular function

---
</details><details>
	<summary> <strong> moserdebruijnsequence </strong> </summary>	

---

##### Functions:

1. [`MoserDeBruijnSequence`](./math/moserdebruijnsequence/sequence.go#L7): No description provided.

---
</details><details>
	<summary> <strong> nested </strong> </summary>	

---

#####  Package nested provides functions for testing strings proper brackets nesting.

---
##### Functions:

1. [`IsBalanced`](./other/nested/nestedbrackets.go#L20):  IsBalanced returns true if provided input string is properly nested. Input is a sequence of brackets: '(', ')', '[', ']', '{', '}'. A sequence of brackets `s` is considered properly nested if any of the following conditions are true: 	- `s` is empty; 	- `s` has the form (U) or [U] or {U} where U is a properly nested string; 	- `s` has the form VW where V and W are properly nested strings. For example, the string "()()[()]" is properly nested but "[(()]" is not. **Note** Providing characters other then brackets would return false, despite brackets sequence in the string. Make sure to filter input before usage.

---
</details><details>
	<summary> <strong> palindrome </strong> </summary>	

---

##### Functions:

1. [`IsPalindrome`](./strings/palindrome/ispalindrome.go#L26): No description provided.

---
</details><details>
	<summary> <strong> password </strong> </summary>	

---

#####  Package password contains functions to help generate random passwords

---
##### Functions:

1. [`Generate`](./other/password/generator.go#L15):  Generate returns a newly generated password

---
</details><details>
	<summary> <strong> permutation </strong> </summary>	

---

##### Functions:

1. [`Heaps`](./math/permutation/heaps.go#L8):  Heap's Algorithm for generating all permutations of n objects
2. [`GenerateElementSet`](./math/permutation/heaps.go#L37): No description provided.

---
</details><details>
	<summary> <strong> pi </strong> </summary>	

---

#####  spigotpi_test.go description: Test for Spigot Algorithm for the Digits of Pi author(s) [red_byte](https://github.com/i-redbyte) see spigotpi.go

---
##### Functions:

1. [`Spigot`](./math/pi/spigotpi.go#L12): No description provided.
2. [`MonteCarloPi`](./math/pi/montecarlopi.go#L15): No description provided.

---
</details><details>
	<summary> <strong> polybius </strong> </summary>	

---

#####  Package polybius is encrypting method with polybius square ref: https://en.wikipedia.org/wiki/Polybius_square#Hybrid_Polybius_Playfair_Cipher

---
##### Functions:

1. [`NewPolybius`](./cipher/polybius/polybius.go#L21):  NewPolybius returns a pointer to object of Polybius. If the size of "chars" is longer than "size", "chars" are truncated to "size".

---
##### Types

1. [`Polybius`](./cipher/polybius/polybius.go#L12): No description provided.


---
</details><details>
	<summary> <strong> power </strong> </summary>	

---

##### Functions:

1. [`IterativePower`](./math/power/fastexponent.go#L4):  IterativePower is iterative O(logn) function for pow(x, y)
2. [`RecursivePower`](./math/power/fastexponent.go#L18):  RecursivePower is recursive O(logn) function for pow(x, y)
3. [`RecursivePower1`](./math/power/fastexponent.go#L30):  RecursivePower1 is recursive O(n) function for pow(x, y)
4. [`UsingLog`](./math/power/powvialogarithm.go#L14): No description provided.

---
</details><details>
	<summary> <strong> prime </strong> </summary>	

---

##### Functions:

1. [`NaiveApproach`](./math/prime/primecheck.go#L8):  NaiveApproach checks if an integer is prime or not. Returns a bool.
2. [`PairApproach`](./math/prime/primecheck.go#L22):  PairApproach checks primality of an integer and returns a bool. More efficient than the naive approach as number of iterations are less.
3. [`Factorize`](./math/prime/primefactorization.go#L5):  Factorize is a function that computes the exponents of each prime in the prime factorization of n
4. [`GenerateChannel`](./math/prime/sieve.go#L9):  Generate generates the sequence of integers starting at 2 and sends it to the channel `ch`
5. [`Sieve`](./math/prime/sieve.go#L16):  Sieve Sieving the numbers that are not prime from the channel - basically removing them from the channels
6. [`Generate`](./math/prime/sieve.go#L26):  Generate returns a int slice of prime numbers up to the limit
7. [`MillerTest`](./math/prime/millerrabinprimalitytest.go#L32):  MillerTest This is the intermediate step that repeats within the miller rabin primality test for better probabilitic chances of receiving the correct result.
8. [`MillerRabinTest`](./math/prime/millerrabinprimalitytest.go#L59):  MillerRabinTest Probabilistic test for primality of an integer based of the algorithm devised by Miller and Rabin.

---
</details><details>
	<summary> <strong> pythagoras </strong> </summary>	

---

##### Functions:

1. [`Distance`](./math/pythagoras/pythagoras.go#L15): Distance calculates the distance between to vectors with the   Pythagoras theorem

---
##### Types

1. [`Vector`](./math/pythagoras/pythagoras.go#L8): No description provided.


---
</details><details>
	<summary> <strong> queue </strong> </summary>	

---

##### Functions:

1. [`EnQueue`](./structure/queue/queuearray.go#L15):  EnQueue it will be added new value into our list
2. [`DeQueue`](./structure/queue/queuearray.go#L20):  DeQueue it will be removed the first value that added into the list
3. [`FrontQueue`](./structure/queue/queuearray.go#L27):  FrontQueue return the Front value
4. [`BackQueue`](./structure/queue/queuearray.go#L32):  BackQueue return the Back value
5. [`LenQueue`](./structure/queue/queuearray.go#L37):  LenQueue will return the length of the queue list
6. [`IsEmptyQueue`](./structure/queue/queuearray.go#L42):  IsEmptyQueue check our list is empty or not

---
##### Types

1. [`Node`](./structure/queue/queuelinkedlist.go#L13): No description provided.

2. [`Queue`](./structure/queue/queuelinkedlist.go#L19): No description provided.

3. [`LQueue`](./structure/queue/queuelinklistwithlist.go#L20): No description provided.


---
</details><details>
	<summary> <strong> rsa </strong> </summary>	

---

#####  Package rsa shows a simple implementation of RSA algorithm

---
##### Functions:

1. [`Encrypt`](./cipher/rsa/rsa.go#L28):  Encrypt encrypts based on the RSA algorithm - uses modular exponentitation in math directory
2. [`Decrypt`](./cipher/rsa/rsa.go#L43):  Decrypt decrypts encrypted rune slice based on the RSA algorithm

---
</details><details>
	<summary> <strong> search </strong> </summary>	

---

##### Functions:

1. [`BoyerMoore`](./strings/search/boyermoore.go#L5):  Implementation of boyer moore string search O(l) where l=len(text)
2. [`Naive`](./strings/search/naive.go#L5):  Implementation of naive string search O(n*m) where n=len(txt) and m=len(pattern)

---
</details><details>
	<summary> <strong> segmenttree </strong> </summary>	

---

##### Functions:

1. [`NewSegmentTree`](./structure/segmenttree/segmenttree.go#L114): No description provided.

---
##### Types

1. [`SegmentTree`](./structure/segmenttree/segmenttree.go#L17): No description provided.


---
</details><details>
	<summary> <strong> set </strong> </summary>	

---

#####  package set implements a Set using a golang map. This implies that only the types that are accepted as valid map keys can be used as set elements. For instance, do not try to Add a slice, or the program will panic. package set implements a Set using a golang map. This implies that only the types that are accepted as valid map keys can be used as set elements. For instance, do not try to Add a slice, or the program will panic. 

---
##### Functions:

1. [`New`](./structure/set/set.go#L4):  New gives new set.

---
</details><details>
	<summary> <strong> sort </strong> </summary>	

---

#####  Package sort a package for demonstrating sorting algorithms in Go package sort provides primitives for sorting slices and user-defined collections

---
##### Functions:

1. [`HeapSort`](./sort/heapsort.go#L121): No description provided.
2. [`Exchange`](./sort/exchangesort.go#L6): No description provided.
3. [`SelectionSort`](./sort/selectionsort.go#L3): No description provided.
4. [`SimpleSort`](./sort/simplesort.go#L11): No description provided.
5. [`ImprovedSimpleSort`](./sort/simplesort.go#L25):  ImprovedSimpleSort is a improve SimpleSort by skipping an unnecessary comparison of the first and last. This improved version is more similar to implementation of insertion sort
6. [`Count`](./sort/countingsort.go#L10): No description provided.
7. [`InsertionSort`](./sort/insertionsort.go#L3): No description provided.
8. [`RadixSort`](./sort/radixsort.go#L35): No description provided.
9. [`Mergesort`](./sort/mergesort.go#L35): Mergesort Perform mergesort on a slice of ints
10. [`Pigeonhole`](./sort/pigeonholesort.go#L12):  Pigeonhole sorts a slice using pigeonhole sorting algorithm.
11. [`QuickSortRange`](./sort/quicksort.go#L24):  QuickSortRange Sorts the specified range within the array
12. [`QuickSort`](./sort/quicksort.go#L37):  QuickSort Sorts the entire array
13. [`ShellSort`](./sort/shellsort.go#L3): No description provided.

---
##### Types

1. [`MaxHeap`](./sort/heapsort.go#L3): No description provided.

2. [`Int`](#L0): 

	Methods:
	1. [`More`](./sort/heapsort.go#L114): No description provided.

---
</details><details>
	<summary> <strong> stack </strong> </summary>	

---

##### Types

1. [`Node`](./structure/stack/stacklinkedlist.go#L13): No description provided.

2. [`Stack`](./structure/stack/stacklinkedlist.go#L19): No description provided.

3. [`SList`](./structure/stack/stacklinkedlistwithlist.go#L18): No description provided.


---
</details><details>
	<summary> <strong> transposition </strong> </summary>	

---

##### Functions:

1. [`Encrypt`](./cipher/transposition/transposition.go#L54): No description provided.
2. [`Decrypt`](./cipher/transposition/transposition.go#L82): No description provided.

---
##### Types

1. [`NoTextToEncryptError`](./cipher/transposition/transposition.go#L15): No description provided.

2. [`KeyMissingError`](./cipher/transposition/transposition.go#L16): No description provided.


---
</details><details>
	<summary> <strong> trie </strong> </summary>	

---

#####  Package trie provides Trie data structures in golang.  Wikipedia: https://en.wikipedia.org/wiki/Trie

---
##### Functions:

1. [`NewNode`](./structure/trie/trie.go#L14):  NewNode creates a new Trie node with initialized children map.

---
##### Types

1. [`Node`](./structure/trie/trie.go#L7): No description provided.


---
</details><details>
	<summary> <strong> xor </strong> </summary>	

---

#####  Package xor is an encryption algorithm that operates the exclusive disjunction(XOR) ref: https://en.wikipedia.org/wiki/XOR_cipher

---
##### Functions:

1. [`Encrypt`](./cipher/xor/xor.go#L10):  Encrypt encrypts with Xor encryption after converting each character to byte The returned value might not be readable because there is no guarantee which is within the ASCII range If using other type such as string, []int, or some other types, add the statements for converting the type to []byte.
2. [`Decrypt`](./cipher/xor/xor.go#L19):  Decrypt decrypts with Xor encryption

---
</details>
<!--- GODOCMD END --->

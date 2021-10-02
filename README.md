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

### Ciphers
* [Caesar](./cipher/caesar/caesar.go)
* [Diffie Hellman Key Exchange](./cipher/diffiehellman/diffiehellmankeyexchange.go)
* [Polybius](./cipher/polybius/polybius.go)
* [Rot13](./cipher/rot13/rot13.go)
* [Rsa](./cipher/rsa/rsa.go)
* [Xor](./cipher/xor/xor.go)

### Conversions
* [Roman Numeral To Integer](./conversion/romantointeger.go)
* [Integer To Roman Numeral](./conversion/integertoroman.go)

### Data Structures
* [AVL Tree](./structure/avl/avl.go)
* [Binary Tree](./structure/binarytree/btree.go)
* [Dynamic Array](./structure/dynamicarray/dynamicarray.go)
* [Hashmap](./structure/hashmap/hashmap.go)
* Linked-List
    * [Doubly Linked List](./structure/linkedlist/doubly.go)
    * [Singly Linked List](./structure/linkedlist/singlylinkedlist.go)
    * [Cyclic Linked List AKA Looped Linked List](./structure/linkedlist/cyclic.go)
* [Queue](./structure/queue)
    * [Array Based](./structure/queue/queuearray.go)
    * [Custom Linked List Based](./structure/queue/queuelinkedlist.go)
    * [Standard Library Container/List Based](./structure/queue/queuelinkedlist.go)
* [Set](./structure/set/set.go)
* [Stack](./structure/stack)
    * [Array Based](./structure/stack/stackarray.go)
    * [Custom Linked List](./structure/stack/stacklinkedlist.go)
    * [Standard Library Container/List Based](structure/stack/stacklinkedlistwithlist.go)
* [Trie](./structure/trie/trie.go)

### Dynamic Programming
* [Binomial Coefficient](./dynamic/binomialcoefficient.go)
* [Fibonacci](./dynamic/fibonacci.go)
* [Knapsack](./dynamic/knapsack.go)
* [Longest Common Subsequence](./dynamic/longestcommonsubsequence.go)
* [Longest Palindromic Subsequence](./dynamic/longestpalindromicsubsequence.go)
* [Matrix Multiplication](./dynamic/matrixmultiplication.go)
* [Cutting a Rod](./dynamic/rodcutting.go)

### Graphs
* [Breadth First Search](./graph/breadthfirstsearch.go)
* [Depth First Search](./graph/depthfirstsearch.go)
* [Floyd Warshall](./graph/floydwarshall.go)
* [Coloring Algorithms](./graph/coloring)
  * [Using BFS](./graph/coloring/bfs.go)
  * [Using Backtracking](./graph/coloring/backtracking.go)
  * [Using Greedy Approach](./graph/coloring/greedy.go)

### Math
* [Gcd](./math/gcd)
    * [Extended](./math/gcd/extended.go)
    * [Recursive](./math/gcd/gcd.go)
    * [Iterative](./math/gcd/gcditerative.go)
* [Lcm](./math/lcm/lcm.go)
* [Max](./math/max/max.go)
* [Modular](./math/modular)
    * [Exponentiation](./math/max/max.go)
    * [Inverse](./math/modular/inverse.go) 
* [Moser-de Bruijn sequence](./math/moserdebruijnsequence/sequence.go)
* [Permutation](./math/permutation/heaps.go)
* [Power](./math/power/fastexponent.go)
* [Prime](./math/prime)
    * [Miller Rabin Primality Test](./math/prime/millerrabinprimalitytest.go)
    * [Naive and Pair Approach](./math/prime/primecheck.go)
    * [Sieve](./math/prime/sieve.go)
* [Pythagoras](./math/pythagoras/pythagoras.go)
* [Sieve](./math/prime/sieve.go)
* [Straight Lines](./math/straightlines/straightlines.go)

### Other
* [Max Subarray Sum](./other/maxsubarraysum/maxsubarraysum.go)
* [Nested Brackets](./other/nested/nestedbrackets.go)
* [Password Generator](./other/password/generator.go)

### Searches
* [Binary Search](./search/binary.go)
* [Linear Search](./search/linear.go)
* [Interpolation Search](./search/interpolation.go)

### Sorts
* [Bubble Sort](./sort/bubblesort.go)
* [Heap Sort](./sort/heapsort.go)
* [Insertion Sort](./sort/insertionsort.go)
* [Merge Sort](./sort/mergesort.go)
* [Quick Sort](./sort/quicksort.go)
* [Radix Sort](./sort/radixsort.go)
* [Selection Sort](./sort/selectionsort.go)
* [Shell Sort](./sort/shellsort.go)


* [Sorts Test](./sort/sorts_test.go)

### Strings
* [Levenshtein Distance](./strings/levenshtein/levenshteindistance.go)
* Multiple String Matching
    * [Advanced Aho Corasick](./strings/ahocorasick/advancedahocorasick.go)
    * [Aho Corasick](./strings/ahocorasick/ahocorasick.go)
* Single String Matching
    * [Bom](./strings/bom/bom.go)
    * [Horspool](./strings/horspool/horspool.go)
    * [Kmp](./strings/kmp/kmp.go)

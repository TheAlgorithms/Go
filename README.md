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
* [Caesar](./cipher/caesar)
* [Diffie Hellman Key Exchange](./cipher/diffiehellman)
* [Polybius](./cipher/polybius)
* [Rot13](./cipher/rot13)
* [Rsa](./cipher/rsa)
* [Xor](./cipher/xor)

### Conversions
* [Roman To Integer](./conversion/romantointeger.go)

### Data Structures
* [AVL Tree](./structure/avl)
* [Binary Tree](./structure/binarytree)
* [Dynamic Array](./structure/dynamicarray)
* [Hashmap](./structure/hashmap)
* Linked-List
    * [Doubly Linked List](./structure/linkedlist/doubly.go)
    * [Singly Linked List](./structure/linkedlist/singlylinkedlist.go)
    * [Cyclic Linked List AKA Looped Linked List](./structure/linkedlist/cyclic.go)
* [Queue](./structure/queue) | Array, Linked-List, STL(Container/List), Test
* [Set](./structure/set)
* [Stack](./structure/stack) | Array, Linked-List, STL(Container/List), Test
* [Trie](./structure/trie)

### Dynamic Programming
* [Binomial Coefficient](./dynamic/binomialcoefficient.go)
* [Fibonacci](./dynamic/fibonacci.go) | [Test](./dynamic/fibonacci_test.go)
* [Knapsack](./dynamic/knapsack.go)
* [Longest Common Subsequence](./dynamic/longestcommonsubsequence.go)
* [Longest Palindromic Subsequence](./dynamic/longestpalindromicsubsequence.go)
* [Matrix Multiplication](./dynamic/matrixmultiplication.go)
* [Cutting a Rod](./dynamic/rodcutting.go)

### Graphs
* [Breadth First Search](./graph/breadthfirstsearch.go)
* [Depth First Search](./graph/depthfirstsearch.go)
* [Floyd Warshall](./graph/floydwarshall.go)
    
### Math
* [Gcd](./math/gcd)
* [Lcm](./math/lcm)
* [Max](./math/max)
* [Modular](./math/modular) 
* [Moser-de Bruijn sequence](./math/moserdebruijnsequence)
* [Permutation](./math/permutation)
* [Power](./math/power)
* [Prime](./math/prime)
* [Pythagoras](./math/pythagoras)
* [Sieve](./math/prime/sieve.go)

### Other
* [Max Subarray Sum](./other/maxsubarraysum)
* [Nested Brackets](./other/nested)
* [Password Generator](./other/password)

### Searches
* [Binary Search](./search/binary.go)
* [Linear Search](./search/linear.go)
* [Interpolation Search](./search/interpolation.go)

## Sorts
* [Bubble Sort](./sort/bubblesort.go)
* [Heap Sort](./sort/heapsort.go)
* [Insertion Sort](./sort/insertionsort.go)
* [Merge Sort](./sort/mergesort.go)
* [Quick Sort](./sort/quicksort.go)
* [Radix Sort](./sort/radixsort.go)
* [Selection Sort](./sort/selectionsort.go)
* [Shell Sort](./sort/shellsort.go)


* [Sorts Test](./sort/sorts_test.go)

## Strings
* [Levenshtein Distance](./strings/levenshtein)
* Multiple String Matching
    * [Advanced Aho Corasick](./strings/ahocorasick/advancedahocorasick.go)
    * [Aho Corasick](./strings/ahocorasick/ahocorasick.go)
* Single String Matching
    * [Bom](./strings/bom)
    * [Horspool](./strings/horspool)
    * [Kmp](./strings/kmp)

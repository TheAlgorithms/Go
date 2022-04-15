# The Algorithms Go Repository Style Guide

This guide contains a set of go idioms and best practices which should be followed while writing
code whenever they are applicable. The instructions in this document should be given precedence
if there is a conflict between it and `CONTRIBUTING.md`. If you find any issue or ambiguity in
this document, the maintainers of this repository should be consulted.

## Formatting

All go code should be formatted with the official formatting tool `gofmt`. This requirement is
implemented and verified by the repository workflows.

It is also required for the code to pass the `go vet` checks.

## Commenting

All exported symbols should be commented with documentation, so that their motivation and use is
clear. All documentation should record any nuances of the symbol so that users are well aware of
them.

Line comments should be preferred over C-style block comments.

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
  
```go
/*
  Unmarshal converts a JSON string into a go interface
  ...
*/
```
  
</td><td>

```go
// Unmarshal converts a JSON string into a go interface
// ...
```
 
</td></tr>
</tbody></table>

Package comments should start with the word "Package" followed by the package name. For packages
with multiple files or with a large documentation, use a separate `doc.go` file for the  comment.

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
  
```go
// sort is a package which implements sorting functions.
```
  
</td><td>

```go
// Package sort implements sorting functions.
```
 
</td></tr>
</tbody></table>

All doc comments for symbols should start with the symbol name.

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
  
```go
// Function Quick is an implementation
// of the Quicksort algorithm.
```
  
</td><td>

```go
// Quick is an implementation
// of the Quicksort algorithm.
```
 
</td></tr>
</tbody></table>

# The Algorithms Go Repository Style Guide

This guide contains a set of go idioms and best practices which should be followed while writing
code whenever they are applicable. The instructions in this document should be given precedence
if there is a conflict between it and `CONTRIBUTING.md`. If you find any issue or ambiguity in
this document, the maintainers of this repository should be consulted.

## Table of Contents

- [Formatting](#formatting)
- [Commenting](#commenting)
  - [Package Comments](#package-comments)
  - [Documentation Comments](#documentation-comments)
  - [Author Comments](#author-comments)
- [Naming](#naming)
  - [Package Naming](#package-naming)
  - [Symbol Naming](#symbol-naming)
  - [Function Naming](#function-naming)
      - [Constructor Naming](#constructor-naming)
      - [Getter Naming](#getter-naming)
      - [Setter Naming](#setter-naming) 
  - [Interface Naming](#interface-naming)

## Formatting

All go code should be formatted with the official formatting tool `gofmt`. This requirement is
implemented and verified by the repository workflows.

It is also required for the code to pass the `go vet` checks.

## Commenting

All exported symbols should be commented with documentation, so that their motivation and use is
clear. All documentation should record any nuances of the symbol so that users are well aware of
them.

C++ style line comments should be preferred over C-style block comments.

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

### Package Comments

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

### Documentation Comments

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

### Author Comments

A comment recording the author of a particular file may also be added. This comment should be
written at the top of the file and it should not be a part of the package documentation.

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
  
```go
// author: Rak Laptudirm (@raklaptudirm)
package sort
```
Comment is a part of the package documentation.
  
</td><td>

```go
// author: Rak Laptudirm (@raklaptudirm)

package sort
```
Comment is not a part of the package documentation.
 
</td></tr>
</tbody></table>

## Naming

### Package Naming

Package names should be short and to the point. Keep in mind that this identifier will be used to
refer to your package, so make it easy to remember. It should be only composed of lower case
letters. Prefer `json` to `JSON` or `Json`. If your package name has two words, merge them
together. Prefer `jsonencodeing` to `jsonEncoding` or `json_encodeing`.

Add the `_test` suffix to your package name to implement black-box testing for your package.

### Symbol Naming

Go symbols should be named in the `camelCase` or `UpperCamelCase`, depending of whether the symbol
is exported or not. The case when using acronymns in names should be consistent. Use `json` or
`JSON` instead of `Json`.

For exported symbols, use the package name to your advantage to write concise symbol names. For
example, if you have a package `png`, instead of defining a `png.PNGFile`, define a `png.File`.
It is much clearer and avoids repetition.

### Function Naming

#### Constructor Naming

Constructors should use the naming format `New<Type>()`. If the package only exports a single
constructor, use the name `New()`. Some valid names are `reader.New()`, `metainfo.NewFile()`.

#### Getter Naming

Getters should use the name of the field that is being fetched. For example, prefer a getter
name like `user.Name()` to `user.GetName()`.

#### Setter Naming

Setters should use names in the format `Set<Field>`. For example, `user.SetName()`.

### Interface Naming

Interfaces should use names in the format of `<Method>er` or `<Method>`. For example, some
valid interface names are:

```go
type Node interface {
  Node()
}
 
type Writer interface {
  Write()
}
```

# CONTRIBUTION GUIDELINES

## Before contributing

Welcome to [TheAlgorithms/Go](https://github.com/TheAlgorithms/Go)! Before submitting pull requests, please make sure that you have **read the whole guidelines**. If you have any doubts about this contribution guide, please open [an issue](https://github.com/TheAlgorithms/Go/issues/new/choose) or ask in our [Discord server](https://the-algorithms.com/discord/) / [Gitter](https://gitter.im/TheAlgorithms), and clearly state your concerns.

## Contributing

### Maintainers
---

Please check the [`CODEOWNERS`](https://github.com/TheAlgorithms/Go/blob/master/.github/CODEOWNERS) file for a list of repository maintainers.

### Contributor
---

Being a contributor at The Algorithms, we request you to follow the points mentioned below:

- You did your own work.
  - No plagiarism is allowed. Any plagiarized work will not be merged.
- Your work will be distributed under the [MIT License](https://github.com/TheAlgorithms/Go/blob/master/LICENSE) once your pull request has been merged.
- Please follow the repository guidelines and standards mentioned below.

#### New implementation

- New implementations are welcome!

- You can add new algorithms or data structures that are **not present in the repository** or that can **improve** the old implementations (**documentation**, **improving test cases**, **removing bugs**, or in any other reasonable sense)

#### Issues

- Please avoid opening issues asking to be "assigned” to a particular algorithm. This merely creates unnecessary noise for maintainers. Instead, please submit your implementation in a pull request, and it will be evaluated by project maintainers.

### Making Changes
---

#### Code

- Please use the directory structure of the repository.
- Make sure the file extensions should be `*.go`.
- Use meaningful variable names.
- Use standard library inside your code and avoid to import packages from other repositories
- If an implementation of the algorithm already exists, please refer to the [filename section below](#new-file-name-guidelines).
- You can suggest reasonable changes to existing algorithms.
- Strictly use snake_case (underscore_separated)  in filenames.
- If you have added or modified code, please make sure the code compiles before submitting.
- Our automated testing runs [LGTM](https://lgtm.com/projects/g/TheAlgorithms/Go/) on all the pull requests, so please be sure that your code passes before submitting.
- Please conform to [GoDoc](https://blog.golang.org/godoc) standard and document the code as much as possible. This not only facilitates the readers but also generates the correct info on the website.
- **Be consistent in the use of these guidelines.**

#### Documentation

- Make sure you put useful comments in your code. Do not comment on obvious things.
- Please avoid creating new directories if at all possible. Try to fit your work into the existing directory structure. If you want to create a new directory, then please check if a similar category has been recently suggested or created by other pull requests.
- If you have modified/added documentation, please ensure that your language is concise and must not contain grammatical errors.
- Do not update [`README.md`](https://github.com/TheAlgorithms/Go/blob/master/README.md) along with other changes. First, create an issue and then link to that issue in your pull request to suggest specific changes required to [`README.md`](https://github.com/TheAlgorithms/Go/blob/master/README.md).
- The repository follows [GoDoc](https://blog.golang.org/godoc) standards and auto-generates the [repository website](https://thealgorithms.github.io/). Please ensure the code is documented in this structure. A sample implementation is given below.

#### Test

- Make sure to add examples and test cases in your `filename_test.go` file.
- If you find an algorithm or document without tests, please feel free to create a pull request or issue describing suggested changes.
- Please try to add one or more `Test` functions that will invoke the algorithm implementation on random test data with the expected output.

### Benchmark
---

- Make sure to add examples and benchmark cases in your `filename_test.go` or `filename_bench.go` if you want separated test and benchmark files.
- If you find an algorithm or document without benchmarks, please feel free to create a pull request or issue describing suggested changes.
- Please try to add one or more `Benchmark` functions that will invoke the algorithm implementation.
- For running the benchmark, you could use this command `go test -bench=.` for more details, read this article [Using Subtests and Sub-benchmarks](https://go.dev/blog/subtests)

#### Typical structure of a program

```go

// filename
// description: Add one line description here
// details:
// This is a multi-line
// description containing links, references,
// math equations, etc.
// author(s) [Name](https://github.com/handle), [Name](https://github.com/handle)
// see relatedfile.go, anotherfile.go, file_test.go
 
// ** Is just an example of how to write description for package and function and other stuff ... **

// Package sort provides primitives for sorting slices and user-defined
// collections.
package sort

// Reasons why you used those packages 
// name package : Add one line description here
// name2 package : Add one line description here
// ...
import (
   ...
)


// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written, and any write error encountered.
func Fprint(w io.Writer, a ...any) (n int, err error) {
    ...
}


```

#### New File Name guidelines

- Use lowercase words without ``"_"`` for the file name
- Use ``"_"`` as a separator only for `_test.go` or `_bench.go`
- For instance

```markdown
MyNewGoFile.GO       is incorrect
my_new_go_file.go    is incorrect
mynewgofile.go       is the correct format
mynewgofile_test.go  is the correct format
```

- It will be used to dynamically create a directory of files and implementation.
- File name validation will run on Docker to ensure validity.
- If an implementation of the algorithm already exists and your version is different from that implemented, please use incremental numeric digit as a suffix. For example: if `binarysearch.go` already exists in the `search` folder, and you are contributing a new implementation, the filename should be `binarysearch2.go` and for a third implementation, `binarysearch3.go`.
- Check out `Go` [Package names](https://go.dev/blog/package-names) roles

#### New Directory guidelines

- We recommend adding files to existing directories as much as possible.
- Use lowercase words with ``"_"`` as separator ( no spaces or ```"-"``` allowed )
- For instance

```markdown
SomeNew Fancy-Category          is incorrect
some_new_fancy_category         is correct
```

- Filepaths will be used to dynamically create a directory of our algorithms.
- Filepath validation will run on GitHub Actions to ensure compliance.

#### Commit Guidelines

- It is recommended to keep your changes grouped logically within individual commits. Maintainers find it easier to understand changes that are logically spilled across multiple commits. Try to modify just one or two files in the same directory. Pull requests that span multiple directories are often rejected.

```bash
git add filexyz.go
git commit -m "your message"
```

Examples of commit messages with semantic prefixes:

```markdown
fix: XYZ algorithm bug
feat: add XYZ algorithm
test: add test for XYZ algorithm
docs: add comments and explanation to XYZ algorithm
```

Common prefixes:

- fix: A bug fix
- feat: A new feature
- docs: Documentation changes
- test: Correct existing tests or add new ones

### For New Gophers
---

#### Installing Go

- Installation (only needs to be installed once.)
  - Mac (using home-brew): `brew install go`
  - Windows (MSYS2 64-bit): `choco install golang` [Chocolatey Package Manager](https://chocolatey.org/)
  - Linux (Debian): `sudo apt-get install golang`
  - Manual Installation: [Downloads - The Go Programming Language](https://golang.org/dl/)
- Running (all platforms): `go run filexyz.go`

> Note: New packages should not be `main`, and never implement `main` functions for any package.

#### Code formatter
To format your code, you can use the gofmt tool directly:

```bash
gofmt -w filexyz.go
```

Or you can use the "go fmt" command:

```bash
go fmt path/to/your/package
```

#### Building

To build you code locally, simply run:

```bash
go build .
```

### Pull Requests
---

- Check out our [pull request template](https://github.com/TheAlgorithms/Go/blob/master/.github/PULL_REQUEST_TEMPLATE/pull_request.md
  )

#### Building Locally

Before submitting a pull request, [build the code locally](#building) or using the convenient [![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/TheAlgorithms/Go) service.

#### GitHub Actions

- Enable GitHub Actions on your fork of the repository.
  After enabling, it will execute `golang_lint_and_test` after every push (not a commit).
- The result can create another commit if the actions made any changes on your behalf.
- Hence, it is better to wait and check the results of GitHub Actions after every push.
- Run `git pull` in your local clone if these actions made many changes to avoid merge conflicts.

Most importantly,

- Happy coding!

[![GoDoc](https://godoc.org/github.com/aadamandersson/multiset?status.svg)](https://godoc.org/github.com/aadamandersson/multiset)
![CI](https://github.com/aadamandersson/multiset/actions/workflows/ci.yml/badge.svg)
[![Codecov](https://codecov.io/github/aadamandersson/multiset/coverage.svg?branch=main)](https://codecov.io/gh/aadamandersson/multiset)

An unordered multiset implementation.

A multiset, unlike a set, allows values to occur multiple times, and keep tracks of the number of duplicates a value holds.

## Installation

Run `go get github.com/aadamandersson/multiset` to add this package as a dependency to your project.

## Example

```go
package main

import (
	"fmt"

	"github.com/aadamandersson/multiset"
)

func main() {
	// Add some books and keep track of all the copies (since we are a library)
	books := multiset.New[string]()

	books.Insert("The Alchemist")
	books.Insert("The Alchemist")
	books.Insert("The Alchemist")
	books.Insert("The Alchemist")
	books.Insert("To Kill a Mockingbird")
	books.Insert("To Kill a Mockingbird")
	books.Insert("To Kill a Mockingbird")
	books.Insert("Metamorphoses")
	books.Insert("Metamorphoses")

	// Duplicates are included
	fmt.Println(books.Len()) // 9

	// Look for a specific book
	if books.Contains("The Alchemist") == 0 {
		fmt.Println("Could not find any copies of The Alchemist.")
	}

	// Remove a book
	copies := books.Remove("Metamorphoses")
	fmt.Println(copies) // 2

	// Iterate over all of the books
	books.Each(func(book string, copies int) bool {
		fmt.Printf("%s (%d copies)\n", book, copies)
		return false
	})
}
```

More examples can be found [here](https://pkg.go.dev/github.com/aadamandersson/multiset#section-documentation).

## License

MIT licensed. See the [LICENSE](./LICENSE) file for details.

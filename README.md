# SkipList

[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)


ref: https://en.wikipedia.org/wiki/Skip_list

## example:

```go
import (
	"fmt"
	"github.com/MaruHyl/go-skiplist"
)

func ExampleSkipList() {
	intCompare := func(a interface{}, b interface{}) int {
		return a.(int) - b.(int)
	}
	l := skiplist.New(intCompare)

	// insert
	l.Insert(3, "value 1")
	l.Insert(1, "value 1")
	l.Insert(2, "value 1")

	// delete
	l.Delete(2)

	// get
	fmt.Println(l.Search(1))

	// foreach
	l.Foreach(func(key interface{}, value interface{}) {
		fmt.Println(key, ":", value)
	})

	// Output:
	// value 1 true
	// 1 : value 1
	// 3 : value 1
}
```

## benchmark:

```go
goos: darwin
goarch: amd64
pkg: github.com/MaruHyl/go-skiplist
BenchmarkSkipList_Insert-12      5000000               435 ns/op
BenchmarkSkipList_Search-12      5000000               285 ns/op
BenchmarkSkipList_Delete-12     10000000               199 ns/op
```

# Indexable skiplist
TODO

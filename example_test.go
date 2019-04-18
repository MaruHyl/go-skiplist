package skiplist_test

import (
	"fmt"
	"github.com/MaruHyl/go-skiplist"
)

func ExampleSkipList() {
	intCompare := func(a interface{}, b interface{}) bool {
		return a.(int) < b.(int)
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

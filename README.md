# SkipList
参考 https://en.wikipedia.org/wiki/Skip_list

example:

```go
func intCompare(a interface{},b interface{}) int{
	inta:=a.(int)
	intb:=b.(int)
	return inta-intb
}
l:=NewSkipList(intCompare)
l.Insert(1,1)
v,ok:=l.Search(1)
v,ok:=l.Delete(1)
```

benchmark:

insert: 3000000	       444 ns/op

search: 5000000	       363 ns/op

delete: 10000000	     186 ns/op



# Indexable skiplist
TODO

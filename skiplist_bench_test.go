package SkipList

import (
	"testing"
)


func BenchmarkSkipList_Insert(b *testing.B) {
	l:=NewSkipList(intCompare)
	for i:=0;i<b.N;i++{
		l.Insert(i,i)
	}
}

func BenchmarkSkipList_Search(b *testing.B) {
	b.StopTimer()
	l:=NewSkipList(intCompare)
	for i:=0;i<b.N;i++{
		l.Insert(i,i)
	}
	b.StartTimer()
	for i:=0;i<b.N;i++{
		l.Search(i)
	}
}

func BenchmarkSkipList_Delete(b *testing.B) {
	b.StopTimer()
	l:=NewSkipList(intCompare)
	for i:=0;i<b.N;i++{
		l.Insert(i,i)
	}
	b.StartTimer()
	for i:=0;i<b.N;i++{
		l.Delete(i)
	}
}

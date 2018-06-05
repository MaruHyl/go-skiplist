package SkipList

import (
	"testing"
	"strings"
	"fmt"
	"math/rand"
	"time"
)



func println(l *SkipList){
	var b strings.Builder
	for i:=0;i<maxLevel;i++{
		x:=l.header
		for x!=nil{
			if x==l.header{
				fmt.Fprintf(&b,"[%d] header -> ",i)
			}else{
				fmt.Fprintf(&b,"(%d:%d) -> ",x.key,x.value)
			}
			x=x.forward[i]
		}
		fmt.Fprintf(&b,"\n")
	}
	fmt.Println(b.String())
}

func intCompare(a interface{},b interface{}) int{
	inta:=a.(int)
	intb:=b.(int)
	return inta-intb
}

func TestSkipList_Insert(t *testing.T) {
	l:=NewSkipList(intCompare)
	l.Insert(1,1)
	l.Insert(2,2)
	l.Insert(3,3)
	l.Insert(0,0)
	println(l)
}

func checkList(l *SkipList,dm map[int]int){
	// 检查键值对
	if l.Len()!=len(dm){
		panic("wrong size")
	}
	for k,v:=range dm{
		lv,ok:=l.Search(k)
		if !ok{
			panic("wrong key")
		}
		if lv!=v{
			panic("wrong value")
		}
	}
	// 检查所有层是否有序
	for i:=0;i<maxLevel;i++{
		x:=l.header
		for x.forward[i]!=nil{
			if x!=l.header{
				if intCompare(x.key,x.forward[i].key)>=0{
					panic("wrong sort")
				}
			}
			x=x.forward[i]
		}
	}
	fmt.Println("check pass")
}

func TestSkipList_Delete(t *testing.T) {
	l:=NewSkipList(intCompare)
	l.Insert(1,1)
	l.Insert(2,2)
	l.Insert(3,3)
	l.Insert(0,0)

	l.Delete(1)
	println(l)
}

// 随机插入测试正确性
func TestSkipList_InsertRandom(t *testing.T) {
	l:=NewSkipList(intCompare)
	testCount:=10000
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	dm:=make(map[int]int,testCount)
	for i:=0;i<testCount;i++{
		k:=r.Intn(testCount)
		v:=i
		dm[k]=v
		l.Insert(k,v)
	}
	checkList(l,dm)
	println(l)
}

// 随机插入和删除测试正确性
func TestSkipList_DeleteRandom(t *testing.T) {
	l:=NewSkipList(intCompare)
	testCount:=10000
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	dm:=make(map[int]int,testCount)
	for i:=0;i<testCount;i++{
		k:=r.Intn(testCount)
		v:=i
		dm[k]=v
		l.Insert(k,v)
	}
	for k,v:=range dm{
		// 随机删除
		if r.Float32()>0.5{
			lv,ok:=l.Delete(k)
			delete(dm,k)
			if !ok || v!=lv{
				panic("wrong delete")
			}
		}
	}
	checkList(l,dm)
	println(l)
}
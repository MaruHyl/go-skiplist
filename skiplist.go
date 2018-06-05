package SkipList

import (
	"math/rand"
	"time"
)

// skiplist的实现,存储k-v对,且key不能重复,以key作降序比较


// 如果a>b 返回正数
// 如果a=b 返回0
// 如果a<b 返回负数
type CompareFunc func(keyA interface{}, keyB interface{}) int

func compare(n *node, key interface{}, f CompareFunc) int {
	// nil永远大于key
	if n == nil {
		return 1
	}
	return f(n.key, key)
}

const maxLevel = 20
const p = 0.5

type SkipList struct {
	header *node
	CompareFunc
	r *rand.Rand
	length int
}

func NewSkipList(f CompareFunc) *SkipList {
	l := &SkipList{
		CompareFunc: f,
		r:           rand.New(rand.NewSource(time.Now().UnixNano())),
		header:      makeNode(maxLevel, nil, nil),
	}
	return l
}

type node struct {
	key     interface{}
	value   interface{}
	forward []*node
}

// 跳表的标准查找
func (l *SkipList) Search(key interface{}) (value interface{}, ok bool) {
	x := l.header
	for i := maxLevel - 1; i >= 0; i-- {
		for compare(x.forward[i], key, l.CompareFunc) < 0 {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if compare(x, key, l.CompareFunc) == 0 {
		value = x.value
		ok = true
	}
	return
}

func (l *SkipList) randomLevel() (level int) {
	level++
	for maxLevel > level && l.r.Float32() >= p {
		level++
	}
	return
}

func makeNode(level int, key interface{}, value interface{}) *node {
	return &node{
		key:     key,
		value:   value,
		forward: make([]*node, level),
	}
}

// 找到该key的所有前置节点
// 如果第0层的key等于该key，则刷新值,否则逐层插入链表
func (l *SkipList) Insert(key interface{}, value interface{}) {
	update := make([]*node, maxLevel)
	x := l.header
	for i := maxLevel - 1; i >= 0; i-- {
		for compare(x.forward[i], key, l.CompareFunc) < 0 {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if compare(x, key, l.CompareFunc) == 0 {
		x.value = value
	} else {
		level := l.randomLevel()
		x := makeNode(level, key, value)
		for i := 0; i < level; i++ {
			x.forward[i] = update[i].forward[i]
			update[i].forward[i] = x
		}
		l.length++
	}
}

// 找到该key的所有前置节点
// 如果第0层的key等于该key，则逐层拆链，否则直接返回
func (l *SkipList) Delete(key interface{}) (value interface{}, ok bool) {
	update := make([]*node, maxLevel)
	x := l.header
	for i := maxLevel - 1; i >= 0; i-- {
		for compare(x.forward[i], key, l.CompareFunc) < 0 {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if compare(x, key, l.CompareFunc) == 0 {
		for i := 0; i < maxLevel; i++ {
			if update[i].forward[i]!=x {
				break
			}
			update[i].forward[i] = x.forward[i]
		}
		value = x.value
		ok = true
		l.length--
	}
	return
}

func (l *SkipList) Len() int{
	return l.length
}




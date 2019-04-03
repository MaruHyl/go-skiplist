package skiplist

import (
	"math/rand"
)

// if a>b then comp>0
// if a=b then comp=0
// if a<b then comp<0
type CompareFunc func(a interface{}, b interface{}) (comp int)

func compare(n *node, key interface{}, f CompareFunc) int {
	if n == nil {
		return 1
	}
	return f(n.key, key)
}

const maxLevel = 64
const p = 0.5

type node struct {
	key     interface{}
	value   interface{}
	forward []*node
}

type SkipList struct {
	header      *node
	compareFunc CompareFunc
	length      int
}

func New(f CompareFunc) *SkipList {
	l := &SkipList{
		compareFunc: f,
		header:      makeNode(maxLevel, nil, nil),
	}
	return l
}

func (l *SkipList) Search(key interface{}) (interface{}, bool) {
	x := l.header
	for i := maxLevel - 1; i >= 0; i-- {
		for compare(x.forward[i], key, l.compareFunc) < 0 {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if compare(x, key, l.compareFunc) == 0 {
		return x.value, true
	}
	return nil, false
}

func (l *SkipList) randomLevel() (level int) {
	level++
	for maxLevel > level && rand.Float32() >= p {
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

func (l *SkipList) Insert(key interface{}, value interface{}) {
	update := make([]*node, maxLevel)
	x := l.header
	for i := maxLevel - 1; i >= 0; i-- {
		for compare(x.forward[i], key, l.compareFunc) < 0 {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if compare(x, key, l.compareFunc) == 0 {
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

func (l *SkipList) Delete(key interface{}) (interface{}, bool) {
	update := make([]*node, maxLevel)
	x := l.header
	for i := maxLevel - 1; i >= 0; i-- {
		for compare(x.forward[i], key, l.compareFunc) < 0 {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if compare(x, key, l.compareFunc) == 0 {
		for i := 0; i < maxLevel; i++ {
			if update[i].forward[i] != x {
				break
			}
			update[i].forward[i] = x.forward[i]
		}
		l.length--
		return x.value, true
	}
	return nil, false
}

func (l *SkipList) Len() int {
	return l.length
}

func (l *SkipList) Foreach(f func(key interface{}, value interface{})) {
	for x := l.header.forward[0]; x != nil; x = x.forward[0] {
		f(x.key, x.value)
	}
}

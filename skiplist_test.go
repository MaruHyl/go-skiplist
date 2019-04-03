package skiplist

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func intCompare(a interface{}, b interface{}) int {
	return a.(int) - b.(int)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func checkSkipList(t *testing.T, sl *SkipList, dataMap map[int]int) {
	// check len
	require.Equal(t, len(dataMap), sl.Len())
	length := 0
	sl.Foreach(func(_ interface{}, _ interface{}) {
		length++
	})
	require.Equal(t, len(dataMap), length)

	// check sorted
	var lastKey = -1
	sl.Foreach(func(key interface{}, _ interface{}) {
		intkey := key.(int)
		require.True(t, intkey > lastKey)
		lastKey = intkey
	})

	// check key-value
	sl.Foreach(func(key interface{}, value interface{}) {
		v, ok := dataMap[key.(int)]
		require.True(t, ok)
		require.Equal(t, v, value)
	})

	for k, v := range dataMap {
		lv, ok := sl.Search(k)
		require.True(t, ok)
		require.Equal(t, v, lv)
	}

	// check every level
	for i := 0; i < maxLevel; i++ {
		x := sl.header
		for x.forward[i] != nil {
			if x != sl.header {
				require.True(t, intCompare(x.key, x.forward[i].key) < 0)
			}
			x = x.forward[i]
		}
	}

	t.Log("check pass")

	if testing.Short() {
		t.Log(print(sl))
	}
}

func TestSkipList(t *testing.T) {
	//
	sl := New(intCompare)
	// random data
	testCount := 10000
	if testing.Short() {
		testCount = 10
	}
	data := make([]int, 0, testCount)
	dataMap := make(map[int]int, testCount)
	for i := 0; i < testCount; i++ {
		d := rand.Intn(testCount / 2)
		data = append(data, d)
		dataMap[d] = d
	}
	// insert
	for _, i := range data {
		sl.Insert(i, i)
	}
	checkSkipList(t, sl, dataMap)
	t.Log("after insert", testCount, sl.Len())
	// random delete
	maxDeleteCount := sl.Len() / 2
	deleteCount := 0
	for k := range dataMap {
		if deleteCount >= maxDeleteCount {
			break
		}
		delete(dataMap, k)
		sl.Delete(k)
		deleteCount++
	}
	checkSkipList(t, sl, dataMap)
	t.Log("after delete", testCount, sl.Len())
}

func print(l *SkipList) string {
	var b strings.Builder
	for i := 0; i < maxLevel; i++ {
		x := l.header
		for x != nil {
			if x == l.header {
				fmt.Fprintf(&b, "[%d] header -> ", i)
			} else {
				fmt.Fprintf(&b, "(%d:%d) -> ", x.key, x.value)
			}
			x = x.forward[i]
		}
		fmt.Fprintf(&b, "\n")
	}
	return b.String()
}

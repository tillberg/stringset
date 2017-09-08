package stringset

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeSets(N int, equal bool) (a, b *StringSet) {
	a, b = New(), New()
	for i := 0; i < N; i++ {
		s := fmt.Sprintf("%X", rand.Int63()) // random string
		a.Add(s)
		if equal {
			b.Add(s)
		} else {
			b.Add(fmt.Sprintf("%X", rand.Int63())) // different random string
		}
	}
	return
}

func TestMakeSet(t *testing.T) {
	N := 3
	a, b := makeSets(N, true)
	t.Logf("makeSet ret a %v, b %v", a, b)
	assert.Len(t, a.strMap, N)
	assert.Len(t, b.strMap, N)
}

func TestEquals(t *testing.T) {
	N := 1000
	a, b := makeSets(N, true)
	assert.Equal(t, a.Len(), N, "makeSets error")
	assert.Equal(t, b.Len(), N, "makeSets error")
	assert.True(t, a.Equal(b), "sets with same elements should be equal")
	assert.True(t, b.Equal(a), "sets with same elements should be equal")
	assert.True(t, a.Equal(a), "set should be equal with itself")

	x, y := New(), New()
	assert.True(t, x.Equal(y), "two empty sets should be equal")
	x.Add("1")
	assert.False(t, x.Equal(y))
	assert.False(t, y.Equal(x))
	y.Add("1")
	assert.True(t, x.Equal(y))
	assert.True(t, y.Equal(x))
}

func TestUnequal(t *testing.T) {
	N := 1000
	a, b := makeSets(N, false)
	assert.Equal(t, a.Len(), N, "makeSets error")
	assert.Equal(t, b.Len(), N, "makeSets error")
	assert.False(t, a.Equal(b), "error checking equality of different sets")

	c, d := makeSets(N*2, true)
	assert.Equal(t, c.Len(), N*2, "makeSets error")
	assert.Equal(t, d.Len(), N*2, "makeSets error")
	assert.True(t, c.Equal(d), "error checking equality of identical sets")
	assert.False(t, a.Equal(c), "differing sets returned equal")
	assert.False(t, a.Equal(d), "differing sets returned equal")
	assert.False(t, b.Equal(c), "differing sets returned equal")
	assert.False(t, b.Equal(d), "differing sets returned equal")
}

func TestDifference(t *testing.T) {
	a := &StringSet{}
	b := &StringSet{}
	a.Clear()
	b.Clear()

	a.Add("a")
	a.Add("b")
	a.Add("c")

	b.Add("b")
	b.Add("c")
	b.Add("d")

	c := a.Difference(b)
	assert.Equal(t, 1, c.Len())
	assert.True(t, c.Has("a"))
}

func TestUnion(t *testing.T) {
	a := &StringSet{}
	b := &StringSet{}
	a.Clear()
	b.Clear()

	a.Add("a")
	a.Add("b")
	a.Add("c")

	b.Add("b")
	b.Add("c")
	b.Add("d")

	c := a.Union(b)
	assert.Equal(t, 4, c.Len())
	assert.True(t, c.Has("a"))
	assert.True(t, c.Has("b"))
	assert.True(t, c.Has("c"))
	assert.True(t, c.Has("d"))
}

func TestIntersection(t *testing.T) {
	a := &StringSet{}
	b := &StringSet{}
	a.Clear()
	b.Clear()

	a.Add("a")
	a.Add("b")
	a.Add("c")

	b.Add("b")
	b.Add("c")
	b.Add("d")

	c := a.Intersection(b)
	assert.Equal(t, 2, c.Len())
	assert.True(t, c.Has("b"))
	assert.True(t, c.Has("c"))
}

var benchSz = 10000

func BenchmarkEqualMaps(b *testing.B) {
	x, y := makeSets(benchSz, true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Equal(y)
	}
}
func BenchmarkUnequalMaps(b *testing.B) {
	x, y := makeSets(benchSz, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Equal(y)
	}
}

// BenchmarkEqualsReflect shows reflect.DeepEqual is slower
func BenchmarkEqualMapsReflect(b *testing.B) {
	x, y := makeSets(benchSz, true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(x.strMap, y.strMap)
	}
}
func BenchmarkUnequalMapsReflect(b *testing.B) {
	x, y := makeSets(benchSz, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(x.strMap, y.strMap)
	}
}

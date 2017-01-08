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
	assert.True(t, a.Equal(b))
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

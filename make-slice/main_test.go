package slice

import (
	"fmt"
	"testing"
)

func BenchmarkMakeSlice(b *testing.B) {
	benchmarks := []struct {
		makeNum  int
		breakNum int
	}{
		{1, 0},
		{1, 1},
		{10, 0},
		{10, 1},
		{10, 10},
		{100, 0},
		{100, 1},
		{100, 10},
		{100, 100},
		{1000, 0},
		{1000, 1},
		{1000, 10},
		{1000, 100},
		{1000, 1000},
		{10000, 0},
		{10000, 1},
		{10000, 10},
		{10000, 100},
		{10000, 1000},
		{10000, 10000},
	}

	for _, bm := range benchmarks {
		name := fmt.Sprintf("len: %dx%d", bm.makeNum, bm.breakNum)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(bm.makeNum, bm.breakNum)
			}
		})
	}

	for _, bm := range benchmarks {
		name := fmt.Sprintf("app: %dx%d", bm.makeNum, bm.breakNum)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				g(bm.makeNum, bm.breakNum)
			}
		})
	}
}

func f(makeNum, breakNum int) {
	a := make([]int, makeNum)
	for i := 0; i < makeNum; i++ {
		a[i] = i
		if i >= breakNum {
			break
		}
	}
}

func g(makeNum, breakNum int) {
	a := make([]int, 0)
	for i := 0; i < makeNum; i++ {
		a = append(a, i)
		if i >= breakNum {
			break
		}
	}
}

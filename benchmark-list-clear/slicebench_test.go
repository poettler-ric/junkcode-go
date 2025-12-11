package slicebench

import (
	"math/rand"
	"testing"
)

func prepSlice(n int) []int {
	s := make([]int, 0, n)
	for i := range n {
		s = append(s, i)
	}
	return s
}

func BenchmarkResetSlice_SliceZero(b *testing.B) {
	s := prepSlice(1_000_000)

	sizes := make([]int, b.N)
	r := rand.New(rand.NewSource(1))
	for i := range sizes {
		sizes[i] = r.Intn(1000)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := range b.N {
		s = s[:0]
		m := sizes[i]
		for j := range m {
			s = append(s, j)
		}
	}
}

func BenchmarkResetSlice_Nil(b *testing.B) {
	s := prepSlice(1_000_000)

	sizes := make([]int, b.N)
	r := rand.New(rand.NewSource(1))
	for i := range sizes {
		sizes[i] = r.Intn(1000)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := range b.N {
		s = nil
		m := sizes[i]
		for j := range m {
			s = append(s, j)
		}
	}
}

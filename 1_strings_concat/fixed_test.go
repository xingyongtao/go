package strings_concat

import (
	"strings"
	"testing"
)
var preventCompilerOptimization int // http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

func BenchmarkAppendEmpty(b *testing.B) {
	sl := make([]byte, 0, 0)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sl = append(sl, 'x')
	}
	b.StopTimer()
	preventCompilerOptimization = len(sl)

	if s := strings.Repeat("x", b.N); string(sl) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(sl), s)
	}
}

func BenchmarkAppendPreAlloc(b *testing.B) {
	sl := make([]byte, 0, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sl = append(sl, 'x')
	}
	b.StopTimer()
	preventCompilerOptimization = len(sl)

	if s := strings.Repeat("x", b.N); string(sl) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(sl), s)
	}
}

func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], "x")
	}
	b.StopTimer()
	preventCompilerOptimization = bl

	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
	}
}

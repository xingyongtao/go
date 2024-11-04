package strings_concat

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	var str string
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		str += "x"
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	var sb strings.Builder

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sb.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); sb.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", sb.String(), s)
	}
}

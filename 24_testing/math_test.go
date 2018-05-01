package math

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"testing/quick"
)

func TestAdder(t *testing.T) {
	result := Adder(4, 7)
	if result != 11 {
		t.Fatal("4 + 7 did not equal 11")
	}
}

func BenchmarkAdder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adder(4, 7)
	}
}

func ExampleAdder() {
	fmt.Println(Adder(4, 7))
	// Output:
	// 11
}

func ExampleAdder_multiple() {
	fmt.Println(Adder(3, 6, 7, 4, 61))
	// Output:
	// 81
}

func TestAdderBlackbox(t *testing.T) {
	err := quick.Check(a, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func a(x, y int) bool {
	return Adder(x, y) == x+y
}

func BenchmarkConcat(b *testing.B) {
	var str string
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
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
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

	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
	}
}

// Go 1.10
func BenchmarkStringBuilder(b *testing.B) {
	var strBuilder strings.Builder

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		strBuilder.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); strBuilder.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", strBuilder.String(), s)
	}
}

package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("no repeat provided", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"
		assertCorrectRepeatCount(t, repeated, expected)
	})

	t.Run("repeat 6 times", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"
		assertCorrectRepeatCount(t, repeated, expected)
	})
}

func assertCorrectRepeatCount(t testing.TB, repeated, expected string) {
	t.Helper()
	if strings.Compare(expected, repeated) != 0 {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 10)
	fmt.Println(repeated)
	// Output: aaaaaaaaaa
}
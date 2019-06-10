package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := RepeatFiveTimes("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' got '%s'", expected, repeated)
	}
}

func BenchMarkRepeatFiveTimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatFiveTimes("a")
	}
}

func ExampleRepeatFiveTimes() {
	repeated := RepeatFiveTimes("g")
	fmt.Println(repeated)
	// Output: "ggggg"
}

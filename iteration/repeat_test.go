package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := RepeatFiveTimes("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' got '%s'", expected, repeated)
	}
}

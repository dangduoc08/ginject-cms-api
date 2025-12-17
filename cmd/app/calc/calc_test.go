package calc

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 5)
	want := 6

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

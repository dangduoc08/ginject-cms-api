package calc

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

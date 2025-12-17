package calc

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 3)
	want := 4

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

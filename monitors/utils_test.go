package monitors

import "testing"

func Test_trimLeft(t *testing.T) {
	in := []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	in = trimLeft(in, 5)
	if len(in) != 5 {
		t.Errorf("expected %d, was %d", 5, len(in))
	}
	if in[0] != 5 {
		t.Errorf("expected %d, was %f", 5, in[0])
	}
	if in[4] != 9 {
		t.Errorf("expected %d, was %f", 9, in[4])
	}
}

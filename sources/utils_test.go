package sources

import (
	"testing"
)

func Test_bInt64(t *testing.T) {
	tests := []struct {
		in   string
		want int64
	}{
		{
			in:   "651438501",
			want: 651438501,
		},
		{
			in:   "0",
			want: 0,
		},
		{
			in:   "110010029399294",
			want: 110010029399294,
		},
		{
			in:   "2 3 4 5",
			want: 2345,
		},
		{
			in:   "thomas",
			want: 0,
		},
	}

	for _, c := range tests {
		if r := bInt64([]byte(c.in)); r != c.want {
			t.Errorf("expected %d, got %d", c.want, r)
		}
	}
}

func Test_readInts(t *testing.T) {
	line := "Mem:         1893616     1078816       70800        1268      744000      774856"
	ints := readInts(line, 6)
	if ints[0] != 1893616 {
		t.Errorf("expected %d, got %d", 1893616, ints[0])
	}
	if ints[1] != 1078816 {
		t.Errorf("expected %d, got %d", 1078816, ints[1])
	}
	if ints[2] != 70800 {
		t.Errorf("expected %d, got %d", 70800, ints[2])
	}
	if ints[3] != 1268 {
		t.Errorf("expected %d, got %d", 1268, ints[3])
	}
	if ints[4] != 744000 {
		t.Errorf("expected %d, got %d", 744000, ints[4])
	}
	if ints[5] != 774856 {
		t.Errorf("expected %d, got %d", 774856, ints[5])
	}

	ints = readInts(line, 1)
	if len(ints) != 1 {
		t.Errorf("expected len %d, got %d", 1, len(ints))
	}
	if ints[0] != 1893616 {
		t.Errorf("expected %d, got %d", 1893616, ints[0])
	}
}

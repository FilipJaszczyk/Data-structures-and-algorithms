package searching

import "testing"

type searchTest struct {
	list   []int
	number int
	want   bool
}

var (
	list = []int{1, 2, 3, 5, 6, 7, 8, 10, 12, 14, 16}
)

func TestBinarySearch(t *testing.T) {
	tests := []searchTest{
		{list, 4, false},
		{list, 12, true},
	}
	for _, tc := range tests {
		if res := BinarySearch(tc.list, tc.number); res != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, res)
		}
	}
}

func TestRecursiveBinarySearch(t *testing.T) {
	tests := []searchTest{
		{list, 4, false},
		{list, 12, true},
	}
	for _, tc := range tests {
		if res := RecursiveBinarySearch(tc.list, tc.number); res != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, res)
		}
	}
}

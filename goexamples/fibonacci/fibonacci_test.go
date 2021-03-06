package fibonacci

import (
	"reflect"
	"testing"
	"time"
)

type fibTest struct {
	n       int
	fibFunc func(n int) []int
	want    []int
}

type fibTimeoutTest struct {
	n       int
	timeout time.Duration
	fibFunc func(n int, timeout time.Duration) []int
	want    []int
}

func TestFibonacciNums(t *testing.T) {
	tests := []fibTest{
		{
			1, fibbonaciNums, []int{0},
		},
		{
			4, fibbonaciNums, []int{0, 1, 1, 2},
		},
		{
			6, fibbonaciNums, []int{0, 1, 1, 2, 3, 5},
		},
	}
	for _, tc := range tests {
		if res := tc.fibFunc(tc.n); !reflect.DeepEqual(tc.want, res) {
			t.Fatalf("expected: %#v, got: %#v", tc.want, res)
		}
	}
}

func TestFibonacciLimit(t *testing.T) {
	tests := []fibTest{
		{
			1, fibbonaciLimit, []int{0, 1, 1},
		},
		{
			4, fibbonaciLimit, []int{0, 1, 1, 2, 3},
		},
		{
			25, fibbonaciLimit, []int{0, 1, 1, 2, 3, 5, 8, 13, 21},
		},
	}
	for _, tc := range tests {
		if res := tc.fibFunc(tc.n); !reflect.DeepEqual(tc.want, res) {
			t.Fatalf("expected: %#v, got: %#v", tc.want, res)
		}
	}
}

func TestFibonacciLimitTimeout(t *testing.T) {
	tests := []fibTimeoutTest{
		{
			4, time.Second * 1, fibbonaciLimitTimeout, []int{0, 1, 1, 2, 3},
		},
		{
			25, time.Nanosecond, fibbonaciLimitTimeout, []int{},
		},
	}
	for _, tc := range tests {
		if res := tc.fibFunc(tc.n, tc.timeout); !reflect.DeepEqual(res, tc.want) {
			t.Fatalf("expected empty list, got: %#v", res)
		}
	}
}

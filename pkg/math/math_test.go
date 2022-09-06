package math

import (
	"reflect"
	"testing"
)

func TestMultipleMin(t *testing.T) {
	tests := []struct {
		in  []float64
		exp float64
	}{
		{
			[]float64{1, 2, 3, 4},
			1,
		},
		{
			[]float64{4, 3, 2, 1},
			1,
		},
		{
			[]float64{0, -1, 1},
			-1,
		},
		{
			[]float64{-999, -888, -777},
			-999,
		},
		{
			[]float64{-.2, -.3},
			-.3,
		},
		{
			[]float64{0},
			0,
		},
	}

	for _, test := range tests {
		res := MultipleMin(test.in[0], test.in[1:]...)
		if !reflect.DeepEqual(res, test.exp) {
			t.Fatalf("expected %v, received %v", test.exp, res)
		}
	}
}

func TestMultipleMax(t *testing.T) {
	tests := []struct {
		in  []int
		exp int
	}{
		{
			[]int{1, 2, 3, 4},
			4,
		},
		{
			[]int{4, 3, 2, 1},
			4,
		},
		{
			[]int{0, -1, 1},
			1,
		},
		{
			[]int{-999, -888, -777},
			-777,
		},
		{
			[]int{-2, -3},
			-2,
		},
		{
			[]int{0},
			0,
		},
	}

	for _, test := range tests {
		res := MultipleMax(test.in[0], test.in[1:]...)
		if !reflect.DeepEqual(res, test.exp) {
			t.Fatalf("expected %v, received %v", test.exp, res)
		}
	}
}

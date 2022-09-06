package array

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	testsInt := []struct {
		input   []int
		compare func(int) bool
		exp     []int
	}{
		{
			[]int{1, 2, 3, 4},
			func(item int) bool { return item > 2 },
			[]int{3, 4},
		}, {
			[]int{1, 2, 3, 4},
			func(item int) bool { return item > 0 },
			[]int{1, 2, 3, 4},
		}, {
			[]int{-1, 0, 1},
			func(item int) bool { return item >= 0 },
			[]int{0, 1},
		}, {
			[]int{9999, 1, 0, -1, -99999},
			func(item int) bool { return item < 0 },
			[]int{-1, -99999},
		}, {
			[]int{0},
			func(item int) bool { return item < 0 },
			[]int{},
		}, {
			[]int{},
			func(item int) bool { return item < 0 },
			[]int{},
		},
	}

	for _, test := range testsInt {
		res := Filter(test.input, test.compare)
		if !reflect.DeepEqual(res, test.exp) {
			t.Fatalf("received %v, expected %v", res, test.exp)
		}
	}

	testsString := []struct {
		input   []string
		compare func(string) bool
		exp     []string
	}{
		{
			[]string{"a", "ab", "abc"},
			func(item string) bool { return len(item) > 1 },
			[]string{"ab", "abc"},
		}, {
			[]string{},
			func(item string) bool { return len(item) > 1 },
			[]string{},
		},
	}

	for _, test := range testsString {
		res := Filter(test.input, test.compare)
		if !reflect.DeepEqual(res, test.exp) {
			t.Fatalf("received %v, expected %v", res, test.exp)
		}
	}
}

func TestMap(t *testing.T) {
	testsInt := []struct {
		input  []int
		mapper func(int) int
		exp    []int
	}{
		{
			[]int{1, 2, 3, 4},
			func(item int) int { return item + 1 },
			[]int{2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3},
			func(item int) int { return item * item },
			[]int{1, 4, 9},
		},
		{
			[]int{},
			func(item int) int { return item + 1 },
			[]int{},
		},
	}

	for _, test := range testsInt {
		res := Map(test.input, test.mapper)
		if !reflect.DeepEqual(res, test.exp) {
			t.Fatalf("received %v, expected %v", res, test.exp)
		}
	}

	testsString := []struct {
		input  []string
		mapper func(string) string
		exp    []string
	}{
		{
			[]string{"a", "b"},
			func(item string) string { return item + "." },
			[]string{"a.", "b."},
		},
		{
			[]string{},
			func(item string) string { return item + "." },
			[]string{},
		},
	}

	for _, test := range testsString {
		res := Map(test.input, test.mapper)
		if !reflect.DeepEqual(res, test.exp) {
			t.Fatalf("received %v, expected %v", res, test.exp)
		}
	}
}

func TestReduce(t *testing.T) {
	testsInt := []struct {
		input   []int
		reducer func(int, int, int) int
		exp     int
	}{
		{
			[]int{1, 2, 3},
			func(acc, item, key int) int {
				return acc + item
			},
			6,
		},
		{
			[]int{1, 2, 3},
			func(acc, item, key int) int {
				if key > 1 {
					return acc
				} else {
					return acc + item
				}
			},
			3,
		},
		{
			[]int{},
			func(acc, item, key int) int {
				return acc + item
			},
			0,
		},
	}

	for _, test := range testsInt {
		res := Reduce(test.input, test.reducer)
		if !reflect.DeepEqual(res, test.exp) {
			t.Fatalf("received %v, expected %v", res, test.exp)
		}
	}
}

func TestContains(t *testing.T) {
	testsInt := []struct {
		input   []int
		compare int
		exp     bool
	}{
		{
			[]int{1, 2},
			2,
			true,
		}, {
			[]int{1, 2},
			3,
			false,
		}, {
			[]int{},
			2,
			false,
		}, {
			[]int{1, 2, 3, 4, 5, 6},
			1,
			true,
		}, {
			[]int{1, 2, 3, 4, 5, 6},
			3,
			true,
		}, {
			[]int{1, 2, 3, 4, 5, 6},
			6,
			true,
		},
	}

	for _, test := range testsInt {
		res := Contains(test.input, test.compare)
		if !reflect.DeepEqual(res, test.exp) {
			t.Fatalf("received %v, expected %v", res, test.exp)
		}
	}

	testsString := []struct {
		input   []string
		compare string
		exp     bool
	}{
		{
			[]string{"a", "ab", "abc"},
			"a",
			true,
		}, {
			[]string{"a", "ab", "abc"},
			"b",
			false,
		}, {
			[]string{},
			"b",
			false,
		},
	}

	for _, test := range testsString {
		res := Contains(test.input, test.compare)
		if !reflect.DeepEqual(res, test.exp) {
			t.Fatalf("received %v, expected %v", res, test.exp)
		}
	}
}

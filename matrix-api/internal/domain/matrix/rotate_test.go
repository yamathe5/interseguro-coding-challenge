package matrix

import (
	"reflect"
	"testing"
)

func TestClockwise(t *testing.T) {
	for _, tc := range []struct{ in, want [][]float64 }{
		{[][]float64{{1, 2, 3}, {4, 5, 6}}, [][]float64{{4, 1}, {5, 2}, {6, 3}}},
		{[][]float64{{4, 1}, {5, 2}, {6, 3}}, [][]float64{{6, 5, 4}, {3, 2, 1}}},
		{[][]float64{{1, 2}}, [][]float64{{1}, {2}}},
		{[][]float64{{1}, {2}}, [][]float64{{2, 1}}},
		{[][]float64{{1}}, [][]float64{{1}}},
	} {
		out := Clockwise(tc.in)
		if !reflect.DeepEqual(out, tc.want) {
			t.Fatalf("got %v want %v", out, tc.want)
		}
		if !reflect.DeepEqual(Clockwise(Clockwise(Clockwise(out))), tc.in) {
			t.Fatal("four turns must restore input")
		}
	}
}

package gonum

import (
	"math"
	"testing"
)

func TestQR(t *testing.T) {
	for _, a := range [][][]float64{
		{{1, 2}, {3, 4}}, {{1, 2}, {3, 4}, {5, 6}}, {{1, 2, 3}, {4, 5, 6}},
		{{0, 0, 3}, {0, 0, -5}}, {{1, 0}, {0, 1}}, {{0, 0}, {0, 0}}, {{-1}, {-2}}, {{1, 2}, {2, 4}, {3, 6}},
		{{0.00001, 2e-5, 3e-5}}, {{1e200, -1e200}, {1e200, 1e200}},
	} {
		m, n := len(a), len(a[0])
		k := min(m, n)
		q, r, err := (QR{}).Factorize(a)
		if err != nil || len(q) != m || len(q[0]) != k || len(r) != k || len(r[0]) != n {
			t.Fatalf("shape or error: %v", err)
		}
		scale := 0.0
		for _, row := range a {
			for _, value := range row {
				scale = math.Max(scale, math.Abs(value))
			}
		}
		if scale == 0 {
			scale = 1
		}
		var residual, norm float64
		for i, row := range a {
			for j, value := range row {
				product := 0.0
				for h := 0; h < k; h++ {
					product += q[i][h] * (r[h][j] / scale)
				}
				residual = math.Hypot(residual, value/scale-product)
				norm = math.Hypot(norm, value/scale)
			}
		}
		if residual > 1e-9*norm {
			t.Fatalf("reconstruction %v: %g", a, residual)
		}
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				dot, want := 0.0, 0.0
				if i == j {
					want = 1
				}
				for h := 0; h < m; h++ {
					dot += q[h][i] * q[h][j]
				}
				if math.Abs(dot-want) > 1e-9 {
					t.Fatalf("not orthogonal: %v", a)
				}
				if i > j && r[i][j] != 0 {
					t.Fatal("R is not upper triangular")
				}
			}
		}
	}
}

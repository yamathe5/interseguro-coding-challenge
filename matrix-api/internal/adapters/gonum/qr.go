package gonum

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mat"
)

type QR struct{}

func (QR) Factorize(values [][]float64) ([][]float64, [][]float64, error) {
	rows, cols := len(values), len(values[0])
	k := min(rows, cols)
	leading := mat.NewDense(rows, k, nil)
	scale := 0.0
	for _, row := range values {
		for _, value := range row {
			scale = math.Max(scale, math.Abs(value))
		}
	}
	if scale == 0 {
		scale = 1
	}
	for i, row := range values {
		for j := 0; j < k; j++ {
			leading.Set(i, j, row[j]/scale)
		}
	}
	var decomposition mat.QR
	decomposition.Factorize(leading)
	var q, r mat.Dense
	decomposition.QTo(&q)
	decomposition.RTo(&r)
	qRows := make([][]float64, rows)
	rRows := make([][]float64, k)
	for i := range qRows {
		qRows[i] = make([]float64, k)
		for j := range qRows[i] {
			qRows[i][j] = q.At(i, j)
		}
	}
	for i := range rRows {
		rRows[i] = make([]float64, cols)
		for j := i; j < cols; j++ {
			value := 0.0
			if j < k {
				value = r.At(i, j)
			} else {
				// For a wide matrix the square Q spans every remaining column.
				for a := 0; a < rows; a++ {
					value += q.At(a, i) * (values[a][j] / scale)
				}
			}
			value *= scale
			if math.IsNaN(value) || math.IsInf(value, 0) {
				return nil, nil, errors.New("QR result exceeds numeric range")
			}
			rRows[i][j] = value
		}
	}
	return qRows, rRows, nil
}

package application

import (
	"context"
	"errors"
)

var ErrOperation = errors.New("Operation must be qr or rotate")
var ErrDependency = errors.New("Statistics service failed")

type Statistics struct {
	Max              float64 `json:"max"`
	Min              float64 `json:"min"`
	Average          float64 `json:"average"`
	Sum              float64 `json:"sum"`
	AnyDiagonal      bool    `json:"anyDiagonal"`
	DiagonalMatrices []int   `json:"diagonalMatrices"`
	MatrixCount      int     `json:"matrixCount"`
}
type StatisticsClient interface {
	Analyze(context.Context, [][][]float64, string) (Statistics, error)
	Ready(context.Context) error
}
type Analysis struct {
	Operation  string     `json:"operation"`
	Result     any        `json:"result"`
	Statistics Statistics `json:"statistics"`
}

func AnalyzeMatrix(ctx context.Context, values [][]float64, operation, id string, f Factorizer, client StatisticsClient) (Analysis, error) {
	var matrices [][][]float64
	var result any
	switch operation {
	case "qr":
		factors, err := FactorizeQR(values, f)
		if err != nil {
			return Analysis{}, err
		}
		result, matrices = factors, [][][]float64{factors.Q, factors.R}
	case "rotate":
		rotated := RotateMatrix(values)
		result, matrices = rotated, [][][]float64{rotated.Matrix}
	default:
		return Analysis{}, ErrOperation
	}
	stats, err := client.Analyze(ctx, matrices, id)
	if err != nil {
		return Analysis{}, ErrDependency
	}
	return Analysis{operation, result, stats}, nil
}

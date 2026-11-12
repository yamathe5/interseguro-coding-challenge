package application

import "interseguro/matrix-api/internal/domain/matrix"

type Rotation struct {
	Matrix  [][]float64 `json:"matrix"`
	Degrees int         `json:"rotationDegrees"`
}

func RotateMatrix(values [][]float64) Rotation { return Rotation{matrix.Clockwise(values), 90} }

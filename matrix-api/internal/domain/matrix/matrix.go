package matrix

import (
	"encoding/json"
	"errors"
	"math"
)

const MaxDimension = 100

// Decode validates the wire representation before converting to float64.
func Decode(raw json.RawMessage) ([][]float64, error) {
	var rows [][]*float64
	invalid := errors.New("Expected a rectangular matrix of finite numbers up to 100×100")
	if json.Unmarshal(raw, &rows) != nil || len(rows) == 0 || len(rows) > MaxDimension || len(rows[0]) == 0 || len(rows[0]) > MaxDimension {
		return nil, invalid
	}
	result := make([][]float64, len(rows))
	for i, row := range rows {
		if len(row) != len(rows[0]) {
			return nil, invalid
		}
		result[i] = make([]float64, len(row))
		for j, value := range row {
			if value == nil || math.IsInf(*value, 0) || math.IsNaN(*value) {
				return nil, invalid
			}
			result[i][j] = *value
		}
	}
	return result, nil
}

package matrix

func Clockwise(values [][]float64) [][]float64 {
	rows, cols := len(values), len(values[0])
	result := make([][]float64, cols)
	for j := range result {
		result[j] = make([]float64, rows)
		for i := range values {
			result[j][rows-1-i] = values[i][j]
		}
	}
	return result
}

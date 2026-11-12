package application

type Factorizer interface {
	Factorize([][]float64) ([][]float64, [][]float64, error)
}

type Shape struct {
	Rows int `json:"rows"`
	Cols int `json:"cols"`
}
type QR struct {
	Q          [][]float64 `json:"q"`
	R          [][]float64 `json:"r"`
	Shape      Shape       `json:"shape"`
	Convention string      `json:"convention"`
}

func FactorizeQR(values [][]float64, f Factorizer) (QR, error) {
	q, r, err := f.Factorize(values)
	return QR{q, r, Shape{len(values), len(values[0])}, "thin"}, err
}

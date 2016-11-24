package mat64f

type TYPE float64

type Matrix interface {
	Dims() (r, c int)

	At(i, j int) float64

	T() Matrix
}

var (
	_	Matrix		= Transpose{}
	_	Untransposer	= Transpose{}
)

type Transpose struct {
	Matrix Matrix
}

func (t Transpose) At(i, j int) float64 {
	return t.Matrix.At(j, i)
}

func (t Transpose) Dims() (r, c int) {
	c, r = t.Matrix.Dims()
	return r, c
}

func (t Transpose) T() Matrix {
	return t.Matrix
}

func (t Transpose) Untranspose() Matrix {
	return t.Matrix
}

type Untransposer interface {
	Untranspose() Matrix
}

type Mutable interface {
	Set(i, j int, v float64)

	Matrix
}
